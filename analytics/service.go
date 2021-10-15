package analytics

import (
	"fmt"
	"sort"

	"github.com/montanaflynn/stats"
	"github.com/msoerjanto/fantasy-helper/bballref"
	"github.com/msoerjanto/fantasy-helper/yahoo"
	"golang.org/x/oauth2"
)

type DataTableResponse struct {
	PlayerAverages []PlayerAverages
	Aggregates     AggregateData
}

type LeagueDataResponse struct {
	Teams []Team
}

type Team struct {
	Manager Manager
	Players []PlayerAverages
}

type Manager struct {
	Nickname       string
	IsCurrentLogin bool
}

type AnalyticsService interface {
	GetPlayerAveragesBySeason(season int, punt PuntCategories, numConsidered int) DataTableResponse
	GetLeagueData(league string, token *oauth2.Token) LeagueDataResponse
}

type analyticsService struct {
	bballrefService bballref.BasketballRefService
	yahooService    yahoo.YahooService
}

func NewAnalyticsService(bballrefService bballref.BasketballRefService,
	yahooService yahoo.YahooService) AnalyticsService {
	return &analyticsService{
		bballrefService: bballrefService,
		yahooService:    yahooService,
	}
}

func (s *analyticsService) GetLeagueData(league string, token *oauth2.Token) LeagueDataResponse {
	client := s.yahooService.NewYahooClient(token)

	var result LeagueDataResponse

	// get teams from league
	teams, err := client.GetAllTeams(league)
	if err != nil {
		fmt.Println("Error fetching teams from Yahoo")
		return LeagueDataResponse{}
	}

	var resultTeams []Team
	// get team rosters for each team
	for i := 0; i < len(teams); i++ {
		currTeam := Team{}

		currTeam.Manager = convertManagerToManager(teams[i].Managers[0])

		teamKey := teams[i].TeamKey
		roster, err := client.GetTeamRoster(teamKey)
		if err != nil {
			fmt.Println("Error fetching roster from Yahoo")
		}

		var currTeamPlayers []PlayerAverages
		for j := 0; j < len(roster); j++ {
			currTeamPlayers = append(currTeamPlayers, convertPlayerToPlayerAverages(roster[j]))
		}
		currTeam.Players = currTeamPlayers
		resultTeams = append(resultTeams, currTeam)
	}
	result.Teams = resultTeams
	return result

}

func convertManagerToManager(manager yahoo.Manager) Manager {
	return Manager{
		Nickname:       manager.Nickname,
		IsCurrentLogin: manager.IsCurrentLogin,
	}
}

func convertPlayerToPlayerAverages(player yahoo.Player) PlayerAverages {
	return PlayerAverages{
		Name: player.Name.Full,
	}
}

func (s *analyticsService) GetPlayerAveragesBySeason(season int, punt PuntCategories, numConsidered int) DataTableResponse {
	playerData := s.bballrefService.ParseData(season)

	aggData, result := InitializeData(playerData, numConsidered)
	processData(&aggData, result, punt)

	sort.Slice(result, func(i, j int) bool {
		return result[j].ZScore < result[i].ZScore
	})
	return DataTableResponse{
		PlayerAverages: result,
		Aggregates:     aggData,
	}
}

func InitializeData(playerData []bballref.PlayerAverages, numConsidered int) (AggregateData, []PlayerAverages) {
	var aggData AggregateData
	var result []PlayerAverages
	for i := 0; i < len(playerData); i++ {
		curr := convertPlayerAverages(playerData[i])
		curr.FantasyPts = computeFantasyPtsForPlayer(curr)
		result = append(result, curr)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[j].FantasyPts < result[i].FantasyPts
	})

	if numConsidered != -1 && numConsidered < len(result) {
		result = result[:numConsidered]
	}

	for i := 0; i < len(result); i++ {
		aggData.Age = append(aggData.Age, result[i].Age)
		aggData.GamesPlayed = append(aggData.GamesPlayed, result[i].GamesPlayed)
		aggData.GamesStarted = append(aggData.GamesStarted, result[i].GamesStarted)
		aggData.MinutesPerGame = append(aggData.MinutesPerGame, result[i].MinutesPerGame)
		aggData.PersonalFouls = append(aggData.PersonalFouls, result[i].PersonalFouls)

		aggData.FGMade = append(aggData.FGMade, result[i].FGMade)
		aggData.FGAttempted = append(aggData.FGAttempted, result[i].FGAttempted)
		aggData.FTMade = append(aggData.FTMade, result[i].FTMade)
		aggData.FTAttempted = append(aggData.FTAttempted, result[i].FTAttempted)

		aggData.FGPercentage = append(aggData.FGPercentage, result[i].FGPercentage)
		aggData.FTPercentage = append(aggData.FTPercentage, result[i].FTPercentage)
		aggData.ThreePTMade = append(aggData.ThreePTMade, result[i].ThreePTMade)
		aggData.TotalRebounds = append(aggData.TotalRebounds, result[i].TotalRebounds)
		aggData.Assists = append(aggData.Assists, result[i].Assists)
		aggData.Steals = append(aggData.Steals, result[i].Steals)
		aggData.Blocks = append(aggData.Blocks, result[i].Blocks)
		aggData.Turnovers = append(aggData.Turnovers, result[i].Turnovers)
		aggData.Points = append(aggData.Points, result[i].Points)
	}

	return aggData, result
}

func processData(aggData *AggregateData, result []PlayerAverages, punt PuntCategories) {
	computeSums(aggData)
	computeAverages(aggData)
	computeStandardDeviations(aggData)
	computeZScores(aggData, result, punt)
}

