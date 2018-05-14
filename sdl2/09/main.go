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
	window   *sdl.Window
	renderer *sdl.Renderer
	texture  *sdl.Texture
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
	window, err = sdl.CreateWindow("The Viewport", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN)
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

func loadTexture(path string) (*sdl.Texture, error) {
	var texture *sdl.Texture
	loadedSurface, err := img.Load(path)
	if err != nil {
		return nil, img.GetError()
	}
	texture, err = renderer.CreateTextureFromSurface(loadedSurface)
	if err != nil {
		return nil, img.GetError()
	}
	loadedSurface.Free()
	return texture, nil
}

func loadMedia() error {
	var err error
	texture, err = loadTexture("assets/09city.png")
	return err
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

		// Top-left viewport
		topLeftViewport := sdl.Rect{X: 0, Y: 0, W: screenWidth / 2, H: screenHeight / 2}
		if err := renderer.SetViewport(&topLeftViewport); err != nil {
			return err
		}
		// Draw top-left viewport to screen
		if err := renderer.Copy(texture, nil, nil); err != nil {
			return err
		}

		// Top-right viewport
		topRightViewport := sdl.Rect{X: screenWidth / 2, Y: 0, W: screenWidth / 2, H: screenHeight / 2}
		if err := renderer.SetViewport(&topRightViewport); err != nil {
			return err
		}
		// Draw top-right viewport to screen
		if err := renderer.Copy(texture, nil, nil); err != nil {
			return err
		}

		// Bottom viewport
		bottomViewport := sdl.Rect{X: 0, Y: screenHeight / 2, W: screenWidth, H: screenHeight / 2}
		if err := renderer.SetViewport(&bottomViewport); err != nil {
			return err
		}
		// Draw bottom viewport to screen
		if err := renderer.Copy(texture, nil, nil); err != nil {
			return err
		}

		// Update the screen
		renderer.Present()
	}
}

func close() {
	texture.Destroy()
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
