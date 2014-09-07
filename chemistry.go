package chemistry

import "strconv"

// AtomicNumber is the number of protons found in the nucleus of an
// atom of that element
type AtomicNumber byte

func (n AtomicNumber) String() string {
	return strconv.Itoa(int(n)) + "e"
}

type ErrInvalidElectronConfiguration uint8

func (e ErrInvalidElectronConfiguration) Error() string {
	return "Invalid electron configuration on shell " +
		strconv.Itoa(int(e))
}

type Orbital uint8

const (
	ORBITAL_S Orbital = iota
	ORBITAL_P
	ORBITAL_D
	ORBITAL_F
	ORBITAL_G
	ORBITAL_H
	ORBITAL_I
	ORBITAL_J
	ORBITAL_K
)

var orbitals = "spdfghijk"

func (o Orbital) String() string {

	if int(o) < len(orbitals) {
		return string(orbitals[o])
	}
	return "[orbital " + strconv.Itoa(int(o)) + "]"
}

// Element is chemical element
type Element byte

func (e Element) String() string {
	if _, ok := names[e]; ok {
		return names[e][0] + "(" + strconv.Itoa(int(e)) + ") " +
			names[e][1] + " " + e.ElectronConfiguration().Short()
	}
	return ""
}

// Symbol returns element symbol
func (e Element) Symbol() string {
	if _, ok := names[e]; ok {
		return names[e][0]
	}
	return ""
}

// Name returns full element name
func (e Element) Name() string {
	if _, ok := names[e]; ok {
		return names[e][1]
	}
	return ""
}

// Weight returns atomic weight of element
func (e Element) Weight() float32 {
	if _, ok := weights[e]; ok {
		return weights[e]
	}
	return 0
}

// AtomicNumber returns number of protons found in the nucleus
func (e Element) AtomicNumber() AtomicNumber {
	return AtomicNumber(e)
}

