package chemistry

import "strconv"

type AtomicNumber byte

func (n AtomicNumber) String() string {
	return strconv.Itoa(int(n)) + "e"
}

type Element byte

const (
	ELEMENT_H Element = 1 + iota
	ELEMENT_He
	ELEMENT_Li
	ELEMENT_Be
	ELEMENT_B
	ELEMENT_C
	ELEMENT_N
	ELEMENT_O
	ELEMENT_F
	ELEMENT_Ne
	ELEMENT_Na
	ELEMENT_Mg
	ELEMENT_Al
	ELEMENT_Si
	ELEMENT_P
	ELEMENT_S
	ELEMENT_Cl
	ELEMENT_Ar
	ELEMENT_K
	ELEMENT_Ca
	ELEMENT_Sc
	ELEMENT_Ti
	ELEMENT_V
	ELEMENT_Cr
	ELEMENT_Mn
	ELEMENT_Fe
	ELEMENT_Co
	ELEMENT_Ni
	ELEMENT_Cu
	ELEMENT_Zn
	ELEMENT_Ga
	ELEMENT_Ge
	ELEMENT_As
	ELEMENT_Se
	ELEMENT_Br
	ELEMENT_Kr
	ELEMENT_Rb
	ELEMENT_Sr
	ELEMENT_Y
	ELEMENT_Zr
	ELEMENT_Nb
	ELEMENT_Mo
	ELEMENT_Tc
	ELEMENT_Ru
	ELEMENT_Rh
	ELEMENT_Pd
	ELEMENT_Ag
	ELEMENT_Cd
	ELEMENT_In
	ELEMENT_Sn
	ELEMENT_Sb
	ELEMENT_Te
	ELEMENT_I
	ELEMENT_Xe
	ELEMENT_Cs
	ELEMENT_Ba
	ELEMENT_La
	ELEMENT_Ce
	ELEMENT_Pr
	ELEMENT_Nd
	ELEMENT_Pm
	ELEMENT_Sm
	ELEMENT_Eu
	ELEMENT_Gd
	ELEMENT_Tb
	ELEMENT_Dy
	ELEMENT_Ho
	ELEMENT_Er
	ELEMENT_Tm
	ELEMENT_Yb
	ELEMENT_Lu
	ELEMENT_Hf
	ELEMENT_Ta
	ELEMENT_W
	ELEMENT_Re
	ELEMENT_Os
	ELEMENT_Ir
	ELEMENT_Pt
	ELEMENT_Au
	ELEMENT_Hg
	ELEMENT_Tl
	ELEMENT_Pb
	ELEMENT_Bi
	ELEMENT_Po
	ELEMENT_At
	ELEMENT_Rn
	ELEMENT_Fr
	ELEMENT_Ra
	ELEMENT_Ac
	ELEMENT_Th
	ELEMENT_Pa
	ELEMENT_U
	ELEMENT_Np
	ELEMENT_Pu
	ELEMENT_Am
	ELEMENT_Cm
	ELEMENT_Bk
	ELEMENT_Cf
	ELEMENT_Es
	ELEMENT_Fm
	ELEMENT_Md
	ELEMENT_No
	ELEMENT_Lr
	ELEMENT_Rf
	ELEMENT_Db
	ELEMENT_Sg
	ELEMENT_Bh
	ELEMENT_Hs
	ELEMENT_Mt
	ELEMENT_Ds
	ELEMENT_Rg
	ELEMENT_Cn
	ELEMENT_Uut
	ELEMENT_Fl
	ELEMENT_Uup
	ELEMENT_Lv
	ELEMENT_Uus
	ELEMENT_Uuo
	ELEMENT_Uue
	ELEMENT_Ubn
	ELEMENT_Ubu
	ELEMENT_Ubb
	ELEMENT_Ubt
	ELEMENT_Ubq
	ELEMENT_Ubp
	ELEMENT_Ubh
)

