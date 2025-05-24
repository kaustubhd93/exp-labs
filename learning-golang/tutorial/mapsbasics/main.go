package main

import (
	"fmt"
)

func main() {
	welcomeMessage := "############## Maps in go. ##############"
	fmt.Println(welcomeMessage)

	// below syntax gives an error: "panic: assignment to entry in nil map". You will have to use make()
	// var languages map[string]string
	// this below initializes an empty map.
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	fmt.Println("Languages:", languages)

	//accessing an element by key
	fmt.Println("JS:", languages["JS"])
	// deleting an element using the key, if key is not found it doesnt throw an error.
	delete(languages, "JS")
	fmt.Println("Modified lanaguages:", languages)

	for key, value := range languages {
		fmt.Printf("%v: %v\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("language: %v\n", value)
	}

	for key, _ := range languages {
		fmt.Printf("language short form: %v\n", key)
	}
}
