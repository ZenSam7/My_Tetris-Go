package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math/rand"
	"time"
)

var (
	// Объявляем фигуры
	c              = cell_size // Для краткости
	list_of_figurs = [7]Figure{
		Figure{[4]Cell{{0, 0, color_yellow}, {c, 0, color_yellow}, {0, c, color_yellow}, {c, c, color_yellow}}},             // O
		Figure{[4]Cell{{0, 0, color_gray}, {0, c, color_gray}, {0, 2 * c, color_gray}, {c, 2 * c, color_gray}}},             // L
		Figure{[4]Cell{{c, 0, color_blue}, {c, c, color_blue}, {c, 2 * c, color_blue}, {0, 2 * c, color_blue}}},             // J
		Figure{[4]Cell{{0, 0, color_goluboy}, {0, c, color_goluboy}, {0, 2 * c, color_goluboy}, {0, 3 * c, color_goluboy}}}, // I
		Figure{[4]Cell{{0, 0, color_violet}, {c, 0, color_violet}, {2 * c, 0, color_violet}, {c, c, color_violet}}},         // T
		Figure{[4]Cell{{0, 0, color_red}, {c, 0, color_red}, {c, c, color_red}, {2 * c, c, color_red}}},                     // Z
		Figure{[4]Cell{{0, c, color_green}, {c, c, color_green}, {c, 0, color_green}, {2 * c, 0, color_green}}},             // S
	}
	figure_now = Figure{}

	// Записываем все упавшие клетки
	fallen_cells = []Cell{}

	// Дополнительно
	timer_for_falling  = time.Now().UnixMilli()
	timer_for_move     = time.Now().UnixMilli()
	timer_for_KeySpace = time.Now().UnixMilli()
)

type Cell struct {
	x, y  int
	color color.RGBA
}

// Каждая фигура состоит из 4х клеток
type Figure struct {
	form [4]Cell
}

// Методы для фигур

// Сдвигаем вниз
func (f *Figure) Move_down() (was_collision bool) {
	possible_move_down := true
	// Если какая-то клетка находится на полу, то не сдвигаем
	for _, figure_cell := range f.form {
		if figure_cell.y == height_wind-cell_size {
			possible_move_down = false
		}
	}

	// Когда фигура на полу, не совершаем лишних операций
	if possible_move_down {
		// Если обнаруживается упавшая клетка под фигурой, то не опускаем
		for _, figure_cell := range f.form {
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
		for i, figure_cell := range f.form {
			moved_cell := figure_cell
			moved_cell.y += cell_size
			f.form[i] = moved_cell
		}
	} else { // Если не двигаем, значит у нас коллизия
		Collision()
		return true
	}
	return false
}

// Сдвигаем влево
func (f *Figure) Move_left() {
	// Если какая-то клетка находится около стены, то не сдвигаем
	possible_move_left := true
	for _, figure_cell := range f.form {
		if figure_cell.x == 0 {
			possible_move_left = false
		}
	}

	// Когда фигура у стены, не совершаем лишних операций
	if possible_move_left {
		// Если обнаруживается упавшая клетка около фигуры, то не двигаем
		for _, figure_cell := range f.form {
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

	// Сдвигаем
	if possible_move_left {
		for i, figure_cell := range f.form {
			moved_cell := figure_cell
			moved_cell.x -= cell_size
			f.form[i] = moved_cell
		}
	}
}

// Сдвигаем вправо
func (f *Figure) Move_right() {
	// Если какая-то клетка находится около стены, то не сдвигаем
	possible_move_right := true
	for _, figure_cell := range f.form {
		if figure_cell.x == width_area-cell_size {
			possible_move_right = false
		}
	}

	// Когда фигура у стены, не совершаем лишних операций
	if possible_move_right {
		// Если обнаруживается упавшая клетка около фигуры, то не двигаем
		for _, figure_cell := range f.form {
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
		for i, figure_cell := range f.form {
			moved_cell := figure_cell
			moved_cell.x += cell_size
			f.form[i] = moved_cell
		}
	}
}

// Создаём новую фигуру для управления
func Random_figure_now() {
	figure_now = list_of_figurs[rand.Intn(7)]

	// Сдвигаем фугуру к центру
	for i := 0; i < (width_area/cell_size)/2-1; i++ {
		figure_now.Move_right()
	}

	// Если новая фигура появляется в уже упавшей, то мы проиграли
	for _, fallen_cell := range fallen_cells {
		for _, figure_cell := range figure_now.form {
			wg.Add(1)
			go func(fallen_cell, figure_cell Cell) {
				if fallen_cell.x == figure_cell.x &&
					fallen_cell.y == figure_cell.y {
					Game_over()
				}
				wg.Done()
			}(fallen_cell, figure_cell)
		}
	}
	wg.Wait()
}

// Всячески двигаем фигуру
func Control_figure() {
	time_now := time.Now().UnixMilli()

	// Постепенно спускаем фигуру
	if time_now >= timer_for_falling+game_speed {
		figure_now.Move_down()
		timer_for_falling = time_now
	}

	// Двигаем влево вправо и вниз по нажатию кнопки
	if (ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft)) &&
		time_now >= timer_for_move+speed_move_figure {
		figure_now.Move_left()
		timer_for_move = time_now
	}
	if (ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight)) &&
		time_now >= timer_for_move+speed_move_figure {
		figure_now.Move_right()
		timer_for_move = time_now
	}
	if (ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown)) &&
		time_now >= timer_for_move+speed_move_figure {
		figure_now.Move_down()
		timer_for_move = time_now
	}
	// Двигаем вниз, пока не обнаружим коллизию (путём замены фигуры)
	if ebiten.IsKeyPressed(ebiten.KeySpace) && time_now >= timer_for_KeySpace+time_keydown_space {
		was_collision := figure_now.Move_down()
		for !was_collision {
			was_collision = figure_now.Move_down()
		}
		timer_for_KeySpace = time_now
	}
}

// Оформляем столкновения
func Collision() {
	// Записываем фигуру как упавшую
	for _, cell := range figure_now.form {
		fallen_cells = append(fallen_cells, cell)
	}

	// Собралась ли полная линия

	Random_figure_now()
}

func Game_over() {
	GAME_OVER = true
}