var (
	symbols = map[Element]string{
		ELEMENT_H:   "H",
		ELEMENT_He:  "He",
		ELEMENT_Li:  "Li",
		ELEMENT_Be:  "Be",
		ELEMENT_B:   "B",
		ELEMENT_C:   "C",
		ELEMENT_N:   "N",
		ELEMENT_O:   "O",
		ELEMENT_F:   "F",
		ELEMENT_Ne:  "Ne",
		ELEMENT_Na:  "Na",
		ELEMENT_Mg:  "Mg",
		ELEMENT_Al:  "Al",
		ELEMENT_Si:  "Si",
		ELEMENT_P:   "P",
		ELEMENT_S:   "S",
		ELEMENT_Cl:  "Cl",
		ELEMENT_Ar:  "Ar",
		ELEMENT_K:   "K",
		ELEMENT_Ca:  "Ca",
		ELEMENT_Sc:  "Sc",
		ELEMENT_Ti:  "Ti",
		ELEMENT_V:   "V",
		ELEMENT_Cr:  "Cr",
		ELEMENT_Mn:  "Mn",
		ELEMENT_Fe:  "Fe",
		ELEMENT_Co:  "Co",
		ELEMENT_Ni:  "Ni",
		ELEMENT_Cu:  "Cu",
		ELEMENT_Zn:  "Zn",
		ELEMENT_Ga:  "Ga",
		ELEMENT_Ge:  "Ge",
		ELEMENT_As:  "As",
		ELEMENT_Se:  "Se",
		ELEMENT_Br:  "Br",
		ELEMENT_Kr:  "Kr",
		ELEMENT_Rb:  "Rb",
		ELEMENT_Sr:  "Sr",
		ELEMENT_Y:   "Y",
		ELEMENT_Zr:  "Zr",
		ELEMENT_Nb:  "Nb",
		ELEMENT_Mo:  "Mo",
		ELEMENT_Tc:  "Tc",
		ELEMENT_Ru:  "Ru",
		ELEMENT_Rh:  "Rh",
		ELEMENT_Pd:  "Pd",
		ELEMENT_Ag:  "Ag",
		ELEMENT_Cd:  "Cd",
		ELEMENT_In:  "In",
		ELEMENT_Sn:  "Sn",
		ELEMENT_Sb:  "Sb",
		ELEMENT_Te:  "Te",
		ELEMENT_I:   "I",
		ELEMENT_Xe:  "Xe",
		ELEMENT_Cs:  "Cs",
		ELEMENT_Ba:  "Ba",
		ELEMENT_La:  "La",
		ELEMENT_Ce:  "Ce",
		ELEMENT_Pr:  "Pr",
		ELEMENT_Nd:  "Nd",
		ELEMENT_Pm:  "Pm",
		ELEMENT_Sm:  "Sm",
		ELEMENT_Eu:  "Eu",
		ELEMENT_Gd:  "Gd",
		ELEMENT_Tb:  "Tb",
		ELEMENT_Dy:  "Dy",
		ELEMENT_Ho:  "Ho",
		ELEMENT_Er:  "Er",
		ELEMENT_Tm:  "Tm",
		ELEMENT_Yb:  "Yb",
		ELEMENT_Lu:  "Lu",
		ELEMENT_Hf:  "Hf",
		ELEMENT_Ta:  "Ta",
		ELEMENT_W:   "W",
		ELEMENT_Re:  "Re",
		ELEMENT_Os:  "Os",
		ELEMENT_Ir:  "Ir",
		ELEMENT_Pt:  "Pt",
		ELEMENT_Au:  "Au",
		ELEMENT_Hg:  "Hg",
		ELEMENT_Tl:  "Tl",
		ELEMENT_Pb:  "Pb",
		ELEMENT_Bi:  "Bi",
		ELEMENT_Po:  "Po",
		ELEMENT_At:  "At",
		ELEMENT_Rn:  "Rn",
		ELEMENT_Fr:  "Fr",
		ELEMENT_Ra:  "Ra",
		ELEMENT_Ac:  "Ac",
		ELEMENT_Th:  "Th",
		ELEMENT_Pa:  "Pa",
		ELEMENT_U:   "U",
		ELEMENT_Np:  "Np",
		ELEMENT_Pu:  "Pu",
		ELEMENT_Am:  "Am",
		ELEMENT_Cm:  "Cm",
		ELEMENT_Bk:  "Bk",
		ELEMENT_Cf:  "Cf",
		ELEMENT_Es:  "Es",
		ELEMENT_Fm:  "Fm",
		ELEMENT_Md:  "Md",
		ELEMENT_No:  "No",
		ELEMENT_Lr:  "Lr",
		ELEMENT_Rf:  "Rf",
		ELEMENT_Db:  "Db",
		ELEMENT_Sg:  "Sg",
		ELEMENT_Bh:  "Bh",
		ELEMENT_Hs:  "Hs",
		ELEMENT_Mt:  "Mt",
		ELEMENT_Ds:  "Ds",
		ELEMENT_Rg:  "Rg",
		ELEMENT_Cn:  "Cn",
		ELEMENT_Uut: "Uut",
		ELEMENT_Fl:  "Fl",
		ELEMENT_Uup: "Uup",
		ELEMENT_Lv:  "Lv",
		ELEMENT_Uus: "Uus",
		ELEMENT_Uuo: "Uuo",
		ELEMENT_Uue: "Uue",
		ELEMENT_Ubn: "Ubn",
		ELEMENT_Ubu: "Ubu",
		ELEMENT_Ubb: "Ubb",
		ELEMENT_Ubt: "Ubt",
		ELEMENT_Ubq: "Ubq",
		ELEMENT_Ubp: "Ubp",
		ELEMENT_Ubh: "Ubh",
	}

	orbitals = "spdfgh"
)

