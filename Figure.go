package main

import (
	"image/color"
)

var (
	// Объявляем фигуры
	c             = cellSize // Для краткости
	listOfFigures = [7]Figure{
		{name: "O",
			rotates: []FigureForm{
				{[4]Cell{{0, 0, colorYellow}, {c, 0, colorYellow}, {0, c, colorYellow}, {c, c, colorYellow}}}, // Поворот 0
			}},

		{name: "L",
			rotates: []FigureForm{
				{[4]Cell{{0, 0, colorGray}, {0, c, colorGray}, {0, 2 * c, colorGray}, {c, 2 * c, colorGray}}}, // Поворот 0
				{[4]Cell{{0, 0, colorGray}, {c, 0, colorGray}, {2 * c, 0, colorGray}, {0, c, colorGray}}},     // Поворот 1
				{[4]Cell{{0, 0, colorGray}, {c, 0, colorGray}, {c, c, colorGray}, {c, 2 * c, colorGray}}},     // Поворот 2
				{[4]Cell{{0, c, colorGray}, {c, c, colorGray}, {2 * c, c, colorGray}, {2 * c, 0, colorGray}}}, // Поворот 3
			}},

		{name: "J",
			rotates: []FigureForm{
				{[4]Cell{{c, 0, colorBlue}, {c, c, colorBlue}, {c, 2 * c, colorBlue}, {0, 2 * c, colorBlue}}}, // Поворот 0
				{[4]Cell{{0, 0, colorBlue}, {0, c, colorBlue}, {c, c, colorBlue}, {2 * c, c, colorBlue}}},     // Поворот 1
				{[4]Cell{{0, 0, colorBlue}, {c, 0, colorBlue}, {0, c, colorBlue}, {0, 2 * c, colorBlue}}},     // Поворот 2
				{[4]Cell{{0, 0, colorBlue}, {c, 0, colorBlue}, {2 * c, 0, colorBlue}, {2 * c, c, colorBlue}}}, // Поворот 3
			}},

		{name: "I",
			rotates: []FigureForm{
				{[4]Cell{{0, 0, colorGoluboy}, {0, c, colorGoluboy}, {0, 2 * c, colorGoluboy}, {0, 3 * c, colorGoluboy}}}, // Поворот 0
				{[4]Cell{{0, 0, colorGoluboy}, {c, 0, colorGoluboy}, {2 * c, 0, colorGoluboy}, {3 * c, 0, colorGoluboy}}}, // Поворот 1
			}},

		{name: "T",
			rotates: []FigureForm{
				{[4]Cell{{0, 0, colorViolet}, {c, 0, colorViolet}, {2 * c, 0, colorViolet}, {c, c, colorViolet}}}, // Поворот 0
				{[4]Cell{{c, 0, colorViolet}, {c, c, colorViolet}, {c, 2 * c, colorViolet}, {0, c, colorViolet}}}, // Поворот 1
				{[4]Cell{{c, 0, colorViolet}, {0, c, colorViolet}, {c, c, colorViolet}, {2 * c, c, colorViolet}}}, // Поворот 2
				{[4]Cell{{0, 0, colorViolet}, {0, c, colorViolet}, {0, 2 * c, colorViolet}, {c, c, colorViolet}}}, // Поворот 3
			}},

		{name: "Z",
			rotates: []FigureForm{
				{[4]Cell{{0, 0, colorRed}, {c, 0, colorRed}, {c, c, colorRed}, {2 * c, c, colorRed}}}, // Поворот 0
				{[4]Cell{{c, 0, colorRed}, {c, c, colorRed}, {0, c, colorRed}, {0, 2 * c, colorRed}}}, // Поворот 1
			}},

		{name: "S",
			rotates: []FigureForm{
				{[4]Cell{{0, c, colorGreen}, {c, c, colorGreen}, {c, 0, colorGreen}, {2 * c, 0, colorGreen}}}, // Поворот 0
				{[4]Cell{{0, 0, colorGreen}, {0, c, colorGreen}, {c, c, colorGreen}, {c, 2 * c, colorGreen}}}, // Поворот 1
			}},
	}

	figureNow = Figure{}

	// Какой сейчас поворот у фигуры
	indexRotate = 0

	// Записываем все упавшие клетки
	fallenCells []Cell
)

type Cell struct {
	x, y  int
	color color.RGBA
}

type FigureForm struct {
	form [4]Cell
}

type Figure struct {
	name    string
	rotates []FigureForm
}

