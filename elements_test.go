package chemistry

import (
	"fmt"
	"testing"
)

type LevelsTest struct {
	Elem        Element
	Full, Short string
}

func TestLevels(t *testing.T) {
	tests := []*LevelsTest{
		&LevelsTest{
			ELEMENT_H,
			"1s^1",
			"1s^1",
		},
		&LevelsTest{
			ELEMENT_Be,
			"1s^2 2s^2",
			"2s^2",
		},
		&LevelsTest{
			ELEMENT_Ag,
			"1s^2 2s^2 2p^6 3s^2 3p^6 3d^10 4s^2 4p^6 4d^10 5s^1",
			"4d^10 5s^1",
		},
		&LevelsTest{
			ELEMENT_Sb,
			"1s^2 2s^2 2p^6 3s^2 3p^6 3d^10 4s^2 4p^6 4d^10 5s^2 5p^3",
			"4d^10 5p^3",
		},
		&LevelsTest{
			ELEMENT_Ga,
			"1s^2 2s^2 2p^6 3s^2 3p^6 3d^10 4s^2 4p^1",
			"4p^1",
		},
		&LevelsTest{
			ELEMENT_Bi,
			"1s^2 2s^2 2p^6 3s^2 3p^6 3d^10 4s^2 4p^6 4d^10 4f^14 5s^2 5p^6 5d^10 6s^2 6p^3",
			"5d^10 6p^3",
		},
		&LevelsTest{
			ELEMENT_Pb,
			"1s^2 2s^2 2p^6 3s^2 3p^6 3d^10 4s^2 4p^6 4d^10 4f^14 5s^2 5p^6 5d^10 6s^2 6p^2",
			"5d^10 6p^2",
		},
		&LevelsTest{
			ELEMENT_Ac,
			"1s^2 2s^2 2p^6 3s^2 3p^6 3d^10 4s^2 4p^6 4d^10 4f^14 5s^2 5p^6 5d^10 6s^2 6p^6 6d^1 7s^2",
			"5d^10 6d^1 7s^2",
		},
		&LevelsTest{
			ELEMENT_Er,
			"1s^2 2s^2 2p^6 3s^2 3p^6 3d^10 4s^2 4p^6 4d^10 4f^12 5s^2 5p^6 6s^2",
			"4f^12 5p^6 6s^2",
		},
		&LevelsTest{
			ELEMENT_Ba,
			"1s^2 2s^2 2p^6 3s^2 3p^6 3d^10 4s^2 4p^6 4d^10 5s^2 5p^6 6s^2",
			"4d^10 5p^6 6s^2",
		},
	}

	for i, test := range tests {
		fmt.Printf("%3d. %3s... ", i+1, test.Elem)
		if LevelsFullString(test.Elem.Levels()) != test.Full {
			fmt.Println("fail")
			t.Fatal("fail")
		}
		fmt.Print("#1[ok] ... ")
		if LevelsShortString(test.Elem.Levels()) != test.Short {
			fmt.Println("fail")
			t.Fatal("fail")
		}
		fmt.Println("#2[ok]")
	}
}
