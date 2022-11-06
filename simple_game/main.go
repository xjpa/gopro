package main

import "fmt"

type Player struct {
	Health      int
	Name        string
	AttackPower int
}

func main() {
	player := Player{
		Health:      100,
		Name:        "Chad",
		AttackPower: 50,
	}
	killPlayer(player)
}

func killPlayer(player Player) {
	player.Health = 0
	fmt.Println("game over")
}
