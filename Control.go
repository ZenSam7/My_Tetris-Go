package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
	"os"
	"time"
)

var (
	// Разные таймеры
	timerForFalling  = time.Now().UnixMilli()
	timerForMove     = time.Now().UnixMilli()
	timerForKeyspace = time.Now().UnixMilli()
	timerForRotate   = time.Now().UnixMilli()
	timerForRestart  = time.Now().UnixMilli()

	// Можно ли использовать кармашек
	canUsePocket = true
)

// RandomFigureNow Создаём новую фигуру для управления
func RandomFigureNow() {

	// Изменение в следующей и текущей фигуре, а значить перерисовываем всё
	changeInArea = true
	changeInMenu = true

	indexRotate = 0

	canUsePocket = true

	// Создаём новую следущую фигуру
	// При условии, что она не равна предыдущей фигуре
	// (я это сделал чтобы небыло последовательностей одинаковых фигур, которые сильно усложняют)
	figureNow = nextFigure
	for nextFigure.name == figureNow.name {
		nextFigure = DeepCopyFigure(listOfFigures[rand.Intn(7)])
	}

	// Сдвигаем фугуру к центру
	for i := 0; i < (widthArea/cellSize)/2-1; i++ {
		figureNow.MoveRight()
	}

	// Если новая фигура появляется в уже упавшей, то мы проиграли
	if RotateInFallenCells(figureNow.rotates[0]) {
		GameOver()
	}
}

// Control Всячески двигаем фигуру
func Control() {
	timeNow := time.Now().UnixMilli()

	// Постепенно спускаем фигуру
	if timeNow >= timerForFalling+int64(gameSpeed) {
		figureNow.MoveDown()
		timerForFalling = timeNow
	}

	// Поворачиваем фигуру
	if (ebiten.IsKeyPressed(ebiten.KeyW) ||
		ebiten.IsKeyPressed(ebiten.KeyArrowUp)) &&
		timeNow >= timerForRotate+timeRotate {
		figureNow.Rotate()
		timerForRotate = timeNow
	}

	// Двигаем влево вправо и вниз по нажатию кнопки
	canMove := timeNow >= timerForMove+speedMoveFigure
	if canMove {
		if ebiten.IsKeyPressed(ebiten.KeyA) ||
			ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
			figureNow.MoveLeft()
			timerForMove = timeNow
		} else if ebiten.IsKeyPressed(ebiten.KeyD) ||
			ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
			figureNow.MoveRight()
			timerForMove = timeNow
		}
		if ebiten.IsKeyPressed(ebiten.KeyS) ||
			ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
			figureNow.MoveDown()
			timerForMove = timeNow
		}
	}

	// Когда нажимаем пробел, двигаем вниз до упора
	if ebiten.IsKeyPressed(ebiten.KeySpace) && timeNow >= timerForKeyspace+timeKeydownSpace {
		wasCollision := false
		for !wasCollision {
			wasCollision = figureNow.MoveDown()
		}
		timerForKeyspace = timeNow
	}

	// Используем кармашек
	if (ebiten.IsKeyPressed(ebiten.KeyShift) ||
		ebiten.IsKeyPressed(ebiten.KeyShiftRight) ||
		ebiten.IsKeyPressed(ebiten.KeyShiftLeft)) &&
		timeNow >= timerForRotate+timeKeydownSpace &&
		canUsePocket {

		changeInMenu = true
		timerForRotate = timeNow
		indexRotate = 0

		// Изменение в фигуре из кармашка, а значить перерисовываем меню
		DrawMenu()

		// Нельзя использовать кармашек второй раз, пока фигура не упала
		canUsePocket = false

		// Если фигуры нет, то добавляем
		if figureInPocket == "" {
			figureInPocket = figureNow.name
			RandomFigureNow()

		} else { // Если фигура есть, то меняем
			// Ищем фигуру из кармашка в списке фигур
			for _, findFigure := range listOfFigures {
				if findFigure.name == figureInPocket {
					// Меняем местами
					figureInPocket = figureNow.name
					figureNow = DeepCopyFigure(findFigure)

					// Сдвигаем фугуру к центру
					for i := 0; i < (widthArea/cellSize)/2-1; i++ {
						figureNow.MoveRight()
					}

					break
				}
			}
		}

		timerForRestart = timeNow
	}

}

// ExitAndRestart Надо отделить от контроля фигуры (см. как работает game_over)
func ExitAndRestart() {
	// Закрываем прогу
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	// Рестартим
	if ebiten.IsKeyPressed(ebiten.KeyR) && (time.Now().UnixMilli() >= timerForRestart+1000) {
		fallenCells = []Cell{}
		gameOver = false
		gameScore = 0
		gameSpeed = constGameSpeed
		figureInPocket = ""
		RandomFigureNow()

		timerForRestart = time.Now().UnixMilli()
	}
}

// Collision Оформляем столкновения
func Collision() {
	// Записываем фигуру как упавшую
	for _, cell := range figureNow.rotates[indexRotate].form {
		fallenCells = append(fallenCells, cell)
	}

	// Собралась ли полная линия
	for numRow := 0; numRow < heightWind; numRow += cellSize {
		numCellsInRow := 0 // Количество клеток в ряду num_row

		for _, fallenCell := range fallenCells {
			if fallenCell.y == numRow {
				numCellsInRow += 1
			}
		}

		// Если собрался полный ряд
		if numCellsInRow == widthArea/cellSize {
			CollectingRow(numRow)
		}
	}

	RandomFigureNow()
}

// GameOver Когда закончили игру
func GameOver() {
	gameOver = true
}

// CollectingRow Что делаем когда ряд собран
func CollectingRow(numRow int) {
	// Ускоряем игру
	gameSpeed = int(float32(gameSpeed) * speedFactor)

	// Увеличиваем счётчик
	gameScore += 1

	var indCellsInRow []int // Индексы всех клеток в полном ряду
	// Добавляем индексы всех клеток на удаление
	for ind, fallenCell := range fallenCells {
		if fallenCell.y == numRow {
			indCellsInRow = append(indCellsInRow, ind)
		}
	}

	// Удаляем
	for bias, ind := range indCellsInRow {
		fallenCells = append(fallenCells[:ind-bias], fallenCells[ind-bias+1:]...)
	}

	// Сдвигаем все клетки сверху
	for i, fallenCell := range fallenCells {
		if fallenCell.y <= numRow {
			newCell := fallenCell
			newCell.y += cellSize
			fallenCells[i] = newCell
		}
	}
}

// DeepCopyFigure Создаём глубокую копию фигуры (фактически новую фигуру)
func DeepCopyFigure(figure Figure) Figure {
	newFigure := Figure{}

	// Делаем копию каждого поворота
	for _, rotate := range figure.rotates {
		newFigureForm := FigureForm{}

		// Делаем копию каждого Figure_form
		for i, cell := range rotate.form {
			newCell := Cell{cell.x, cell.y, cell.color}
			newFigureForm.form[i] = newCell
		}

		// Добавляем поворот
		newFigure.rotates = append(newFigure.rotates, newFigureForm)
	}
	newFigure.name = figure.name

	return newFigure
}
