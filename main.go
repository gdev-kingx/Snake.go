package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	width  = 30
	height = 15
)

type Point struct {
	x int
	y int
}

type Snake struct {
	body []Point
	dir  Point
}

type Food struct {
	position Point
}

func (s *Snake) move() {
	head := s.body[0]
	newHead := Point{head.x + s.dir.x, head.y + s.dir.y}
	s.body = append([]Point{newHead}, s.body[:len(s.body)-1]...)
}

func (s *Snake) grow() {
	head := s.body[0]
	newHead := Point{head.x + s.dir.x, head.y + s.dir.y}
	s.body = append([]Point{newHead}, s.body...)
}

func (f *Food) spawn(snake Snake) {
	for {
		f.position = Point{rand.Intn(width-2) + 1, rand.Intn(height-2) + 1} // Spawn food within borders
		if !f.isOnSnake(snake) {
			break
		}
	}
}

func (f *Food) isOnSnake(snake Snake) bool {
	for _, p := range snake.body {
		if p == f.position {
			return true
		}
	}
	return false
}

func (s *Snake) changeDirection(newDir Point) {
	if (s.dir.x == -newDir.x && s.dir.y == 0) || (s.dir.y == -newDir.y && s.dir.x == 0) {
		return
	}
	s.dir = newDir
}

func draw(snake Snake, food Food, score int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// Draw borders
	for x := 0; x < width; x++ {
		termbox.SetCell(x, 0, '#', termbox.ColorWhite, termbox.ColorDefault) // Top border
		termbox.SetCell(x, height-1, '#', termbox.ColorWhite, termbox.ColorDefault) // Bottom border
	}
	for y := 0; y < height; y++ {
		termbox.SetCell(0, y, '#', termbox.ColorWhite, termbox.ColorDefault) // Left border
		termbox.SetCell(width-1, y, '#', termbox.ColorWhite, termbox.ColorDefault) // Right border
	}

	// Draw snake
	for _, p := range snake.body {
		termbox.SetCell(p.x, p.y, 'O', termbox.ColorGreen, termbox.ColorDefault)
	}

	// Draw food
	termbox.SetCell(food.position.x, food.position.y, 'X', termbox.ColorRed, termbox.ColorDefault)

	// Draw score
	scoreText := fmt.Sprintf("Score: %d", score)
	for i, c := range scoreText {
		termbox.SetCell(i+1, height-2, c, termbox.ColorYellow, termbox.ColorDefault)
	}

	termbox.Flush()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	snake := Snake{
		body: []Point{{5, 5}, {4, 5}, {3, 5}},
		dir:  Point{1, 0},
	}
	food := Food{}
	food.spawn(snake)

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	gameOver := false
	score := 0

	go func() {
		for !gameOver {
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyArrowUp:
					snake.changeDirection(Point{0, -1})
				case termbox.KeyArrowDown:
					snake.changeDirection(Point{0, 1})
				case termbox.KeyArrowLeft:
					snake.changeDirection(Point{-1, 0})
				case termbox.KeyArrowRight:
					snake.changeDirection(Point{1, 0})
				case termbox.KeyEsc:
					gameOver = true
				}
			case termbox.EventError:
				panic(ev.Err)
			}
		}
	}()

	for !gameOver {
		select {
	case <-ticker.C:
			snake.move()

			// Check for collision with food
			if snake.body[0] == food.position {
				snake.grow()
				food.spawn(snake)
				score++ // Increase score when food is eaten
			}

			// Check for collision with borders
			head := snake.body[0]
			if head.x <= 0 || head.x >= width-1 || head.y <= 0 || head.y >= height-1 {
				gameOver = true
			}

			draw(snake, food, score)
		}
	}

	termbox.Close()
	fmt.Println("Game Over! Your final score is:", score)
}