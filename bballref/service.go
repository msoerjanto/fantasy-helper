package bballref

import (
	"fmt"
	"strconv"

	"github.com/gocolly/colly"
)

type PlayerAverages struct {
	Name           string
	Position       string
	Age            int
	Team           string
	GamesPlayed    int
	GamesStarted   int
	MinutesPerGame float32
	PersonalFouls  float32

	FGMade      float32
	FGAttempted float32
	FTMade      float32
	FTAttempted float32

	FGPercentage  float32
	FTPercentage  float32
	ThreePTMade   float32
	TotalRebounds float32
	Assists       float32
	Steals        float32
	Blocks        float32
	Turnovers     float32
	Points        float32
}

type BasketballRefService interface {
	GetPlayerAveragesBySeason(season int) []PlayerAverages
}

type basketballRefService struct{}

func NewBasketballRefService() BasketballRefService {
	return &basketballRefService{}
}

func ParseFloatFromString(str string) float32 {
	res, err := strconv.ParseFloat(str, 32)
	if err != nil {
		fmt.Println(err)
		return 0.0
	}
	return float32(res)
}

func (s *basketballRefService) GetPlayerAveragesBySeason(season int) []PlayerAverages {
	c := colly.NewCollector()

	var result []PlayerAverages

	c.OnHTML(".full_table", func(e *colly.HTMLElement) {
		playerName := e.ChildText("[data-stat = 'player']")
		position := e.ChildText("[data-stat = 'pos']")
		age, err := strconv.Atoi(e.ChildText("[data-stat = 'age']"))
		if err != nil {
			fmt.Println(err)
		}
		team := e.ChildText("[data-stat = 'team_id']")
		gamesPlayed, err := strconv.Atoi(e.ChildText("[data-stat = 'g']"))
		if err != nil {
			fmt.Println(err)
		}
		gamesStarted, err := strconv.Atoi(e.ChildText("[data-stat = 'gs']"))
		if err != nil {
			fmt.Println(err)
		}
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

		// Print link
		fmt.Printf("Player: %s \n", playerName)
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
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		// c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	err := c.Visit(fmt.Sprintf("https://www.basketball-reference.com/leagues/NBA_%d_per_game.html", season))
	if err != nil {
		return nil
	}

	return result
}
