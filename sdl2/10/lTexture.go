package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// LTexture is a texture wrapper classes
type LTexture struct {
	mTexture *sdl.Texture
	mWidth   int32
	mHeight  int32
}

// Free releases the texture from memory
func (l *LTexture) Free() error {
	if l.mTexture != nil {
		if err := l.mTexture.Destroy(); err != nil {
			l.mTexture = nil
			return err
		}
	}
	return nil
}

// LoadFromFile loads a texture from a file path
func (l *LTexture) LoadFromFile(path string) error {
	// get rid of preexisting texture
	if err := l.Free(); err != nil {
		return err
	}

	var newTexture *sdl.Texture
	loadedSurface, err := img.Load(path)
	if err != nil {
		return err
	}

	// Color key image to blue
	if err := loadedSurface.SetColorKey(true, sdl.MapRGB(loadedSurface.Format, 0, 0, 255)); err != nil {
		return err
	}

	// Create texture from surface pixels
	newTexture, err = renderer.CreateTextureFromSurface(loadedSurface)
	if err != nil {
		return err
	}
	// Get image dimensions
	l.mWidth = loadedSurface.W
	l.mHeight = loadedSurface.H
	l.mTexture = newTexture

	loadedSurface.Free()

	return nil

}

// Render sets the rendering space and renders the texture to the screen
func (l *LTexture) Render(x, y int32) error {
	renderQuad := sdl.Rect{X: x, Y: y, W: l.mWidth, H: l.mHeight}
	return renderer.Copy(l.mTexture, nil, &renderQuad)
}

// GetHeight returns the height of the LTexture
func (l *LTexture) GetHeight() int32 {
	return l.mHeight
}

// GetWidth returns the width of the LTexture
func (l *LTexture) GetWidth() int32 {
	return l.mWidth
}

// NewLTexture returns a default LTexture
func NewLTexture() LTexture {
	return LTexture{mTexture: nil, mHeight: 0, mWidth: 0}
}
