/******************************************
Names: Apurva Gandhi and Margaret Haley
Course: CSCI324
Professor King
Creative Program
Sample execution: go run images.go
*****************************************/

package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"strings"
	"time"
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

// start() gets the images the user wants to put in their collage, and places their names in an array
func getNames() []string {
	fmt.Println("Please enter the names of fours images separated by a space (png only). For example: ")
	fmt.Println("image1.png image2.png image3.png image4.png")
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	image_names := strings.Split(response, " ")
	for len(image_names) != 4 {
		fmt.Println("Sorry, you have to input 4 names. Please try again.")
		reader := bufio.NewReader(os.Stdin)
		response, _ := reader.ReadString('\n')
		image_names = strings.Split(response, " ")
	}
	// Remove the extra carriage return from the last element of the array
	changeName := image_names[len(image_names)-1]
	image_names[len(image_names)-1] = ""
	for i := 0; i < len(changeName)-1; i++ {
		image_names[len(image_names)-1] += string(changeName[i])
	}
	return image_names
}

func main() {

	// Get names of images from user
	image_names := getNames()

	// Start a timer
	starting := time.Now()

	// initialize new image to all black
	background := image.NewRGBA(image.Rect(0, 0, 480, 480))
	black := color.RGBA{0, 0, 0, 255}
	draw.Draw(background, background.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)

	for i := 0; i < len(image_names); i++ {

		// Input an image
		picture, err := os.Open(image_names[i])
		if err != nil {
			fmt.Println("Sadly, an error has occured with this image: ", err)
		}
		defer picture.Close()

		// Decode it
		theImage, _, err := image.Decode(picture)
		if err != nil {
			fmt.Println("Sadly, an error has occured with this image: ", err)
		}

		// Calculate the top left and bottom right of the mini-square we're pasting to
		var sp image.Point
		var ep image.Point
		if i == 0 {
			sp = image.Point{0, 0}
			ep = image.Point{240, 240}
		} else if i == 1 {
			sp = image.Point{240, 0}
			ep = image.Point{480, 240}
		} else if i == 2 {
			sp = image.Point{0, 240}
			ep = image.Point{240, 480}
		} else {
			sp = image.Point{240, 240}
			ep = image.Point{480, 480}
		}
		// Paste a circle on the background
		draw.DrawMask(background, image.Rectangle{sp, ep}, theImage, image.ZP, &circle{120, image.Point{120, 120}}, image.ZP, draw.Over)
	}
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

	totaltime := time.Since(starting)
	fmt.Printf("Making your collage took %s", totaltime)
	fmt.Println()
}
