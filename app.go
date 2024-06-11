package main

import (
	"context"
	"fmt"
	"go-weather-now/parser"
)

// App struct
type App struct {
	ctx         context.Context
	WeatherInfo parser.WeatherInfo
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.WeatherInfo.Update()
}

func (a *App) GetWeather() parser.WeatherInfo {
	a.WeatherInfo.Json("weather.json")
	return a.WeatherInfo
}

func (a *App) UpdateWeather() parser.WeatherInfo {
	a.WeatherInfo.Json("weather.json")
	a.WeatherInfo.Update()
	return a.WeatherInfo
}

func (a *App) DownloadImage(path string) string {
	a.WeatherInfo.Image.Download(path)
	return a.WeatherInfo.Image.File
}

func (a *App) GetDay() string {
	day := a.WeatherInfo.Time.Day()
	return fmt.Sprintf("%d", day)
}

func (a *App) GetMonth() string {
	month := a.WeatherInfo.Time.Month()
	return month.String()
}

func (a *App) GetTime() string {
	return a.WeatherInfo.Time.Format("15:04")
}
