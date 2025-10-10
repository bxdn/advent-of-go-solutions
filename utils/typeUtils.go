package utils

func StringsToRunes(strings []string) [][]rune {
	toRet := [][]rune{}
	for _, s := range strings {
		toRet = append(toRet, []rune(s))
	}
	return toRet
}

func Flatten[T any](slices [][]T) []T {
	var toRet []T
	for _, s := range slices {
		toRet = append(toRet, s...)
	}
	return toRet
}