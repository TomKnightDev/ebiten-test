package scenes

import (
	"bytes"
	_ "embed"
	"image"
	"image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed resources/enemy1.png
	enemy1Png    []byte
	enemy1Sprite *ebiten.Image
	enemy1       Character
)

type EnemyScene struct {
	count       int
	childScenes []Scene
}

func init() {
	img, err := png.Decode(bytes.NewReader(enemy1Png))
	if err != nil {
		log.Fatal(err)
	}

	enemy1Sprite = ebiten.NewImageFromImage(img)

	enemy1 = *NewCharacter(22*tileSize, 18*tileSize, "down", 40)
	enemy1.MapSprites("up", 9)
	enemy1.MapSprites("down", 3)
	enemy1.MapSprites("right", 0)
	enemy1.MapSprites("left", 6)
}

func (s *EnemyScene) Draw(r *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(enemy1.xPos), float64(enemy1.yPos))

	x := enemy1.sprites[enemy1.direction].x
	y := enemy1.sprites[enemy1.direction].y

	r.DrawImage(enemy1Sprite.SubImage(image.Rect(x, y, x+tileSize, y+tileSize)).(*ebiten.Image), op)
}

func (s *EnemyScene) Update(state *GameState) error {
	enemy1.currentTurnTime++

	if enemy1.currentTurnTime < enemy1.turnTimer {
		return nil
	}

	nextXPos := enemy1.xPos
	nextYPos := enemy1.yPos
	nextDirection := enemy1.direction

	// Move towards player
	if enemy1.xPos < player.xPos {
		nextXPos += tileSize
		nextDirection = "right"
	} else if enemy1.xPos > player.xPos {
		nextXPos -= tileSize
		nextDirection = "left"
	} else if enemy1.yPos < player.yPos {
		nextYPos += tileSize
		nextDirection = "down"
	} else if enemy1.yPos > player.yPos {
		nextYPos -= tileSize
		nextDirection = "up"
	}

	if (nextXPos != player.xPos || nextYPos != player.yPos) && CanTraverse(Tile{x: nextXPos, y: nextYPos}) {
		enemy1.xPos = nextXPos
		enemy1.yPos = nextYPos
		enemy1.direction = nextDirection
	}

	enemy1.currentTurnTime = 0

	return nil
}

func NewEnemyScene() *EnemyScene {
	return &EnemyScene{}
}

func (s *EnemyScene) GetChildScenes() []Scene {
	return s.childScenes
}
