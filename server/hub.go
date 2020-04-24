package server

import (
	"github.com/vetusbs/gomino/models"
)

type Hub struct {
	Games map[string]*models.Game
}
