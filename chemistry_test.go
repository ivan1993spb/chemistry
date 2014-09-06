package chemistry

import (
	"fmt"
	"testing"
)

type ElectronConfTest struct {
	Elem        Element
	Full, Short string
}

func TestElectronConf(t *testing.T) {
	fmt.Println("TestElectronConf:")
	tests := []*ElectronConfTest{
		&ElectronConfTest{
			ELEMENT_H,
			"1s^1",
			"1s^1",
		},
		&ElectronConfTest{
			ELEMENT_Be,
			"1s^2 2s^2",
			"2s^2",
		},
		&ElectronConfTest{
			ELEMENT_Ag,
			"1s^2 2s^2 2p^6 3s^2 3p^6 3d^10 4s^2 4p^6 4d^10 5s^1",
			"4d^10 5s^1",
		},
		&ElectronConfTest{
			ELEMENT_Sb,
			"1s^2 2s^2 2p^6 3s^2 3p^6 3d^10 4s^2 4p^6 4d^10 5s^2 5p^3",
			"4d^10 5p^3",
		},
		&ElectronConfTest{
			ELEMENT_Ga,
			"1s^2 2s^2 2p^6 3s^2 3p^6 3d^10 4s^2 4p^1",
			"4p^1",
		},
		&ElectronConfTest{
			ELEMENT_Bi,
			"1s^2 2s^2 2p^6 3s^2 3p^6 3d^10 4s^2 4p^6 4d^10 4f^14 5s^2 5p^6 5d^10 6s^2 6p^3",
			"5d^10 6p^3",
		},
		&ElectronConfTest{
			ELEMENT_Pb,
			"1s^2 2s^2 2p^6 3s^2 3p^6 3d^10 4s^2 4p^6 4d^10 4f^14 5s^2 5p^6 5d^10 6s^2 6p^2",
			"5d^10 6p^2",
		},
		&ElectronConfTest{
			ELEMENT_Ac,
			"1s^2 2s^2 2p^6 3s^2 3p^6 3d^10 4s^2 4p^6 4d^10 4f^14 5s^2 5p^6 5d^10 6s^2 6p^6 6d^1 7s^2",
			"5d^10 6d^1 7s^2",
		},
		&ElectronConfTest{
			ELEMENT_Er,
			"1s^2 2s^2 2p^6 3s^2 3p^6 3d^10 4s^2 4p^6 4d^10 4f^12 5s^2 5p^6 6s^2",
			"4f^12 5p^6 6s^2",
		},
		&ElectronConfTest{
			ELEMENT_Ba,
			"1s^2 2s^2 2p^6 3s^2 3p^6 3d^10 4s^2 4p^6 4d^10 5s^2 5p^6 6s^2",
			"4d^10 5p^6 6s^2",
		},
	}

	for i, test := range tests {
		fmt.Printf("%3d. %3s... ", i+1, test.Elem)
		if test.Elem.ElectronConfiguration().String() != test.Full {
			fmt.Println("fail")
			t.Fatal("fail")
		}
		fmt.Print("#1[ok] ... ")
		if test.Elem.ElectronConfiguration().Short() != test.Short {
			fmt.Println("fail")
			t.Fatal("fail")
		}
		fmt.Println("#2[ok]")
	}
}

type LevelTest struct {
	elem   Element
	result []uint8
}

func TestLevels(t *testing.T) {
	fmt.Println("TestLevels:")
	tests := []*LevelTest{
		{ELEMENT_Au, []uint8{2, 8, 18, 32, 18, 1}},
		{ELEMENT_Yb, []uint8{2, 8, 18, 32, 8, 2}},
		{ELEMENT_Ca, []uint8{2, 8, 8, 2}},
		{ELEMENT_Pa, []uint8{2, 8, 18, 32, 20, 9, 2}},
		{ELEMENT_Fm, []uint8{2, 8, 18, 32, 30, 8, 2}},
		{ELEMENT_Tc, []uint8{2, 8, 18, 13, 2}},
		{ELEMENT_W, []uint8{2, 8, 18, 32, 12, 2}},
		{ELEMENT_Dy, []uint8{2, 8, 18, 28, 8, 2}},
		{ELEMENT_Uup, []uint8{2, 8, 18, 32, 32, 18, 5}},
		{ELEMENT_Fr, []uint8{2, 8, 18, 32, 18, 8, 1}},
	}
	for i, test := range tests {
		fmt.Printf("%3d. %s ...", i+1, test.elem)
		levels := test.elem.ElectronConfiguration().EnergyPerShell()
		if len(levels) != len(test.result) {
			fmt.Println("fail")
			t.Fatal("fail")
		}
		for j, _ := range levels {
			if levels[j] != test.result[j] {
				fmt.Println("fail")
				t.Fatal("fail")
			}
		}
		fmt.Println("ok")
	}
}

func Test_(t *testing.T) {
	fmt.Println(ELEMENT_Usu)
	fmt.Println(ELEMENT_Uoq)
}
