package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type TitleScene struct {
	count int
}

func (s *TitleScene) Draw(r *ebiten.Image) {
	ebitenutil.DebugPrint(r, "Title scene")

	// s.drawTitleBackground(r, s.count)
	// drawLogo(r, "BLOCKS")

	// message := "PRESS SPACE TO START"
	// x := 0
	// y := ScreenHeight - 48
	// drawTextWithShadowCenter(r, message, x, y, 1, color.NRGBA{0x80, 0, 0, 0xff}, ScreenWidth)
}

func (s *TitleScene) Update(state *GameState) error {
	s.count++
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		state.SceneManager.GoTo(NewMainScene())
		return nil
	}

	return nil
}
