package main

import (
	"github.com/murInJ/SDAS-plugin/kitex_gen/api"
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
func Pipeline(entry map[string]*api.SourceMsg) (map[string]*api.SourceMsg, error) {
	res := &api.SourceMsg{}
	l := int(math.Ceil(math.Sqrt(float64(len(entry)))))
	height := 0
	width := 0
	index := 0
	var resultImage *image.RGBA
	for _, msg := range entry {
		var img image.RGBA
		err := msgpack.Unmarshal(msg.Data, &img)
		if err != nil {
			return nil, err
		}

		if index == 0 {
			res.Ntp = msg.Ntp
			res.DataType = msg.DataType
			width = img.Bounds().Dx()
			height = img.Bounds().Dy()
			resultImage = image.NewRGBA(image.Rect(0, 0, l*width, l*height))
		}

		ny := width * (index / l)
		nx := height * (index % l)
		draw.Draw(resultImage, resultImage.Bounds().Add(image.Pt(nx, ny)), &img, image.ZP, draw.Src)

		index += 1
	}

	bin, err := msgpack.Marshal(resultImage)
	if err != nil {
		return nil, err
	}
	res.Data = bin
	out := make(map[string]*api.SourceMsg)
	out["Exist"] = res
	return out, nil
}
