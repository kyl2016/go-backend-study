package main

import (
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"log"
)

func main() {
	// Open a test image.
	src, err := imaging.Open("testdata/flowers.jpg")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	if src.Bounds().Max.Y-src.Bounds().Min.Y == src.Bounds().Max.X-src.Bounds().Min.X {
		fmt.Println("是正方形")
	}

	resized := imaging.Resize(src, 150, 0, imaging.Lanczos)
	err = imaging.Save(resized, "testdata/flowers_150.jpg")

	resized = imaging.Resize(src, 60, 0, imaging.Box)
	err = imaging.Save(resized, "testdata/flowers_60_Box.jpg")

	resized = imaging.Resize(src, 60, 0, imaging.Bartlett)
	err = imaging.Save(resized, "testdata/flowers_60_Bartlett.jpg")

	resized = imaging.Resize(src, 60, 0, imaging.Blackman)
	err = imaging.Save(resized, "testdata/flowers_60_Blackman.jpg")

	resized = imaging.Resize(src, 60, 0, imaging.BSpline)
	err = imaging.Save(resized, "testdata/flowers_60_BSpline.jpg")

	resized = imaging.Resize(src, 60, 0, imaging.CatmullRom)
	err = imaging.Save(resized, "testdata/flowers_60_CatmullRom.jpg")

	resized = imaging.Resize(src, 60, 0, imaging.Cosine)
	err = imaging.Save(resized, "testdata/flowers_60_Cosine.jpg")

	resized = imaging.Resize(src, 60, 0, imaging.Gaussian)
	err = imaging.Save(resized, "testdata/flowers_60_Gaussian.jpg")

	resized = imaging.Resize(src, 60, 0, imaging.Hamming)
	err = imaging.Save(resized, "testdata/flowers_60_Hamming.jpg")

	resized = imaging.Resize(src, 60, 0, imaging.Hann)
	err = imaging.Save(resized, "testdata/flowers_60_Hann.jpg")

	resized = imaging.Resize(src, 60, 0, imaging.Hermite)
	err = imaging.Save(resized, "testdata/flowers_60_Hermite.jpg")

	resized = imaging.Resize(src, 60, 0, imaging.Linear)
	err = imaging.Save(resized, "testdata/flowers_60_Linear.jpg")

	resized = imaging.Resize(src, 60, 0, imaging.MitchellNetravali)
	err = imaging.Save(resized, "testdata/flowers_60_MitchellNetravali.jpg")

	resized = imaging.Resize(src, 160, 0, imaging.Lanczos)
	err = imaging.Save(resized, "testdata/flowers_160.jpg")

	// Crop the original image to 300x300px size using the center anchor.
	src = imaging.CropAnchor(src, 300, 300, imaging.Center)

	// Resize the cropped image to width = 200px preserving the aspect ratio.
	src = imaging.Resize(src, 200, 0, imaging.Lanczos)

	// Create a blurred version of the image.
	img1 := imaging.Blur(src, 5)

	// Create a grayscale version of the image with higher contrast and sharpness.
	img2 := imaging.Grayscale(src)
	img2 = imaging.AdjustContrast(img2, 20)
	img2 = imaging.Sharpen(img2, 2)

	// Create an inverted version of the image.
	img3 := imaging.Invert(src)

	// Create an embossed version of the image using a convolution filter.
	img4 := imaging.Convolve3x3(
		src,
		[9]float64{
			-1, -1, 0,
			-1, 1, 1,
			0, 1, 1,
		},
		nil,
	)

	// Create a new image and paste the four produced images into it.
	dst := imaging.New(400, 400, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, img1, image.Pt(0, 0))
	dst = imaging.Paste(dst, img2, image.Pt(0, 200))
	dst = imaging.Paste(dst, img3, image.Pt(200, 0))
	dst = imaging.Paste(dst, img4, image.Pt(200, 200))

	// Save the resulting image as JPEG.
	err = imaging.Save(dst, "testdata/out_example.jpg")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}
