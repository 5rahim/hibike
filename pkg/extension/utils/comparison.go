package utils

import "github.com/adrg/strutil/metrics"

type SorensenDiceResult struct {
	OriginalValue string
	Value         string
	Rating        float64
}

func CompareWithSorensenDice(v string, vals []string) []*SorensenDiceResult {

	dice := metrics.NewSorensenDice()
	dice.CaseSensitive = false

	res := make([]*SorensenDiceResult, len(vals))

	for _, val := range vals {
		res = append(res, &SorensenDiceResult{
			OriginalValue: v,
			Value:         val,
			Rating:        dice.Compare(v, val),
		})
	}

	return res
}

func FindBestMatchWithSorensenDice(v string, vals []string) (*SorensenDiceResult, bool) {
	res := CompareWithSorensenDice(v, vals)

	if len(res) == 0 {
		return nil, false
	}

	var bestResult *SorensenDiceResult
	for _, result := range res {
		if bestResult == nil || result.Rating > bestResult.Rating {
			bestResult = result
		}
	}

	return bestResult, true
}
