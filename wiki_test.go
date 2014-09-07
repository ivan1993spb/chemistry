package chemistry

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

var srcUrl = `http://en.wikipedia.org/wiki/Electron_configurations_of_the_elements_(data_page)`

var elementList = []Element{
	ELEMENT_H,
	ELEMENT_He,
	ELEMENT_Li,
	ELEMENT_Be,
	ELEMENT_B,
	ELEMENT_C,
	ELEMENT_N,
	ELEMENT_O,
	ELEMENT_F,
	ELEMENT_Ne,
	ELEMENT_Na,
	ELEMENT_Mg,
	ELEMENT_Al,
	ELEMENT_Si,
	ELEMENT_P,
	ELEMENT_S,
	ELEMENT_Cl,
	ELEMENT_Ar,
	ELEMENT_K,
	ELEMENT_Ca,
	ELEMENT_Sc,
	ELEMENT_Ti,
	ELEMENT_V,
	ELEMENT_Cr,
	ELEMENT_Mn,
	ELEMENT_Fe,
	ELEMENT_Co,
	ELEMENT_Ni,
	ELEMENT_Cu,
	ELEMENT_Zn,
	ELEMENT_Ga,
	ELEMENT_Ge,
	ELEMENT_As,
	ELEMENT_Se,
	ELEMENT_Br,
	ELEMENT_Kr,
	ELEMENT_Rb,
	ELEMENT_Sr,
	ELEMENT_Y,
	ELEMENT_Zr,
	ELEMENT_Nb,
	ELEMENT_Mo,
	ELEMENT_Tc,
	ELEMENT_Ru,
	ELEMENT_Rh,
	ELEMENT_Pd,
	ELEMENT_Ag,
	ELEMENT_Cd,
	ELEMENT_In,
	ELEMENT_Sn,
	ELEMENT_Sb,
	ELEMENT_Te,
	ELEMENT_I,
	ELEMENT_Xe,
	ELEMENT_Cs,
	ELEMENT_Ba,
	ELEMENT_La,
	ELEMENT_Ce,
	ELEMENT_Pr,
	ELEMENT_Nd,
	ELEMENT_Pm,
	ELEMENT_Sm,
	ELEMENT_Eu,
	ELEMENT_Gd,
	ELEMENT_Tb,
	ELEMENT_Dy,
	ELEMENT_Ho,
	ELEMENT_Er,
	ELEMENT_Tm,
	ELEMENT_Yb,
	ELEMENT_Lu,
	ELEMENT_Hf,
	ELEMENT_Ta,
	ELEMENT_W,
	ELEMENT_Re,
	ELEMENT_Os,
	ELEMENT_Ir,
	ELEMENT_Pt,
	ELEMENT_Au,
	ELEMENT_Hg,
	ELEMENT_Tl,
	ELEMENT_Pb,
	ELEMENT_Bi,
	ELEMENT_Po,
	ELEMENT_At,
	ELEMENT_Rn,
	ELEMENT_Fr,
	ELEMENT_Ra,
	ELEMENT_Ac,
	ELEMENT_Th,
	ELEMENT_Pa,
	ELEMENT_U,
	ELEMENT_Np,
	ELEMENT_Pu,
	ELEMENT_Am,
	ELEMENT_Cm,
	ELEMENT_Bk,
	ELEMENT_Cf,
	ELEMENT_Es,
	ELEMENT_Fm,
	ELEMENT_Md,
	ELEMENT_No,
	ELEMENT_Lr,
	ELEMENT_Rf,
	ELEMENT_Db,
	ELEMENT_Sg,
	ELEMENT_Bh,
	ELEMENT_Hs,
	ELEMENT_Mt,
	ELEMENT_Ds,
	ELEMENT_Rg,
	ELEMENT_Cn,
	ELEMENT_Uut,
	ELEMENT_Fl,
	ELEMENT_Uup,
	ELEMENT_Lv,
	ELEMENT_Uus,
	ELEMENT_Uuo,
	ELEMENT_Uue,
	ELEMENT_Ubn,
	ELEMENT_Ubu,
	ELEMENT_Ubb,
	ELEMENT_Ubt,
	ELEMENT_Ubq,
	ELEMENT_Ubp,
	ELEMENT_Ubh,
	ELEMENT_Ubs,
	ELEMENT_Ubo,
	ELEMENT_Ube,
	ELEMENT_Utn,
	ELEMENT_Utu,
	ELEMENT_Utb,
	ELEMENT_Utt,
	ELEMENT_Utq,
	ELEMENT_Utp,
	ELEMENT_Uth,
	ELEMENT_Uts,
	ELEMENT_Uto,
	ELEMENT_Ute,
	ELEMENT_Uqn,
	ELEMENT_Uqu,
	ELEMENT_Uqb,
	ELEMENT_Uqt,
	ELEMENT_Uqq,
	ELEMENT_Uqp,
	ELEMENT_Uqh,
	ELEMENT_Uqs,
	ELEMENT_Uqo,
	ELEMENT_Uqe,
	ELEMENT_Upn,
	ELEMENT_Upu,
	ELEMENT_Upb,
	ELEMENT_Upt,
	ELEMENT_Upq,
	ELEMENT_Upp,
	ELEMENT_Uph,
	ELEMENT_Ups,
	ELEMENT_Upo,
	ELEMENT_Upe,
	ELEMENT_Uhn,
	ELEMENT_Uhu,
	ELEMENT_Uhb,
	ELEMENT_Uht,
	ELEMENT_Uhq,
	ELEMENT_Uhp,
	ELEMENT_Uhh,
	ELEMENT_Uhs,
	ELEMENT_Uho,
	ELEMENT_Uhe,
	ELEMENT_Usn,
	ELEMENT_Usu,
	ELEMENT_Usb,
}

