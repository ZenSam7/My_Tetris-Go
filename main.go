package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
	"sync"
)

// Настраиваемые параметры

const (
	cell_size = 50 // Ширина одной клетки

	width_area  = cell_size * 9  // Ширина поля для фигур
	width_menu  = cell_size * 5  // Ширина поля для дополнительй информации
	height_wind = cell_size * 16 // Высота экрана

	text_size = 40 // Разрмер отображаемого текста
	text_dpi  = 50 // Типа тоже размер текста (что это?)

	const_game_speed = 500  // корость игры (миллисекунд на 1 автоматическое опускание)
	speed_factor     = 0.99 // ВО сколько раз уменьшаем скорость

	speed_move_figure  = 120 // Скорость движения фигуры по нажатию на кнопки (в миллисекундах)
	time_keydown_space = 240 // Отдельно для space (из-за его применения)
	time_rotate        = 230 // Отдельно для поворота (я так хочу)
)

var (
	// Начальная скорость игры (каждые 500 миллисекунд спускаем фигуру вниз)
	game_speed = const_game_speed

	game_score = 0 // Сколько рядов собрали

	game_over = false
)

// Добавляем асинхронность (это же Go)
var wg = sync.WaitGroup{}

// Надо для библиотеки
type Game struct{}

// Логика игры
func (g *Game) Update() error {
	if !game_over {
		Control()
	}

	Exit_and_restart()

	return nil
}

// Запускаем
func main() {
	next_figure = Deep_copy_figure(list_of_figures[rand.Intn(7)])
	Random_figure_now()

	game := &Game{}

	ebiten.SetWindowSize(width_area+width_menu, height_wind)
	ebiten.SetWindowTitle("(*^ω^) Tetris on Golang")

	// Т.к. я хочу всё оптимизировать, то стирать, а затем рисовать такие-же кадры
	// каждый фрейм я не собираюсь. А рисовать новые кадры я буду только
	// когда происходят видимые изменения
	ebiten.SetScreenClearedEveryFrame(false)

	err := ebiten.RunGame(game)
	if err != nil {
		return
	}
}
