package models

import (
	"errors"
	"fmt"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPerson(t *testing.T) {
	assert.Error(t, errors.New("panic"))
}

func TestPerson2(t *testing.T) {
	// panic("error")
	log.Println("ok")
}

func TestPerson3(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	fmt.Println(mock)
}
