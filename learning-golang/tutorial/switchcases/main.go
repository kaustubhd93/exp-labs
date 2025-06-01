package main

import (
	"fmt"
	"math/rand"
)

func main() {
	welcomeMessage := "############## SwitchCase in go. ##############"
	fmt.Println(welcomeMessage)

	diceNumber := rand.Intn(6) + 1
	fmt.Println("Number on dice is:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Move 1 step. Cannot open if pawn is still inside")
	case 2:
		fmt.Println("Move 2 steps. Cannot open if pawn is still inside")
	case 3:
		fmt.Println("Move 3 steps. Cannot open if pawn is still inside")
	case 4:
		fmt.Println("Move 4 steps. Cannot open if pawn is still inside")
	case 5:
		fmt.Println("Move 5 steps. Cannot open if pawn is still inside")
	case 6:
		fmt.Println("Move 6 steps or you can open if pawn is still inside. Roll the dice again. You get one more chance.")
	default:
		fmt.Println("Wait! what just happened ?")
	}
}
