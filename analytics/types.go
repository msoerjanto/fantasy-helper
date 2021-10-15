package analytics

import "github.com/msoerjanto/fantasy-helper/bballref"

type PuntCategories struct {
	FGPercentage  bool
	FTPercentage  bool
	ThreePTMade   bool
	TotalRebounds bool
	Assists       bool
	Steals        bool
	Blocks        bool
	Turnovers     bool
	Points        bool
}

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

	FantasyPts float64
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

	Totals             PlayerAverages
	Averages           PlayerAverages
	StandardDeviations bballref.PlayerAverages
}
