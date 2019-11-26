package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	tex *sdl.Texture
}

func newPlayer(renderer *sdl.Renderer) (p player, err error) {
	img, err := sdl.LoadBMP("sprites/necro.bmp")
	if err != nil {
		return player{}, fmt.Errorf("Loading player sprite: %v", err)
	}
	defer img.Free()

	p.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return player{}, fmt.Errorf("Creating player texture: %v", err)
	}
	return p, nil
}

func (p *player) draw(renderer *sdl.Renderer) {
	renderer.Copy(p.tex, &sdl.Rect{X: 0, Y: 0, W: 80, H: 80}, &sdl.Rect{X: 0, Y: 0, W: 80, H: 80})
}
