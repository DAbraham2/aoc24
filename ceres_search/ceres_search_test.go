package ceressearch

import (
	"fmt"
	"slices"
	"testing"
)

func TestGetAll(t *testing.T) {
	out := get_possible_subtexts(point{0, 0}, 4, 2)

	if len(out) != 2 {
		t.Errorf("Len expected %d; len got %d", 2, len(out))
	}
	fmt.Printf("out: %d", out)
}

func TestGetAll2(t *testing.T) {
	out := get_possible_subtexts(point{0, 0}, 4, 4)

	if len(out) != 6 {
		t.Errorf("Len expected %d; len got %d", 6, len(out))
	}
	fmt.Printf("out: %d", out)
}

func TestGetFromMatrix(t *testing.T) {
	m := ByteMatrix{
		[]byte("MMMSXXMASM"),
	}

	xmas := get_slice_from_matrix(&m, [4]point{{5, 0}, {6, 0}, {7, 0}, {8, 0}})

	if len(xmas) != 4 {
		t.Errorf("Len expected %d, len got %d", 4, len(xmas))
	}

	if !slices.Equal(xmas, []byte("XMAS")) {
		t.Errorf("Wanted XMAS, got %s", xmas)
	}
}

func TestHasMas(t *testing.T) {
	m := [3][3]byte{
		{'M', '.', 'S'},
		{'.', 'A', '.'},
		{'M', '.', 'S'},
	}

	if !submatrix_has_mas(m) {
		t.Errorf("Wanted true, got false")
	}
}

func TestHasMas2(t *testing.T) {
	m := [3][3]byte{
		{'M', '.', 'M'},
		{'.', 'A', '.'},
		{'S', '.', 'S'},
	}

	if !submatrix_has_mas(m) {
		t.Errorf("Wanted true, got false")
	}
}
func TestHasMas3(t *testing.T) {
	m := [3][3]byte{
		{'S', '.', 'M'},
		{'.', 'A', '.'},
		{'S', '.', 'M'},
	}

	if !submatrix_has_mas(m) {
		t.Errorf("Wanted true, got false")
	}
}
func TestSampleText(t *testing.T) {

	m := ByteMatrix{
		[]byte("MMMSXXMASM"),
		[]byte("MSAMXMSMSA"),
		[]byte("AMXSXMAAMM"),
		[]byte("MSAMASMSMX"),
		[]byte("XMASAMXAMM"),
		[]byte("XXAMMXXAMA"),
		[]byte("SMSMSASXSS"),
		[]byte("SAXAMASAAA"),
		[]byte("MAMMMXMMMM"),
		[]byte("MXMXAXMASX"),
	}

	c := Count(m)

	if c != 18 {
		t.Errorf("Wanted 18, got %d", c)
	}
}
func TestSampleText2(t *testing.T) {

	m := ByteMatrix{
		[]byte("MMMSXXMASM"),
		[]byte("MSAMXMSMSA"),
		[]byte("AMXSXMAAMM"),
		[]byte("MSAMASMSMX"),
		[]byte("XMASAMXAMM"),
		[]byte("XXAMMXXAMA"),
		[]byte("SMSMSASXSS"),
		[]byte("SAXAMASAAA"),
		[]byte("MAMMMXMMMM"),
		[]byte("MXMXAXMASX"),
	}

	c := CountMas(m)

	if c != 9 {
		t.Errorf("Wanted 9, got %d", c)
	}
}
