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
			return err
		}
		l.mTexture = nil
		l.mWidth = 0
		l.mHeight = 0
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
	if err := loadedSurface.SetColorKey(true, sdl.MapRGB(loadedSurface.Format, 0, 255, 255)); err != nil {
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
func (l *LTexture) Render(x, y int32, clip *sdl.Rect) error {
	renderQuad := sdl.Rect{X: x, Y: y, W: l.mWidth, H: l.mHeight}

	if clip != nil {
		renderQuad.W = clip.W
		renderQuad.H = clip.H
	}

	return renderer.Copy(l.mTexture, clip, &renderQuad)
}

// Height returns the height of the LTexture
func (l *LTexture) Height() int32 {
	return l.mHeight
}

// Width returns the width of the LTexture
func (l *LTexture) Width() int32 {
	return l.mWidth
}

// NewLTexture returns a default LTexture
func NewLTexture() LTexture {
	return LTexture{mTexture: nil, mHeight: 0, mWidth: 0}
}
