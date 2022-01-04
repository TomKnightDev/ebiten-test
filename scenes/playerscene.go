package scenes

import (
	"bytes"
	_ "embed"
	"image"
	"image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	//go:embed resources/player.png
	playerPng    []byte
	playerSprite *ebiten.Image
)

type PlayerScene struct {
	count       int
	childScenes []Scene
	player      Character
}

func init() {
	// Get image here
	img, err := png.Decode(bytes.NewReader(playerPng))
	if err != nil {
		log.Fatal(err)
	}

	playerSprite = ebiten.NewImageFromImage(img)

}

func (s *PlayerScene) Draw(r *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.player.xPos), float64(s.player.yPos))

	x := s.player.sprites[s.player.direction].x
	y := s.player.sprites[s.player.direction].y

	r.DrawImage(playerSprite.SubImage(image.Rect(x, y, x+tileSize, y+tileSize)).(*ebiten.Image), op)
}

func (s *PlayerScene) Update(state *GameState) error {

	ActionUpdate(s)
	MoveUpdate(s)

	return nil
}

func ActionUpdate(s *PlayerScene) {
	s.player.currentActionTurnTime++

	if s.player.currentActionTurnTime < s.player.actionTurnTimer {
		return
	}

	// Fire
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		s.childScenes = append(s.childScenes, NewBulletScene(
			Tile{
				x: s.player.xPos,
				y: s.player.yPos},
			s.player.direction))
		s.player.currentActionTurnTime = 0
	}
}

func MoveUpdate(s *PlayerScene) {
	s.player.currentMoveTurnTime++

	if s.player.currentMoveTurnTime < s.player.moveTurnTimer {
		return
	}

	nextXPos := s.player.xPos
	nextYPos := s.player.yPos
	nextDirection := s.player.direction

	for _, pk := range inpututil.AppendPressedKeys(nil) {
		if pk == ebiten.KeyW {
			nextYPos -= tileSize
			nextDirection = "up"
		}
		if pk == ebiten.KeyS {
			nextYPos += tileSize
			nextDirection = "down"
		}
		if pk == ebiten.KeyA {
			nextXPos -= tileSize
			nextDirection = "left"
		}
		if pk == ebiten.KeyD {
			nextXPos += tileSize
			nextDirection = "right"
		}
	}

	if nextDirection != s.player.direction {
		s.player.direction = nextDirection
		s.player.currentMoveTurnTime = 0
		return
	}

	if (nextXPos != s.player.xPos || nextYPos != s.player.yPos) && CanTraverse(Tile{x: nextXPos, y: nextYPos}) {
		s.player.xPos = nextXPos
		s.player.yPos = nextYPos
		s.player.currentMoveTurnTime = 0
	}
}

func NewPlayerScene() *PlayerScene {
	p := &PlayerScene{
		player: *NewCharacter(18*tileSize, 15*tileSize, "down", 20, 10),
	}

	p.player.MapSprites("up", 9, 8)
	p.player.MapSprites("down", 3, 8)
	p.player.MapSprites("right", 0, 8)
	p.player.MapSprites("left", 6, 8)
	return p
}

func (s *PlayerScene) GetChildScenes() []Scene {
	return s.childScenes
}
