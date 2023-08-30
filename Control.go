package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
	"os"
	"time"
)

var (
	// Дополнительно
	timer_for_falling  = time.Now().UnixMilli()
	timer_for_move     = time.Now().UnixMilli()
	timer_for_KeySpace = time.Now().UnixMilli()
	timer_for_rotate   = time.Now().UnixMilli()
	timer_for_restart  = time.Now().UnixMilli()
)

// Создаём новую фигуру для управления
func Random_figure_now() {
	figure_now = list_of_figurs[rand.Intn(7)]
	index_rotate = 0

	// Сдвигаем фугуру к центру
	for i := 0; i < (width_area/cell_size)/2-1; i++ {
		figure_now.Move_right()
	}

	// Если новая фигура появляется в уже упавшей, то мы проиграли
	for _, fallen_cell := range fallen_cells {
		for _, figure_cell := range figure_now.rotates[index_rotate].form {
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
		was_collision := figure_now.Move_down()
		for !was_collision {
			was_collision = figure_now.Move_down()
		}
		timer_for_KeySpace = time_now
	}
}

// Надо отделить от контроля фигуры (см. как работает GAME_OVER)
func Exit_and_restart() {
	// Закрываем прогу
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	// Рестартим
	if ebiten.IsKeyPressed(ebiten.KeyR) && (time.Now().UnixMilli() >= timer_for_restart+1000) {
		fallen_cells = []Cell{}
		Random_figure_now()
		GAME_OVER = false

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
	// ...

	Random_figure_now()
}

// Когда закончили игру
func Game_over() {
	GAME_OVER = true
}
