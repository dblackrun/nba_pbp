package nba_pbp

import (
	"testing"
)

func TestIntInSlice(t *testing.T) {
	slice_to_test := []int64{0, 15, 60, 74}
	in_slice := IntInSlice(int64(15), slice_to_test)
	if in_slice != true {
		t.Errorf("Expected true. Got %t.", in_slice)
	}

	not_in_slice := IntInSlice(int64(18), slice_to_test)
	if not_in_slice != false {
		t.Errorf("Expected false. Got %t.", not_in_slice)
	}
}

func TestGetYearForGame(t *testing.T) {
	game_id := "0021700579"
	expected_year := "2017"
	year := GetYearForGame(game_id)
	if year != expected_year {
		t.Errorf("Expected %s. Got %s.", expected_year, year)
	}

	game_id = "0020000579"
	expected_year = "2000"
	year = GetYearForGame(game_id)
	if year != expected_year {
		t.Errorf("Expected %s. Got %s.", expected_year, year)
	}

	game_id = "0029800579"
	expected_year = "1998"
	year = GetYearForGame(game_id)
	if year != expected_year {
		t.Errorf("Expected %s. Got %s.", expected_year, year)
	}
}
