package config

type Subbranchformat struct {

	// defining struct variables
	Name string
}
type Branchformat struct {

	// defining struct variables
	Name     string
	Branches []Subbranchformat
}
