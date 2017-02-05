package main

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"

	"log"

	"path/filepath"

	"strings"

	"path"

	"math"

	"github.com/oliamb/cutter"
)

// Crop main crop function
func Crop(imagefile string) {
	_, destpath := filepath.Split(imagefile)
	destpath = strings.Split(destpath, ".")[0]
	cropImageBy9(imagefile, destpath)
}

func cropImageBy9(imagefile string, destpath string) {
	imageAbsPath, _ := filepath.Abs(imagefile)
	log.Println("Image ABS: " + imageAbsPath)
	f, err := os.Open(imageAbsPath)
	defer f.Close()
	ProcErr("opening image", err)
	img, _, err := image.Decode(f)
	ProcErr("decoding image", err)
	imgcfg, _, err := image.DecodeConfig(f)
	ProcErr("decoding image config", err)
	height, width := imgcfg.Height, imgcfg.Width

	if float64(height)/float64(width) > 1.05 || float64(width)/float64(height) > 1.05 {
		log.Printf("Warning: image not a square, cropping it to one.")
		aside := int(math.Min(float64(height), float64(width)))
		img, _ = cutter.Crop(img, cutter.Config{
			Height:  aside,
			Width:   aside,
			Mode:    cutter.Centered,
			Anchor:  image.Point{width / 2, height / 2},
			Options: 0,
		})
		height, width = aside, aside
	}

	encodeImage(cropImageXY(img, 0, 0, width/3, height/3), destpath, destpath+"1.png")
	encodeImage(cropImageXY(img, width/3, 0, width/3, height/3), destpath, destpath+"2.png")
	encodeImage(cropImageXY(img, width/3*2, 0, width/3, height/3), destpath, destpath+"3.png")

	encodeImage(cropImageXY(img, 0, height/3, width/3, height/3), destpath, destpath+"4.png")
	encodeImage(cropImageXY(img, width/3, height/3, width/3, height/3), destpath, destpath+"5.png")
	encodeImage(cropImageXY(img, width/3*2, height/3, width/3, height/3), destpath, destpath+"6.png")

	encodeImage(cropImageXY(img, 0, height/3*2, width/3, height/3), destpath, destpath+"7.png")
	encodeImage(cropImageXY(img, width/3, height/3*2, width/3, height/3), destpath, destpath+"8.png")
	encodeImage(cropImageXY(img, width/3*2, height/3*2, width/3, height/3), destpath, destpath+"9.png")
}

func cropImageXY(img image.Image, x, y, dx, dy int) image.Image {
	cImg, err := cutter.Crop(img, cutter.Config{
		Height:  dy,
		Width:   dx,
		Mode:    cutter.TopLeft,
		Anchor:  image.Point{x, y},
		Options: 0,
	})
	if err != nil {
		log.Fatalf("Error cropping image: %s", err.Error())
	}
	return cImg
}

func encodeImage(img image.Image, destpath string, fn string) {
	fo, err := os.Create(path.Join(destpath, fn))
	ProcErr("creating destination image", err)
	png.Encode(fo, img)
}
