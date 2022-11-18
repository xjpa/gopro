package main

import "fmt"

type Player struct {
	Health      int
	Name        string
	AttackPower int
}

func main() {
	player := &Player{
		//&Player: pointer
		//& operator yields the address of a variable
		Health:      100,
		Name:        "Chad",
		AttackPower: 50,
	}
	killPlayer(player)
	fmt.Println("player health is ", player.Health)
}

//*Player: pointer
//* operator retrieves the variable that the pointer refers to
func killPlayer(player *Player) {
	player.Health = 0
	fmt.Println("game over")
}
