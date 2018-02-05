package display

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var renderContext *sdl.Renderer

// SetRenderContext for the engine to render and draw
func SetRenderContext(renderer *sdl.Renderer) {
	renderContext = renderer
}

// GBLTRenderGame - Renders the core game
func GBLTRenderGame(sprite *Sprite, renderer *sdl.Renderer) error {
	if renderer != renderContext {
		return fmt.Errorf("you should not call _RenderGame from outside the Engine")
	}

	sdl.Do(func() {
		renderer.Clear()

		sprite.GBLTDrawChildren()

		renderer.Present()
	})

	return nil
}
