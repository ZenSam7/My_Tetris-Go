package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
)

var (
	screen       *ebiten.Image
	changeInArea = true
	changeInMenu = true

	// Цвета
	colorBackground = color.RGBA{R: 30, G: 40, B: 50, A: 255}
	colorShadow     = color.RGBA{R: 40, G: 50, B: 60, A: 255}
	colorGrid       = color.RGBA{R: 120, G: 120, B: 130, A: 255}
	colorText       = color.RGBA{R: 150, G: 170, B: 170, A: 255}

	colorRed     = color.RGBA{R: 150, G: 50, B: 60, A: 255}
	colorBlue    = color.RGBA{R: 50, G: 80, B: 120, A: 255}
	colorYellow  = color.RGBA{R: 130, G: 130, B: 50, A: 255}
	colorGreen   = color.RGBA{R: 40, G: 110, B: 80, A: 255}
	colorViolet  = color.RGBA{R: 110, G: 40, B: 120, A: 255}
	colorGoluboy = color.RGBA{R: 60, G: 130, B: 150, A: 255}
	colorGray    = color.RGBA{R: 80, G: 80, B: 100, A: 255}
)

// DrawSquare Рисуем квадрат (клетку)
func DrawSquare(cell Cell) {
	ebitenutil.DrawRect(screen, float64(cell.x), float64(cell.y), cellSize, cellSize, cell.color)
}

// DrawShadowColomn Рисуем вертикальную тень
func DrawShadowColomn(cell Cell) {
	// Вытягиваем клетку тени на весь экран
	ebitenutil.DrawRect(screen, float64(cell.x), float64(cell.y), cellSize, heightWind, colorShadow)
}

// DrawGrid Рисуем сетку
func DrawGrid() {
	// Вертикальные полосы
	for x := 0.0; x <= widthArea; x += cellSize {
		ebitenutil.DrawLine(screen, x, 0.0, x, heightWind, colorGrid)
	}

	// Горизонтальные полосы
	for y := 0.0; y < heightWind-1; y += cellSize {
		ebitenutil.DrawLine(screen, 0.0, y, widthArea, y, colorGrid)
	}
}

// DisplayText Выводим текст (точка находится слева снизу, а не слева сверху)
func DisplayText(TEXT string, x, y int, color color.RGBA) {
	tt, _ := opentype.Parse(fonts.MPlus1pRegular_ttf)
	mplusNormalFont, _ := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(textSize),
		DPI:     float64(textDpi),
		Hinting: font.HintingVertical,
	})

	text.Draw(screen, TEXT, mplusNormalFont, x, y, color)
}

// DrawFigure Рисуем фигуру полностью (сразу вместе с тенью)
func DrawFigure(figure Figure) {
	// Сначала тень
	for _, cell := range figure.rotates[indexRotate].form {
		DrawShadowColomn(cell)
	}

	// Потом клетки
	for _, cell := range figure.rotates[indexRotate].form {
		DrawSquare(cell)
	}
}

// Draw Выводим всё на экран
func (g *Game) Draw(display *ebiten.Image) {
	screen = display

	// Рисуем весь экран не 60 раз в секуну, а только когда происходят изменения
	// Плчему бы не вызывать функции сразу, когда что-то меняется?
	// Потому что библиотека от этого ломается

	if changeInArea {
		changeInArea = false
		DrawGameArea()
	}
	if changeInMenu {
		changeInMenu = false
		DrawMenu()
	}
}

// DrawGameArea Рисуем весь экран не 60 раз в секуну, а только когда происходят изменения
func DrawGameArea() {
	// Заливка
	ebitenutil.DrawRect(screen, 0.0, 0.0, widthArea, heightWind, colorBackground)

	DrawFigure(figureNow)

	// Рисуем упавшие клетки
	for _, fallenCell := range fallenCells {
		DrawSquare(fallenCell)
	}

	DrawGrid()

	// Когда игра закончилась
	if gameOver {
		DisplayGameOver()
	}
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return widthArea + widthMenu, heightWind
}
