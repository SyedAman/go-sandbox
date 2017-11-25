package main

import "fmt"

// User ... the user playing the game.
type User struct {
	firstName string
}

// Weather defines F and C temperatures
type Weather struct {
	temperatureF float32
	temperatureC float32
}

// GetWeather gets the F and C temperatures
func GetWeather(fahrenheit float32) Weather {
	return Weather{fahrenheit, (fahrenheit - 32) * .5556}
}

// CreateNewUser creates a new user
func CreateNewUser(firstName string) User {
	return User{firstName}
}

// DisplayWeather prints out the weather in Fahrenheit and Celsius
func DisplayWeather(weather Weather) {
  fmt.Println("Weather in Fahrenheit:", weather.temperatureF)
  fmt.Println("Weather in Celsius:", weather.temperatureC)
}

func main() {
  // create user
	newUser := CreateNewUser("Blargo")
	fmt.Println("Name:", newUser.firstName)

  // get the weather today
	var outsideTemperature = GetWeather(50)
  DisplayWeather(outsideTemperature)
}
