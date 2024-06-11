package parser

import (
	"encoding/json"
	strUtils "go-weather-now/utils"
	"os"
	"strconv"
	"time"

	"github.com/gocolly/colly"
)

type Percent int32

type Speed int32

type Celcius int32

type WeatherInfo struct {
	Time         time.Time `json:"time"`
	Temperature  Celcius   `json:"temperature"`
	RainlyChance Percent   `json:"rainly_chance"`
	Wet          Percent   `json:"wet"`
	Wind         Speed     `json:"wind"`
	Image        Image     `json:"image"`
}

func (weather *WeatherInfo) Update() {
	url := "https://www.google.com/search?q=погода+новосибирск+октябрьский+район&client=opera-gx&hs=9JG&sca_esv=b1bdec1098d0cba3&sxsrf=ADLYWIJ0zBCm7om2Lyq_7WEHVpaGXJ8H2g%3A1716801522720&ei=8k9UZrvQK6K79u8PueK9yAE&udm=&oq=погода+новосибирск+октя&gs_lp=Egxnd3Mtd2l6LXNlcnAaAhgCIizQv9C-0LPQvtC00LAg0L3QvtCy0L7RgdC40LHQuNGA0YHQuiDQvtC60YLRjyoCCAEyBRAAGIAEMgUQABiABDIGEAAYFhgeMgYQABgWGB4yBhAAGBYYHjIGEAAYFhgeMgYQABgWGB4yBhAAGBYYHjIGEAAYFhgeMgYQABgWGB5IviBQ8AFYxxhwAXgBkAEAmAGbBKABnBeqAQswLjUuMS4xLjIuMbgBA8gBAPgBAZgCDKAC8B3CAgoQABiwAxjWBBhHwgINEAAYgAQYsAMYQxiKBcICEhAjGIAEGCcYigUYnQIYRhiAAsICDhAAGIAEGLEDGIMBGMkDwgIIEAAYgAQYsQPCAgsQABiABBiSAxiKBcICCBAAGIAEGJIDwgIOEAAYgAQYsQMYgwEYigXCAhwQABiABBiKBRidAhhGGIACGJcFGIwFGN0E2AEBwgILEAAYgAQYsQMYgwGYAwCIBgGQBgq6BgYIARABGBOSBwsxLjUuMS4xLjMuMaAHu3E&sclient=gws-wiz-serp#ip=1"
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	c.OnHTML(".UQt4rd", func(h *colly.HTMLElement) {
		weather.Time = time.Now()
		temperature, err := strconv.Atoi(h.ChildText("span.wob_t.q8U8x"))

		if err != nil {
			panic("error convert temperature to int\n" + err.Error())
		}
		weather.Temperature = Celcius(temperature)

		rainly_ch, err := strconv.Atoi(strUtils.TrimSuffix(h.ChildText("span#wob_pp"), "%"))
		if err != nil {
			panic("error convert rainly chance to int\n" + err.Error())
		}
		weather.RainlyChance = Percent(rainly_ch)

		wet, err := strconv.Atoi(strUtils.TrimSuffix(h.ChildText("span#wob_hm"), "%"))
		if err != nil {
			panic("error convert wet to int\n" + err.Error())
		}
		weather.Wet = Percent(wet)

		wind, err := strconv.Atoi(strUtils.TrimSuffix(h.ChildText("span#wob_ws"), " м/с"))
		if err != nil {
			panic("error convert wind to int\n" + err.Error())
		}
		weather.Wind = Speed(wind)
	})
	c.OnHTML(".wob_tci", func(h *colly.HTMLElement) {
		weather.Image = *NewImage(Url(h.Attr("src")))
	})
	c.Visit(url)
}

func (weather WeatherInfo) Json(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic("error create file\n" + err.Error())
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.Encode(weather)
}
