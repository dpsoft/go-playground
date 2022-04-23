package functions

import (
	"fmt"
	"time"
)

func MapSlice[T any, M any](slice []T, f func(T) M) []M {
	result := make([]M, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

func Filter[T any](slice []T, f func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func Timed(f func()) {
	start := time.Now()
	defer func() { fmt.Printf("execution took %s\n", time.Since(start)) }()
	f()
}

func TimedFunction[T any](f func() T) T {
	start := time.Now()
	defer func() { fmt.Printf("execution took %s\n", time.Since(start)) }()
	return f()
}
