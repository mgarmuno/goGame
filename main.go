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

	img, err := sdl.LoadBMP("sprites/necro.bmp")
	if err != nil {
		fmt.Println("loading player sprite", err)
		return
	}
	defer img.Free()
	playerTex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		fmt.Println("creating player texture:", err)
		return
	}
	defer playerTex.Destroy()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		renderer.Copy(playerTex, &sdl.Rect{X: 0, Y: 0, W: 80, H: 80}, &sdl.Rect{X: 0, Y: 0, W: 80, H: 80})
		renderer.Present()
	}
}
