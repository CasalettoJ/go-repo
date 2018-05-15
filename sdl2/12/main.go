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
	window           *sdl.Window
	renderer         *sdl.Renderer
	modulatedTexture = NewLTexture()
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
	window, err = sdl.CreateWindow("Color Modulation", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN)
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
	return modulatedTexture.LoadFromFile("assets/12texture.png")
}

func run() error {
	var e sdl.Event
	// Modulation components
	var r, g, b uint8 = 255, 255, 255
	for {
		for e = sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			fmt.Printf("event %+v\n", e)
			switch t := e.(type) {
			case *sdl.QuitEvent:
				return nil
			case *sdl.KeyboardEvent:
				if t.Type == sdl.KEYDOWN {
					switch t.Keysym.Sym {
					case sdl.K_q:
						r = r + 32
					case sdl.K_w:
						g = g + 32
					case sdl.K_e:
						b = b + 32
					case sdl.K_a:
						r = r - 32
					case sdl.K_s:
						g = g - 32
					case sdl.K_d:
						b = b - 32
					}
				}
			}
		}

		// Clear the screen
		if err := renderer.SetDrawColor(255, 255, 255, 255); err != nil {
			return err
		}
		if err := renderer.Clear(); err != nil {
			return err
		}

		// Modulate aand render the texture
		modulatedTexture.SetColor(r, g, b)
		modulatedTexture.Render(0, 0, nil)

		// Update the screen
		renderer.Present()
	}
}

func close() {
	modulatedTexture.Free()
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
