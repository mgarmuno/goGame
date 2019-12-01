package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type keyboardMover struct {
	container *element
	speed     float64
	sr        *spriteRenderer
}

func newKeyboardMover(container *element, speed float64) *keyboardMover {
	return &keyboardMover{
		container: container,
		speed:     speed,
		sr:        container.getComponent(&spriteRenderer{}).(*spriteRenderer),
	}
}

func (mover *keyboardMover) onUpdate() error {
	keys := sdl.GetKeyboardState()

	cont := mover.container

	if keys[sdl.SCANCODE_A] == 1 {
		if cont.position.x-(mover.sr.spriteWidth/2.0) > 0 {
			cont.position.x -= mover.speed
		}
	} else if keys[sdl.SCANCODE_D] == 1 {
		if cont.position.x+(mover.sr.spriteHeight/2.0) < screenWidth {
			cont.position.x += mover.speed
		}
	} else if keys[sdl.SCANCODE_W] == 1 {

	} else if keys[sdl.SCANCODE_S] == 1 {

	}

	return nil
}

func (mover *keyboardMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

type keyboardCaster struct {
	container *element
	cooldown  time.Duration
	lastCast  time.Time
}

func newKeyboardCaster(container *element, cooldown time.Duration) *keyboardCaster {
	return &keyboardCaster{
		container: container,
		cooldown:  cooldown,
	}
}

func (caster *keyboardCaster) onUpdate() error {
	keys := sdl.GetKeyboardState()

	pos := caster.container.position

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(caster.lastCast) >= caster.cooldown {
			caster.castFireball(pos.x, pos.y)

			caster.lastCast = time.Now()
		}
	}

	return nil
}

func (caster *keyboardCaster) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (caster *keyboardCaster) castFireball(x, y float64) {
	if fb, ok := fireballFromPool(); ok {
		fb.active = true
		fb.x = x
		fb.y = y
		fb.angle = 360 * (math.Pi / 180)
	}
}
