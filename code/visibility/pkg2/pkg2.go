package pkg2

import "GoCasts/code/visibility/pkg1"

func ReadPkg1Var() (string, *int) {
	return pkg1.Name, pkg1.Age
}
