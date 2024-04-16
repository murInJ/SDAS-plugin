package main

import (
	"github.com/vmihailenco/msgpack/v5"
	"image"
	"image/draw"
	"math"
)

func Init(args string) error {
	return nil
}

// EntryName: any
// ExistName: `Exist`
func Pipeline(entry map[string]map[string]any) (map[string]map[string]any, error) {
	res := make(map[string]map[string]any)
	res["Exist"] = make(map[string]any)
	l := int(math.Ceil(math.Sqrt(float64(len(entry)))))
	row := int(math.Ceil(float64(len(entry)) / float64(l)))
	height := 0
	width := 0
	index := 0
	var resultImage *image.RGBA
	for _, msg := range entry {
		var img image.RGBA
		err := msgpack.Unmarshal(msg["Data"].([]byte), &img)

		if err != nil {

			return nil, err
		}

		if index == 0 {
			res["Exist"]["Ntp"] = msg["Ntp"].(int64)
			res["Exist"]["DataType"] = msg["DataType"].(string)
			width = img.Bounds().Dx()
			height = img.Bounds().Dy()
			resultImage = image.NewRGBA(image.Rect(0, 0, l*width, row*height))
		}
		nx := width * (index % l)
		ny := height * (index / l)
		draw.Draw(resultImage, img.Bounds().Add(image.Pt(nx, ny)), &img, image.ZP, draw.Over)

		index += 1
	}
	bin, err := msgpack.Marshal(resultImage)
	if err != nil {

		return nil, err
	}
	res["Exist"]["Data"] = bin
	return res, nil
}
