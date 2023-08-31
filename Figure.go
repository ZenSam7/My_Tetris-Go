package main

import (
	"image/color"
)

var (
	// Объявляем фигуры
	c              = cell_size // Для краткости
	list_of_figurs = [7]Figure{
		// O
		{[]Figure_form{
			{[4]Cell{{0, 0, color_yellow}, {c, 0, color_yellow}, {0, c, color_yellow}, {c, c, color_yellow}}}, // Поворот 0
		}},

		// L
		{[]Figure_form{
			{[4]Cell{{0, 0, color_gray}, {0, c, color_gray}, {0, 2 * c, color_gray}, {c, 2 * c, color_gray}}}, // Поворот 0
			{[4]Cell{{0, 0, color_gray}, {c, 0, color_gray}, {2 * c, 0, color_gray}, {0, c, color_gray}}},     // Поворот 1
			{[4]Cell{{0, 0, color_gray}, {c, 0, color_gray}, {c, c, color_gray}, {c, 2 * c, color_gray}}},     // Поворот 2
			{[4]Cell{{0, c, color_gray}, {c, c, color_gray}, {2 * c, c, color_gray}, {2 * c, 0, color_gray}}}, // Поворот 3
		}},

		// J
		{[]Figure_form{
			{[4]Cell{{c, 0, color_blue}, {c, c, color_blue}, {c, 2 * c, color_blue}, {0, 2 * c, color_blue}}}, // Поворот 0
			{[4]Cell{{0, 0, color_blue}, {0, c, color_blue}, {c, c, color_blue}, {2 * c, c, color_blue}}},     // Поворот 1
			{[4]Cell{{0, 0, color_blue}, {c, 0, color_blue}, {0, c, color_blue}, {0, 2 * c, color_blue}}},     // Поворот 2
			{[4]Cell{{0, 0, color_blue}, {c, 0, color_blue}, {2 * c, 0, color_blue}, {2 * c, c, color_blue}}}, // Поворот 3
		}},

		// I
		{[]Figure_form{
			{[4]Cell{{0, 0, color_goluboy}, {0, c, color_goluboy}, {0, 2 * c, color_goluboy}, {0, 3 * c, color_goluboy}}}, // Поворот 0
			{[4]Cell{{0, 0, color_goluboy}, {c, 0, color_goluboy}, {2 * c, 0, color_goluboy}, {3 * c, 0, color_goluboy}}}, // Поворот 1
		}},

		// T
		{[]Figure_form{
			{[4]Cell{{0, 0, color_violet}, {c, 0, color_violet}, {2 * c, 0, color_violet}, {c, c, color_violet}}}, // Поворот 0
			{[4]Cell{{c, 0, color_violet}, {c, c, color_violet}, {c, 2 * c, color_violet}, {0, c, color_violet}}}, // Поворот 1
			{[4]Cell{{c, 0, color_violet}, {0, c, color_violet}, {c, c, color_violet}, {2 * c, c, color_violet}}}, // Поворот 2
			{[4]Cell{{0, 0, color_violet}, {0, c, color_violet}, {0, 2 * c, color_violet}, {c, c, color_violet}}}, // Поворот 3
		}},

		// Z
		{[]Figure_form{
			{[4]Cell{{0, 0, color_red}, {c, 0, color_red}, {c, c, color_red}, {2 * c, c, color_red}}}, // Поворот 0
			{[4]Cell{{c, 0, color_red}, {c, c, color_red}, {0, c, color_red}, {0, 2 * c, color_red}}}, // Поворот 1
		}},

		// S
		{[]Figure_form{
			{[4]Cell{{0, c, color_green}, {c, c, color_green}, {c, 0, color_green}, {2 * c, 0, color_green}}}, // Поворот 0
			{[4]Cell{{0, 0, color_green}, {0, c, color_green}, {c, c, color_green}, {c, 2 * c, color_green}}}, // Поворот 1
		}},
	}

	figure_now = Figure{}

	// Какой сейчас поворот у фигуры
	index_rotate = 0

	// Записываем все упавшие клетки
	fallen_cells = []Cell{}
)

type Cell struct {
	x, y  int
	color color.RGBA
}

type Figure_form struct {
	form [4]Cell
}

type Figure struct {
	rotates []Figure_form
}

