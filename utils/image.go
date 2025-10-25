package utils

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/png"
)

func GridToPngString[T ~int](grid FinGrid[T]) (string, error) {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{grid.width, grid.height}})
	for i, score := range grid.items {
		x, y := i%grid.width, i/grid.width
		var colr color.Color
		if score == 1 {
			colr = color.White
		} else {
			colr = color.Black
		}
		img.Set(x, y, colr)
	}
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return "", err
	}
	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
	return encoded, nil
}