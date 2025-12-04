package utils

import "errors"

type Option[T any] struct {
	value    T
	hasValue bool
}

func (o Option[T]) Or(alternative T) T {
	if o.hasValue {
		return o.value
	}
	return alternative
}

func (o Option[T]) OrDefault() T {
	var t T
	if o.hasValue {
		return o.value
	}
	return t
}

func (o Option[T]) OrErr(message string) (T, error) {
	if o.hasValue {
		return o.value, nil
	}
	return o.value, errors.New(message)
}

func (o Option[T]) OrPanic(message string) T {
	if o.hasValue {
		return o.value
	}
	panic(message)
}

func (o Option[T]) Get() (T, bool) {
	return o.value, o.hasValue
}

func Some[T any](value T) Option[T] {
	return Option[T]{value, true}
}

func None[T any]() Option[T] {
	var t T
	return Option[T]{t, false}
}

func MapO[T, U any](o Option[T], f func(T) U) Option[U] {
	if o.hasValue {
		return Some(f(o.value))
	}
	return None[U]()
}

func FlatMapO[T, U any](o Option[T], f func(T) Option[U]) Option[U] {
	if o.hasValue {
		return f(o.value)
	}
	return None[U]()
}

type CmpOption[T comparable] struct {
	Option[T]
}

func (co CmpOption[T]) Is(check T) bool {
	return co.hasValue && co.value == check
}

func CSome[T comparable](t T) CmpOption[T] {
	return CmpOption[T]{Some(t)}
}

func CNone[T comparable]() CmpOption[T] {
	return CmpOption[T]{None[T]()}
}
