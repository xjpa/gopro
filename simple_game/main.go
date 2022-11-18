package main

import "fmt"

type Game struct {
}

type Player struct {
	Health      uint //uint -- unsigned integer, w/c means it cant be negative
	Name        string
	AttackPower int
}

func NewPlayer(name string, hp uint, ap int) *Player {
	//mde our own constructor
	return &Player{
		Name:        name,
		Health:      hp,
		AttackPower: ap,
	}
}

func (p *Player) die() {
	//*Player: pointer
	//the "*" operator retrieves the variable that the pointer refers to
	p.Health = 0
}

func main() {
	playerA := NewPlayer("playerA", 100, 50)
	playerB := NewPlayer("player B", 100, 50)
	/*
		playerA := &Player{
			//&Player: pointer
			//& operator yields the address of a variable
			Health:      100,
			Name:        "player A",
			AttackPower: 50,
		}
	*/
	playerA.die()
	playerB.die()
	fmt.Println("player health is ", playerA.Health)
}

/*
func killPlayer(player *Player) {
	player.Health = 0
	fmt.Println("game over")
}
*/