// MoveDown Сдвигаем вниз
func (f *Figure) MoveDown() (wasCollision bool) {
	// Изменение текущей фигуры, перерисовываем рабочий экран
	changeInArea = true

	wasCollision = false
	// Сдвигаем все повороты вниз
	// Все, кроме текущего поворота, делаем невидимым для упавших клеток
	for i, rotate := range f.rotates {
		ghost := true
		if i == indexRotate {
			ghost = false
		}

		if MoveDownRotate(&rotate) && !ghost {
			wasCollision = true
		}

		f.rotates[i] = rotate
	}

	if wasCollision {
		Collision()
	}

	return wasCollision
}

// MoveDownRotate Сдвигаем один из поворотов фигуры
func MoveDownRotate(rotate *FigureForm) (wasCollision bool) {
	// Если какая-то клетка находится на полу, то у нас коллизия
	for _, figureCell := range rotate.form {
		if figureCell.y == heightWind-cellSize {
			return true
		}
	}

	// Опускаем
	for i, figureCell := range rotate.form {
		movedCell := figureCell
		movedCell.y += cellSize
		rotate.form[i] = movedCell
	}

	// Если обнаруживается упавшая клетка внутри (уже опущенной) фигуры,
	// то поднимаем фигуру назад
	if RotateInFallenCells(*rotate) {
		for i, figureCell := range rotate.form {
			movedCell := figureCell
			movedCell.y -= cellSize
			rotate.form[i] = movedCell
		}
		return true
	}

	return false
}

// MoveLeft Сдвигаем влево
func (f *Figure) MoveLeft() {
	// Изменение текущей фигуры, перерисовываем рабочий экран
	changeInArea = true

	// Сдвигаем каждый поворот влево
	for i, rotate := range f.rotates {
		possibleMoveLeft := true

		// Если какая-то клетка находится около стены, то не сдвигаем
		for _, figureCell := range rotate.form {
			if figureCell.x == 0 {
				possibleMoveLeft = false
			}
		}

		if possibleMoveLeft {
			// Сдвигаем все клетки
			for i, figureCell := range rotate.form {
				movedCell := figureCell
				movedCell.x -= cellSize
				rotate.form[i] = movedCell
			}

			// Если обнаруживается упавшая клетка внутри (уже сдвинутой) фигуры,
			// то возвращаем фигуру назад
			if RotateInFallenCells(rotate) {
				// Сдвигаем все клетки
				for i, figureCell := range rotate.form {
					movedCell := figureCell
					movedCell.x += cellSize
					rotate.form[i] = movedCell
				}
			}

			f.rotates[i] = rotate
		}
	}
}

// MoveRight Сдвигаем вправо
func (f *Figure) MoveRight() {
	// Изменение текущей фигуры, перерисовываем рабочий экран
	changeInArea = true

	// Сдвигаем все повороты вправо
	for i, rotate := range f.rotates {
		possibleMoveRight := true

		// Если какая-то клетка находится около стены, то не сдвигаем
		for _, figureCell := range rotate.form {
			if figureCell.x == widthArea-cellSize {
				possibleMoveRight = false
			}
		}

		if possibleMoveRight {
			// Сдвигаем все клетки
			for i, figureCell := range rotate.form {
				movedCell := figureCell
				movedCell.x += cellSize
				rotate.form[i] = movedCell
			}

			// Если обнаруживается упавшая клетка внутри (уже сдвинутой) фигуры,
			// то возвращаем фигуру назад
			if RotateInFallenCells(rotate) {
				// Сдвигаем все клетки
				for i, figureCell := range rotate.form {
					movedCell := figureCell
					movedCell.x -= cellSize
					rotate.form[i] = movedCell
				}
			}

			f.rotates[i] = rotate
		}
	}
}

// RotateInFallenCells Проверяем, находится ли фигура в упавших клетках
func RotateInFallenCells(rotate FigureForm) bool {
	rotateInCells := false
	for _, figureCell := range rotate.form {
		for _, fallenCell := range fallenCells {
			wg.Add(1)
			go func(figureCell, fallenCell Cell) {
				if figureCell.x == fallenCell.x &&
					figureCell.y == fallenCell.y {
					rotateInCells = true
				}
				wg.Done()
			}(figureCell, fallenCell)
		}
	}
	wg.Wait()

	return rotateInCells
}

// Rotate Поворачиваем фигуру
func (f *Figure) Rotate() {
	// Изменение текущей фигуры, перерисовываем рабочий экран
	changeInArea = true

	rotatedFigure := figureNow.rotates[(indexRotate+1)%len(figureNow.rotates)]

	// Поворачиваем, если не задевает упавшие клетки
	if !RotateInFallenCells(rotatedFigure) {
		indexRotate = (indexRotate + 1) % len(figureNow.rotates)
	}
}
