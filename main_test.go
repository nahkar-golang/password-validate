package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassword_Validate(t *testing.T) {
	password := Password("12")
	err := password.Validate()
	assert.Error(t, err)
}
