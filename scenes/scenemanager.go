package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

var (
	transitionFrom = ebiten.NewImage(384, 384)
	transitionTo   = ebiten.NewImage(384, 384)
)

type Scene interface {
	Update(state *GameState) (bool, error)
	Draw(screen *ebiten.Image)
	GetChildScenes() []Scene
}

const transitionMaxCount = 20

type GameState struct {
	SceneManager *SceneManager
	// Input        *Input
}

type SceneManager struct {
	current         Scene
	next            Scene
	transitionCount int
	Space           *resolv.Space
}

func (s *SceneManager) Update() (bool, error) {
	if s.transitionCount == 0 {
		_, mainErr := s.current.Update(&GameState{
			SceneManager: s,
			// Input:        input,
		})

		UpdateChildren(s, s.current)

		return false, mainErr
	}

	s.transitionCount--
	if s.transitionCount > 0 {
		return false, nil
	}

	s.current = s.next
	s.next = nil
	return false, nil
}

func UpdateChildren(sceneManager *SceneManager, currentScene Scene) error {
	cs := currentScene.GetChildScenes()
	for i, scene := range cs {
		dispose, err := scene.Update(&GameState{
			SceneManager: sceneManager,
		})
		if dispose {
			cs = RemoveScene(cs, i)
			continue
		}
		if err != nil {
			return err
		}

		UpdateChildren(sceneManager, scene)
	}

	return nil
}

func RemoveScene(s []Scene, i int) []Scene {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]

	// return append(scenes[:s], scenes[s+1:]...)
}

func (s *SceneManager) Draw(r *ebiten.Image) {

	if s.transitionCount == 0 {
		s.current.Draw(r)
		DrawChildren(r, s.current.GetChildScenes())
		return
	}

	transitionFrom.Clear()
	s.current.Draw(transitionFrom)

	transitionTo.Clear()
	s.next.Draw(transitionTo)

	r.DrawImage(transitionFrom, nil)

	alpha := 1 - float64(s.transitionCount)/float64(transitionMaxCount)
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, alpha)
	r.DrawImage(transitionTo, op)
}

func DrawChildren(r *ebiten.Image, children []Scene) {
	for _, scene := range children {
		scene.Draw(r)
		DrawChildren(r, scene.GetChildScenes())
	}
}

func (s *SceneManager) GoTo(scene Scene) {
	if s.current == nil {
		s.current = scene
	} else {
		s.next = scene
		s.transitionCount = transitionMaxCount
	}
}
