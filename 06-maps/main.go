package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to class on maps")
	language := make(map[string]string)
	language["en"] = "English"
	language["es"] = "Spanish"
	language["fr"] = "French"
	language["de"] = "German"
	language["it"] = "Italian"
	language["ml"] = "Malayalam"
	language["hi"] = "Hindi"
	language["bn"] = "Bengali"
	language["gu"] = "Gujarati"
	fmt.Println(language)
	fmt.Println(language["en"])

	// delete english from the map
	delete(language, "en")
	fmt.Println(language)
	// looping through the map
	for key, value := range language {
		// fmt.Println(key, value)
		fmt.Printf("key is %s and value is %s\n", key, value)
	}
	// looping through the map value only
	for _, value := range language {
		// fmt.Println(key, value)
		fmt.Println(value)
	}
}
