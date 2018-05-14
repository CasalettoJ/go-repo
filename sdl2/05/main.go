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
	window           *sdl.Window
	screen           *sdl.Surface
	stretchedSurface *sdl.Surface
)

func setupWindow() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	window, err = sdl.CreateWindow("Optimized Surface Loading and Soft Stretching", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	screen, err = window.GetSurface()
	return err
}

func loadSurface(path string) (*sdl.Surface, error) {
	var err error
	var optimizedSurface, loadedSurface *sdl.Surface
	loadedSurface, err = sdl.LoadBMP(path)
	if err != nil {
		return nil, err
	}
	if optimizedSurface, err = loadedSurface.Convert(screen.Format, 0); err != nil {
		return nil, err
	}
	loadedSurface.Free()
	return optimizedSurface, nil
}

func loadMedia() error {
	var err error
	stretchedSurface, err = loadSurface("assets/05boots.bmp")
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
		screen.FillRect(nil, 0)
		stretchRect := sdl.Rect{X: 0, Y: 0, W: screenWidth, H: screenHeight}
		err := stretchedSurface.BlitScaled(nil, screen, &stretchRect)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to blit image: %s\n", err)
			os.Exit(1)
		}
		window.UpdateSurface()
	}
}

func close() {
	stretchedSurface.Free()
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
