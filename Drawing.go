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
	screen *ebiten.Image

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
func Display_text(TEXT string, x, y int) {
	tt, _ := opentype.Parse(fonts.MPlus1pRegular_ttf)
	mplusNormalFont, _ := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(text_size),
		DPI:     float64(text_dpi),
		Hinting: font.HintingVertical,
	})

	text.Draw(screen, TEXT, mplusNormalFont, x, y, color_text)
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

	screen.Fill(color_background)
	Draw_figure(figure_now)

	// Рисуем упавшие клетки
	for _, fallen_cell := range fallen_cells {
		Draw_square(fallen_cell)
	}

	Display_text("Hello, World♥", width_area, 2*cell_size)
	Draw_grid()

	// Когда игра закончилась
	if GAME_OVER {
		Display_text("GAME OVER", width_area/3, height_wind/2)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width_area + width_menu, height_wind
}
