package display

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// LoadTextureFromFile with a given paths
func LoadTextureFromFile(path string) (*sdl.Texture, error) {
	var texture *sdl.Texture
	var err error

	sdl.Do(func() {
		texture, err = img.LoadTexture(renderContext, path)
	})

	return texture, err
}
