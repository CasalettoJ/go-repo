package main

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenHeight = 480
	screenWidth  = 640
)

func main() {
	var window *sdl.Window
	var surface *sdl.Surface

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println("Failure to initialize SDL")
		log.Panic(err)
	}
	defer sdl.Quit()

	window, err = sdl.CreateWindow("Title", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println("Failure to create window")
		log.Panic(err)
	}
	defer window.Destroy()

	surface, err = window.GetSurface()
	if err != nil {
		fmt.Println("Failed to get surface of window")
		log.Panic(err)
	}
	surface.FillRect(nil, 0)
	window.UpdateSurface()
	sdl.Delay(5000)
}
