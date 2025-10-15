package utils

func Pop[T any](slice []T) ([]T, T) {
	return slice[:len(slice)-1], Last(slice)
}

func Dequeue[T any](slice []T) ([]T, T) {
	return slice[1:], slice[0]
}

func IsEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

func Last[T any](slice []T) T {
	return slice[len(slice)-1]
}

func At[T any](slice []T, idx int) T {
	if idx < 0 {
		idx = len(slice) + idx
	}
	return slice[idx]
}

func Set[T any](slice []T, idx int, val T) {
	if idx < 0 {
		idx = len(slice) + idx
	}
	slice[idx] = val
}

func Rev[T any](slice []T) []T {
	length := len(slice)
	toRet := make([]T, length)
	for i, item := range slice {
		toRet[length-(i+1)] = item
	}
	return toRet
}