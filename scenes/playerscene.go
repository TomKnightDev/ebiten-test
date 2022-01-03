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
	player       Character
)

type PlayerScene struct {
	count       int
	childScenes []Scene
}

func init() {
	// Get image here
	img, err := png.Decode(bytes.NewReader(playerPng))
	if err != nil {
		log.Fatal(err)
	}

	playerSprite = ebiten.NewImageFromImage(img)

	player = *NewCharacter(18*tileSize, 15*tileSize, "down", 20)
	player.MapSprites("up", 9)
	player.MapSprites("down", 3)
	player.MapSprites("right", 0)
	player.MapSprites("left", 6)
}

func (s *PlayerScene) Draw(r *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(player.xPos), float64(player.yPos))

	x := player.sprites[player.direction].x
	y := player.sprites[player.direction].y

	r.DrawImage(playerSprite.SubImage(image.Rect(x, y, x+tileSize, y+tileSize)).(*ebiten.Image), op)
}

func (s *PlayerScene) Update(state *GameState) error {
	player.currentTurnTime++

	if player.currentTurnTime < player.turnTimer {
		return nil
	}

	nextXPos := player.xPos
	nextYPos := player.yPos
	nextDirection := player.direction

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

	if (nextXPos != player.xPos || nextYPos != player.yPos) && CanTraverse(Tile{x: nextXPos, y: nextYPos}) {
		player.xPos = nextXPos
		player.yPos = nextYPos
		player.direction = nextDirection

		player.currentTurnTime = 0
	}

	return nil
}

func NewPlayerScene() *PlayerScene {
	return &PlayerScene{}
}

func (s *PlayerScene) GetChildScenes() []Scene {
	return s.childScenes
}
