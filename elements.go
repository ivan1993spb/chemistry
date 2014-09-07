package chemistry

/*

References:
----------

    1. https://en.wikipedia.org/wiki/Extended_periodic_table
    2. http://en.wikipedia.org/wiki/Electron_configurations_of_the_elements_(data_page)

*/

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
	ELEMENT_Ubs
	ELEMENT_Ubo
	ELEMENT_Ube
	ELEMENT_Utn
	ELEMENT_Utu
	ELEMENT_Utb
	ELEMENT_Utt
	ELEMENT_Utq
	ELEMENT_Utp
	ELEMENT_Uth
	ELEMENT_Uts
	ELEMENT_Uto
	ELEMENT_Ute
	ELEMENT_Uqn
	ELEMENT_Uqu
	ELEMENT_Uqb
	ELEMENT_Uqt
	ELEMENT_Uqq
	ELEMENT_Uqp
	ELEMENT_Uqh
	ELEMENT_Uqs
	ELEMENT_Uqo
	ELEMENT_Uqe
	ELEMENT_Upn
	ELEMENT_Upu
	ELEMENT_Upb
	ELEMENT_Upt
	ELEMENT_Upq
	ELEMENT_Upp
	ELEMENT_Uph
	ELEMENT_Ups
	ELEMENT_Upo
	ELEMENT_Upe
	ELEMENT_Uhn
	ELEMENT_Uhu
	ELEMENT_Uhb
	ELEMENT_Uht
	ELEMENT_Uhq
	ELEMENT_Uhp
	ELEMENT_Uhh
	ELEMENT_Uhs
	ELEMENT_Uho
	ELEMENT_Uhe
	ELEMENT_Usn
	ELEMENT_Usu
	ELEMENT_Usb
	ELEMENT_Uoq Element = 184
)

