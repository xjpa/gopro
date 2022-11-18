package main

import "fmt"

type Player struct {
	Health      int
	Name        string
	AttackPower int
}

func (p *Player) die() {
	//*Player: pointer
	//the "*" operator retrieves the variable that the pointer refers to
	p.Health = 0
}

func main() {
	player := &Player{
		//&Player: pointer
		//& operator yields the address of a variable
		Health:      100,
		Name:        "Chad",
		AttackPower: 50,
	}
	player.die()
	fmt.Println("player health is ", player.Health)
}

/*
func killPlayer(player *Player) {
	player.Health = 0
	fmt.Println("game over")
}
*/
