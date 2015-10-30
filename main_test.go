package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	val, _ := ToJSON(5)
	assert.Equal(t, val, "5")
}
