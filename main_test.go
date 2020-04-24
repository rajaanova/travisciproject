package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFib(t *testing.T){
	//TODO add the value
	assert.Equal(t,5,Fib(5))
	assert.Equal(t,8,Fib(5))
	assert.Equal(t,8,Fib(6))


}
