package models

import (
	"errors"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerson(t *testing.T) {
	assert.Error(t, errors.New("panic"))
}

func TestPerson2(t *testing.T) {
	// panic("error")
	log.Println("ok")
}
