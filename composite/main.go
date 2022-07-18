package main

import (
	"./pkg"
)

var (
		cpu = pkg.Cpu {
			Name: "CP-1",
			Description: "Processor-1",
		}

		cpu2 = pkg.Cpu {
			Name: "CP-2",
			Description: "Processor-2",
		}

		card = pkg.GraphicsCard{
			Name: "Radeon",
			Description: "VideoCard-1",
		}

		card2 = pkg.GraphicsCard{
			Name: "NVidia",
			Description: "VideoCard-2",
		}

		mb = pkg.Motherboard {
			Name: "Asus",
			Description: "Motherboard-1",
			Components: []pkg.Component{
				cpu, cpu2, card, card2,
			},
		}

		pc = pkg.Pc {
			Name: "Dell",
			Description: "PC-1",
			Components: []pkg.Component{
				mb,
			},
		}

)

func main() {

	pc.Search("Radeon")

}