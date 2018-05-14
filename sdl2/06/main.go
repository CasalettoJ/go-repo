package main

import (
	"fmt"
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
	screen     *sdl.Surface
	pngSurface *sdl.Surface
)

func setupWindow() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	window, err = sdl.CreateWindow("Loading PNGs with SDL_image", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	screen, err = window.GetSurface()
	imgFlags := img.INIT_PNG
	if (img.Init(imgFlags) & imgFlags) == 0 {
		fmt.Println("Failed to initialize SDL_IMG PNG")
		err = img.GetError()
	}
	return err
}

func loadSurface(path string) (*sdl.Surface, error) {
	loadedSurface, err := img.Load(path)
	if err != nil {
		return nil, err
	}
	optimizedSurface, err := loadedSurface.Convert(screen.Format, 0)
	if err != nil {
		return nil, err
	}
	loadedSurface.Free()
	return optimizedSurface, nil

}

func loadMedia() error {
	var err error
	pngSurface, err = loadSurface("assets/06fish.png")
	return err
}

func run() {
	var e sdl.Event
	for {
		for e = sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			fmt.Printf("event %+v\n", e)
			switch e.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		screen.FillRect(nil, 0x0000FFFF) // Blue to show transparency working
		err := pngSurface.Blit(nil, screen, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to blit image: %s\n", err)
			os.Exit(1)
		}
		window.UpdateSurface()
	}
}

func close() {
	pngSurface.Free()
	window.Destroy()
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
	run()
	close()
}
