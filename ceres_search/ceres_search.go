package ceressearch

import "slices"

type point struct{ x, y int }
type ByteMatrix [][]byte

func has_xmas(text []byte) bool {
	xmas := []byte("XMAS")
	if len(text) == 4 {
		return slices.Equal(xmas, text)
	}
	return false
}

func get_possible_subtexts(p point, xSiz, ySiz int) [][4]point {
	ret := make([][4]point, 0)

	if p.x+4 <= xSiz { //left-right
		ret = append(ret, [4]point{{p.x, p.y}, {p.x + 1, p.y}, {p.x + 2, p.y}, {p.x + 3, p.y}})
		ret = append(ret, [4]point{{p.x + 3, p.y}, {p.x + 2, p.y}, {p.x + 1, p.y}, {p.x, p.y}})
	}

	if p.y+4 <= ySiz { //down-up
		ret = append(ret, [4]point{{p.x, p.y}, {p.x, p.y + 1}, {p.x, p.y + 2}, {p.x, p.y + 3}})
		ret = append(ret, [4]point{{p.x, p.y + 3}, {p.x, p.y + 2}, {p.x, p.y + 1}, {p.x, p.y}})
	}

	if p.x+4 <= xSiz && p.y+4 <= ySiz { //cross forward
		ret = append(ret, [4]point{{p.x, p.y}, {p.x + 1, p.y + 1}, {p.x + 2, p.y + 2}, {p.x + 3, p.y + 3}})
		ret = append(ret, [4]point{{p.x + 3, p.y + 3}, {p.x + 2, p.y + 2}, {p.x + 1, p.y + 1}, {p.x, p.y}})
	}

	if p.x > 2 && p.y+4 <= ySiz {
		ret = append(ret, [4]point{{p.x, p.y}, {p.x - 1, p.y + 1}, {p.x - 2, p.y + 2}, {p.x - 3, p.y + 3}})
		ret = append(ret, [4]point{{p.x - 3, p.y + 3}, {p.x - 2, p.y + 2}, {p.x - 1, p.y + 1}, {p.x, p.y}})
	}

	return ret
}

func get_slice_from_matrix(matrix *ByteMatrix, line [4]point) []byte {
	m := *matrix
	ret := []byte("")
	for _, v := range line {
		ret = append(ret, m[v.y][v.x])
	}

	return ret
}

func Count(matrix ByteMatrix) int {
	y := len(matrix)
	x := len(matrix[0])
	count := 0
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			index := point{i, j}
			ps := get_possible_subtexts(index, x, y)

			for _, v := range ps {
				if has_xmas(get_slice_from_matrix(&matrix, v)) {
					count++
				}
			}

		}
	}
	return count
}

func submatrix_has_mas(matrix [3][3]byte) bool {
	if matrix[1][1] != 'A' {
		return false
	}
	d_f := matrix[0][0] == 'M' && matrix[2][2] == 'S'
	d_b := matrix[0][0] == 'S' && matrix[2][2] == 'M'

	u_f := matrix[2][0] == 'M' && matrix[0][2] == 'S'
	u_b := matrix[2][0] == 'S' && matrix[0][2] == 'M'

	return (d_f || d_b) && (u_f || u_b)
}

func CountMas(matrix ByteMatrix) int {
	count := 0

	y := len(matrix)
	x := len(matrix[0])

	for j := 0; j < y-2; j++ {
		for i := 0; i < x-2; i++ {
			if matrix[j][i] != 'M' && matrix[j][i] != 'S' {
				continue
			}
			a := matrix[j]
			b := matrix[j+1]
			c := matrix[j+2]
			sub := [3][3]byte{
				{a[i], a[i+1], a[i+2]},
				{b[i], b[i+1], b[i+2]},
				{c[i], c[i+1], c[i+2]},
			}

			if submatrix_has_mas(sub) {
				count++
			}
		}
	}

	return count
}
