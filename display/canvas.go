package display

import (
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

// Canvas - A Canvas Entity with a Texture
type Canvas struct {
	mu sync.RWMutex

	Texture             *sdl.Texture
	X, Y, Width, Height int32
}

// NewCanvas -
func NewCanvas() *Canvas {
	return &Canvas{
		X:      0,
		Y:      0,
		Width:  0,
		Height: 0,
	}
}
