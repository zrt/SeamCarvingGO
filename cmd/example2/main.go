package main

// example2
// by zrt

import (
	"fmt"
	. "github.com/zrt/SeamCarvingGO/pkg/sc"
	//"log"
	//"net/http"
	//_ "net/http/pprof"
)

func main() {
	//go func() {
	//	log.Println(http.ListenAndServe("localhost:10000", nil))
	//}()

	img := LoadJPEG("input.jpg")
	fmt.Printf("%v\n", img.Bounds().Size()) // show img size

	energyFunc := L2NormEnergyFunc
	energyImg := SliceToImg(energyFunc(img, 1+2))
	SavePNG(energyImg, "img_energy.png") // show energy img

	CarvingX(img, energyFunc, true)
	CarvingY(img, energyFunc, true) // show a seam example

	newImg := Resize(img, energyFunc, 300, 200)

	fmt.Printf("%v\n", newImg.Bounds().Size()) // show newImg size
	SavePNG(newImg, "output.png")              // show newImg
}
