/// Golang Snake Game && NoTalk
/// Below is a simple example of snake game code in Go using the github.com/nsf/termbox-go
/// package to handle terminal-based graphics. This program creates a simple game where
/// the snake moves based on keyboard arrow keys, and you can quit the game by pressing q.

package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

type Point struct {
	x, y int
}

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	// Initial snake and food
	snake := []Point{{4, 4}, {4, 3}, {4, 2}}
	food := Point{7, 4}

	// Draw initial state
	draw(snake, food)

	// Snake direction
	dx, dy := 1, 0

	// Event loop
	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

loop:
	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch ev.Key {
				case termbox.KeyArrowUp:
					dx, dy = 0, -1
				case termbox.KeyArrowDown:
					dx, dy = 0, 1
				case termbox.KeyArrowLeft:
					dx, dy = -1, 0
				case termbox.KeyArrowRight:
					dx, dy = 1, 0
				case termbox.KeyEsc, termbox.KeyCtrlC, termbox.KeyCtrlD:
					break loop
				}
			}
		default:
			head := snake[0]
			newHead := Point{head.x + dx, head.y + dy}

			// Eat food
			if newHead == food {
				snake = append([]Point{newHead}, snake...)
				food = Point{newHead.x + 2, newHead.y + 2}
			} else {
				snake = append([]Point{newHead}, snake[:len(snake)-1]...)
			}

			draw(snake, food)
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func draw(snake []Point, food Point) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// Draw snake
	for _, point := range snake {
		termbox.SetCell(point.x, point.y, 'O', termbox.ColorGreen, termbox.ColorDefault)
	}

	// Draw food
	termbox.SetCell(food.x, food.y, 'X', termbox.ColorRed, termbox.ColorDefault)

	termbox.Flush()
}
