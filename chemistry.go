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

func (e Element) String() string {
	if _, ok := names[e]; ok {
		return names[e][0] + "(" + strconv.Itoa(int(e)) + ") " +
			names[e][1] + " " + e.ElectronConfiguration().Short()
	}
	return ""
}

// Symbol returns symbol
func (e Element) Symbol() string {
	if _, ok := names[e]; ok {
		return names[e][0]
	}
	return ""
}

// Name returns full name
func (e Element) Name() string {
	if _, ok := names[e]; ok {
		return names[e][1]
	}
	return ""
}

// Weight returns atomic weight
func (e Element) Weight() float32 {
	if _, ok := weights[e]; ok {
		return weights[e]
	}
	return 0
}

// AtomicNumber returns element atomic number
func (e Element) AtomicNumber() AtomicNumber {
	return AtomicNumber(e)
}

// ElectronConfiguration calculates atomic electron configuration
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

var orbitals = "spdfghigk"

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
