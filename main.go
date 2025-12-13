package main

import "fmt"

type WeatherCondition int

const (
    Clear WeatherCondition = iota
    Rain
    HeavyRain
    Snow
)

type WeatherData struct {
    Condition WeatherCondition
    WindSpeed int
}

const basePrise = 1

func GetWeatherMultiplier(weather WeatherData) float64{
	switch{
	case weather.Condition == Clear && weather.WindSpeed > 15:
		return basePrise + 0.1
	case weather.Condition == HeavyRain && weather.WindSpeed > 15:
		return basePrise + 0.3
	case weather.Condition == Rain && weather.WindSpeed > 15:
		return  basePrise + 0.225
	case weather.Condition == Snow && weather.WindSpeed > 15:
		return  basePrise + 0.25	
	case weather.Condition == HeavyRain:
		return basePrise + 0.2
	case weather.Condition == Rain:
		return  basePrise + 0.125
	case weather.Condition == Snow:
		return  basePrise + 0.15
	default:
		return basePrise
	}	
}

func main() {
	fmt.Println(GetWeatherMultiplier(WeatherData{HeavyRain, 10}))
	fmt.Println(GetWeatherMultiplier(WeatherData{HeavyRain, 20}))

	fmt.Println(GetWeatherMultiplier(WeatherData{Rain, 10}))
	fmt.Println(GetWeatherMultiplier(WeatherData{Rain, 20}))

	fmt.Println(GetWeatherMultiplier(WeatherData{Clear, 10}))
	fmt.Println(GetWeatherMultiplier(WeatherData{Clear, 20}))

	fmt.Println(GetWeatherMultiplier(WeatherData{Snow, 10}))
	fmt.Println(GetWeatherMultiplier(WeatherData{Snow, 20}))
}