func TestWikiComparing(t *testing.T) {
	fmt.Println("TestWikiComparing:")
	doc, err := goquery.NewDocument(srcUrl)
	if err != nil {
		t.Fatal(err)
	}

	for i, elem := range elementList {
		fmt.Printf("%4d. %s ...", i+1, elem)

		titleField := doc.Find(fmt.Sprintf(`th[colspan="26"] > a[title^=%q]`, elem.Name()))
		if titleField.Length() > 0 {

			electronConfWikiRow := titleField.Parent().Parent().Next()

			if electronConfWikiRow.Children().Length() == 26 {
				electronConfWiki := elem.ElectronConfiguration()
				for n, _ := range electronConfWiki {
					for l := 0; l < len(electronConfWiki[n]); l++ {
						electronConfWiki[n][l] = 0
					}
				}

				electronConfWikiRow.Children().Each(func(j int, s *goquery.Selection) {

					if len(s.Text()) > 1 {
						if i, err := strconv.Atoi(string(s.Text()[0])); err == nil {
							var n = uint8(i)
							var l uint8
							switch s.Text()[1] {
							case 's':
								l = 0
							case 'p':
								l = 1
							case 'd':
								l = 2
							case 'f':
								l = 3
							case 'g':
								l = 4
							case 'h':
								l = 5
							}
							i, err = strconv.Atoi(s.Find("sup").Text())
							if err == nil {

								electronConfWiki[n][l] = uint8(i)
							}
						}
					}
				})

				electronConfCalc := elem.ElectronConfiguration()

				for n := uint8(1); int(n) <= len(electronConfCalc); n++ {
					for l := uint8(0); l < n-1; l++ {
						if electronConfCalc[n][l] != electronConfWiki[n][l] {
							fmt.Printf("\ncalc: %s\nwiki: %s\n", electronConfCalc, electronConfWiki)
							goto fail
						}
					}
				}

				goto success

			} else {
				fmt.Println("bad field count")
			}
		} else {
			fmt.Println("not found")
		}

	fail:
		t.Fatal("fail")

	success:
		fmt.Println("ok")

	}
}