func computeSums(aggData *AggregateData) {
	// Need following totals for complex categories
	aggData.Totals.FGMade, _ = stats.Sum(aggData.FGMade)
	aggData.Totals.FGAttempted, _ = stats.Sum(aggData.FGAttempted)
	aggData.Totals.FTMade, _ = stats.Sum(aggData.FTMade)
	aggData.Totals.FTAttempted, _ = stats.Sum(aggData.FTAttempted)
}

func computeAverages(aggData *AggregateData) {
	// Averages for complex categories
	aggData.Averages.FGPercentage = aggData.Totals.FGMade / aggData.Totals.FGAttempted
	aggData.Averages.FTPercentage = aggData.Totals.FTMade / aggData.Totals.FTAttempted
	aggData.Averages.FGAttempted, _ = stats.Mean(aggData.FGAttempted)
	aggData.Averages.FTAttempted, _ = stats.Mean(aggData.FTAttempted)

	// Other categories
	aggData.Averages.Points, _ = stats.Mean(aggData.Points)
	aggData.Averages.TotalRebounds, _ = stats.Mean(aggData.TotalRebounds)
	aggData.Averages.Assists, _ = stats.Mean(aggData.Assists)
	aggData.Averages.Steals, _ = stats.Mean(aggData.Steals)
	aggData.Averages.Blocks, _ = stats.Mean(aggData.Blocks)
	aggData.Averages.ThreePTMade, _ = stats.Mean(aggData.ThreePTMade)
	aggData.Averages.Turnovers, _ = stats.Mean(aggData.Turnovers)
}

func computeStandardDeviations(aggData *AggregateData) {
	aggData.StandardDeviations.Points, _ = stats.StandardDeviation(aggData.Points)
	aggData.StandardDeviations.TotalRebounds, _ = stats.StandardDeviation(aggData.TotalRebounds)
	aggData.StandardDeviations.Assists, _ = stats.StandardDeviation(aggData.Assists)
	aggData.StandardDeviations.Steals, _ = stats.StandardDeviation(aggData.Steals)
	aggData.StandardDeviations.Blocks, _ = stats.StandardDeviation(aggData.Blocks)
	aggData.StandardDeviations.ThreePTMade, _ = stats.StandardDeviation(aggData.ThreePTMade)
	aggData.StandardDeviations.Turnovers, _ = stats.StandardDeviation(aggData.Turnovers)

	aggData.StandardDeviations.FGPercentage, _ = stats.StandardDeviation(aggData.FGPercentage)
	aggData.StandardDeviations.FTPercentage, _ = stats.StandardDeviation(aggData.FTPercentage)
}

func computeZScores(aggData *AggregateData, playerData []PlayerAverages, punt PuntCategories) {
	for i := 0; i < len(playerData); i++ {
		playerData[i].PtZ = computeZScoreBasic(playerData[i].Points, aggData.Averages.Points, aggData.StandardDeviations.Points, false)
		playerData[i].AstZ = computeZScoreBasic(playerData[i].Assists, aggData.Averages.Assists, aggData.StandardDeviations.Assists, false)
		playerData[i].RebZ = computeZScoreBasic(playerData[i].TotalRebounds, aggData.Averages.TotalRebounds, aggData.StandardDeviations.TotalRebounds, false)
		playerData[i].StlZ = computeZScoreBasic(playerData[i].Steals, aggData.Averages.Steals, aggData.StandardDeviations.Steals, false)
		playerData[i].BlkZ = computeZScoreBasic(playerData[i].Blocks, aggData.Averages.Blocks, aggData.StandardDeviations.Blocks, false)
		playerData[i].ThreeZ = computeZScoreBasic(playerData[i].ThreePTMade, aggData.Averages.ThreePTMade, aggData.StandardDeviations.ThreePTMade, false)
		playerData[i].ToZ = computeZScoreBasic(playerData[i].Turnovers, aggData.Averages.Turnovers, aggData.StandardDeviations.Turnovers, true)

		playerData[i].FgpZ = computeZScoreComplex(playerData[i].FGMade, playerData[i].FGAttempted,
			aggData.Averages.FGPercentage, aggData.StandardDeviations.FGPercentage, aggData.Averages.FGAttempted)
		playerData[i].FtpZ = computeZScoreComplex(playerData[i].FTMade, playerData[i].FTAttempted,
			aggData.Averages.FTPercentage, aggData.StandardDeviations.FTPercentage, aggData.Averages.FTAttempted)

		zscoreSum := getZscoreSum(playerData[i], punt)
		numCategories := getNumCategories(punt)

		playerData[i].ZScore = zscoreSum / float64(numCategories)
	}
}

func convertPlayerAverages(source bballref.PlayerAverages) PlayerAverages {
	return PlayerAverages{
		Name:           source.Name,
		Position:       source.Position,
		Age:            source.Age,
		Team:           source.Team,
		GamesPlayed:    source.GamesPlayed,
		GamesStarted:   source.GamesStarted,
		MinutesPerGame: source.MinutesPerGame,
		PersonalFouls:  source.PersonalFouls,

		FGMade:      source.FGMade,
		FGAttempted: source.FGAttempted,
		FTMade:      source.FTMade,
		FTAttempted: source.FTAttempted,

		FGPercentage:  source.FGPercentage,
		FTPercentage:  source.FTPercentage,
		ThreePTMade:   source.ThreePTMade,
		TotalRebounds: source.TotalRebounds,
		Assists:       source.Assists,
		Steals:        source.Steals,
		Blocks:        source.Blocks,
		Turnovers:     source.Turnovers,
		Points:        source.Points,
	}
}
