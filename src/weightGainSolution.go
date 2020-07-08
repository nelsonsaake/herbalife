package src

func GetWG4Pack() (pkg Package) {
	pkg = GetPackage([]int{
		NutritionalShakeMix,
		PersonalisedProteinPowder,
		Multivitamin,
		ThermojeticsCellActivator,
	})
	return
}

func GetWG3Pack() (pkg Package) {
	pkg = GetPackage([]int{
		NutritionalShakeMix,
		PersonalisedProteinPowder,
		Multivitamin,
	})
	return
}

func GetWG2Pack() (pkg Package) {
	pkg = GetPackage([]int{
		NutritionalShakeMix,
		PersonalisedProteinPowder,
	})
	return
}

func weightGainSolution() (sln Solution) {
	pkgs := []Package{
		GetWG4Pack(),
		GetWG3Pack(),
		GetWG2Pack(),
	}

	sln = Solution{
		Description: "Muscle/Weight Gain",
		Packages:    pkgs,
	}

	return
}
