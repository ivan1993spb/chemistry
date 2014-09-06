/*
Package chemistry contains chemistry elements and methods for
generation of objects describing of electron configuration of
elements.

For example you can get electron configuration scheme as map:

	var electron_configuration = chemistry.ELEMENT_Kr.ElectronConfiguration()

And then print it in two forms:

	// Full
	fmt.Printf("for %q electron configuration is %q\n",
		chemistry.ELEMENT_Kr,
		chemistry.ELEMENT_Kr.ElectronConfiguration())

	// Short
	fmt.Printf("short form: %q\n",
		chemistry.ELEMENT_Kr.ElectronConfiguration().Short())

	// Electron count per shell
	fmt.Println(chemistry.ELEMENT_Kr.ElectronConfiguration().EnergyPerShell())

*/
package chemistry
