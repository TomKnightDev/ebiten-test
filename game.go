package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
	"github.com/tomknightdev/ebiten-test/scenes"
)

const (
	ScreenWidth  = 384
	ScreenHeight = 384
)

type Game struct {
	sceneManager *scenes.SceneManager
	// 	input        Input
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	if g.sceneManager == nil {
		g.sceneManager = &scenes.SceneManager{
			Space: resolv.NewSpace(384, 384, 8, 8),
		}
		g.sceneManager.GoTo(&scenes.TitleScene{})
	}

	// g.input.Update()
	if _, err := g.sceneManager.Update(); err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}