// Сдвигаем вниз
func (f *Figure) Move_down() (was_collision bool) {
	was_collision = false
	// Сдвигаем все повороты вниз
	// Все, кроме текущего поворота, делаем невидимым для упавших клеток
	for i, rotate := range f.rotates {
		ghost := true
		if i == index_rotate {
			ghost = false
		}

		if Move_down_rotate(&rotate) && !ghost {
			was_collision = true
		}

		f.rotates[i] = rotate
	}

	if was_collision {
		Collision()
	}

	return was_collision
}

// Сдвигаем один из поворотов фигуры
func Move_down_rotate(rotate *Figure_form) (was_collision bool) {
	// Если какая-то клетка находится на полу, то у нас коллизия
	for _, figure_cell := range rotate.form {
		if figure_cell.y == height_wind-cell_size {
			return true
		}
	}

	// Опускаем
	for i, figure_cell := range rotate.form {
		moved_cell := figure_cell
		moved_cell.y += cell_size
		rotate.form[i] = moved_cell
	}

	// Если обнаруживается упавшая клетка внутри (уже опущенной) фигуры,
	// то поднимаем фигуру назад
	if Rotate_in_fallen_cells(*rotate) {
		for i, figure_cell := range rotate.form {
			moved_cell := figure_cell
			moved_cell.y -= cell_size
			rotate.form[i] = moved_cell
		}
		return true
	}

	return false
}

// Сдвигаем влево
func (f *Figure) Move_left() {
	// Сдвигаем каждый поворот влево
	for i, rotate := range f.rotates {
		possible_move_left := true

		// Если какая-то клетка находится около стены, то не сдвигаем
		for _, figure_cell := range rotate.form {
			if figure_cell.x == 0 {
				possible_move_left = false
			}
		}

		if possible_move_left {
			// Сдвигаем все клетки
			for i, figure_cell := range rotate.form {
				moved_cell := figure_cell
				moved_cell.x -= cell_size
				rotate.form[i] = moved_cell
			}

			// Если обнаруживается упавшая клетка внутри (уже сдвинутой) фигуры,
			// то возвращаем фигуру назад
			if Rotate_in_fallen_cells(rotate) {
				// Сдвигаем все клетки
				for i, figure_cell := range rotate.form {
					moved_cell := figure_cell
					moved_cell.x += cell_size
					rotate.form[i] = moved_cell
				}
			}

			f.rotates[i] = rotate
		}
	}
}

// Сдвигаем вправо
func (f *Figure) Move_right() {
	// Сдвигаем все повороты вправо
	for i, rotate := range f.rotates {
		possible_move_right := true

		// Если какая-то клетка находится около стены, то не сдвигаем
		for _, figure_cell := range rotate.form {
			if figure_cell.x == width_area-cell_size {
				possible_move_right = false
			}
		}

		if possible_move_right {
			// Сдвигаем все клетки
			for i, figure_cell := range rotate.form {
				moved_cell := figure_cell
				moved_cell.x += cell_size
				rotate.form[i] = moved_cell
			}

			// Если обнаруживается упавшая клетка внутри (уже сдвинутой) фигуры,
			// то возвращаем фигуру назад
			if Rotate_in_fallen_cells(rotate) {
				// Сдвигаем все клетки
				for i, figure_cell := range rotate.form {
					moved_cell := figure_cell
					moved_cell.x -= cell_size
					rotate.form[i] = moved_cell
				}
			}

			f.rotates[i] = rotate
		}
	}
}

// Проверяем, находится ли фигура в упавших клетках
func Rotate_in_fallen_cells(rotate Figure_form) bool {
	rotate_in_cells := false
	for _, figure_cell := range rotate.form {
		for _, fallen_cell := range fallen_cells {
			wg.Add(1)
			go func(figure_cell, fallen_cell Cell) {
				if figure_cell.x == fallen_cell.x &&
					figure_cell.y == fallen_cell.y {
					rotate_in_cells = true
				}
				wg.Done()
			}(figure_cell, fallen_cell)
		}
	}
	wg.Wait()

	return rotate_in_cells
}

// Поворачиваем фигуру
func (f *Figure) Rotate() {
	rotated_figure := figure_now.rotates[(index_rotate+1)%len(figure_now.rotates)]

	// Поворачиваем, если не задевает упавшие клетки
	if !Rotate_in_fallen_cells(rotated_figure) {
		index_rotate = (index_rotate + 1) % len(figure_now.rotates)
	}
}
