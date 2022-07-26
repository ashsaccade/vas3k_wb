package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Simple(t *testing.T) {
	in := []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}
	k := 2
	assert.Equal(t, 16, getNiceSubArraysCount(in, k))
}

func Test_OnlyOdds(t *testing.T) {
	in := []int{1, 1, 1, 7, 3}
	k := 5
	assert.Equal(t, 1, getNiceSubArraysCount(in, k))
}

func Test_OnlyOdds_1(t *testing.T) {
	in := []int{1, 3, 7, 11, 13}
	k := 2
	assert.Equal(t, 4, getNiceSubArraysCount(in, k))
}

func Test_Hybrid(t *testing.T) {
	in := []int{1, 1, 2, 1, 1}
	k := 3
	assert.Equal(t, 2, getNiceSubArraysCount(in, k))
}

func Test_Only_Even_Numbers(t *testing.T) {
	in := []int{2, 4, 6}
	k := 1
	assert.Equal(t, 0, getNiceSubArraysCount(in, k))
}
