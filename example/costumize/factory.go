package costumize

var CastMap = &map[rune]float64{'A': 1}

func UpdateCastMap(key rune, val float64) {
	temp := *CastMap
	temp[key] = val

	*CastMap = temp
}
