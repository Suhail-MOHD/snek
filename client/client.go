package client

import (
	"fmt"
	"snek/common"

	"github.com/veandco/go-sdl2/sdl"
)

func DrawSnake(s *common.snake, surface *sdl.Surface) {
	rects := make([]sdl.Rect, 0)
	fmt.Println(s.pos[0])
	for _, c := range s.pos {
		rects = append(rects, sdl.Rect{int32(c.x), int32(c.y), 1, 1})
	}
	surface.FillRects(rects, 0xffffffff)
}

func DrawLevel(surface *sdl.Surface, l *common.level) {
	rect := sdl.Rect{0, 0, 1080, 900}
	surface.FillRect(&rect, 0)
	DrawSnake(l.snake, surface)
}

func MoveSnake(l *common.level, key int) {
	switch key {
	case int('w'):
		l.snake.changeDirection(vector{0, -1})
	case int('s'):
		l.snake.changeDirection(vector{0, 1})
	case int('a'):
		l.snake.changeDirection(vector{-1, 0})
	case int('d'):
		l.snake.changeDirection(vector{1, 0})
	}
}

func NewSnake(initial_pos coord, length int, direction vector) (*snake, error) {
	pos := make([]coord, length)
	for i := 0; i < len(pos); i++ {
		pos[i] = coord{x: initial_pos.x - i*direction.x, y: initial_pos.y - i*direction.y}
	}
	snake := snake{direction: direction, pos: pos}
	return &snake, nil
}

func (s *common.Snake) Move() {
	s.pos = append([]coord{coord{x: s.pos[0].x + s.direction.x, y: s.pos[0].y + s.direction.y}}, s.pos...)
	s.pos = s.pos[:len(s.pos)-1]
}

func (s *snake) changeDirection(direction vector) {
	s.direction = direction
}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	rect := sdl.Rect{0, 0, 1080, 900}
	surface.FillRect(&rect, 0xffff0000)
	window.UpdateSurface()

	s, err := NewSnake(coord{100, 100}, 50, vector{1, 0})
	l := level{s}

	running := true
	// ss
	// directions s:= []vector{vector{1 ,0}, vector{0, 1}, vector{-1, 0}, vector{0, -1}}
	// for true {
	// 	DrawLevel(surface, &l)
	// 	s.Move()
	// 	fmt.Println(s.pos[0])
	// 	window.UpdateSurface()
	// 	if num_steps == 0{
	// 		s.changeDirection(directions[num1])
	// 		num1 = (num1 + 1) % 4
	// 	}
	//   time.Sleep(time.Duration(17)*time.Millisecond)
	// 	num_steps = (num_steps + 1) % 50
	// }
	for running {
		for event := sdl.WaitEventTimeout(17); true; event = sdl.WaitEventTimeout(17) {
			switch e := event.(type) {
			case *sdl.KeyboardEvent:
				MoveSnake(&l, int(e.Keysym.Sym))
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
			s.Move()
			DrawLevel(surface, &l)
			window.UpdateSurface()
		}
	}
}
