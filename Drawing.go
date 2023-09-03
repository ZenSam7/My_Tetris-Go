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
	screen         *ebiten.Image
	change_in_area = true
	change_in_menu = true

	// Цвета
	color_background = color.RGBA{R: 30, G: 40, B: 50, A: 255}
	color_shadow     = color.RGBA{R: 40, G: 50, B: 60, A: 255}
	color_grid       = color.RGBA{R: 120, G: 120, B: 130, A: 255}
	color_text       = color.RGBA{R: 150, G: 170, B: 170, A: 255}

	color_red     = color.RGBA{R: 150, G: 50, B: 60, A: 255}
	color_blue    = color.RGBA{R: 50, G: 80, B: 120, A: 255}
	color_yellow  = color.RGBA{R: 130, G: 130, B: 50, A: 255}
	color_green   = color.RGBA{R: 40, G: 110, B: 80, A: 255}
	color_violet  = color.RGBA{R: 110, G: 40, B: 120, A: 255}
	color_goluboy = color.RGBA{R: 60, G: 130, B: 150, A: 255}
	color_gray    = color.RGBA{R: 80, G: 80, B: 100, A: 255}
)

// Рисуем квадрат (клетку)
func Draw_square(cell Cell) {
	ebitenutil.DrawRect(screen, float64(cell.x), float64(cell.y), cell_size, cell_size, cell.color)
}

// Рисуем вертикальную тень
func Draw_shadow_colomn(cell Cell) {
	// Вытягиваем клетку тени на весь экран
	ebitenutil.DrawRect(screen, float64(cell.x), float64(cell.y), cell_size, height_wind, color_shadow)
}

// Рисуем сетку
func Draw_grid() {
	// Вертикальные полосы
	for x := 0.0; x <= width_area; x += cell_size {
		ebitenutil.DrawLine(screen, x, 0.0, x, height_wind, color_grid)
	}

	// Горизонтальные полосы
	for y := 0.0; y < height_wind-1; y += cell_size {
		ebitenutil.DrawLine(screen, 0.0, y, width_area, y, color_grid)
	}
}

// Выводим текст (точка находится слева снизу, а не слева сверху)
func Display_text(TEXT string, x, y int, color color.RGBA) {
	tt, _ := opentype.Parse(fonts.MPlus1pRegular_ttf)
	mplusNormalFont, _ := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(text_size),
		DPI:     float64(text_dpi),
		Hinting: font.HintingVertical,
	})

	text.Draw(screen, TEXT, mplusNormalFont, x, y, color)
}

// Рисуем фигуру полностью (сразу вместе с тенью)
func Draw_figure(figure Figure) {
	// Сначала тень
	for _, cell := range figure.rotates[index_rotate].form {
		Draw_shadow_colomn(cell)
	}

	// Потом клетки
	for _, cell := range figure.rotates[index_rotate].form {
		Draw_square(cell)
	}
}

// Выводим всё на экран
func (g *Game) Draw(display *ebiten.Image) {
	screen = display

	// Рисуем весь экран не 60 раз в секуну, а только когда происходят изменения
	// Плчему бы не вызывать функции сразу, когда что-то меняется?
	// Потому что библиотека от этого ломается

	if change_in_area {
		change_in_area = false
		Draw_game_area()
	}
	if change_in_menu {
		change_in_menu = false
		Draw_menu()
	}
}

// Рисуем весь экран не 60 раз в секуну, а только когда происходят изменения
func Draw_game_area() {
	// Заливка
	ebitenutil.DrawRect(screen, 0.0, 0.0, width_area, height_wind, color_background)

	Draw_figure(figure_now)

	// Рисуем упавшие клетки
	for _, fallen_cell := range fallen_cells {
		Draw_square(fallen_cell)
	}

	Draw_grid()

	// Когда игра закончилась
	if game_over {
		Display_game_over()
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width_area + width_menu, height_wind
}
