package bballref

import (
	"fmt"

	"github.com/gocolly/colly"
)

type PlayerAverages struct {
	DataRow int

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

	c.OnHTML(".full_table, .partial_table", func(e *colly.HTMLElement) {
		currPlayer := parseRow(e)
		result = append(result, currPlayer)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	err := c.Visit(fmt.Sprintf("https://www.basketball-reference.com/leagues/NBA_%d_per_game.html", season))
	if err != nil {
		return nil
	}

	return removeDuplicates(result)
}

func parseRow(e *colly.HTMLElement) PlayerAverages {
	dataRow := e.Index
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

	return PlayerAverages{
		DataRow: dataRow,

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
	}
}
