package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
	"os"
	"time"
)

var (
	// Разные таймеры
	timer_for_falling  = time.Now().UnixMilli()
	timer_for_move     = time.Now().UnixMilli()
	timer_for_KeySpace = time.Now().UnixMilli()
	timer_for_rotate   = time.Now().UnixMilli()
	timer_for_restart  = time.Now().UnixMilli()

	// Можно ли использовать кармашек
	can_use_pocket = true
)

// Создаём новую фигуру для управления
func Random_figure_now() {

	// Изменение в следующей и текущей фигуре, а значить перерисовываем всё
	change_in_area = true
	change_in_menu = true

	index_rotate = 0

	can_use_pocket = true

	// Создаём новую следущую фигуру
	// При условии, что она не равна предыдущей фигуре
	// (я это сделал чтобы небыло последовательностей одинаковых фигур, которые сильно усложняют)
	figure_now = next_figure
	for next_figure.name == figure_now.name {
		next_figure = Deep_copy_figure(list_of_figures[rand.Intn(7)])
	}

	// Сдвигаем фугуру к центру
	for i := 0; i < (width_area/cell_size)/2-1; i++ {
		figure_now.Move_right()
	}

	// Если новая фигура появляется в уже упавшей, то мы проиграли
	if Rotate_in_fallen_cells(figure_now.rotates[0]) {
		Game_over()
	}
}

// Всячески двигаем фигуру
func Control() {
	time_now := time.Now().UnixMilli()

	// Постепенно спускаем фигуру
	if time_now >= timer_for_falling+int64(game_speed) {
		figure_now.Move_down()
		timer_for_falling = time_now
	}

	// Поворачиваем фигуру
	if (ebiten.IsKeyPressed(ebiten.KeyW) ||
		ebiten.IsKeyPressed(ebiten.KeyArrowUp)) &&
		time_now >= timer_for_rotate+time_rotate {
		figure_now.Rotate()
		timer_for_rotate = time_now
	}

	// Двигаем влево вправо и вниз по нажатию кнопки
	can_move := time_now >= timer_for_move+speed_move_figure
	if can_move {
		if ebiten.IsKeyPressed(ebiten.KeyA) ||
			ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
			figure_now.Move_left()
			timer_for_move = time_now
		} else if ebiten.IsKeyPressed(ebiten.KeyD) ||
			ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
			figure_now.Move_right()
			timer_for_move = time_now
		}
		if ebiten.IsKeyPressed(ebiten.KeyS) ||
			ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
			figure_now.Move_down()
			timer_for_move = time_now
		}
	}

	// Когда нажимаем пробел, двигаем вниз до упора
	if ebiten.IsKeyPressed(ebiten.KeySpace) && time_now >= timer_for_KeySpace+time_keydown_space {
		was_collision := false
		for !was_collision {
			was_collision = figure_now.Move_down()
		}
		timer_for_KeySpace = time_now
	}

	// Используем кармашек
	if (ebiten.IsKeyPressed(ebiten.KeyShift) ||
		ebiten.IsKeyPressed(ebiten.KeyShiftRight) ||
		ebiten.IsKeyPressed(ebiten.KeyShiftLeft)) &&
		time_now >= timer_for_rotate+time_keydown_space &&
		can_use_pocket {

		change_in_menu = true
		timer_for_rotate = time_now
		index_rotate = 0

		// Изменение в фигуре из кармашка, а значить перерисовываем меню
		Draw_menu()

		// Нельзя использовать кармашек второй раз, пока фигура не упала
		can_use_pocket = false

		// Если фигуры нет, то добавляем
		if figure_in_pocket == "" {
			figure_in_pocket = figure_now.name
			Random_figure_now()

		} else { // Если фигура есть, то меняем
			// Ищем фигуру из кармашка в списке фигур
			for _, find_figure := range list_of_figures {
				if find_figure.name == figure_in_pocket {
					// Меняем местами
					figure_in_pocket = figure_now.name
					figure_now = Deep_copy_figure(find_figure)

					// Сдвигаем фугуру к центру
					for i := 0; i < (width_area/cell_size)/2-1; i++ {
						figure_now.Move_right()
					}

					break
				}
			}
		}

		timer_for_restart = time_now
	}

}

// Надо отделить от контроля фигуры (см. как работает game_over)
func Exit_and_restart() {
	// Закрываем прогу
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	// Рестартим
	if ebiten.IsKeyPressed(ebiten.KeyR) && (time.Now().UnixMilli() >= timer_for_restart+1000) {
		fallen_cells = []Cell{}
		game_over = false
		game_score = 0
		game_speed = const_game_speed
		figure_in_pocket = ""
		Random_figure_now()

		timer_for_restart = time.Now().UnixMilli()
	}
}

// Оформляем столкновения
func Collision() {
	// Записываем фигуру как упавшую
	for _, cell := range figure_now.rotates[index_rotate].form {
		fallen_cells = append(fallen_cells, cell)
	}

	// Собралась ли полная линия
	for num_row := 0; num_row < height_wind; num_row += cell_size {
		num_cells_in_row := 0 // Количество клеток в ряду num_row

		for _, fallen_cell := range fallen_cells {
			if fallen_cell.y == num_row {
				num_cells_in_row += 1
			}
		}

		// Если собрался полный ряд
		if num_cells_in_row == width_area/cell_size {
			Collecting_row(num_row)
		}
	}

	Random_figure_now()
}

// Когда закончили игру
func Game_over() {
	game_over = true
}

// Что делаем когда ряд собран
func Collecting_row(num_row int) {
	// Ускоряем игру
	game_speed = int(float32(game_speed) * speed_factor)

	// Увеличиваем счётчик
	game_score += 1

	ind_cells_in_row := []int{} // Индексы всех клеток в полном ряду
	// Добавляем индексы всех клеток на удаление
	for ind, fallen_cell := range fallen_cells {
		if fallen_cell.y == num_row {
			ind_cells_in_row = append(ind_cells_in_row, ind)
		}
	}

	// Удаляем
	for bias, ind := range ind_cells_in_row {
		fallen_cells = append(fallen_cells[:ind-bias], fallen_cells[ind-bias+1:]...)
	}

	// Сдвигаем все клетки сверху
	for i, fallen_cell := range fallen_cells {
		if fallen_cell.y <= num_row {
			new_cell := fallen_cell
			new_cell.y += cell_size
			fallen_cells[i] = new_cell
		}
	}
}

// Создаём глубокую копию фигуры (фактически новую фигуру)
func Deep_copy_figure(figure Figure) Figure {
	new_figure := Figure{}

	// Делаем копию каждого поворота
	for _, rotate := range figure.rotates {
		new_figure_form := Figure_form{}

		// Делаем копию каждого Figure_form
		for i, cell := range rotate.form {
			new_cell := Cell{cell.x, cell.y, cell.color}
			new_figure_form.form[i] = new_cell
		}

		// Добавляем поворот
		new_figure.rotates = append(new_figure.rotates, new_figure_form)
	}
	new_figure.name = figure.name

	return new_figure
}
