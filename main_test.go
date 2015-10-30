package main_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/thejsj/go-json/"
	"testing"
)

func Test1(t *testing.T) {
	val, _ := main.ToJSON(5)
	assert.Equal(t, val, "5")
}