var names = map[Element][2]string{
	ELEMENT_H:   {"H", "Hydrogen"},
	ELEMENT_He:  {"He", "Helium"},
	ELEMENT_Li:  {"Li", "Lithium"},
	ELEMENT_Be:  {"Be", "Beryllium"},
	ELEMENT_B:   {"B", "Boron"},
	ELEMENT_C:   {"C", "Carbon"},
	ELEMENT_N:   {"N", "Nitrogen"},
	ELEMENT_O:   {"O", "Oxygen"},
	ELEMENT_F:   {"F", "Fluorine"},
	ELEMENT_Ne:  {"Ne", "Neon"},
	ELEMENT_Na:  {"Na", "Sodium"},
	ELEMENT_Mg:  {"Mg", "Magnesium"},
	ELEMENT_Al:  {"Al", "Aluminium"},
	ELEMENT_Si:  {"Si", "Silicon"},
	ELEMENT_P:   {"P", "Phosphorus"},
	ELEMENT_S:   {"S", "Sulfur"},
	ELEMENT_Cl:  {"Cl", "Chlorine"},
	ELEMENT_Ar:  {"Ar", "Argon"},
	ELEMENT_K:   {"K", "Potassium"},
	ELEMENT_Ca:  {"Ca", "Calcium"},
	ELEMENT_Sc:  {"Sc", "Scandium"},
	ELEMENT_Ti:  {"Ti", "Titanium"},
	ELEMENT_V:   {"V", "Vanadium"},
	ELEMENT_Cr:  {"Cr", "Chromium"},
	ELEMENT_Mn:  {"Mn", "Manganese"},
	ELEMENT_Fe:  {"Fe", "Iron"},
	ELEMENT_Co:  {"Co", "Cobalt"},
	ELEMENT_Ni:  {"Ni", "Nickel"},
	ELEMENT_Cu:  {"Cu", "Copper"},
	ELEMENT_Zn:  {"Zn", "Zinc"},
	ELEMENT_Ga:  {"Ga", "Gallium"},
	ELEMENT_Ge:  {"Ge", "Germanium"},
	ELEMENT_As:  {"As", "Arsenic"},
	ELEMENT_Se:  {"Se", "Selenium"},
	ELEMENT_Br:  {"Br", "Bromine"},
	ELEMENT_Kr:  {"Kr", "Krypton"},
	ELEMENT_Rb:  {"Rb", "Rubidium"},
	ELEMENT_Sr:  {"Sr", "Strontium"},
	ELEMENT_Y:   {"Y", "Yttrium"},
	ELEMENT_Zr:  {"Zr", "Zirconium"},
	ELEMENT_Nb:  {"Nb", "Niobium"},
	ELEMENT_Mo:  {"Mo", "Molybdenum"},
	ELEMENT_Tc:  {"Tc", "Technetium"},
	ELEMENT_Ru:  {"Ru", "Ruthenium"},
	ELEMENT_Rh:  {"Rh", "Rhodium"},
	ELEMENT_Pd:  {"Pd", "Palladium"},
	ELEMENT_Ag:  {"Ag", "Silver"},
	ELEMENT_Cd:  {"Cd", "Cadmium"},
	ELEMENT_In:  {"In", "Indium"},
	ELEMENT_Sn:  {"Sn", "Tin"},
	ELEMENT_Sb:  {"Sb", "Antimony"},
	ELEMENT_Te:  {"Te", "Tellurium"},
	ELEMENT_I:   {"I", "Iodine"},
	ELEMENT_Xe:  {"Xe", "Xenon"},
	ELEMENT_Cs:  {"Cs", "Caesium"},
	ELEMENT_Ba:  {"Ba", "Barium"},
	ELEMENT_La:  {"La", "Lanthanum"},
	ELEMENT_Ce:  {"Ce", "Cerium"},
	ELEMENT_Pr:  {"Pr", "Praseodymium"},
	ELEMENT_Nd:  {"Nd", "Neodymium"},
	ELEMENT_Pm:  {"Pm", "Promethium"},
	ELEMENT_Sm:  {"Sm", "Samarium"},
	ELEMENT_Eu:  {"Eu", "Europium"},
	ELEMENT_Gd:  {"Gd", "Gadolinium"},
	ELEMENT_Tb:  {"Tb", "Terbium"},
	ELEMENT_Dy:  {"Dy", "Dysprosium"},
	ELEMENT_Ho:  {"Ho", "Holmium"},
	ELEMENT_Er:  {"Er", "Erbium"},
	ELEMENT_Tm:  {"Tm", "Thulium"},
	ELEMENT_Yb:  {"Yb", "Ytterbium"},
	ELEMENT_Lu:  {"Lu", "Lutetium"},
	ELEMENT_Hf:  {"Hf", "Hafnium"},
	ELEMENT_Ta:  {"Ta", "Tantalum"},
	ELEMENT_W:   {"W", "Tungsten"},
	ELEMENT_Re:  {"Re", "Rhenium"},
	ELEMENT_Os:  {"Os", "Osmium"},
	ELEMENT_Ir:  {"Ir", "Iridium"},
	ELEMENT_Pt:  {"Pt", "Platinum"},
	ELEMENT_Au:  {"Au", "Gold"},
	ELEMENT_Hg:  {"Hg", "Mercury"},
	ELEMENT_Tl:  {"Tl", "Thallium"},
	ELEMENT_Pb:  {"Pb", "Lead"},
	ELEMENT_Bi:  {"Bi", "Bismuth"},
	ELEMENT_Po:  {"Po", "Polonium"},
	ELEMENT_At:  {"At", "Astatine"},
	ELEMENT_Rn:  {"Rn", "Radon"},
	ELEMENT_Fr:  {"Fr", "Francium"},
	ELEMENT_Ra:  {"Ra", "Radium"},
	ELEMENT_Ac:  {"Ac", "Actinium"},
	ELEMENT_Th:  {"Th", "Thorium"},
	ELEMENT_Pa:  {"Pa", "Protactinium"},
	ELEMENT_U:   {"U", "Uranium"},
	ELEMENT_Np:  {"Np", "Neptunium"},
	ELEMENT_Pu:  {"Pu", "Plutonium"},
	ELEMENT_Am:  {"Am", "Americium"},
	ELEMENT_Cm:  {"Cm", "Curium"},
	ELEMENT_Bk:  {"Bk", "Berkelium"},
	ELEMENT_Cf:  {"Cf", "Californium"},
	ELEMENT_Es:  {"Es", "Einsteinium"},
	ELEMENT_Fm:  {"Fm", "Fermium"},
	ELEMENT_Md:  {"Md", "Mendelevium"},
	ELEMENT_No:  {"No", "Nobelium"},
	ELEMENT_Lr:  {"Lr", "Lawrencium"},
	ELEMENT_Rf:  {"Rf", "Rutherfordium"},
	ELEMENT_Db:  {"Db", "Dubnium"},
	ELEMENT_Sg:  {"Sg", "Seaborgium"},
	ELEMENT_Bh:  {"Bh", "Bohrium"},
	ELEMENT_Hs:  {"Hs", "Hassium"},
	ELEMENT_Mt:  {"Mt", "Meitnerium"},
	ELEMENT_Ds:  {"Ds", "Darmstadtium"},
	ELEMENT_Rg:  {"Rg", "Roentgenium"},
	ELEMENT_Cn:  {"Cn", "Copernicium"},
	ELEMENT_Uut: {"Uut", "Ununtrium"},
	ELEMENT_Fl:  {"Fl", "Flerovium"},
	ELEMENT_Uup: {"Uup", "Ununpentium"},
	ELEMENT_Lv:  {"Lv", "Livermorium"},
	ELEMENT_Uus: {"Uus", "Ununseptium"},
	ELEMENT_Uuo: {"Uuo", "Ununoctium"},
	ELEMENT_Uue: {"Uue", "Ununennium"},
	ELEMENT_Ubn: {"Ubn", "Unbinilium"},
	ELEMENT_Ubu: {"Ubu", "Unbiunium"},
	ELEMENT_Ubb: {"Ubb", "Unbibium"},
	ELEMENT_Ubt: {"Ubt", "Unbitrium"},
	ELEMENT_Ubq: {"Ubq", "Unbiquadium"},
	ELEMENT_Ubp: {"Ubp", "Unbipentium"},
	ELEMENT_Ubh: {"Ubh", "Unbihexium"},
	ELEMENT_Ubs: {"Ubs", "Unbiseptium"},
	ELEMENT_Ubo: {"Ubo", "Unbioctium"},
	ELEMENT_Ube: {"Ube", "Unbiennium"},
	ELEMENT_Utn: {"Utn", "Untrinilium"},
	ELEMENT_Utu: {"Utu", "Untriunium"},
	ELEMENT_Utb: {"Utb", "Untribium"},
	ELEMENT_Utt: {"Utt", "Untritrium"},
	ELEMENT_Utq: {"Utq", "Untriquadium"},
	ELEMENT_Utp: {"Utp", "Untripentium"},
	ELEMENT_Uth: {"Uth", "Untrihexium"},
	ELEMENT_Uts: {"Uts", "Untriseptium"},
	ELEMENT_Uto: {"Uto", "Untrioctium"},
	ELEMENT_Ute: {"Ute", "Untriennium"},
	ELEMENT_Uqn: {"Uqn", "Unquadnilium"},
	ELEMENT_Uqu: {"Uqu", "Unquadunium"},
	ELEMENT_Uqb: {"Uqb", "Unquadbium"},
	ELEMENT_Uqt: {"Uqt", "Unquadtrium"},
	ELEMENT_Uqq: {"Uqq", "Unquadquadium"},
	ELEMENT_Uqp: {"Uqp", "Unquadpentium"},
	ELEMENT_Uqh: {"Uqh", "Unquadhexium"},
	ELEMENT_Uqs: {"Uqs", "Unquadseptium"},
	ELEMENT_Uqo: {"Uqo", "Unquadoctium"},
	ELEMENT_Uqe: {"Uqe", "Unquadennium"},
	ELEMENT_Upn: {"Upn", "Unpentnilium"},
	ELEMENT_Upu: {"Upu", "Unpentunium"},
	ELEMENT_Upb: {"Upb", "Unpentbium"},
	ELEMENT_Upt: {"Upt", "Unpenttrium"},
	ELEMENT_Upq: {"Upq", "Unpentquadium"},
	ELEMENT_Upp: {"Upp", "Unpentpentium"},
	ELEMENT_Uph: {"Uph", "Unpenthexium"},
	ELEMENT_Ups: {"Ups", "Unpentseptium"},
	ELEMENT_Upo: {"Upo", "Unpentoctium"},
	ELEMENT_Upe: {"Upe", "Unpentennium"},
	ELEMENT_Uhn: {"Uhn", "Unhexnilium"},
	ELEMENT_Uhu: {"Uhu", "Unhexunium"},
	ELEMENT_Uhb: {"Uhb", "Unhexbium"},
	ELEMENT_Uht: {"Uht", "Unhextrium"},
	ELEMENT_Uhq: {"Uhq", "Unhexquadium"},
	ELEMENT_Uhp: {"Uhp", "Unhexpentium"},
	ELEMENT_Uhh: {"Uhh", "Unhexhexium"},
	ELEMENT_Uhs: {"Uhs", "Unhexseptium"},
	ELEMENT_Uho: {"Uho", "Unhexoctium"},
	ELEMENT_Uhe: {"Uhe", "Unhexennium"},
	ELEMENT_Usn: {"Usn", "Unseptnilium"},
	ELEMENT_Usu: {"Usu", "Unseptunium"},
	ELEMENT_Usb: {"Usb", "Unseptbium"},
	ELEMENT_Uoq: {"Uoq", "Unoctquadium"},
}

