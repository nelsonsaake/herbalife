package src

type Solution struct {
	Description string
	Packages    []Package
}

func GetSolution(pkgsIds [][]int) (sln Solution) {
	pkgs := []Package{}

	for _, ids := range pkgsIds {
		pkg := GetPackage(ids)
		pkgs = append(pkgs, pkg)
	}

	sln = Solution{Packages: pkgs}
	return
}
