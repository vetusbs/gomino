package models_test

import (
	. "gomino/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

const ANY_NAME = "vetussao"

func TestPlayer_GetName(t *testing.T) {
	player := NewPlayer(ANY_NAME, nil, nil)

	name := player.GetName()

	assert.Equal(t, name, ANY_NAME)
}
