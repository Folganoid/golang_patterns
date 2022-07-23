package main

import "./pkg"

var types = []string{
	pkg.PCType, pkg.NotebookType, pkg.ServerType, "monoblock",
}

func main() {

	for _, typeName := range types {
		computer := pkg.New(typeName)
		if computer == nil {
			continue
		}
		computer.PrintDetails()
	}


}
