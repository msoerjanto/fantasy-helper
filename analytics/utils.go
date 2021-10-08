package analytics

// TODO determine if punting should be placed here
// for now dont think its necessary
func computeFantasyPtsForPlayer(player PlayerAverages) float64 {
	return (player.Points * 0.5) + (player.Assists * 2) + (player.TotalRebounds * 1.5) +
		(player.Blocks * 3) + (player.Steals * 3) + (player.ThreePTMade * 3) +
		(player.Turnovers * -2) + (player.FGAttempted * -0.45) + (player.FGMade * 1) +
		(player.FTAttempted * -0.75) + (player.FTMade * 1)
}

func getNumCategories(punt PuntCategories) int {
	result := 0
	if !punt.Points {
		result++
	}
	if !punt.Assists {
		result++
	}
	if !punt.TotalRebounds {
		result++
	}
	if !punt.Steals {
		result++
	}
	if !punt.ThreePTMade {
		result++
	}
	if !punt.Turnovers {
		result++
	}
	if !punt.Blocks {
		result++
	}
	if !punt.FGPercentage {
		result++
	}
	if !punt.FTPercentage {
		result++
	}
	return result
}

func getZscoreSum(player PlayerAverages, punt PuntCategories) float64 {
	var zscoreSum float64

	if !punt.Points {
		zscoreSum += player.PtZ
	}
	if !punt.Assists {
		zscoreSum += player.AstZ
	}
	if !punt.TotalRebounds {
		zscoreSum += player.RebZ
	}
	if !punt.Steals {
		zscoreSum += player.StlZ
	}
	if !punt.ThreePTMade {
		zscoreSum += player.ThreeZ
	}
	if !punt.Turnovers {
		zscoreSum += player.ToZ
	}
	if !punt.Blocks {
		zscoreSum += player.BlkZ
	}
	if !punt.FGPercentage {
		zscoreSum += player.FgpZ
	}
	if !punt.FTPercentage {
		zscoreSum += player.FtpZ
	}
	return zscoreSum
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
