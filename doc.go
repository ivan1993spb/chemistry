/*
Package chemistry contains chemistry elements and methods for
generation of objects describing of electron configuration of
elements.

For example you can get electron configuration scheme as map:

	var electron_configuration = chemistry.ELEMENT_Kr.Levels()

And then print it in two forms:

	// Full
	fmt.Printf("for %q electron configuration is %q\n",
		chemistry.ELEMENT_Kr,
		chemistry.LevelsFullString(electron_configuration))

	// Short
	fmt.Printf("short form: %q\n",
		chemistry.LevelsShortString(electron_configuration))

*/
package chemistry
