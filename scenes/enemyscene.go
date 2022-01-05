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
	enemy       Entity
	target      *Entity
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
	op.GeoM.Translate(float64(s.enemy.xPos), float64(s.enemy.yPos))

	x := s.enemy.sprites[s.enemy.direction].x
	y := s.enemy.sprites[s.enemy.direction].y

	r.DrawImage(enemy1Sprite.SubImage(image.Rect(x, y, x+tileSize, y+tileSize)).(*ebiten.Image), op)
}

func (s *EnemyScene) Update(state *GameState) (bool, error) {
	s.enemy.currentMoveTurnTime++

	if s.enemy.currentMoveTurnTime < s.enemy.moveTurnTimer {
		return false, nil
	}

	nextXPos := s.enemy.xPos
	nextYPos := s.enemy.yPos
	nextDirection := s.enemy.direction

	// Move towards player
	if s.enemy.xPos < s.target.xPos {
		nextXPos += tileSize
		nextDirection = "right"
	} else if s.enemy.xPos > s.target.xPos {
		nextXPos -= tileSize
		nextDirection = "left"
	} else if s.enemy.yPos < s.target.yPos {
		nextYPos += tileSize
		nextDirection = "down"
	} else if s.enemy.yPos > s.target.yPos {
		nextYPos -= tileSize
		nextDirection = "up"
	}

	if (nextXPos != s.target.xPos || nextYPos != s.target.yPos) && CanTraverse(Tile{x: nextXPos, y: nextYPos}) {
		s.enemy.xPos = nextXPos
		s.enemy.yPos = nextYPos
		s.enemy.entityObj.X = float64(nextXPos)
		s.enemy.entityObj.Y = float64(nextYPos)
		s.enemy.entityObj.Update()
		s.enemy.direction = nextDirection
	}

	s.enemy.currentMoveTurnTime = 0

	return false, nil
}

func NewEnemyScene(s *SceneManager, target *Entity) *EnemyScene {
	e := &EnemyScene{
		enemy:  *NewEntity(s, 22*tileSize, 18*tileSize, "down", 40, 20, "enemy"),
		target: target,
	}

	e.enemy.MapSprites("up", 9, 8)
	e.enemy.MapSprites("down", 3, 8)
	e.enemy.MapSprites("right", 0, 8)
	e.enemy.MapSprites("left", 6, 8)
	return e
}

func (s *EnemyScene) GetChildScenes() []Scene {
	return s.childScenes
}
