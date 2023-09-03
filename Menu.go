package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

var (
	next_figure      = Figure{}
	figure_in_pocket = Figure{}.name

	//time_start_game = time.Now()
)

// Запускается когда игра закончилась (выводим GAME OVER)
func Display_game_over() {
	Display_text("GAME OVER", width_area/3, height_wind/2, color.RGBA{R: 255, G: 255, B: 255, A: 255})
}

// Выводим всю информацию для меню
func Draw_menu() {
	// Заливка
	ebitenutil.DrawRect(screen, width_area+1, 0.0, width_menu, height_wind, color_background)

	// Выводим следущую фигуру
	Display_text("Next figure:", width_area+cell_size, cell_size, color_text)
	// Рисуем следущую фигуру
	for _, next_figure_cell := range next_figure.rotates[0].form {
		// Сдвигаем клетки фигуры (не саму фигуру) в окошко для меню
		cell_in_menu := next_figure_cell
		cell_in_menu.x += width_area + width_menu/4
		cell_in_menu.y += cell_size * 2
		Draw_square(cell_in_menu)
	}

	// Рисуем фигуру в кармашке
	Display_text("Figure in pocket:", width_area+cell_size/3, cell_size*8, color_text)
	// Ищем фигуру из кармашка в списке фигур
	for _, find_figure := range list_of_figures {
		if find_figure.name == figure_in_pocket {
			// Рисуем фигуру из кармашка
			for _, next_figure_cell := range find_figure.rotates[0].form {
				// Сдвигаем клетки фигуры (не саму фигуру) в окошко для меню
				cell_in_menu := next_figure_cell
				cell_in_menu.x += width_area + width_menu/4
				cell_in_menu.y += cell_size * 9
				Draw_square(cell_in_menu)
			}
			break
		}
	}

	// Выводим количество собранных рядов
	Display_text(fmt.Sprintf("Score: %d", game_score), width_area+cell_size, height_wind-2*cell_size, color_text)

	//// Выводим время
	//time_in_game := time.Now().Sub(time_start_game)
	//Display_text(fmt.Sprintf("Time: %2.0d:%2.0d", int(time_in_game.Seconds())/60, int(time_in_game.Seconds())%60),
	//	width_area+cell_size, height_wind-cell_size, color_text)
}