// ElectronConfiguration calculates atomic electron configuration
func (e Element) ElectronConfiguration() ElectronConfiguration {
	var (
		n     uint8 = 1        // principal quantum number (n > 0)
		l     uint8 = 0        // azimuthal quantum number (n - 1)
		count       = uint8(e) // count of electron
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

	// create level 9 for elements wich has electron skipping to 9 lev
	switch e {
	case ELEMENT_Upe, ELEMENT_Uhn, ELEMENT_Uhu, ELEMENT_Uhp,
		ELEMENT_Uhh, ELEMENT_Uhs, ELEMENT_Uho:
		ec[9] = make([]uint8, 9)
	}

	// electron skipping
	switch e {
	case ELEMENT_Cr, ELEMENT_Cu:
		ec[4][ORBITAL_S] -= 1
		ec[3][ORBITAL_D] += 1

	case ELEMENT_Nb, ELEMENT_Mo, ELEMENT_Ru, ELEMENT_Rh, ELEMENT_Ag:
		ec[5][ORBITAL_S] -= 1
		ec[4][ORBITAL_D] += 1

	case ELEMENT_Pd:
		ec[5][ORBITAL_S] -= 2
		ec[4][ORBITAL_D] += 2

	case ELEMENT_Pt, ELEMENT_Au:
		ec[6][ORBITAL_S] -= 1
		ec[5][ORBITAL_D] += 1

	case ELEMENT_La, ELEMENT_Ce, ELEMENT_Gd:
		ec[4][ORBITAL_F] -= 1
		ec[5][ORBITAL_D] += 1

	case ELEMENT_Ac, ELEMENT_Pa, ELEMENT_U, ELEMENT_Np, ELEMENT_Cm:
		ec[5][ORBITAL_F] -= 1
		ec[6][ORBITAL_D] += 1

	case ELEMENT_Th:
		ec[5][ORBITAL_F] -= 2
		ec[6][ORBITAL_D] += 2

	case ELEMENT_Lr:
		ec[6][ORBITAL_D] -= 1
		ec[7][ORBITAL_P] += 1

	case ELEMENT_Ubu:
		ec[5][ORBITAL_G] -= 1
		ec[8][ORBITAL_P] += 1

	case ELEMENT_Ubb:
		ec[5][ORBITAL_G] -= 2
		ec[7][ORBITAL_D] += 1
		ec[8][ORBITAL_P] += 1

	case ELEMENT_Ubt:
		ec[5][ORBITAL_G] -= 3
		ec[6][ORBITAL_F] += 1
		ec[7][ORBITAL_D] += 1
		ec[8][ORBITAL_P] += 1

	case ELEMENT_Ubq, ELEMENT_Ubp:
		ec[5][ORBITAL_G] -= 4
		ec[6][ORBITAL_F] += 3
		ec[8][ORBITAL_P] += 1

	case ELEMENT_Ubh:
		ec[5][ORBITAL_G] -= 4
		ec[6][ORBITAL_F] += 2
		ec[7][ORBITAL_D] += 1
		ec[8][ORBITAL_P] += 1

	case ELEMENT_Ubs, ELEMENT_Ubo, ELEMENT_Ube, ELEMENT_Utn,
		ELEMENT_Utu, ELEMENT_Utb:
		ec[5][ORBITAL_G] -= 4
		ec[6][ORBITAL_F] += 2
		ec[8][ORBITAL_P] += 2

	case ELEMENT_Utt:
		ec[5][ORBITAL_G] -= 5
		ec[6][ORBITAL_F] += 3
		ec[8][ORBITAL_P] += 2

	case ELEMENT_Utq, ELEMENT_Utp, ELEMENT_Uth:
		ec[5][ORBITAL_G] -= 6
		ec[6][ORBITAL_F] += 4
		ec[8][ORBITAL_P] += 2

	case ELEMENT_Uts, ELEMENT_Uto:
		ec[5][ORBITAL_G] -= 6
		ec[6][ORBITAL_F] += 3
		ec[7][ORBITAL_D] += 1
		ec[8][ORBITAL_P] += 2

	case ELEMENT_Ute:
		ec[5][ORBITAL_G] -= 5
		ec[6][ORBITAL_F] += 2
		ec[7][ORBITAL_D] += 1
		ec[8][ORBITAL_P] += 2

	case ELEMENT_Uqn:
		ec[5][ORBITAL_G] -= 4
		ec[6][ORBITAL_F] += 1
		ec[7][ORBITAL_D] += 1
		ec[8][ORBITAL_P] += 2

	case ELEMENT_Uqu:
		ec[5][ORBITAL_G] -= 3
		ec[6][ORBITAL_F] -= 1
		ec[7][ORBITAL_D] += 2
		ec[8][ORBITAL_P] += 2

	case ELEMENT_Uqb:
		ec[5][ORBITAL_G] -= 2
		ec[6][ORBITAL_F] -= 2
		ec[7][ORBITAL_D] += 2
		ec[8][ORBITAL_P] += 2

	case ELEMENT_Uqt:
		ec[5][ORBITAL_G] -= 1
		ec[6][ORBITAL_F] -= 3
		ec[7][ORBITAL_D] += 2
		ec[8][ORBITAL_P] += 2

	case ELEMENT_Uqq, ELEMENT_Uqe:
		ec[6][ORBITAL_F] -= 5
		ec[7][ORBITAL_D] += 3
		ec[8][ORBITAL_P] += 2

	case ELEMENT_Uqp, ELEMENT_Uqh, ELEMENT_Uqs, ELEMENT_Uqo:
		ec[6][ORBITAL_F] -= 4
		ec[7][ORBITAL_D] += 2
		ec[8][ORBITAL_P] += 2

	case ELEMENT_Upn:
		ec[6][ORBITAL_F] -= 6
		ec[7][ORBITAL_D] += 4
		ec[8][ORBITAL_P] += 2

	case ELEMENT_Upu, ELEMENT_Upb:
		ec[6][ORBITAL_F] -= 5
		ec[7][ORBITAL_D] += 3
		ec[8][ORBITAL_P] += 2

	case ELEMENT_Upt:
		ec[6][ORBITAL_F] -= 3
		ec[7][ORBITAL_D] += 1
		ec[8][ORBITAL_P] += 2

	case ELEMENT_Upq:
		ec[6][ORBITAL_F] -= 2
		ec[8][ORBITAL_P] += 2

	case ELEMENT_Upp:
		ec[6][ORBITAL_F] -= 1
		ec[7][ORBITAL_D] -= 1
		ec[8][ORBITAL_P] += 2

	case ELEMENT_Uph, ELEMENT_Ups, ELEMENT_Upo:
		ec[7][ORBITAL_D] -= 2
		ec[8][ORBITAL_P] += 2

	case ELEMENT_Uhb:
		ec[7][ORBITAL_D] -= 2
		ec[8][ORBITAL_P] += 2

	case ELEMENT_Uht:
		ec[7][ORBITAL_D] -= 1
		ec[8][ORBITAL_P] += 1

	case ELEMENT_Uhp:
		ec[8][ORBITAL_P] -= 1
		ec[9][ORBITAL_S] += 1

	case ELEMENT_Uhh:
		ec[8][ORBITAL_P] -= 2
		ec[9][ORBITAL_S] += 2

	case ELEMENT_Upe, ELEMENT_Uhn, ELEMENT_Uhu:
		ec[7][ORBITAL_D] -= 3
		ec[8][ORBITAL_P] += 2
		ec[9][ORBITAL_S] += 1

	case ELEMENT_Uhs:
		ec[8][ORBITAL_P] -= 3
		ec[9][ORBITAL_S] += 2
		ec[9][ORBITAL_P] += 1

	case ELEMENT_Uho:
		ec[8][ORBITAL_P] -= 4
		ec[9][ORBITAL_S] += 2
		ec[9][ORBITAL_P] += 2

	case ELEMENT_Uhe:
		ec[8][ORBITAL_P] -= 3
		ec[9][ORBITAL_S] += 1
		ec[9][ORBITAL_P] += 2

	case ELEMENT_Usn:
		ec[8][ORBITAL_P] -= 2
		ec[9][ORBITAL_P] += 2

	case ELEMENT_Usu:
		ec[6][ORBITAL_G] -= 1
		ec[8][ORBITAL_P] -= 1
		ec[9][ORBITAL_P] += 2

	case ELEMENT_Usb:
		ec[6][ORBITAL_G] -= 2
		ec[9][ORBITAL_P] += 2
	}

	return ec
}

// ElectronConfiguration object describes distribution of electrons
// of an atom in atomic orbitals
type ElectronConfiguration map[uint8][]uint8

func (ec ElectronConfiguration) String() string {
	var (
		n   uint8 = 1
		str string
	)

	for int(n) <= len(ec) {
		_, ok := ec[n]

		if !ok || int(n) != len(ec[n]) {
			panic(ErrInvalidElectronConfiguration(n))
		}

		var delim string
		if n > 1 {
			delim = " "
		}

		var l uint8
		for l < n {
			if ec[n][l] > 0 {
				str += delim + strconv.Itoa(int(n)) + Orbital(l).String() +
					"^" + strconv.Itoa(int(ec[n][l]))
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

	sublevstr = strconv.Itoa(int(n)) + Orbital(l).String() + "^" +
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
