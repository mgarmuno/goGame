package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed                    = 0.05
	playerSize                     = 80
	playerPngSizeX, playerPngSizeY = 408, 424
)

type player struct {
	tex  *sdl.Texture
	x, y float64
}

func newPlayer(renderer *sdl.Renderer) (p player, err error) {
	playerImg, err := img.Load("sprites/wizard_fire/idle_1.png")
	if err != nil {
		return player{}, fmt.Errorf("Loading player sprite: %v", err)
	}
	defer playerImg.Free()

	p.tex, err = renderer.CreateTextureFromSurface(playerImg)
	if err != nil {
		return player{}, fmt.Errorf("Creating player texture: %v", err)
	}

	p.x = screenWidth * 0.1
	p.y = screenHeight*0.9 - playerSize/2.0

	return p, nil
}

func (p *player) draw(renderer *sdl.Renderer) {
	x := p.x - playerSize/2.0
	y := p.y - playerSize/2.0
	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: 408, H: 424},
		&sdl.Rect{X: int32(x), Y: int32(y), W: playerSize, H: playerSize})
}

func (p *player) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_A] == 1 {
		p.x -= playerSpeed
	} else if keys[sdl.SCANCODE_D] == 1 {
		p.x += playerSpeed
	} else if keys[sdl.SCANCODE_W] == 1 {

	} else if keys[sdl.SCANCODE_S] == 1 {

	} else if keys[sdl.SCANCODE_SPACE] == 1 {
		p.fire()
	}
}

func (p *player) fire() {

}
