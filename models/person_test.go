package models

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerson(t *testing.T) {
	assert.Error(t, errors.New("panic"))
}

func TestPerson2(t *testing.T) {
	panic("error")
}
