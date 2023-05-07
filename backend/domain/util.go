package domain

import (
	"math/rand"
	"strconv"
	"time"
)

func randomID(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return strconv.FormatInt(time.Now().UnixNano(), 10) + string(b)
}

func remove[T any](s []T, index int) []T {
	if index < 0 || index >= len(s) {
		return s
	}
	return append(s[:index], s[index+1:]...)
}

func insert[T any](s []T, index int, value T) []T {
	if index < 0 || index > len(s) {
		return s
	}
	s = append(s, value)
	copy(s[index+1:], s[index:])
	s[index] = value
	return s
}

func findIndex[T any](s []T, find func(T) bool) int {
	for i, v := range s {
		if find(v) {
			return i
		}
	}
	return -1
}
