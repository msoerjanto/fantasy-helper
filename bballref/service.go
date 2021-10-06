package bballref

import (
	"fmt"
	"strconv"

	"github.com/gocolly/colly"
)

type PlayerAverages struct {
	Name           string
	Position       string
	Age            float64
	Team           string
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
}

// var dataStatKeyFloat = [...]string{"mp_per_g", "fg_per_g", "fga_per_g", "ft_per_g", "fta_per_g", "pf_per_g", "fg_pct", "fg3_per_g", "ft_pct", "trb_per_g", "ast_per_g", "stl_per_g", "blk_per_g", "tov_per_g", "pts_per_g"}

type BasketballRefService interface {
	ParseData(season int) []PlayerAverages
}

type basketballRefService struct{}

func NewBasketballRefService() BasketballRefService {
	return &basketballRefService{}
}

func (s *basketballRefService) ParseData(season int) []PlayerAverages {
	c := colly.NewCollector()

	var result []PlayerAverages

	c.OnHTML(".full_table", func(e *colly.HTMLElement) {
		playerName := e.ChildText("[data-stat = 'player']")
		position := e.ChildText("[data-stat = 'pos']")
		age := ParseFloatFromString(e.ChildText("[data-stat = 'age']"))
		team := e.ChildText("[data-stat = 'team_id']")
		gamesPlayed := ParseFloatFromString(e.ChildText("[data-stat = 'g']"))
		gamesStarted := ParseFloatFromString(e.ChildText("[data-stat = 'gs']"))

		minutesPerGame := ParseFloatFromString(e.ChildText("[data-stat = 'mp_per_g']"))
		fgPerGame := ParseFloatFromString(e.ChildText("[data-stat = 'fg_per_g']"))
		fgaPerGame := ParseFloatFromString(e.ChildText("[data-stat = 'fga_per_g']"))
		ftPerGame := ParseFloatFromString(e.ChildText("[data-stat = 'ft_per_g']"))
		ftaPerGame := ParseFloatFromString(e.ChildText("[data-stat = 'fta_per_g']"))
		personalFouls := ParseFloatFromString(e.ChildText("[data-stat = 'pf_per_g']"))

		fgPercentage := ParseFloatFromString(e.ChildText("[data-stat = 'fg_pct']"))
		threePm := ParseFloatFromString(e.ChildText("[data-stat = 'fg3_per_g']"))
		ftPercentage := ParseFloatFromString(e.ChildText("[data-stat = 'ft_pct']"))
		totalRebounds := ParseFloatFromString(e.ChildText("[data-stat = 'trb_per_g']"))
		assists := ParseFloatFromString(e.ChildText("[data-stat = 'ast_per_g']"))
		steals := ParseFloatFromString(e.ChildText("[data-stat = 'stl_per_g']"))
		blocks := ParseFloatFromString(e.ChildText("[data-stat = 'blk_per_g']"))
		turnovers := ParseFloatFromString(e.ChildText("[data-stat = 'tov_per_g']"))
		points := ParseFloatFromString(e.ChildText("[data-stat = 'pts_per_g']"))

		result = append(result, PlayerAverages{
			Name:           playerName,
			Position:       position,
			Age:            age,
			Team:           team,
			GamesPlayed:    gamesPlayed,
			GamesStarted:   gamesStarted,
			MinutesPerGame: minutesPerGame,
			FGMade:         fgPerGame,
			FGAttempted:    fgaPerGame,
			FTMade:         ftPerGame,
			FTAttempted:    ftaPerGame,
			PersonalFouls:  personalFouls,

			FGPercentage:  fgPercentage,
			ThreePTMade:   threePm,
			FTPercentage:  ftPercentage,
			TotalRebounds: totalRebounds,
			Assists:       assists,
			Steals:        steals,
			Blocks:        blocks,
			Turnovers:     turnovers,
			Points:        points,
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	err := c.Visit(fmt.Sprintf("https://www.basketball-reference.com/leagues/NBA_%d_per_game.html", season))
	if err != nil {
		return nil
	}

	// now compute z-score and sort by rank
	// numPlayer := len(result)
	// ComputeAverages(&data, numPlayer)
	// standardDeviations := ComputeStandardDeviation(numPlayer, result, &data)
	// ComputeZScores(numPlayer, result, averages, standardDeviations)
	// fmt.Printf("averages for season: %+v\n", averages)
	// fmt.Printf("sd for season: %+v\n", standardDeviations)

	// sort.Slice(result, func(i, j int) bool {
	// 	return result[j].ZScore < result[i].ZScore
	// })

	return result
}

func ParseFloatFromString(str string) float64 {
	res, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
		return 0.0
	}
	return res
}

// func ParseIntFromString(str string) int {
// 	res, err := strconv.Atoi(str)
// 	if err != nil {
// 		fmt.Println(err, str)
// 		return 0
// 	}
// 	return res
// }

// func ComputeZScores(
// 	len int,
// 	playerAverages []PlayerAverages,
// 	averages PlayerAverages,
// 	standardDeviations PlayerAverages,
// ) {
// 	for i := 0; i < len; i++ {
// 		astZ := ((playerAverages[i].Assists - averages.Assists) / standardDeviations.Assists)
// 		ptZ := ((playerAverages[i].Points - averages.Points) / standardDeviations.Points)
// 		rebZ := ((playerAverages[i].TotalRebounds - averages.TotalRebounds) / standardDeviations.TotalRebounds)
// 		stlZ := ((playerAverages[i].Steals - averages.Steals) / standardDeviations.Steals)
// 		blkZ := ((playerAverages[i].Blocks - averages.Blocks) / standardDeviations.Blocks)
// 		fgpZ := ((playerAverages[i].FGPercentage - averages.FGPercentage) / standardDeviations.FGPercentage) * (playerAverages[i].FGAttempted / averages.FGAttempted)
// 		ftpZ := ((playerAverages[i].FTPercentage - averages.FTPercentage) / standardDeviations.FTPercentage) * (playerAverages[i].FTAttempted / averages.FTAttempted)
// 		toZ := ((playerAverages[i].Turnovers - averages.Turnovers) / standardDeviations.Turnovers) * (-1)
// 		threeZ := ((playerAverages[i].ThreePTMade - averages.ThreePTMade) / standardDeviations.ThreePTMade)

// 		playerAverages[i].AstZ = astZ
// 		playerAverages[i].PtZ = ptZ
// 		playerAverages[i].RebZ = rebZ
// 		playerAverages[i].StlZ = stlZ
// 		playerAverages[i].BlkZ = blkZ
// 		playerAverages[i].FgpZ = fgpZ
// 		playerAverages[i].FtpZ = ftpZ
// 		playerAverages[i].ToZ = toZ
// 		playerAverages[i].ThreeZ = threeZ
// 		playerAverages[i].ZScore = (astZ + ptZ + rebZ + stlZ + blkZ + fgpZ + ftpZ + toZ + threeZ) / 9
// 	}
// }

// func ComputeStandardDeviation(
// 	len int,
// 	playerAverages []PlayerAverages,
// 	data *AggregateData) PlayerAverages {
// 	result := PlayerAverages{}

// 	var totalAst float64
// 	var totalPts float64
// 	var totalReb float64
// 	var totalStl float64
// 	var totalBlk float64
// 	var totalFgp float64
// 	var totalFtp float64
// 	var totalTo float64
// 	var total3pm float64
// 	for i := 0; i < len; i++ {
// 		totalAst += math.Pow(float64(playerAverages[i].Assists-data.Averages.Assists), 2)
// 		totalPts += math.Pow(float64(playerAverages[i].Points-data.Averages.Points), 2)
// 		totalReb += math.Pow(float64(playerAverages[i].TotalRebounds-data.Averages.TotalRebounds), 2)
// 		totalStl += math.Pow(float64(playerAverages[i].Steals-data.Averages.Steals), 2)
// 		totalBlk += math.Pow(float64(playerAverages[i].Blocks-data.Averages.Blocks), 2)

// 		totalTo += math.Pow(float64(playerAverages[i].Turnovers-data.Averages.Turnovers), 2)
// 		total3pm += math.Pow(float64(playerAverages[i].ThreePTMade-data.Averages.ThreePTMade), 2)

// 		totalFgp += math.Pow(float64(playerAverages[i].FGPercentage-data.Averages.FGPercentage), 2)
// 		totalFtp += math.Pow(float64(playerAverages[i].FTPercentage-data.Averages.FTPercentage), 2)
// 	}

// 	result.Assists = float32(math.Sqrt(totalAst / float64(len)))
// 	result.Points = float32(math.Sqrt(totalPts / float64(len)))
// 	result.TotalRebounds = float32(math.Sqrt(totalReb / float64(len)))
// 	result.Steals = float32(math.Sqrt(totalStl / float64(len)))
// 	result.Blocks = float32(math.Sqrt(totalBlk / float64(len)))
// 	result.FGPercentage = float32(math.Sqrt(totalFgp / float64(len)))
// 	result.FTPercentage = float32(math.Sqrt(totalFtp / float64(len)))
// 	result.Turnovers = float32(math.Sqrt(totalTo / float64(len)))
// 	result.ThreePTMade = float32(math.Sqrt(total3pm / float64(len)))

// 	return result
// }

// func ComputeAverages(result *AggregateData, numPlayers int) {
// 	result.Averages.Assists = result.AssistsTotal / float32(numPlayers)
// 	result.Averages.Points = result.PointsTotal / float32(numPlayers)
// 	result.Averages.TotalRebounds = result.TotalReboundsTotal / float32(numPlayers)
// 	result.Averages.Steals = result.StealsTotal / float32(numPlayers)
// 	result.Averages.Blocks = result.BlocksTotal / float32(numPlayers)
// 	result.Averages.Turnovers = result.TurnoversTotal / float32(numPlayers)
// 	result.Averages.ThreePTMade = result.ThreePTMadeTotal / float32(numPlayers)

// 	// calculate FG and FT percentage using sum of made and attempts
// 	result.Averages.FGPercentage = result.FGMadeTotal / result.FGAttemptedTotal
// 	result.Averages.FTPercentage = result.FTMadeTotal / result.FTAttemptedTotal

// 	result.Averages.FTAttempted = result.FTAttemptedTotal / float32(numPlayers)
// 	result.Averages.FGAttempted = result.FGAttemptedTotal / float32(numPlayers)
// }