var weights = map[Element]float32{
	ELEMENT_H:  1.00794,
	ELEMENT_He: 4.002602,
	ELEMENT_Li: 6.941,
	ELEMENT_Be: 9.012182,
	ELEMENT_B:  10.811,
	ELEMENT_C:  12.0107,
	ELEMENT_N:  14.0067,
	ELEMENT_O:  15.9994,
	ELEMENT_F:  18.9984032,
	ELEMENT_Ne: 20.1797,
	ELEMENT_Na: 22.989770,
	ELEMENT_Mg: 24.3050,
	ELEMENT_Al: 26.981538,
	ELEMENT_Si: 28.0855,
	ELEMENT_P:  30.973761,
	ELEMENT_S:  32.065,
	ELEMENT_Cl: 35.453,
	ELEMENT_Ar: 39.948,
	ELEMENT_K:  39.0983,
	ELEMENT_Ca: 40.078,
	ELEMENT_Sc: 44.955910,
	ELEMENT_Ti: 47.867,
	ELEMENT_V:  50.9415,
	ELEMENT_Cr: 51.9961,
	ELEMENT_Mn: 54.938049,
	ELEMENT_Fe: 55.845,
	ELEMENT_Co: 58.933200,
	ELEMENT_Ni: 58.6934,
	ELEMENT_Cu: 63.546,
	ELEMENT_Zn: 65.409,
	ELEMENT_Ga: 69.723,
	ELEMENT_Ge: 72.64,
	ELEMENT_As: 74.92160,
	ELEMENT_Se: 78.96,
	ELEMENT_Br: 79.904,
	ELEMENT_Kr: 83.798,
	ELEMENT_Rb: 85.4678,
	ELEMENT_Sr: 87.62,
	ELEMENT_Y:  88.90585,
	ELEMENT_Zr: 91.224,
	ELEMENT_Nb: 92.90638,
	ELEMENT_Mo: 95.94,
	ELEMENT_Tc: 97.9072,
	ELEMENT_Ru: 101.07,
	ELEMENT_Rh: 102.90550,
	ELEMENT_Pd: 106.42,
	ELEMENT_Ag: 107.8682,
	ELEMENT_Cd: 112.411,
	ELEMENT_In: 114.818,
	ELEMENT_Sn: 118.710,
	ELEMENT_Sb: 121.760,
	ELEMENT_Te: 127.60,
	ELEMENT_I:  126.90447,
	ELEMENT_Xe: 131.293,
	ELEMENT_Cs: 132.90545,
	ELEMENT_Ba: 137.327,
	ELEMENT_La: 138.9055,
	ELEMENT_Ce: 140.116,
	ELEMENT_Pr: 140.90765,
	ELEMENT_Nd: 144.24,
	ELEMENT_Pm: 144.9127,
	ELEMENT_Sm: 150.36,
	ELEMENT_Eu: 151.964,
	ELEMENT_Gd: 157.25,
	ELEMENT_Tb: 158.92534,
	ELEMENT_Dy: 162.500,
	ELEMENT_Ho: 164.93032,
	ELEMENT_Er: 167.259,
	ELEMENT_Tm: 168.93421,
	ELEMENT_Yb: 173.054,
	ELEMENT_Lu: 174.967,
	ELEMENT_Hf: 178.49,
	ELEMENT_Ta: 180.9479,
	ELEMENT_W:  183.84,
	ELEMENT_Re: 186.207,
	ELEMENT_Os: 190.23,
	ELEMENT_Ir: 192.217,
	ELEMENT_Pt: 195.078,
	ELEMENT_Au: 196.96655,
	ELEMENT_Hg: 200.59,
	ELEMENT_Tl: 204.3833,
	ELEMENT_Pb: 207.2,
	ELEMENT_Bi: 208.98038,
	ELEMENT_Po: 208.9824,
	ELEMENT_At: 209.9871,
	ELEMENT_Rn: 222.0176,
	ELEMENT_Fr: 223.0197,
	ELEMENT_Ra: 226.0254,
	ELEMENT_Ac: 227.0278,
	ELEMENT_Th: 232.0381,
	ELEMENT_Pa: 231.03588,
	ELEMENT_U:  238.02891,
	ELEMENT_Np: 237.048,
	ELEMENT_Pu: 244.0642,
	ELEMENT_Am: 243.0614,
	ELEMENT_Cm: 247.0703,
	ELEMENT_Bk: 247.0703,
	ELEMENT_Cf: 251.0796,
	ELEMENT_Es: 252.083,
	ELEMENT_Fm: 257.0951,
	ELEMENT_Md: 258.1,
	ELEMENT_No: 259.1009,
	ELEMENT_Lr: 262.0,
	ELEMENT_Rf: 261.0,
	ELEMENT_Db: 262.0,
	ELEMENT_Sg: 266.0,
	ELEMENT_Hs: 264.0,
	ELEMENT_Bh: 277.0,
	ELEMENT_Mt: 268.0,
}
