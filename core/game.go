package core

import (
	"github.com/justjoeyuk/gobolt/display"
	"github.com/veandco/go-sdl2/sdl"
)

// Game - The root object
type Game struct {
	*display.Sprite
}

// GameRunner -
type GameRunner interface {
	GetSprite() *display.Sprite
	Start()
	Update()
	HandleEvent(e sdl.Event)
}

// NewGame -
func NewGame() *Game {
	return &Game{
		Sprite: display.NewSprite(),
	}
}

// Start - Start the game
func (g *Game) Start() {}

// Update - Runs each frame
func (g *Game) Update() {}

// HandleEvent - Handle any events here
func (g *Game) HandleEvent(e sdl.Event) {}

// GetSprite -
func (g *Game) GetSprite() *display.Sprite {
	return g.Sprite
}
