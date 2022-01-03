package scenes

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"image"
	"image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	tileSize = 8
	tileXNum = 32
)

var (
	//go:embed resources/map001.png
	sprite_sheet []byte
	//go:embed resources/map001.png.json
	map001     []byte
	tilesImage *ebiten.Image
	tileMap    TileMap
	Obstacles  []Tile
)

type Tile struct {
	x int
	y int
}

type MainScene struct {
	count       int
	childScenes []Scene
}

type TileMap struct {
	Tileshigh int `json:"tileshigh"`
	Layers    []struct {
		Tiles []struct {
			X     int  `json:"x"`
			Rot   int  `json:"rot"`
			Y     int  `json:"y"`
			Index int  `json:"index"`
			FlipX bool `json:"flipX"`
			Tile  int  `json:"tile"`
		} `json:"tiles"`
		Name   string `json:"name"`
		Number int    `json:"number"`
	} `json:"layers"`
	Tileheight int `json:"tileheight"`
	Tileswide  int `json:"tileswide"`
	Tilewidth  int `json:"tilewidth"`
}

func init() {
	img, err := png.Decode(bytes.NewReader(sprite_sheet))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)

	if err := json.Unmarshal(map001, &tileMap); err != nil {
		panic(err)
	}

}

func (s *MainScene) Draw(r *ebiten.Image) {
	for l := len(tileMap.Layers) - 1; l >= 0; l-- {
		for _, t := range tileMap.Layers[l].Tiles {
			if tileMap.Layers[l].Name == "Obstacles" && t.Tile >= 0 {
				Obstacles = append(Obstacles, Tile{x: t.X * tileSize, y: t.Y * tileSize})
				continue
			}

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(t.X)*tileSize, float64(t.Y)*tileSize)

			sx := (t.Tile % tileXNum) * tileSize
			sy := (t.Tile / tileXNum) * tileSize
			r.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
		}
	}
}

func (s *MainScene) Update(state *GameState) error {
	for _, scene := range s.childScenes {
		if err := scene.Update(&GameState{
			SceneManager: state.SceneManager,
		}); err != nil {
			return err
		}
	}

	return nil
}

func NewMainScene() *MainScene {
	p := NewPlayerScene()
	e := NewEnemyScene()
	m := &MainScene{
		childScenes: []Scene{},
	}

	m.childScenes = append(m.childScenes, e, p)

	return m
}

func (s *MainScene) GetChildScenes() []Scene {
	return s.childScenes
}
