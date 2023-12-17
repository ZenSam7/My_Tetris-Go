package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
	"sync"
)

// Настраиваемые параметры

const (
	cellSize = 50 // Ширина одной клетки

	widthArea  = cellSize * 9  // Ширина поля для фигур
	widthMenu  = cellSize * 5  // Ширина поля для дополнительй информации
	heightWind = cellSize * 16 // Высота экрана

	textSize = 40 // Разрмер отображаемого текста
	textDpi  = 50 // Типа тоже размер текста (что это?)

	constGameSpeed = 500  // корость игры (миллисекунд на 1 автоматическое опускание)
	speedFactor    = 0.99 // ВО сколько раз уменьшаем скорость

	speedMoveFigure  = 120 // Скорость движения фигуры по нажатию на кнопки (в миллисекундах)
	timeKeydownSpace = 240 // Отдельно для space (из-за его применения)
	timeRotate       = 230 // Отдельно для поворота (я так хочу)
)

var (
	// Начальная скорость игры (каждые 500 миллисекунд спускаем фигуру вниз)
	gameSpeed = constGameSpeed

	gameScore = 0 // Сколько рядов собрали

	gameOver = false
)

// Добавляем асинхронность (это же Go)
var wg = sync.WaitGroup{}

// Game Надо для библиотеки
type Game struct{}

// Update Логика игры
func (g *Game) Update() error {
	if !gameOver {
		Control()
	}

	ExitAndRestart()

	return nil
}

// Запускаем
func main() {
	nextFigure = DeepCopyFigure(listOfFigures[rand.Intn(7)])
	RandomFigureNow()

	game := &Game{}

	ebiten.SetWindowSize(widthArea+widthMenu, heightWind)
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