func (e Element) String() string {
	return symbols[e]
}

func (e Element) AtomicNumber() AtomicNumber {
	return AtomicNumber(e)
}

func (e Element) Levels() (levels map[uint8][]uint8) {
	var (
		n     = uint8(1) // main quantum number (n > 0)
		l     = uint8(0) // orbital quantum number (n - 1)
		count = uint8(e) // electron count
	)

	levels = make(map[uint8][]uint8)

level:
	{
		// create energy sublevels
		if len(levels[n]) == 0 {
			levels[n] = make([]uint8, n)
		}
	}

sublevel:
	{
		sublevelLength := 2 * (2*l + 1)
		if count > sublevelLength {
			levels[n][l] += sublevelLength
			count -= sublevelLength
		} else {
			levels[n][l] += count
			count = 0
		}
	}

	// if elctrons isn't depleted calculate new main and orbital
	// quantum numbers
	if count > 0 {

		if l > 0 {
			// without growth of sum of n and l
			n += 1
			l -= 1

			goto level
		}

		var (
			// start values n and l
			new_n uint8
			new_l = n

			// growth of energy for 1
			energy = n + 1
		)

		for {
			new_n = energy - new_l

			if new_n > new_l {
				break
			} else {
				new_l--
			}
		}

		// level router
		if n != new_n {
			n, l = new_n, new_l
			goto level
		}

		l = new_l
		goto sublevel
	}

	// electron skipping

	switch e {
	case ELEMENT_Cr, ELEMENT_Cu:
		levels[4][0] -= 1
		levels[3][2] += 1

	case ELEMENT_Nb, ELEMENT_Mo, ELEMENT_Ru, ELEMENT_Rh, ELEMENT_Ag:
		levels[5][0] -= 1
		levels[4][2] += 1

	case ELEMENT_Pd:
		levels[5][0] -= 2
		levels[4][2] += 2

	case ELEMENT_Pt, ELEMENT_Au:
		levels[6][0] -= 1
		levels[5][2] += 1

	case ELEMENT_Ds, ELEMENT_Rg:
		levels[7][0] -= 1
		levels[6][2] += 1

	case ELEMENT_La, ELEMENT_Gd:
		levels[4][3] -= 1
		levels[5][2] += 1

	case ELEMENT_Ac, ELEMENT_Pa, ELEMENT_U, ELEMENT_Np, ELEMENT_Cm:
		levels[5][3] -= 1
		levels[6][2] += 1

	case ELEMENT_Th:
		levels[5][3] -= 2
		levels[6][2] += 2
	}

	return
}

func LevelsFullString(l map[uint8][]uint8) (res string) {
	for n := uint8(1); int(n) <= len(l); n++ {
		_, ok := l[n]
		if !ok || int(n) != len(l[n]) || int(n)-1 > len(orbitals) {
			panic("invalid formula")
		}

		var delim string
		if n > 1 {
			delim = " "
		}

		for m := uint8(0); m < n; m++ {
			if l[n][m] > 0 {
				res += delim + strconv.Itoa(int(n)) +
					string(orbitals[m]) + "^" +
					strconv.Itoa(int(l[n][m]))
			}
		}

	}
	return
}

func LevelsShortString(l map[uint8][]uint8) (res string) {
	n := uint8(len(l))
	if n == 0 {
		return
	}

	var (
		m         = n - 1
		sublevstr string
	)

findsublevel:
	for {
		if l[n][m] > 0 {
			break
		}
		if m == 0 {
			goto nextlevel
		}
		m--
	}

	sublevstr = strconv.Itoa(int(n)) + string(orbitals[m]) + "^" +
		strconv.Itoa(int(l[n][m]))
	if len(res) > 0 {
		res = sublevstr + " " + res
	} else {
		res = sublevstr
	}

nextlevel:
	if n -= 1; n == 0 || len(l[n]) != int(n) {
		return
	}
	if l[n][n-1] == 4*n-2 {
		goto nextlevel
	}

	m = n - 1
	goto findsublevel
}
