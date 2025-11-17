package utils

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/color"
	"image/png"

	vision "cloud.google.com/go/vision/apiv1"
)

func GridToPng[T ~int](grid FinGrid[T]) ([]byte, error) {
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
		return nil, err
	}
	return buf.Bytes(), nil
}

func DetectText(pngData []byte) (string, error) {
	ctx := context.Background()

	client, e := vision.NewImageAnnotatorClient(ctx)
	if e != nil {
		return "", fmt.Errorf("error creating image annotator client: %w", e)
	}
	defer client.Close()
	image, e := vision.NewImageFromReader(bytes.NewReader(pngData))
	if e != nil {
		return "", fmt.Errorf("error creating image from reader: %w", e)
	}
	annotations, e := client.DetectTexts(ctx, image, nil, 1)
	if e != nil {
		return "", fmt.Errorf("error detecting text: %w", e)
	}
	if len(annotations) == 0 {
		return "", fmt.Errorf("no text detected in image")
	}
	return annotations[0].Description, nil
}
