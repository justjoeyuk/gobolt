package display

import "github.com/veandco/go-sdl2/sdl"

// Sprite -
type Sprite struct {
	*Canvas

	Children []*Sprite
	Parent   *Sprite
}

// Spritable -
type Spritable interface {
	AddChild(child Spritable)
}

// NewSprite -
func NewSprite() *Sprite {
	return &Sprite{
		Canvas: NewCanvas(),
	}
}

// AddChild - Adds a child to a Sprite
func (s *Sprite) AddChild(child Spritable) {
	s.mu.Lock()
	defer s.mu.Unlock()

	childSprite := child.(*Sprite)

	childSprite.mu.Lock()
	defer childSprite.mu.Unlock()

	s.Children = append(s.Children, childSprite)
	childSprite.Parent = s
}

// GBLTDrawChildren of this Sprite
func (s *Sprite) GBLTDrawChildren() {
	s.mu.RLock()
	defer s.mu.RUnlock()

	s.draw()

	for _, childSprite := range s.Children {
		childSprite.GBLTDrawChildren()
	}
}

func (s *Sprite) draw() error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.Texture == nil {
		return nil
	}

	var err error
	xPos := s.X
	yPos := s.Y

	if s.Parent != nil {
		xPos = s.X + s.Parent.X
		yPos = s.Y + s.Parent.Y
	}

	destRect := &sdl.Rect{X: xPos, Y: yPos, W: s.Width, H: s.Height}
	err = renderContext.Copy(s.Texture, nil, destRect)

	return err
}
