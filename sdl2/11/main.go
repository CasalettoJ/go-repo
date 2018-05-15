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
	window             *sdl.Window
	renderer           *sdl.Renderer
	spriteSheetTexture = NewLTexture()
	spriteClips        = [4]sdl.Rect{}
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
	window, err = sdl.CreateWindow("Clip Rendering and Sprite Sheets", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN)
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
	if err := spriteSheetTexture.LoadFromFile("assets/11spritesheet.png"); err != nil {
		return err
	}
	// Set top left sprite
	spriteClips[0] = sdl.Rect{X: 0, Y: 0, W: 100, H: 100}
	// Set top right sprite
	spriteClips[1] = sdl.Rect{X: 100, Y: 0, W: 100, H: 100}
	// Set bottom left sprite
	spriteClips[2] = sdl.Rect{X: 0, Y: 100, W: 100, H: 100}
	// Set bottom right sprite
	spriteClips[3] = sdl.Rect{X: 100, Y: 100, W: 100, H: 100}

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

		// Render top left sprite
		spriteSheetTexture.Render(0, 0, &spriteClips[0])
		// Render top right sprite
		spriteSheetTexture.Render(screenWidth-spriteClips[1].W, 0, &spriteClips[1])
		// Render bottom left sprite
		spriteSheetTexture.Render(0, screenHeight-spriteClips[2].H, &spriteClips[2])
		// Render bottom right sprite
		spriteSheetTexture.Render(screenWidth-spriteClips[3].W, screenHeight-spriteClips[3].H, &spriteClips[3])

		// Update the screen
		renderer.Present()
	}
}

func close() {
	spriteSheetTexture.Free()
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
	err := run()
	close()
	if err != nil {
		log.Panic(err)
	}
}
