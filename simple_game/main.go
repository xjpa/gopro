package main

import (
	"fmt"
	"time"
)

type Game struct {
	isRunning bool
}

func NewGame() *Game {
	//made our own constructor as golang has no in built
	return &Game{
		isRunning: false,
		//this is not needed tho as default value of booelans in Golang is false
	}
}

func (g *Game) Start() {
	g.isRunning = true
	g.gameLoop()
}

func (g *Game) gameLoop() {
	interval := 1 * time.Second
	//this for-loop keeps looping til we quit the program
	for {
		fmt.Println("the game is rnning")
		time.Sleep(interval) //time.Sleep -- belongs in standard library, makes it sleep which in this case is for 1 second
	}
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
	game := NewGame()
	game.Start()
	fmt.Printf("%+v", game) //shows everything inside the struct, the "+v" means verbose
}
