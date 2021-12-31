package main

import "fmt"

func main() {
	// var colors map[string]string

	// colors := make(map[string]string)

	// colors["white"] = "#FFFFFF"

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
		"white": "#FFFFFF",
	}

	// delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Hex code for %s is %s \n", color, hex)
	}
}
