package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	fireballSizeX, fireballSizeY       = 432, 144
	fireballPngSizeX, fireballPngSizeY = 50, 10
	fireballSpeed                      = 0.14
)

type fireball struct {
	tex    *sdl.Texture
	x, y   float64
	active bool
	angle  float64
}

func newFireball(renderer *sdl.Renderer) (fb fireball) {
	fb.tex = textureFromPNG(renderer, "sprites/fueguito/efecto_fuego_00020.png")
	return fb
}

func (fb *fireball) draw(renderer *sdl.Renderer) {
	x := fb.x - fireballPngSizeX/2.0
	y := fb.y - fireballPngSizeY/2.0
	renderer.Copy(fb.tex,
		&sdl.Rect{X: 0, Y: 0, W: fireballSizeX, H: fireballSizeY},
		&sdl.Rect{X: int32(x), Y: int32(y), W: fireballPngSizeX, H: fireballPngSizeY})
}

func (fb *fireball) update() {
	fb.x += fireballSpeed * math.Cos(fb.angle)
	fb.y += fireballSpeed * math.Sin(fb.angle)
}

var fireballPool []*fireball

func initFireballPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		fb := newFireball(renderer)
		fireballPool = append(fireballPool, &fb)
	}
}

func fireballFromPool() (*fireball, bool) {
	for _, fb := range fireballPool {
		if !fb.active {
			return fb, true
		}
	}

	return nil, false
}
