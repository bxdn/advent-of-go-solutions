package utils

func Pop[T any](slice []T) ([]T, Option[T]) {
	if IsEmpty(slice) {
		return slice, None[T]()
	}
	return slice[:len(slice)-1], Last(slice)
}

func Dequeue[T any](slice []T) ([]T, Option[T]) {
	if IsEmpty(slice) {
		return slice, None[T]()
	}
	return slice[1:], Some(slice[0])
}

func IsEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

func Last[T any](slice []T) Option[T] {
	if IsEmpty(slice) {
		return None[T]()
	}
	return Some(slice[len(slice)-1])
}

func At[T any](slice []T, idx int) Option[T] {
	if idx < 0 {
		idx = len(slice) + idx
	}
	if idx < 0 || idx >= len(slice) {
		return None[T]()
	}
	return Some(slice[idx])
}

func Set[T any](slice []T, idx int, val T) bool {
	if idx < 0 {
		idx = len(slice) + idx
	}
	if idx < 0 || idx >= len(slice) {
		return false
	}
	slice[idx] = val
	return true
}

func Rev[T any](slice []T) []T {
	length := len(slice)
	toRet := make([]T, length)
	for i, item := range slice {
		toRet[length-(i+1)] = item
	}
	return toRet
}

func ContainsSubslice[T comparable](s, sub []T) bool {
	if len(sub) == 0 {
		return false
	}
	if len(sub) > len(s) {
		return false
	}

	pi := make([]int, len(sub))
	for i, j := 1, 0; i < len(sub); i++ {
		for j > 0 && sub[i] != sub[j] {
			j = pi[j-1]
		}
		if sub[i] == sub[j] {
			j++
		}
		pi[i] = j
	}

	for i, j := 0, 0; i < len(s); i++ {
		for j > 0 && s[i] != sub[j] {
			j = pi[j-1]
		}
		if s[i] == sub[j] {
			j++
		}
		if j == len(sub) {
			return true
		}
	}
	return false
}
