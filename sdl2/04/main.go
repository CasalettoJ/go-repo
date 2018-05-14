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

type keyPressSurface int

const (
	keyPressSurfaceDefault keyPressSurface = iota
	keyPressSurfaceUp
	keyPressSurfaceDown
	keyPressSurfaceLeft
	keyPressSurfaceRight
	keyPressSurfaceTotal
)

var (
	window           *sdl.Window
	screen           *sdl.Surface
	currentSurface   *sdl.Surface
	keyPressSurfaces [keyPressSurfaceTotal]*sdl.Surface
)

func setupWindow() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	window, err = sdl.CreateWindow("Reacting to Button Presses (Up, Down, Left, Right)", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	screen, err = window.GetSurface()
	return err
}

func loadMedia() error {
	var err error
	keyPressSurfaces[keyPressSurfaceDefault], err = sdl.LoadBMP("assets/04default.bmp")
	keyPressSurfaces[keyPressSurfaceUp], err = sdl.LoadBMP("assets/04up.bmp")
	keyPressSurfaces[keyPressSurfaceDown], err = sdl.LoadBMP("assets/04down.bmp")
	keyPressSurfaces[keyPressSurfaceLeft], err = sdl.LoadBMP("assets/04left.bmp")
	keyPressSurfaces[keyPressSurfaceRight], err = sdl.LoadBMP("assets/04right.bmp")
	return err
}

func run() {
	var e sdl.Event
	currentSurface = keyPressSurfaces[keyPressSurfaceDefault]
	for {
		for e = sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			fmt.Printf("event %+v\n", e)
			switch t := e.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.KeyboardEvent:
				if t.Type == sdl.KEYDOWN {
					switch t.Keysym.Sym {
					case sdl.K_UP:
						currentSurface = keyPressSurfaces[keyPressSurfaceUp]
					case sdl.K_DOWN:
						currentSurface = keyPressSurfaces[keyPressSurfaceDown]
					case sdl.K_LEFT:
						currentSurface = keyPressSurfaces[keyPressSurfaceLeft]
					case sdl.K_RIGHT:
						currentSurface = keyPressSurfaces[keyPressSurfaceRight]
					default:
						currentSurface = keyPressSurfaces[keyPressSurfaceDefault]
					}
				}
			}
		}
		screen.FillRect(nil, 0)
		err := currentSurface.Blit(nil, screen, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to blit image: %s\n", err)
			os.Exit(1)
		}
		window.UpdateSurface()
	}
}

func close() {
	for _, surface := range keyPressSurfaces {
		surface.Free()
	}
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
