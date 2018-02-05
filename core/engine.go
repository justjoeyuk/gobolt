package core

import (
	"os"
	"runtime"
	"time"

	"github.com/justjoeyuk/gobolt/display"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Engine - The core of the engine
type Engine struct {
	FPS uint64

	WindowTitle string
	ResX        int32
	ResY        int32

	window   *sdl.Window
	renderer *sdl.Renderer
	game     GameRunner
}

// Initialize the engine
func (e *Engine) Initialize() error {
	window, renderer, err := e.initializeSDL()
	display.SetRenderContext(renderer)

	e.window = window
	e.renderer = renderer

	return err
}

// Run - start the game running
func (e *Engine) Run(g GameRunner) error {
	defer sdl.Quit()
	defer ttf.Quit()
	defer e.window.Destroy()

	e.game = g

	events := make(chan sdl.Event)
	errc := e.runLoop(events)

	e.game.Start()

	runtime.LockOSThread()

	for {
		select {
		case events <- sdl.WaitEvent():
		case err := <-errc:
			return err
		}
	}
}

func (e *Engine) runLoop(events <-chan sdl.Event) chan error {
	errc := make(chan error)

	go func() {
		defer close(errc)

		tick := time.Tick(time.Second / 60.0)

		for {
			select {
			case event := <-events:
				if exit := e.handleEvent(event); exit {
					os.Exit(0)
				}
			case <-tick:
				e.game.Update()
				if err := display.GBLTRenderGame(e.game.GetSprite(), e.renderer); err != nil {
					errc <- err
				}
			}
		}
	}()

	return errc
}

func (e *Engine) handleEvent(event sdl.Event) bool {
	switch event.(type) {
	case *sdl.QuitEvent:
		return true
	}

	e.game.HandleEvent(event)
	return false
}

func (e *Engine) initializeSDL() (*sdl.Window, *sdl.Renderer, error) {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var err error

	sdl.Do(func() {
		if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
			return
		}

		if err := ttf.Init(); err != nil {
			return
		}

		if window, renderer, err = sdl.CreateWindowAndRenderer(e.ResX, e.ResY, sdl.WINDOW_SHOWN); err != nil {
			return
		}

		window.SetTitle(e.WindowTitle)
	})

	return window, renderer, err
}
