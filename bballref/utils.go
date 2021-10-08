package bballref

import (
	"fmt"
	"strconv"
)

func removeDuplicates(players []PlayerAverages) []PlayerAverages {
	keys := make(map[string]PlayerAverages)

	for i := 0; i < len(players); i++ {
		// put to map so for duplicates we keep only the latest one
		if _, ok := keys[players[i].Name]; !ok {
			// first time seeing player
			keys[players[i].Name] = players[i]
		} else {
			// not first time, update value based on order
			if players[i].DataRow > keys[players[i].Name].DataRow {
				keys[players[i].Name] = players[i]
			}
		}
	}

	// Convert map to slice of values.
	result := []PlayerAverages{}
	for _, value := range keys {
		result = append(result, value)
	}

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
