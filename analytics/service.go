package analytics

import (
	"sort"

	"github.com/montanaflynn/stats"
	"github.com/msoerjanto/fantasy-helper/bballref"
)

type PlayerAverages struct {
	Name     string
	Position string
	Team     string

	Age            float64
	GamesPlayed    float64
	GamesStarted   float64
	MinutesPerGame float64
	PersonalFouls  float64

	FGMade      float64
	FGAttempted float64
	FTMade      float64
	FTAttempted float64

	FGPercentage  float64
	FTPercentage  float64
	ThreePTMade   float64
	TotalRebounds float64
	Assists       float64
	Steals        float64
	Blocks        float64
	Turnovers     float64
	Points        float64

	AstZ   float64
	PtZ    float64
	RebZ   float64
	StlZ   float64
	BlkZ   float64
	FgpZ   float64
	FtpZ   float64
	ToZ    float64
	ThreeZ float64
	ZScore float64
}

type AggregateData struct {
	Age            []float64
	GamesPlayed    []float64
	GamesStarted   []float64
	MinutesPerGame []float64
	PersonalFouls  []float64

	FGMade      []float64
	FGAttempted []float64
	FTMade      []float64
	FTAttempted []float64

	FGPercentage  []float64
	FTPercentage  []float64
	ThreePTMade   []float64
	TotalRebounds []float64
	Assists       []float64
	Steals        []float64
	Blocks        []float64
	Turnovers     []float64
	Points        []float64

	Totals      PlayerAverages
	Averages    PlayerAverages
	StandardDev bballref.PlayerAverages
}

type AnalyticsService interface {
	GetPlayerAveragesBySeason(season int, categories ...string) []PlayerAverages
}

type analyticsService struct {
	bballrefService bballref.BasketballRefService
}

func NewAnalyticsService(bballrefService bballref.BasketballRefService) AnalyticsService {
	return &analyticsService{
		bballrefService: bballrefService,
	}
}

func (s *analyticsService) GetPlayerAveragesBySeason(season int, categories ...string) []PlayerAverages {
	playerData := s.bballrefService.ParseData(season)

	aggData, result := InitializeData(playerData)
	processData(&aggData, result)

	sort.Slice(result, func(i, j int) bool {
		return result[j].ZScore < result[i].ZScore
	})
	return result
}

func processData(aggData *AggregateData, result []PlayerAverages) {
	computeSums(aggData)
	computeAverages(aggData)
	computeStandardDeviations(aggData)
	computeZScores(aggData, result)
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
	aggData.StandardDev.Points, _ = stats.StandardDeviation(aggData.Points)
	aggData.StandardDev.TotalRebounds, _ = stats.StandardDeviation(aggData.TotalRebounds)
	aggData.StandardDev.Assists, _ = stats.StandardDeviation(aggData.Assists)
	aggData.StandardDev.Steals, _ = stats.StandardDeviation(aggData.Steals)
	aggData.StandardDev.Blocks, _ = stats.StandardDeviation(aggData.Blocks)
	aggData.StandardDev.ThreePTMade, _ = stats.StandardDeviation(aggData.ThreePTMade)
	aggData.StandardDev.Turnovers, _ = stats.StandardDeviation(aggData.Turnovers)

	aggData.StandardDev.FGPercentage, _ = stats.StandardDeviation(aggData.FGPercentage)
	aggData.StandardDev.FTPercentage, _ = stats.StandardDeviation(aggData.FTPercentage)
}

