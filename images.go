/******************************************
Names: Apurva Gandhi and Margaret Haley
Course: CSCI324
Professor King
Creative Program
Sample execution: go run images.go
*****************************************/

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

type circle struct {
	radius int
	center image.Point
}

func (c *circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *circle) Bounds() image.Rectangle {
	return image.Rect(c.center.X-c.radius, c.center.Y-c.radius, c.center.X+c.radius, c.center.Y+c.radius)
}

func (c *circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.center.X)+0.5, float64(y-c.center.Y)+0.5, float64(c.radius)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}

func main() {
	// input an image
	picture1, err := os.Open("covid2.png")
	if err != nil {
		fmt.Println("Sadly, an error has occured with this image: ", err)
	}
	defer picture1.Close()

	theImage, format, err := image.Decode(picture1)
	if err != nil {
		fmt.Println("Sadly, an error has occured with this image: ", err)
	}
	fmt.Println("format is: ", format)

	// convert to RGBA
	bounds := theImage.Bounds()
	m := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(m, m.Bounds(), theImage, bounds.Min, draw.Src)
	//fmt.Println("bounds are: ", bounds)

	// initialize new image to all black
	background := image.NewRGBA(image.Rect(0, 0, 640, 480))
	black := color.RGBA{0, 0, 0, 255}
	draw.Draw(background, background.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)

	// paste on background
	//sr := bounds
	//dp := image.Point{50, 50}
	// converting source rectangle into destination's coordinates
	//collage := image.Rectangle{dp, dp.Add(sr.Size())}
	//draw.Draw(background, collage, theImage, sr.Min, draw.Src)

	//Paste on cropped background circle
	draw.DrawMask(background, background.Bounds(), theImage, image.ZP, &circle{100, image.Point{100, 100}}, image.ZP, draw.Over)

	/*
		func Width(i image.Image) int {
			return i.Bounds().Max.X — i.Bounds().Min.X
		  }
		  func Height(i image.Image) int {
			return i.Bounds().Max.Y — i.Bounds().Min.Y
		  }

	*/

	// output the image into a file named output
	out, err := os.Create("output.png")
	if err != nil {
		fmt.Println("Sadly, an error has occured with this image: ", err)
	}
	defer out.Close()

	err = png.Encode(out, background)
	if err != nil {
		fmt.Println("Sadly, an error has occured with this image: ", err)
	}

	/*
		// Specify the quality, between 0-100
		// Higher is better
		opt := jpeg.Options{
			Quality: 90,
		}
		err = jpeg.Encode(picture, theImage, &opt)
		if err != nil {
			// Handle error
			fmt.Println("error!")
			fmt.Println(err)
		}*/

}
