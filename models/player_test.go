package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vetusbs/gomino/models"
)

const ANY_NAME = "vetussao"

func TestPlayer_GetName(t *testing.T) {
	player := models.NewPlayer(ANY_NAME, nil, nil)

	name := player.GetName()

	assert.Equal(t, name, ANY_NAME)
}
