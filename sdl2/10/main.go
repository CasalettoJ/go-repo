package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenHeight = 480
	screenWidth  = 640
)

var (
	window     *sdl.Window
	renderer   *sdl.Renderer
	sprite     = NewLTexture()
	background = NewLTexture()
)

func setupWindow() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}
	//Set texture filtering to linear
	if !(sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")) {
		fmt.Println("Warning! Linear texture filtering not enabled!")
	}
	var err error
	window, err = sdl.CreateWindow("Color Keying", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return err
	}
	imgFlags := img.INIT_PNG
	if (img.Init(imgFlags) & imgFlags) == 0 {
		fmt.Println("Failed to initialize SDL_IMG PNG")
		err = img.GetError()
	}
	return nil
}

func loadMedia() error {
	if err := sprite.LoadFromFile("assets/10bowblue.png"); err != nil {
		return err
	}
	if err := background.LoadFromFile("assets/10bg.png"); err != nil {
		return err
	}
	return nil
}

func run() error {
	var e sdl.Event
	for {
		for e = sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			fmt.Printf("event %+v\n", e)
			switch e.(type) {
			case *sdl.QuitEvent:
				return nil
			}
		}

		// Clear the screen
		if err := renderer.SetDrawColor(255, 255, 255, 255); err != nil {
			return err
		}
		if err := renderer.Clear(); err != nil {
			return err
		}

		// Render background texture to screen
		if err := background.Render((screenWidth/2)-(background.Width()/2), (screenHeight/2)-(background.Height()/2)); err != nil {
			return err
		}

		// Render sprite to screen
		if err := sprite.Render((screenWidth/2)-(sprite.Width()/2), (screenHeight/2)-(sprite.Height()/2)); err != nil {
			return err
		}

		// Update the screen
		renderer.Present()
	}
}

func close() {
	sprite.Free()
	background.Free()
	renderer.Destroy()
	window.Destroy()
	img.Quit()
	sdl.Quit()
}

func init() {
	runtime.LockOSThread()
	err := setupWindow()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to setup window: %s\n", err)
		os.Exit(1)
	}
	err = loadMedia()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load media: %s\n", err)
		os.Exit(1)
	}
}

func main() {
	if err := run(); err != nil {
		close()
		log.Panic(err)
	}
	close()
}
