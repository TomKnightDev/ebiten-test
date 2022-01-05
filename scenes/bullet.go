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
	bullet      Entity
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

func (s *BulletScene) Update(state *GameState) (bool, error) {
	s.bullet.currentMoveTurnTime++

	if s.bullet.currentMoveTurnTime < s.bullet.moveTurnTimer {
		return false, nil
	}

	s.bullet.currentMoveTurnTime = 0

	x := 0
	y := 0

	if s.bullet.direction == "down" {
		y++
	} else if s.bullet.direction == "up" {
		y--
	} else if s.bullet.direction == "right" {
		x++
	} else if s.bullet.direction == "left" {
		x--
	}

	eo := s.bullet.entityObj
	if collision := eo.Check(float64(x), float64(y), "enemy"); collision != nil {
		return true, nil
	}

	s.bullet.xPos += x
	s.bullet.yPos += y
	s.bullet.entityObj.X = float64(s.bullet.xPos)
	s.bullet.entityObj.Y = float64(s.bullet.yPos)
	s.bullet.entityObj.Update()

	return false, nil
}

func NewBulletScene(s *SceneManager, start Tile, dir string, ignoreTag string) *BulletScene {
	b := &BulletScene{
		bullet: *NewEntity(s, start.x, start.y, dir, 0, 0, "bullet"),
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
