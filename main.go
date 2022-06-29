package main

import (
	"mytool/form"

	"github.com/ying32/govcl/vcl"
)

func main() {
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&form.MainForm)
	vcl.Application.Run()
}
