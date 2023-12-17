package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

var (
	nextFigure     = Figure{}
	figureInPocket = Figure{}.name

	//time_start_game = time.Now()
)

// DisplayGameOver Запускается когда игра закончилась (выводим GAME OVER)
func DisplayGameOver() {
	DisplayText("GAME OVER", widthArea/3, heightWind/2, color.RGBA{R: 255, G: 255, B: 255, A: 255})
}

// DrawMenu Выводим всю информацию для меню
func DrawMenu() {
	// Заливка
	ebitenutil.DrawRect(screen, widthArea+1, 0.0, widthMenu, heightWind, colorBackground)

	// Выводим следущую фигуру
	DisplayText("Next figure:", widthArea+cellSize, cellSize, colorText)
	// Рисуем следущую фигуру
	for _, nextFigureCell := range nextFigure.rotates[0].form {
		// Сдвигаем клетки фигуры (не саму фигуру) в окошко для меню
		cellInMenu := nextFigureCell
		cellInMenu.x += widthArea + widthMenu/4
		cellInMenu.y += cellSize * 2
		DrawSquare(cellInMenu)
	}

	// Рисуем фигуру в кармашке
	DisplayText("Figure in pocket:", widthArea+cellSize/3, cellSize*8, colorText)
	// Ищем фигуру из кармашка в списке фигур
	for _, findFigure := range listOfFigures {
		if findFigure.name == figureInPocket {
			// Рисуем фигуру из кармашка
			for _, nextFigureCell := range findFigure.rotates[0].form {
				// Сдвигаем клетки фигуры (не саму фигуру) в окошко для меню
				cellInMenu := nextFigureCell
				cellInMenu.x += widthArea + widthMenu/4
				cellInMenu.y += cellSize * 9
				DrawSquare(cellInMenu)
			}
			break
		}
	}

	// Выводим количество собранных рядов
	DisplayText(fmt.Sprintf("Score: %d", gameScore), widthArea+cellSize, heightWind-2*cellSize, colorText)

	//// Выводим время
	//time_in_game := time.Now().Sub(time_start_game)
	//Display_text(fmt.Sprintf("Time: %2.0d:%2.0d", int(time_in_game.Seconds())/60, int(time_in_game.Seconds())%60),
	//	width_area+cell_size, height_wind-cell_size, color_text)
}
