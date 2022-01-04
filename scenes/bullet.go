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
	//go:embed resources/bullets.png
	bulletsPng   []byte
	bulletSprite *ebiten.Image
)

type BulletScene struct {
	count       int
	childScenes []Scene
	bullet      Character
}

func init() {
	img, err := png.Decode(bytes.NewReader(bulletsPng))
	if err != nil {
		log.Fatal(err)
	}

	bulletSprite = ebiten.NewImageFromImage(img)

}

func (s *BulletScene) Draw(r *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.bullet.xPos), float64(s.bullet.yPos))

	x := s.bullet.sprites[s.bullet.direction].x
	y := s.bullet.sprites[s.bullet.direction].y

	r.DrawImage(bulletSprite.SubImage(image.Rect(x, y, x+tileSize, y+tileSize)).(*ebiten.Image), op)
}

func (s *BulletScene) Update(state *GameState) error {
	s.bullet.currentMoveTurnTime++

	if s.bullet.currentMoveTurnTime < s.bullet.moveTurnTimer {
		return nil
	}

	s.bullet.currentMoveTurnTime = 0

	if s.bullet.direction == "down" {
		s.bullet.yPos++
	} else if s.bullet.direction == "up" {
		s.bullet.yPos--
	} else if s.bullet.direction == "right" {
		s.bullet.xPos++
	} else if s.bullet.direction == "left" {
		s.bullet.xPos--
	}

	return nil
}

func NewBulletScene(start Tile, dir string) *BulletScene {
	b := &BulletScene{
		bullet: *NewCharacter(start.x, start.y, dir, 0, 0),
	}
	// bullet := Character{}
	// bullet.xPos = start.x
	// bullet.yPos = start.y
	// bullet.direction = dir
	// bullet.turnTimer = 2
	b.bullet.MapSprites("up", 1, 12)
	b.bullet.MapSprites("down", 25, 12)
	b.bullet.MapSprites("right", 14, 12)
	b.bullet.MapSprites("left", 12, 12)

	return b
}

func (s *BulletScene) GetChildScenes() []Scene {
	return s.childScenes
}

// func (s *BulletScene) SetBulletVals(x int, y int, dir string) {
// 	s.bullet.xPos = x
// 	s.bullet.yPos = y
// 	s.bullet.direction = dir
// }
