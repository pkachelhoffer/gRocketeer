package entity

import (
	"fmt"
	"gRocketeer/core/config"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type System struct {
	entities map[string]Entity

	init []string
}

func NewSystem() *System {
	return &System{
		entities: make(map[string]Entity),
		init:     make([]string, 0),
	}
}

func (s *System) Update() {
	if len(s.init) > 0 {
		for _, k := range s.init {
			e, ok := s.Get(k)
			if !ok {
				continue
			}

			if config.Debug {
				fmt.Println("sys init", k)
			}

			e.Init()
		}
		s.init = s.init[:0]
	}

	var toRemove []string
	for k, l := range s.entities {
		if l.Update() {
			l.Terminate()
			toRemove = append(toRemove, k)
		}
	}

	for _, key := range toRemove {
		if config.Debug {
			fmt.Println("sys del", key)
		}
		delete(s.entities, key)
	}
}

func (s *System) Draw(target *ebiten.Image) {
	for _, l := range s.entities {
		l.Draw(target)
	}
}

func (s *System) Add(key string, ent Entity) {
	if config.Debug {
		fmt.Println("sys add", key)
	}

	s.entities[key] = ent
	s.init = append(s.init, key)
}

func (s *System) Remove(key string) {
	_, ok := s.entities[key]
	if ok {
		if config.Debug {
			fmt.Println("sys remove", key)
		}
		delete(s.entities, key)
	}
}

func (s *System) Get(key string) (Entity, bool) {
	e, ok := s.entities[key]
	if ok {
		return e, true
	}

	return nil, false
}

func Get[T Entity](s *System, key string) T {
	entity, ok := s.Get(key)
	if !ok {
		log.Fatal(fmt.Sprintf("failed getting %s", key))
	}

	tp, ok := entity.(T)
	if !ok {
		log.Fatal(fmt.Sprintf("incorrect type cast %s", key))
	}

	return tp
}
