package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed = 0.05
	playerSize  = 80
)

type player struct {
	tex  *sdl.Texture
	x, y float64
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

	p.y = screenWidth / 2.0
	p.x = screenHeight - playerSize/2.0

	return p, nil
}

func (p *player) draw(renderer *sdl.Renderer) {
	x := p.x - playerSize/2.0
	y := p.y - playerSize/2.0
	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: 80, H: 80},
		&sdl.Rect{X: int32(x), Y: int32(y), W: 80, H: 80})
}

func (p *player) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_A] == 1 {
		p.x -= playerSpeed
	} else if keys[sdl.SCANCODE_D] == 1 {
		p.x += playerSpeed
	} else if keys[sdl.SCANCODE_W] == 1 {

	} else if keys[sdl.SCANCODE_S] == 1 {

	}
}
