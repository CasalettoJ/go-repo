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
	err := sdl.Init(sdl.INIT_EVERYTHING)
	//Set texture filtering to linear
	if !(sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")) {
		fmt.Println("Warning! Linear texture filtering not enabled!")
	}
	window, err = sdl.CreateWindow("Texture Loading and Rendering", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	renderer.SetDrawColor(255, 255, 255, 255)
	imgFlags := img.INIT_PNG
	if (img.Init(imgFlags) & imgFlags) == 0 {
		fmt.Println("Failed to initialize SDL_IMG PNG")
		err = img.GetError()
	}
	return err
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
	texture, err = loadTexture("assets/07cake.png")
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
		if err := renderer.Clear(); err != nil {
			return err
		}
		if err := renderer.Copy(texture, nil, nil); err != nil {
			return err
		}
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
