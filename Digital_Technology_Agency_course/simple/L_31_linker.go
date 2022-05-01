// package main

// import "simple/L_31_linker"

// var (
// 	cpu = L_31_linker.Cpu{
// 		Name:		"CP-1",
// 		Description: "Protsessor",
// 	}
// 	cpu2 = L_31_linker.Cpu{
// 		Name:		"CP-66",
// 		Description: "Protsessor-M",
// 	}
// 	card = L_31_linker.GraphicsCard{
// 		Name:		"Radeon",
// 		Description: "Video-card",
// 	}
// 	card2 = L_31_linker.GraphicsCard{
// 		Name:		"GeForse",
// 		Description: "Video-card6",
// 	}
// 	motherboad = L_31_linker.Motherboard{
// 		Name:  		"Gigabyte",
// 		Description: "Motherboard",
// 		Components: []L_31_linker.Component{
// 			cpu, cpu2, card, card2,
// 		},
// 	}
// 	pc = L_31_linker.Pc{
// 		Name:  		"PC",
// 		Description: "Personal Computer",
// 		Components: []L_31_linker.Component{
// 			motherboad,
// 		},
// 	}
// )

// func main() {
// 	pc.Search("CP-66")
// }
