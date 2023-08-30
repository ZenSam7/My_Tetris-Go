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

		{[]Figure_form{{[4]Cell{{0, 0, color_gray}, {0, c, color_gray}, {0, 2 * c, color_gray}, {c, 2 * c, color_gray}}}, {[4]Cell{{0, 0, color_gray}, {c, 0, color_gray}, {2 * c, 0, color_gray}, {0, c, color_gray}}}, {[4]Cell{{0, 0, color_gray}, {c, 0, color_gray}, {c, c, color_gray}, {c, 2 * c, color_gray}}}, {[4]Cell{{0, c, color_gray}, {c, c, color_gray}, {2 * c, c, color_gray}, {2 * c, 0, color_gray}}}}},
		{[]Figure_form{{[4]Cell{{0, 0, color_gray}, {0, c, color_gray}, {0, 2 * c, color_gray}, {c, 2 * c, color_gray}}}, {[4]Cell{{0, 0, color_gray}, {c, 0, color_gray}, {2 * c, 0, color_gray}, {0, c, color_gray}}}, {[4]Cell{{0, 0, color_gray}, {c, 0, color_gray}, {c, c, color_gray}, {c, 2 * c, color_gray}}}, {[4]Cell{{0, c, color_gray}, {c, c, color_gray}, {2 * c, c, color_gray}, {2 * c, 0, color_gray}}}}},
		{[]Figure_form{{[4]Cell{{0, 0, color_gray}, {0, c, color_gray}, {0, 2 * c, color_gray}, {c, 2 * c, color_gray}}}, {[4]Cell{{0, 0, color_gray}, {c, 0, color_gray}, {2 * c, 0, color_gray}, {0, c, color_gray}}}, {[4]Cell{{0, 0, color_gray}, {c, 0, color_gray}, {c, c, color_gray}, {c, 2 * c, color_gray}}}, {[4]Cell{{0, c, color_gray}, {c, c, color_gray}, {2 * c, c, color_gray}, {2 * c, 0, color_gray}}}}},
		{[]Figure_form{{[4]Cell{{0, 0, color_gray}, {0, c, color_gray}, {0, 2 * c, color_gray}, {c, 2 * c, color_gray}}}, {[4]Cell{{0, 0, color_gray}, {c, 0, color_gray}, {2 * c, 0, color_gray}, {0, c, color_gray}}}, {[4]Cell{{0, 0, color_gray}, {c, 0, color_gray}, {c, c, color_gray}, {c, 2 * c, color_gray}}}, {[4]Cell{{0, c, color_gray}, {c, c, color_gray}, {2 * c, c, color_gray}, {2 * c, 0, color_gray}}}}},
		{[]Figure_form{{[4]Cell{{0, 0, color_gray}, {0, c, color_gray}, {0, 2 * c, color_gray}, {c, 2 * c, color_gray}}}, {[4]Cell{{0, 0, color_gray}, {c, 0, color_gray}, {2 * c, 0, color_gray}, {0, c, color_gray}}}, {[4]Cell{{0, 0, color_gray}, {c, 0, color_gray}, {c, c, color_gray}, {c, 2 * c, color_gray}}}, {[4]Cell{{0, c, color_gray}, {c, c, color_gray}, {2 * c, c, color_gray}, {2 * c, 0, color_gray}}}}},

		//Figure{[4]Cell{{c, 0, color_blue}, {c, c, color_blue}, {c, 2 * c, color_blue}, {0, 2 * c, color_blue}}},             // J
		//Figure{[4]Cell{{0, 0, color_goluboy}, {0, c, color_goluboy}, {0, 2 * c, color_goluboy}, {0, 3 * c, color_goluboy}}}, // I
		//Figure{[4]Cell{{0, 0, color_violet}, {c, 0, color_violet}, {2 * c, 0, color_violet}, {c, c, color_violet}}},         // T
		//Figure{[4]Cell{{0, 0, color_red}, {c, 0, color_red}, {c, c, color_red}, {2 * c, c, color_red}}},                     // Z
		//Figure{[4]Cell{{0, c, color_green}, {c, c, color_green}, {c, 0, color_green}, {2 * c, 0, color_green}}},             // S
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

//

// Сдвигаем вниз
func (f *Figure) Move_down() (was_collision bool) {
	// Сдвигаем все повороты влево
	// А текущий поворот делаем видимым для упавших клеток
	for i, rotate := range f.rotates {
		ghost := true
		if i == index_rotate {
			ghost = false
		}

		was_collision = Move_down_rotate(&rotate, ghost)
		if was_collision {
			return true
		}

		f.rotates[i] = rotate
	}
	return false
}

