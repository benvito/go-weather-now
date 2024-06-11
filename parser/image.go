package parser

import (
	"io"
	"net/http"
	"os"
	"strings"
)

type Url string

func (url *Url) HttpAdd() {
	if !strings.HasPrefix(string(*url), "http:") && !strings.HasPrefix(string(*url), "https:") {
		*url = Url("http:" + string(*url))
	}
}

// const filepath = "misc/images/weather.jpg"

type Image struct {
	URL  Url
	File string
}

func NewImage(url Url) *Image {
	url.HttpAdd()
	img := &Image{URL: url}
	// img.Download()
	return img
}

func (image *Image) Download(filepath string) {
	file, err := os.Create(filepath)
	if err != nil {
		panic("error create file\n" + err.Error())
	}
	defer file.Close()

	response, err := http.Get(string(image.URL))
	if err != nil {
		panic("error http get image\n" + err.Error())
	}
	defer response.Body.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		panic("error download image\n" + err.Error())
	}

	image.File = filepath
}
