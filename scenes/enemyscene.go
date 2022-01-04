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
)

type EnemyScene struct {
	count       int
	childScenes []Scene
	enemy1      Character
	target      *Character
}

func init() {
	img, err := png.Decode(bytes.NewReader(enemy1Png))
	if err != nil {
		log.Fatal(err)
	}

	enemy1Sprite = ebiten.NewImageFromImage(img)

}

func (s *EnemyScene) Draw(r *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.enemy1.xPos), float64(s.enemy1.yPos))

	x := s.enemy1.sprites[s.enemy1.direction].x
	y := s.enemy1.sprites[s.enemy1.direction].y

	r.DrawImage(enemy1Sprite.SubImage(image.Rect(x, y, x+tileSize, y+tileSize)).(*ebiten.Image), op)
}

func (s *EnemyScene) Update(state *GameState) error {
	s.enemy1.currentMoveTurnTime++

	if s.enemy1.currentMoveTurnTime < s.enemy1.moveTurnTimer {
		return nil
	}

	nextXPos := s.enemy1.xPos
	nextYPos := s.enemy1.yPos
	nextDirection := s.enemy1.direction

	// Move towards player
	if s.enemy1.xPos < s.target.xPos {
		nextXPos += tileSize
		nextDirection = "right"
	} else if s.enemy1.xPos > s.target.xPos {
		nextXPos -= tileSize
		nextDirection = "left"
	} else if s.enemy1.yPos < s.target.yPos {
		nextYPos += tileSize
		nextDirection = "down"
	} else if s.enemy1.yPos > s.target.yPos {
		nextYPos -= tileSize
		nextDirection = "up"
	}

	if (nextXPos != s.target.xPos || nextYPos != s.target.yPos) && CanTraverse(Tile{x: nextXPos, y: nextYPos}) {
		s.enemy1.xPos = nextXPos
		s.enemy1.yPos = nextYPos
		s.enemy1.direction = nextDirection
	}

	s.enemy1.currentMoveTurnTime = 0

	return nil
}

func NewEnemyScene(target *Character) *EnemyScene {
	e := &EnemyScene{
		enemy1: *NewCharacter(22*tileSize, 18*tileSize, "down", 40, 20),
		target: target,
	}

	e.enemy1.MapSprites("up", 9, 8)
	e.enemy1.MapSprites("down", 3, 8)
	e.enemy1.MapSprites("right", 0, 8)
	e.enemy1.MapSprites("left", 6, 8)
	return e
}

func (s *EnemyScene) GetChildScenes() []Scene {
	return s.childScenes
}