// Сдвигаем один из поворотив фигуры
func Move_down_rotate(rotate *Figure_form, ghost bool) (was_collision bool) {
	// Если какая-то клетка находится около стены, то не сдвигаем
	possible_move_down := true
	for _, figure_cell := range rotate.form {
		if figure_cell.y == height_wind-cell_size {
			possible_move_down = false
		}
	}

	// Если фигура прозрачная (ghost), то не учитываем другие клетки
	// Когда фигура на полу, не совершаем лишних операций
	if possible_move_down && !ghost {
		// Если обнаруживается упавшая клетка под фигурой, то не опускаем
		for _, figure_cell := range rotate.form {
			for _, fallen_cell := range fallen_cells {
				wg.Add(1)
				go func(figure_cell, fallen_cell Cell) {
					if figure_cell.x == fallen_cell.x &&
						figure_cell.y == fallen_cell.y-cell_size {
						possible_move_down = false
					}
					wg.Done()
				}(figure_cell, fallen_cell)
			}
		}
		wg.Wait()
	}

	// Сдвигаем
	if possible_move_down {
		for i, figure_cell := range rotate.form {
			moved_cell := figure_cell
			moved_cell.y += cell_size
			rotate.form[i] = moved_cell
		}
	} else { // Если не двигаем, значит у нас коллизия
		Collision()
		return true
	}
	return false
}

//

// Сдвигаем влево
func (f *Figure) Move_left() {
	// Сдвигаем все повороты влево
	// А текущий поворот делаем видимым для упавших клеток
	for i, rotate := range f.rotates {
		ghost := true
		if i == index_rotate {
			ghost = false
		}

		Move_left_rotate(&rotate, ghost)
		f.rotates[i] = rotate
	}

}

// Сдвигаем один из поворотив фигуры
func Move_left_rotate(rotate *Figure_form, ghost bool) {
	// Если какая-то клетка находится около стены, то не сдвигаем
	possible_move_left := true
	for _, figure_cell := range rotate.form {
		if figure_cell.x == 0 {
			possible_move_left = false
		}
	}

	// Если фигура прозрачная (ghost), то не учитываем другие клетки
	// Когда фигура уже у стены, не совершаем лишних операций
	if possible_move_left && !ghost {
		// Если обнаруживается упавшая клетка около фигуры, то не двигаем
		for _, figure_cell := range rotate.form {
			for _, fallen_cell := range fallen_cells {
				wg.Add(1)
				go func(figure_cell, fallen_cell Cell) {
					if figure_cell.x == fallen_cell.x+cell_size &&
						figure_cell.y == fallen_cell.y {
						possible_move_left = false
					}
					wg.Done()
				}(figure_cell, fallen_cell)
			}
		}
		wg.Wait()
	}

	// Сдвигаем все клетки
	if possible_move_left {
		for i, figure_cell := range rotate.form {
			moved_cell := figure_cell
			moved_cell.x -= cell_size
			rotate.form[i] = moved_cell
		}
	}
}

//

// Сдвигаем вправо
func (f *Figure) Move_right() {
	// Сдвигаем все повороты вправо
	// А текущий поворот делаем видимым для упавших клеток
	for i, rotate := range f.rotates {
		ghost := true
		if i == index_rotate {
			ghost = false
		}

		Move_right_rotate(&rotate, ghost)
		f.rotates[i] = rotate
	}

}

// Сдвигаем один из поворотив фигуры
func Move_right_rotate(rotate *Figure_form, ghost bool) {
	// Если какая-то клетка находится около стены, то не сдвигаем
	possible_move_right := true
	for _, figure_cell := range rotate.form {
		if figure_cell.x == width_area-cell_size {
			possible_move_right = false
		}
	}

	// Если фигура прозрачная (ghost), то не учитываем другие клетки
	// Когда фигура уже у стены, не совершаем лишних операций
	if possible_move_right && !ghost {
		// Если обнаруживается упавшая клетка около фигуры, то не двигаем
		for _, figure_cell := range rotate.form {
			for _, fallen_cell := range fallen_cells {
				wg.Add(1)
				go func(figure_cell, fallen_cell Cell) {
					if figure_cell.x+cell_size == fallen_cell.x &&
						figure_cell.y == fallen_cell.y {
						possible_move_right = false
					}
					wg.Done()
				}(figure_cell, fallen_cell)
			}
		}
		wg.Wait()
	}

	// Сдвигаем
	if possible_move_right {
		for i, figure_cell := range rotate.form {
			moved_cell := figure_cell
			moved_cell.x += cell_size
			rotate.form[i] = moved_cell
		}
	}
}

//

// Поворачиваем фигуру
func (f *Figure) Rotate() {
	index_rotate = (index_rotate + 1) % len(figure_now.rotates)
}
