package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	basicEnemySizeX, basicEnemySizeY = 331, 299
	enemySize                        = 80
)

type basicEnemy struct {
	tex  *sdl.Texture
	x, y float64
}

func newBasicEnemy(renderer *sdl.Renderer, x, y float64) (be basicEnemy) {
	be.tex = textureFromPNG(renderer, "sprites/monsters/3/3_enemies_1_attack_003.png")

	be.x = x
	be.y = y

	return be
}

func (be *basicEnemy) draw(renderer *sdl.Renderer) {
	x := be.x - enemySize/2.0
	y := be.y - enemySize/2.0
	renderer.Copy(be.tex,
		&sdl.Rect{X: 0, Y: 0, W: basicEnemySizeX, H: basicEnemySizeY},
		&sdl.Rect{X: int32(x), Y: int32(y), W: enemySize, H: enemySize})
}
