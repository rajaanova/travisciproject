package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFib(t *testing.T){
	assert.Equal(t,5,Fib(5))

}
