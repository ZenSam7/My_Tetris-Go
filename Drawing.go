package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
)

var (
	text_size = 30
	text_dpi  = 80 // Типа размер текста (что это?)

	// Цвета
	color_background = color.RGBA{30, 40, 50, 255}
	color_shadow     = color.RGBA{40, 50, 60, 255}
	color_grid       = color.RGBA{120, 120, 130, 255}
	color_text       = color.RGBA{150, 170, 170, 255}

	color_red     = color.RGBA{150, 50, 60, 255}
	color_blue    = color.RGBA{50, 80, 120, 255}
	color_yellow  = color.RGBA{130, 130, 50, 255}
	color_green   = color.RGBA{40, 110, 80, 255}
	color_violet  = color.RGBA{110, 40, 120, 255}
	color_goluboy = color.RGBA{60, 130, 150, 255}
	color_gray    = color.RGBA{80, 80, 100, 255}
)

// Рисуем квадрат (клетку)
func Draw_square(screen *ebiten.Image, cell Cell) {
	for x := cell.x; x < cell.x+cell_size; x++ {
		for y := cell.y; y < cell.y+cell_size; y++ {
			screen.Set(x, y, cell.color)
		}
	}
}

// Рисуем вертикальную тень
func Draw_shadow_colomn(screen *ebiten.Image, cell Cell) {
	shadow_cell := Cell{cell.x, cell.y, color_shadow}

	// "Спускаем" клетку тени
	for shadow_cell.y < height_wind {
		shadow_cell.y += cell_size
		Draw_square(screen, shadow_cell)
	}
}

// Рисуем сетку
func Draw_grid(screen *ebiten.Image) {
	// Вертикальные полосы
	for x := -1; x < width_area; x += cell_size {
		for y := -1; y < height_wind; y++ {
			screen.Set(x, y, color_grid)
		}
	}

	// Горизонтальные полосы
	for x := -1; x < width_area; x++ {
		for y := -1; y < height_wind-1; y += cell_size {
			screen.Set(x, y, color_grid)
		}
	}
}

// Выводим текст (точка находится слева снизу, а не слева сверху)
func Display_text(screen *ebiten.Image, TEXT string, x, y int) {
	tt, _ := opentype.Parse(fonts.MPlus1pRegular_ttf)
	mplusNormalFont, _ := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(text_size),
		DPI:     float64(text_dpi),
		Hinting: font.HintingVertical,
	})

	text.Draw(screen, TEXT, mplusNormalFont, x, y, color_text)
}

// Рисуем фигуру полностью (сразу вместе с тенью)
func Draw_figure(screen *ebiten.Image, figure Figure) {
	// Сначала тень
	for _, cell := range figure.form {
		Draw_shadow_colomn(screen, cell)
	}
	// Потом клетки
	for _, cell := range figure.form {
		Draw_square(screen, cell)
	}
}

// Выводим всё на экран
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color_background)

	Draw_figure(screen, figure_now)

	// Рисуем упавшие клетки
	for _, fallen_cell := range fallen_cells {
		Draw_square(screen, fallen_cell)
	}

	Display_text(screen, "Hello, World♥", width_area, 2*cell_size)
	Draw_grid(screen)

	// Когда игра закончилась
	if GAME_OVER {
		Display_text(screen, "GAME OVER", width_area/3, height_wind/2)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width_area + width_menu, height_wind
}
