package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed                    = 0.05
	playerSize                     = 80
	playerPngSizeX, playerPngSizeY = 408, 424
	payerCastCooldown              = time.Millisecond * 500
)

func newPlayer(renderer *sdl.Renderer) *element {
	player := &element{}

	player.position = vector{
		x: screenWidth * 0.1,
		y: screenHeight*0.9 - playerSize/2.0,
	}

	player.active = true

	sr := newSpriteRenderer(player, renderer, "sprites/wizard_fire/idle_1.png")
	player.addComponent(sr)

	mover := newKeyboardMover(player, playerSpeed)
	player.addComponent(mover)

	caster := newKeyboardCaster(player, payerCastCooldown)
	player.addComponent(caster)

	return player
}
