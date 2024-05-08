package islands

import "testing"

func TestInput1(t *testing.T) {

	m := [][]bool{
		{false, false, true, true},
		{true, false, false, true},
		{true, true, false, false},
		{false, false, false, true},
	}

	count := countIslands(m)
	if count != 3 {
		t.Fail()
	}
}

func TestInput2(t *testing.T) {

	m := [][]bool{
		{false, false, true, true},
		{true, false, false, true},
		{true, true, false, true},
		{false, false, false, true},
	}

	count := countIslands(m)
	if count != 2 {
		t.Fail()
	}
}

func TestInput3(t *testing.T) {

	m := [][]bool{
		{false, false, true, true},
		{true, false, false, true},
		{true, true, false, true},
		{false, false, false, false},
		{false, false, false, true},
		{false, false, false, true},
		{false, false, false, true},
		{false, false, false, true},
		{false, false, false, false},
		{false, false, false, true},
		{true, false, false, true},
		{false, false, false, true},
	}

	count := countIslands(m)
	if count != 5 {
		t.Fail()
	}
}

func TestInput4(t *testing.T) {

	m := [][]bool{
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}

	count := countIslands(m)
	if count != 0 {
		t.Fail()
	}
}

func TestInput5(t *testing.T) {

	m := [][]bool{
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, true, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, true, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}

	count := countIslands(m)
	if count != 2 {
		t.Fail()
	}
}

func TestInput6(t *testing.T) {

	m := [][]bool{
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, true, true, false, false},
		{false, false, false, false, false, false, true, true, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, true, false, false, false, false, false, false, false, false},
		{false, true, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}

	count := countIslands(m)
	if count != 2 {
		t.Fail()
	}
}

func TestInput7(t *testing.T) {

	m := [][]bool{
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, true, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, true, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, true, true, true, false, false},
	}

	count := countIslands(m)
	if count != 3 {
		t.Fail()
	}
}
