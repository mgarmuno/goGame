package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type spriteRenderer struct {
	container                 *element
	tex                       *sdl.Texture
	width, height             float64
	spriteWidth, spriteHeight float64
}

func newSpriteRenderer(container *element, renderer *sdl.Renderer, filename string) *spriteRenderer {
	tex := textureFromPNG(renderer, filename)

	_, _, width, height, err := tex.Query()
	if err != nil {
		panic(fmt.Errorf("querying texture: %v", err))
	}

	return &spriteRenderer{
		container:    container,
		tex:          textureFromPNG(renderer, filename),
		width:        float64(width),
		height:       float64(height),
		spriteWidth:  float64(width) / 5,
		spriteHeight: float64(width) / 5,
	}
}

func (sr *spriteRenderer) onDraw(renderer *sdl.Renderer) error {

	x := sr.container.position.x - sr.spriteWidth/2.0
	y := sr.container.position.y - sr.spriteHeight/2.0

	renderer.CopyEx(
		sr.tex,
		&sdl.Rect{X: 0, Y: 0, W: int32(sr.width), H: int32(sr.height)},
		&sdl.Rect{X: int32(x), Y: int32(y), W: int32(sr.spriteWidth), H: int32(sr.spriteHeight)},
		sr.container.rotation,
		&sdl.Point{X: int32(sr.width) / 2, Y: int32(sr.height) / 2},
		sdl.FLIP_NONE,
	)

	return nil
}

func (sr *spriteRenderer) onUpdate() error {
	return nil
}

func textureFromPNG(renderer *sdl.Renderer, filename string) *sdl.Texture {
	image, err := img.Load(filename)
	if err != nil {
		panic(fmt.Errorf("loading %v: %v", filename, err))
	}
	defer image.Free()
	tex, err := renderer.CreateTextureFromSurface(image)
	if err != nil {
		panic(fmt.Errorf("Creating texture from: %v %v", filename, err))
	}

	return tex
}
