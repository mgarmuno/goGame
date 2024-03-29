package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 1500
	screenHeight = 900
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initialining SDL:", err)
	}

	window, err := sdl.CreateWindow(
		"Gaming in go",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("initializing window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}
	defer renderer.Destroy()

	plr := newPlayer(renderer)
	if err != nil {
		fmt.Println("Creating player: ", err)
		return
	}

	enemy := newBasicEnemy(renderer, screenWidth*0.9, screenHeight*0.9-enemySize/2.0)
	if err != nil {
		fmt.Println("Creating basic enemy:", err)
		return
	}

	initFireballPool(renderer)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		err = plr.draw(renderer)
		if err != nil {
			fmt.Println("drawing palyer:", err)
			return
		}
		err = plr.update()
		if err != nil {
			fmt.Println("updating player:", err)
			return
		}

		for _, fb := range fireballPool {
			fb.draw(renderer)
			fb.update()
		}

		enemy.draw(renderer)
		renderer.Present()
	}
}
