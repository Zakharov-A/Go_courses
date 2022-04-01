package main

import dv "simple/devices"

var devices = []dv.Device{
	dv.Pc{dv.Cpu{
		Prod: "Intel",
		Core: 8,
	}},
	dv.Mobile{dv.Cpu{
		Prod: "ARM",
		Core: 4,
	}},
}

func main() {

	for _, dvc := range devices {
		dvc.On()
	}
}
