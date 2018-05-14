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
	screenWidth  = 640
	screenHeight = 480
)

var (
	window   *sdl.Window
	renderer *sdl.Renderer
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
	window, err = sdl.CreateWindow("Geometry Rendering", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return err
	}
	if err := renderer.SetDrawColor(255, 255, 255, 255); err != nil {
		return err
	}
	imgFlags := img.INIT_PNG
	if (img.Init(imgFlags) & imgFlags) == 0 {
		fmt.Println("Failed to initialize SDL_IMG PNG")
		err = img.GetError()
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
		// Clear Screen
		if err := renderer.SetDrawColor(255, 255, 255, 255); err != nil {
			return err
		}
		if err := renderer.Clear(); err != nil {
			return err
		}

		// Render red-filled quad
		fillRect := sdl.Rect{X: screenWidth / 4, Y: screenHeight / 4, W: screenWidth / 2, H: screenHeight / 2}
		if err := renderer.SetDrawColor(255, 0, 0, 255); err != nil {
			return err
		}
		if err := renderer.FillRect(&fillRect); err != nil {
			return err
		}

		// Render green-outlined quad
		outlineRect := sdl.Rect{X: screenWidth / 6, Y: screenHeight / 6, W: screenWidth * 2 / 3, H: screenHeight * 2 / 3}
		if err := renderer.SetDrawColor(0, 255, 0, 255); err != nil {
			return err
		}
		if err := renderer.DrawRect(&outlineRect); err != nil {
			return err
		}

		// Draw blue horizontal line
		if err := renderer.SetDrawColor(0, 0, 255, 255); err != nil {
			return err
		}
		if err := renderer.DrawLine(0, screenHeight/2, screenWidth, screenHeight/2); err != nil {
			return err
		}

		// Draw vertical line of yellow dots
		if err := renderer.SetDrawColor(255, 255, 0, 255); err != nil {
			return err
		}
		for i := 0; i < screenHeight; i = i + 4 {
			if err := renderer.DrawPoint(screenWidth/2, int32(i)); err != nil {
				return err
			}
		}

		// Update screen
		renderer.Present()
	}
}

func close() {
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
}

func main() {
	if err := run(); err != nil {
		close()
		log.Panic(err)
	}
	close()
}
