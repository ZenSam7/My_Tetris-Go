package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
)

// Настраиваемые параметры

const (
	cell_size = 50 // Ширина одной клетки

	width_area  = cell_size * 10 // Ширина поля для фигур
	width_menu  = cell_size * 5  // Ширина поля для дополнительй информации
	height_wind = cell_size * 16 // Высота экрана

	// Начальная скорость игры (количество секунд на 1 движение вниз)
	game_speed   = 0.1
	speed_factor = 0.001 // На сколько уменьшаем скорость

)

var (
	text_size = 30
	text_dpi  = 80 // Типа размер текста (что это?)

	// Цвета
	color_background = color.RGBA{30, 35, 45, 255}
	color_shadow     = color.RGBA{40, 50, 60, 255}
	color_grid       = color.RGBA{100, 100, 110, 255}
	color_text       = color.RGBA{150, 170, 170, 255}

	color_red     = color.RGBA{150, 50, 70, 255}
	color_blue    = color.RGBA{50, 80, 120, 255}
	color_yellow  = color.RGBA{130, 130, 50, 255}
	color_green   = color.RGBA{40, 130, 90, 255}
	color_violet  = color.RGBA{130, 60, 140, 255}
	color_goluboy = color.RGBA{80, 150, 170, 255}
	color_gray    = color.RGBA{110, 110, 120, 255}
)

// Логика игры

type Cell struct {
	x, y  int
	color color.RGBA
}

type Game struct{}

func (g *Game) Update() error {
	// Write your game's logical update.
	return nil
}

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

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color_background)

	all_colors := []color.RGBA{color_background, color_shadow, color_grid, color_text,
		color_red, color_blue, color_yellow, color_green, color_violet, color_goluboy, color_gray}

	// Рисуем сначала тени, потом клетки, потом сетку
	//Draw_shadow_colomn(screen, square)
	x_color := 0
	for _, color := range all_colors {
		square := Cell{x_color, 0, color}
		Draw_square(screen, square)
		x_color += cell_size
	}

	Draw_grid(screen)

	Display_text(screen, "Hello, World♥", width_area, 2*cell_size)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width_area + width_menu, height_wind
}

// Запускаем

func main() {
	game := &Game{}

	ebiten.SetWindowSize(width_area+width_menu, height_wind)
	ebiten.SetWindowTitle("(*^ω^) Tetris on Golang")

	ebiten.RunGame(game)
}
