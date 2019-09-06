// Package pngimg allows for loading png images and applying
// image flitering effects on them
package pngimg

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	//"fmt"
)


// The PNGImage represents a structure for working with PNG images.
type PNGImage struct {
	in  image.Image
	out *image.RGBA64
}

//
// Public functions
//

// This function basically put the data from "out" to "in" in order for later reuse
func (img *PNGImage)ReLoad(){
	inBounds := img.out.Bounds()
	img.in = img.out.SubImage(inBounds)
	img.out = image.NewRGBA64(inBounds)
}

// Load returns a PNGImage that was loaded based on the filePath parameter
func Load(filePath string) (*PNGImage, error) {

	inReader, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}
	defer inReader.Close()

	inImg, err := png.Decode(inReader)

	if err != nil {
		return nil, err
	}

	inBounds := inImg.Bounds()

	outImg := image.NewRGBA64(inBounds)

	return &PNGImage{inImg, outImg}, nil
}

// Save saves the image to the given file
func (img *PNGImage) Save(filePath string) error {

	outWriter, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outWriter.Close()

	err = png.Encode(outWriter, img.in)
	if err != nil {
		return err
	}
	return nil
}

//clamp will clamp the comp parameter to zero if it is less than zero or to 65535 if the comp parameter
// is greater than 65535.
func clamp(comp float64) uint16 {
	return uint16(math.Min(65535, math.Max(0, comp)))
}

// Grayscale applies a grayscale filtering effect to the image
func (img *PNGImage) Grayscale(minX int, maxX int) {
	bounds := img.out.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := minX; x < maxX; x++ {
			//Returns the pixel (i.e., RGBA) value at a (x,y) position
			// Note: These get returned as int32 so based on the math you'll
			// be performing you'll need to do a conversion to float64(..)
			r, g, b, a := img.in.At(x, y).RGBA()

			//Note: The values for r,g,b,a for this assignment will range between [0, 65535].
			//For certain computations (i.e., convolution) the values might fall outside this
			// range so you need to clamp them between those values.
			greyC := clamp(float64(r+g+b) / 3)

			//Note: The values need to be stored back as uint16 (I know weird..but there's valid reasons
			// for this that I won't get into right now).
			img.out.Set(x, y, color.RGBA64{greyC, greyC, greyC, uint16(a)})
		}
	}
}

func (img *PNGImage) getPixel(x int, y int) (int, int, int, int) {
	bounds := img.out.Bounds()
	// check for boundaries. Only return when accessible
	if x >= bounds.Min.X && x < bounds.Max.X {
		if y >= bounds.Min.Y && y < bounds.Max.Y{
			r,g,b,a := img.in.At(x, y).RGBA()
			return int(r),int(g),int(b),int(a)	
		}
	}
	return 0, 0, 0, 0
}

func (img *PNGImage) calculatePixel(effect string, x int, y int)(uint16, uint16, uint16, uint16) {
	var rTotal float64
	var gTotal float64
	var bTotal float64
	//row 1
	r1, g1, b1, _ := img.getPixel(x-1, y-1)
	r2, g2, b2, _ := img.getPixel(x, y-1)
	r3, g3, b3, _ := img.getPixel(x+1, y-1)

	//row 2
	r4, g4, b4, _ := img.getPixel(x-1, y)
	r5, g5, b5, a := img.getPixel(x, y)
	r6, g6, b6, _ := img.getPixel(x+1, y)

	//row 3
	r7, g7, b7, _ := img.getPixel(x-1, y+1)
	r8, g8, b8, _ := img.getPixel(x, y+1)
	r9, g9, b9, _ := img.getPixel(x+1, y+1)

	switch effect {
		//Sharpen
		case "S":
			rTotal = float64(5*r5) - float64(r2 + r4 + r6 + r8)
			gTotal = float64(5*g5) - float64(g2 + g4 + g6 + g8)
			bTotal = float64(5*b5) - float64(b2 + b4 + b6 + b8)
		//Edge Detection
		case "E":
			rTotal = float64(8*r5) - float64(r1 + r2 + r3 + r4 + r6 + r7 + r8 + r9)
			gTotal = float64(8*g5) - float64(g1 + g2 + g3 + g4 + g6 + g7 + g8 + g9)
			bTotal = float64(8*b5) - float64(b1 + b2 + b3 + b4 + b6 + b7 + b8 + b9)
		//Blur
		case "B":
			rTotal = float64(r5 + r1 + r2 + r3 + r4 + r6 + r7 + r8 + r9)/9
			gTotal = float64(g5 + g1 + g2 + g3 + g4 + g6 + g7 + g8 + g9)/9
			bTotal = float64(b5 + b1 + b2 + b3 + b4 + b6 + b7 + b8 + b9)/9
	}
	return clamp(rTotal), clamp(gTotal), clamp(bTotal), uint16(a)
}

func (img *PNGImage) ApplyEffect(effect string, minX int, maxX int) {
	if effect == "G"{
		img.Grayscale(minX, maxX)
	} else {
		bounds := img.out.Bounds()
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := minX; x < maxX; x++ {
				// Get the value after applying the effect.
				r, g, b, a := img.calculatePixel(effect, x, y)
				// Set it to the "out" of the image
				img.out.Set(x, y, color.RGBA64{r, g, b, a})
			}
		}
	}
}

func (img *PNGImage) ProcessImg(effect string, isPara bool, total int, seq int){
	bounds := img.out.Bounds()
	minX := bounds.Min.X
	maxX := bounds.Max.X

	// if it is called with parallelism, vertically partition it.
	if isPara{
		if total != 1{
			seqf := float64(seq)
			intervalX := math.Ceil(float64(bounds.Max.X/total))
			minX = int(seqf*intervalX)
			maxX = int(math.Min(float64((seqf+1)*intervalX), float64(bounds.Max.X)))
		}
	}

	//fmt.Println("start: ", seq)
	img.ApplyEffect(effect, minX, maxX)

}

