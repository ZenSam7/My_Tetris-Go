package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

// Настраиваемые параметры

const (
	width_area = 500 // Ширина поля для фигур
	width_menu = 200 // Ширина поля для дополнительй информации
	height     = 700 // Высота экрана

	cell_size = 20 // Ширина одной клетки

	// Начальная скорость игры (количество секунд на 1 движение вниз)
	game_speed   = 0.1
	speed_factor = 0.001 // На сколько уменьшаем скорость
)

var (
	// Цвета
	color_background = color.RGBA{30, 35, 45, 255}
	color_cell       = color.RGBA{90, 100, 120, 255}
)

// Логика игры

type Game struct{}

func (g *Game) Update() error {
	// Write your game's logical update.
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width_area + width_menu, height
}

// Запускаем

func main() {
	game := &Game{}

	ebiten.SetWindowSize(width_area+width_menu, height)
	ebiten.SetWindowTitle("(*^ω^) Tetris on Golang")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
