package chemistry

import "strconv"

type AtomicNumber byte

func (n AtomicNumber) String() string {
	return strconv.Itoa(int(n)) + "e"
}

type ErrInvalidElectronConfiguration uint8

func (e ErrInvalidElectronConfiguration) Error() string {
	return "Invalid electron configuration on shell " +
		strconv.Itoa(int(e))
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
	if _, ok := symbols[e]; ok {
		return symbols[e] + "(" + strconv.Itoa(int(e)) + ") " +
			e.ElectronConfiguration().Short()
	}
	return ""
}

func (e Element) AtomicNumber() AtomicNumber {
	return AtomicNumber(e)
}

func (e Element) ElectronConfiguration() ElectronConfiguration {
	var (
		n     uint8 = 1        // principal quantum number (n > 0)
		l     uint8 = 0        // azimuthal quantum number (n - 1)
		count       = uint8(e) // electron count
	)

	ec := make(ElectronConfiguration)

level:
	{
		// create energy sublevel slice
		if len(ec[n]) == 0 {
			ec[n] = make([]uint8, n)
		}
	}

sublevel:
	{
		sublevelLength := 2 * (2*l + 1)
		if count > sublevelLength {
			ec[n][l] += sublevelLength
			count -= sublevelLength
		} else {
			ec[n][l] += count
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
		ec[4][0] -= 1
		ec[3][2] += 1

	case ELEMENT_Nb, ELEMENT_Mo, ELEMENT_Ru, ELEMENT_Rh, ELEMENT_Ag:
		ec[5][0] -= 1
		ec[4][2] += 1

	case ELEMENT_Pd:
		ec[5][0] -= 2
		ec[4][2] += 2

	case ELEMENT_Pt, ELEMENT_Au:
		ec[6][0] -= 1
		ec[5][2] += 1

	case ELEMENT_Ds, ELEMENT_Rg:
		ec[7][0] -= 1
		ec[6][2] += 1

	case ELEMENT_La, ELEMENT_Gd:
		ec[4][3] -= 1
		ec[5][2] += 1

	case ELEMENT_Ac, ELEMENT_Pa, ELEMENT_U, ELEMENT_Np, ELEMENT_Cm:
		ec[5][3] -= 1
		ec[6][2] += 1

	case ELEMENT_Th:
		ec[5][3] -= 2
		ec[6][2] += 2
	}

	return ec
}

type ElectronConfiguration map[uint8][]uint8

func (ec ElectronConfiguration) String() string {
	var (
		n   uint8 = 1
		str string
	)

	for int(n) <= len(ec) {
		_, ok := ec[n]

		if !ok || int(n) != len(ec[n]) || int(n)-1 > len(orbitals) {
			panic(ErrInvalidElectronConfiguration(n))
		}

		var delim string
		if n > 1 {
			delim = " "
		}

		var l uint8
		for l < n {
			if ec[n][l] > 0 {
				str += delim + strconv.Itoa(int(n)) +
					string(orbitals[l]) + "^" +
					strconv.Itoa(int(ec[n][l]))
			}
			l += 1
		}

		n += 1

	}

	return str
}

func (ec ElectronConfiguration) Short() (res string) {
	var n uint8
	if n = uint8(len(ec)); n == 0 {
		return
	}

	var (
		l         uint8
		sublevstr string
	)

findsublevel:
	l = n - 1
	for {
		if ec[n][l] > 0 {
			break
		}
		if l == 0 {
			goto nextlevel
		}
		l -= 1
	}

	sublevstr = strconv.Itoa(int(n)) + string(orbitals[l]) + "^" +
		strconv.Itoa(int(ec[n][l]))
	if len(res) > 0 {
		res = sublevstr + " " + res
	} else {
		res = sublevstr
	}

nextlevel:
	if n -= 1; n == 0 {
		return
	}
	if int(n) != len(ec[n]) {
		panic(ErrInvalidElectronConfiguration(n))
	}
	if ec[n][n-1] == 4*n-2 {
		goto nextlevel
	}

	goto findsublevel
}

func (ec ElectronConfiguration) EnergyPerShell() []uint8 {
	var (
		levels = make([]uint8, 0)
		n      uint8
	)

	for n = 1; int(n) <= len(ec); n++ {
		if int(n) != len(ec[n]) {
			panic(ErrInvalidElectronConfiguration(n))
		}

		var (
			energy uint8
			l      uint8
		)

		for int(l) < len(ec[n]) {
			energy += ec[n][l]
			l += 1
		}
		levels = append(levels, energy)
	}

	return levels
}

// func Valence(l map[uint8][]uint8) (v uint8) {
// fmt.Println(LevelsFullString(l))
// fmt.Println(LevelsShortString(l))
// n := uint8(len(l))
// if n == 0 {
// 	return 0
// }

// var (
// 	m           = n
// 	count uint8 = 0
// )

// for m > 0 {
// 	m -= 1
// 	count += l[n][m] % 2
// }

// fmt.Println(n, 4*n-2, count)
// return 0
// if n := uint8(len(l)); n > 0 {
// 	if l[n][n-1] != 4*n-2 {
// 		var m = n - 1
// 		for {
// 			if l[n][m] > 0 {
// 				break
// 			}
// 			if m == 0 {
// 				break
// 			}
// 			m--
// 		}
// 		l[n][m] -= 1
// 		l[n][m+1] += 1
// 	}
// }
// 	for n := uint8(len(l)); n > 0; n-- {
// 		var m = n
// 		for m > 0 {
// 			m -= 1
// 			if 2*(2*m+1) != l[n][m] {
// 				v += l[n][m]
// 			}
// 		}
// 	}

// 	return
// }
