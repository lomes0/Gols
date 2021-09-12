package main

import (
	"fmt"
	"math/rand"
)

const (
	width   = 158
	height  = 10
	_v_     = 'm'
	___     = ' '
	density = 3
)

type Universe [][]rune

func NewUniverse() Universe {

	uni := make(Universe, height)

	for i := range uni {
		uni[i] = make([]rune, width)
	}
	return uni
}

func (u Universe) Print() {

	for _, row := range u {

		var s string = ""
		for _, c := range row {
			s += string(c)
		}
		fmt.Println(s)
	}
}

/*
*@Desc: Randomly populates the board.
 */
func (u Universe) Seed() {

	for i := range u {

		for j := range u[i] {

			if rand.Intn(density) > 0 {
				u[i][j] = ' '
			} else {
				u[i][j] = _v_
			}
		}
	}
}

func (u Universe) is_alive(i int, j int) bool {

	return u[i][j] != ' '
}

/*
*@Desc: Return a number of live cells, according to a list of coordinates.
 */
func (u Universe) count_by_indecies(coords ...int) int {

	count := 0
	for l := 0; l < len(coords); l += 2 {

		i := coords[l]
		j := coords[l+1]

		if u.is_alive(i, j) {
			count++
		}
	}

	return count
}

/*
*@Desc: Helper functions to get the exact list of neighboors.
*	tl = top left
*	bl = bottom left
*	tr = top right
*	br = bottom right
*	tm = top middle
*	bm = bottom middle
*	lm = left middle
*	rm = right middle
 */
func (u Universe) ne_count_tl(i int, j int) int {

	return u.count_by_indecies(i, j+1,
		i+1, j,
		i+1, j+1)
}

func (u Universe) ne_count_bl(i int, j int) int {

	return u.count_by_indecies(i-1, j,
		i-1, j+1,
		i, j+1)
}

func (u Universe) ne_count_tr(i int, j int) int {

	return u.count_by_indecies(i, j-1,
		i+1, j-1,
		i+1, j)
}

func (u Universe) ne_count_br(i int, j int) int {

	return u.count_by_indecies(i-1, j,
		i-1, j-1,
		i, j-1)
}

func (u Universe) ne_count_tm(i int, j int) int {

	return u.count_by_indecies(i, j-1,
		i, j+1,
		i+1, j-1,
		i+1, j,
		i+1, j+1)
}

func (u Universe) ne_count_bm(i int, j int) int {

	return u.count_by_indecies(i, j-1,
		i, j+1,
		i-1, j-1,
		i-1, j,
		i-1, j+1)
}

func (u Universe) ne_count_ml(i int, j int) int {
	return u.count_by_indecies(i-1, j,
		i-1, j+1,
		i, j+1,
		i+1, j,
		i+1, j+1)
}

func (u Universe) ne_count_mr(i int, j int) int {
	return u.count_by_indecies(i-1, j, i-1,
		j-1, i, j-1,
		i+1, j,
		i+1, j-1)
}

func (u Universe) ne_count_mm(i int, j int) int {

	return u.count_by_indecies(i-1, j-1,
		i-1, j,
		i-1, j+1,
		i, j-1,
		i, j+1,
		i+1, j-1,
		i+1, j,
		i+1, j+1)
}

/*
*@Desc: Get number of neighboors for a given cell.
 */
func (u Universe) ne_count(i int, j int) int {

	switch i {
	case 0:
		switch j {
		case 0:
			return u.ne_count_tl(i, j)
		case width - 1:
			return u.ne_count_tr(i, j)
		default:
			return u.ne_count_tm(i, j)
		}

	case height - 1:
		switch j {
		case 0:
			return u.ne_count_bl(i, j)
		case width - 1:
			return u.ne_count_br(i, j)
		default:
			return u.ne_count_bm(i, j)
		}
	}

	switch j {
	case 0:
		return u.ne_count_ml(i, j)

	case width - 1:
		return u.ne_count_mr(i, j)
	default:
		return u.ne_count_mm(i, j)
	}
}

/*
*@Desc: Apply Game of Life rules on a dead cell.
 */
func (u Universe) dead_cell_test(i int, j int) rune {

	ne_count := u.ne_count(i, j)

	if ne_count == 3 {
		return _v_
	} else {
		return ___
	}
}

/*
*@Desc: Apply Game of Life rules on a living cell.
 */
func (u Universe) alive_cell_test(i int, j int) rune {

	ne_count := u.ne_count(i, j)

	if ne_count < 2 {
		return ___
	} else if ne_count <= 3 {
		return _v_
	} else {
		return ___
	}
}

func (u Universe) calc_next(i int, j int) rune {

	switch u[i][j] {
	case _v_:
		return u.alive_cell_test(i, j)
	case ___:
		return u.dead_cell_test(i, j)
	default:
		return u.dead_cell_test(i, j)
	}
}

/*
* @Desc: For each cell, get next state.
 */
func (u *Universe) Next() {

	u2 := NewUniverse()

	for i := range *u {
		for j := range (*u)[i] {
			u2[i][j] = u.calc_next(i, j)
		}
	}

	*u = u2
}
