/******************************************
Names: Apurva Gandhi and Margaret Haley
Course: CSCI324
Professor King
Creative Program - making and image collage
Sample execution: go run collage.go
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
)

// Implementing a circle struct and image functions on it, in order to create a circular crop

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

// initialize vars to create black background
var background = image.NewRGBA(image.Rect(0, 0, 480, 480))
var black = color.RGBA{0, 0, 0, 255}

// initialize array that will contain our four images
var ourImages []image.Image

// start() gets the images the user wants to put in their collage, and places their names in an array
func getNames() []string {
	fmt.Println("Please enter the names of fours images separated by a space (png only). For example: ")
	fmt.Println("image1.png image2.png image3.png image4.png")
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	image_names := strings.Split(response, " ")

	// if user did not input names correctly, ask them to try again
	for len(image_names) != 4 {
		fmt.Println("Sorry, you have to input 4 names. Please try again.")
		reader := bufio.NewReader(os.Stdin)
		response, _ := reader.ReadString('\n')
		image_names = strings.Split(response, " ")
	}

	/***********************************
	CHANGE: we had forgotten to add the
	slicing to the comparison program
	that we had in the collage program.
	Here we have added it.
	***********************************/
	// Remove the extra carriage return from the last element of the array
	changeName := image_names[len(image_names)-1]
	image_names[len(image_names)-1] = changeName[0:(len(changeName) - 1)]

	return image_names
}

// getCoords() calculates the top left and bottom right coordinate of the mini-square we're pasting to
func getCoords(i int) (image.Point, image.Point) {
	var start image.Point
	var end image.Point
	if i == 0 {
		start = image.Point{0, 0}
		end = image.Point{240, 240}
	} else if i == 1 {
		start = image.Point{240, 0}
		end = image.Point{480, 240}
	} else if i == 2 {
		start = image.Point{0, 240}
		end = image.Point{240, 480}
	} else {
		start = image.Point{240, 240}
		end = image.Point{480, 480}
	}
	return start, end
}

// readImage() reads in one of the image files and returns it
func readImage(file_name string, i int) image.Image {

	// Input the image
	fmt.Println("Starting to read image: ", i)
	picture, err := os.Open(file_name)
	if err != nil {
		fmt.Println("Sadly, an error has occured with this image: ", err)
	}

	// Close the connection when we're done
	defer picture.Close()

	// Decode the image
	theImage, _, err := image.Decode(picture)
	if err != nil {
		fmt.Println("Sadly, an error has occured with this image: ", err)
	}

	return theImage
}

// drawOneImage() takes in an image, calls edit to change the correct portion of
// the background to a different color, and draws the image to that part of the background
func drawOneImage(sp image.Point, ep image.Point, theImage image.Image, i int) {
	// First, edit this image's background (change the color)
	editOneImage(sp, ep, i)

	// Draw the cropped image on the background
	fmt.Println("Starting to draw image: ", i)
	draw.DrawMask(background, image.Rectangle{sp, ep}, theImage, image.ZP, &circle{120, image.Point{120, 120}}, image.ZP, draw.Over)
}

// editOneImage() takes in two sets of x,y coordinates and changes the color of
// the background in the rectangle that those coordinates specify. The color is
// chosen based off of the index of the image in the array of image names. The
// first image in the array gets the red color, for ex.
func editOneImage(sp image.Point, ep image.Point, i int) {
	// loop through every pixel for this image's part of the background and change the color
	for x := sp.X; x < ep.X; x++ {
		for y := sp.Y; y < ep.Y; y++ {
			// ex: the first image to come in will have a red background
			if i == 0 {
				background.Set(x, y, color.RGBA{255, 0, 0, 255}) // red
			} else if i == 1 {
				background.Set(x, y, color.RGBA{0, 255, 0, 255}) // green
			} else if i == 2 {
				background.Set(x, y, color.RGBA{0, 0, 255, 255}) // blue
			} else if i == 3 {
				background.Set(x, y, color.RGBA{128, 0, 128, 255}) // purple
			}
		}
	}
}

func main() {

	// Get names of images from user
	image_names := getNames()

	// Initialize black background to draw on
	draw.Draw(background, background.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)

	for i := 0; i < len(image_names); i++ {
		// Calculate the coordinates this image will be printed at
		sp, ep := getCoords(i)
		// Read an image
		image := readImage(image_names[i], i)
		// Draw it
		drawOneImage(sp, ep, image, i)
	}

	// creates output.png file where our collage will go
	out, err := os.Create("output.png")
	if err != nil {
		fmt.Println("Sadly, an error has occured with this image: ", err)
	}
	defer out.Close()

	// encodes the collage to the .png format
	err = png.Encode(out, background)
	if err != nil {
		fmt.Println("Sadly, an error has occured with this image: ", err)
	}
}