func computeZScores(aggData *AggregateData, playerData []PlayerAverages) {
	for i := 0; i < len(playerData); i++ {
		playerData[i].PtZ = computeZScoreBasic(playerData[i].Points, aggData.Averages.Points, aggData.StandardDev.Points, false)
		playerData[i].AstZ = computeZScoreBasic(playerData[i].Assists, aggData.Averages.Assists, aggData.StandardDev.Assists, false)
		playerData[i].RebZ = computeZScoreBasic(playerData[i].TotalRebounds, aggData.Averages.TotalRebounds, aggData.StandardDev.TotalRebounds, false)
		playerData[i].StlZ = computeZScoreBasic(playerData[i].Steals, aggData.Averages.Steals, aggData.StandardDev.Steals, false)
		playerData[i].BlkZ = computeZScoreBasic(playerData[i].Blocks, aggData.Averages.Blocks, aggData.StandardDev.Blocks, false)
		playerData[i].ThreeZ = computeZScoreBasic(playerData[i].ThreePTMade, aggData.Averages.ThreePTMade, aggData.StandardDev.ThreePTMade, false)
		playerData[i].ToZ = computeZScoreBasic(playerData[i].Turnovers, aggData.Averages.Turnovers, aggData.StandardDev.Turnovers, true)

		playerData[i].FgpZ = computeZScoreComplex(playerData[i].FGMade, playerData[i].FGAttempted,
			aggData.Averages.FGPercentage, aggData.StandardDev.FGPercentage, aggData.Averages.FGAttempted)
		playerData[i].FtpZ = computeZScoreComplex(playerData[i].FTMade, playerData[i].FTAttempted,
			aggData.Averages.FTPercentage, aggData.StandardDev.FTPercentage, aggData.Averages.FTAttempted)

		playerData[i].ZScore = (playerData[i].PtZ +
			playerData[i].AstZ +
			playerData[i].RebZ +
			playerData[i].StlZ +
			playerData[i].BlkZ +
			playerData[i].ThreeZ +
			playerData[i].FgpZ +
			playerData[i].FtpZ +
			playerData[i].ToZ) / 9
	}
}

func computeZScoreBasic(series float64, mean float64, std float64, invert bool) float64 {
	zscore := (series - mean) / std
	if invert {
		zscore = zscore * -1
	}
	return zscore
}

func computeZScoreComplex(seriesMake float64, seriesAttempted float64, mean float64, std float64, meanAttempts float64) float64 {
	var seriesPercentMade float64
	if seriesAttempted == 0 {
		seriesPercentMade = 0
	} else {
		seriesPercentMade = seriesMake / seriesAttempted
	}
	seriesDeltaFromAvg := seriesPercentMade - mean
	seriesZScore := seriesDeltaFromAvg / std
	seriesVolumeMult := seriesAttempted / meanAttempts
	return seriesZScore * seriesVolumeMult

}

func InitializeData(playerData []bballref.PlayerAverages) (AggregateData, []PlayerAverages) {
	var aggData AggregateData
	var result []PlayerAverages
	for i := 0; i < len(playerData); i++ {
		aggData.Age = append(aggData.Age, playerData[i].Age)
		aggData.GamesPlayed = append(aggData.GamesPlayed, playerData[i].GamesPlayed)
		aggData.GamesStarted = append(aggData.GamesStarted, playerData[i].GamesStarted)
		aggData.MinutesPerGame = append(aggData.MinutesPerGame, playerData[i].MinutesPerGame)
		aggData.PersonalFouls = append(aggData.PersonalFouls, playerData[i].PersonalFouls)

		aggData.FGMade = append(aggData.FGMade, playerData[i].FGMade)
		aggData.FGAttempted = append(aggData.FGAttempted, playerData[i].FGAttempted)
		aggData.FTMade = append(aggData.FTMade, playerData[i].FTMade)
		aggData.FTAttempted = append(aggData.FTAttempted, playerData[i].FTAttempted)

		aggData.FGPercentage = append(aggData.FGPercentage, playerData[i].FGPercentage)
		aggData.FTPercentage = append(aggData.FTPercentage, playerData[i].FTPercentage)
		aggData.ThreePTMade = append(aggData.ThreePTMade, playerData[i].ThreePTMade)
		aggData.TotalRebounds = append(aggData.TotalRebounds, playerData[i].TotalRebounds)
		aggData.Assists = append(aggData.Assists, playerData[i].Assists)
		aggData.Steals = append(aggData.Steals, playerData[i].Steals)
		aggData.Blocks = append(aggData.Blocks, playerData[i].Blocks)
		aggData.Turnovers = append(aggData.Turnovers, playerData[i].Turnovers)
		aggData.Points = append(aggData.Points, playerData[i].Points)

		result = append(result, convertPlayerAverages(playerData[i]))
	}
	return aggData, result
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
