package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/davidbyttow/govips/v2/vips"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

type size struct {
	k string
	v int
}

func main() {
	vips.Startup(nil)
	defer vips.Shutdown()

	image1, err := vips.NewImageFromFile("input.heic")
	ep := vips.NewDefaultJPEGExportParams()
	checkError(err)

	// Check image type
	log.Println(image1.Format().FileExt())

	// Crop it
	err = image1.ExtractArea(100, 100, 1500, 1500)
	checkError(err)

	sizes := []*size{
		&size{
			k: "small",
			v: 100,
		},
		&size{
			k: "medium",
			v: 500,
		},
		&size{
			k: "large",
			v: 1000,
		},
	}
	for _, v := range sizes {
		// create a copy of original image
		s, err := image1.Copy()
		checkError(err)

		// crop it
		err = s.Thumbnail(v.v, v.v, 6)
		checkError(err)

		sBytes, _, err := s.Export(ep)
		checkError(err)

		// write in file
		f := fmt.Sprintf("output/%s.jpeg", v.k)
		err = ioutil.WriteFile(f, sBytes, 0644)
		checkError(err)
	}
}
