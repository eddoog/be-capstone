package pkg

import (
	"math"
	"strconv"
)

func DecideHexColor(value float64) string {
	red, green, blue := float64ToRGB(value)

	hexColor := rgbToHex(red, green, blue)

	return hexColor
}

func float64ToRGB(value float64) (red, green, blue float64) {
	value = math.Max(0, math.Min(1, value))

	green = 255 * (1 - value)

	red = 255 * value

	blue = 0

	return red, green, blue
}

func rgbToHex(red, green, blue float64) string {
	// Convert each component to hexadecimal string
	redHex := strconv.FormatInt(int64(red), 16)
	greenHex := strconv.FormatInt(int64(green), 16)
	blueHex := strconv.FormatInt(int64(blue), 16)

	// Ensure two-digit representation for each color
	if len(redHex) == 1 {
		redHex = "0" + redHex
	}
	if len(greenHex) == 1 {
		greenHex = "0" + greenHex
	}
	if len(blueHex) == 1 {
		blueHex = "0" + blueHex
	}

	// Combine into hexadecimal color string
	hexColor := "#" + redHex + greenHex + blueHex

	return hexColor
}
