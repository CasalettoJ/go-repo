package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenHeight = 480
	screenWidth  = 640
)

var (
	window *sdl.Window
	screen *sdl.Surface
	image  *sdl.Surface
)

func setupWindow() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	window, err = sdl.CreateWindow("Title", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	screen, err = window.GetSurface()
	return err
}

func loadMedia() error {
	var err error
	image, err = sdl.LoadBMP("assets/02water.bmp")
	return err
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
	defer window.Destroy()
	defer sdl.Quit()

	screen.FillRect(nil, 0)
	window.UpdateSurface()
	err := image.Blit(nil, screen, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to blit image: %s\n", err)
		os.Exit(1)
	}
	defer image.Free()
	window.UpdateSurface()
	sdl.Delay(5000)
}
