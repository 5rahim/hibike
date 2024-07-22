package similarity

import "github.com/adrg/strutil/metrics"

type LevenshteinResult struct {
	OriginalValue string
	Value         string
	Distance      int
}

func CompareWithLevenstein(v string, vals []string) []*LevenshteinResult {

	lev := metrics.NewLevenshtein()
	lev.CaseSensitive = false

	res := make([]*LevenshteinResult, len(vals))

	for _, val := range vals {
		res = append(res, &LevenshteinResult{
			OriginalValue: v,
			Value:         val,
			Distance:      lev.Distance(v, val),
		})
	}

	return res
}

func FindBestMatchWithLevenstein(v string, vals []string) (*LevenshteinResult, bool) {
	res := CompareWithLevenstein(v, vals)

	if len(res) == 0 {
		return nil, false
	}

	var bestResult *LevenshteinResult
	for _, result := range res {
		if bestResult == nil || result.Distance < bestResult.Distance {
			bestResult = result
		}
	}

	return bestResult, true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

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

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func EliminateLeastSimilarValue(arr []string) []string {
	if len(arr) < 3 {
		return arr
	}

	sd := metrics.NewSorensenDice()
	sd.CaseSensitive = false

	leastSimilarIndex := -1
	leastSimilarScore := 2.0

	for i := 0; i < len(arr); i++ {
		totalSimilarity := 0.0

		for j := 0; j < len(arr); j++ {
			if i != j {
				score := sd.Compare(arr[i], arr[j])
				totalSimilarity += score
			}
		}

		if totalSimilarity < leastSimilarScore {
			leastSimilarScore = totalSimilarity
			leastSimilarIndex = i
		}
	}

	if leastSimilarIndex != -1 {
		arr = append(arr[:leastSimilarIndex], arr[leastSimilarIndex+1:]...)
	}

	return arr

}
