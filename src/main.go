package main

import "fmt"

func main() {
	newUser := createNewUser("Blargo")
	fmt.Println("Name:", newUser.firstName)

	var outsideTemperature = getWeather(50)
	fmt.Println("Weather in Fahrenheit:", outsideTemperature.temperatureF)
	fmt.Println("Weather in Celsius:", outsideTemperature.temperatureC)
}

// User ... the user playing the game.
type User struct {
	firstName string
}

type Weather struct {
	temperatureF float32
	temperatureC float32
}

// gets the F and C temperatures
func getWeather(fahrenheit float32) Weather {
	return Weather{fahrenheit, (fahrenheit - 32) * .5556}
}

func createNewUser(firstName string) User {
	return User{firstName}
}
