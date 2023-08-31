package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"sync"
)

// Настраиваемые параметры

const (
	cell_size = 50 // Ширина одной клетки

	width_area  = cell_size * 16 // Ширина поля для фигур
	width_menu  = cell_size * 5  // Ширина поля для дополнительй информации
	height_wind = cell_size * 16 // Высота экрана

	speed_factor = 0.98 // ВО сколько раз уменьшаем скорость

	speed_move_figure  = 120 // Скорость движения фигуры по нажатию на кнопки (в миллисекундах)
	time_keydown_space = 200 // Отдельно для space (из-за его применения)
	time_rotate        = 200 // Отдельно для поворота (я так хочу)
)

var (
	// Начальная скорость игры (каждые 500 миллисекунд спускаем фигуру вниз)
	game_speed = 500

	game_score = 0 // Сколько рядов собрали

	GAME_OVER = false
)

// Добавляем асинхронность (это же Go)
var wg = sync.WaitGroup{}

// Надо для библиотеки
type Game struct{}

// Логика игрыыввфывц
func (g *Game) Update() error {
	if !GAME_OVER {
		Control_figure()
	}

	Exit_and_restart()

	return nil
}

// Запускаем
func main() {
	Random_figure_now()

	game := &Game{}

	ebiten.SetWindowSize(width_area+width_menu, height_wind)
	ebiten.SetWindowTitle("(*^ω^) Tetris on Golang")

	ebiten.RunGame(game)
}
