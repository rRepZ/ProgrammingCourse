package main

import "fmt"

func main() {
	// пример встраивания
	b := B{}
	b.SomeMethod()

	// полиморфизм

	// p := &Player{}

	// game := NewGame(p)
	// game.DoMoveCurrent()

	// p1 := &Player{Name: "Bot_Easy"}

	// game1 := NewGame(p1)
	// game1.DoMoveCurrent()

	// p2 := &Player{Name: "Bot_Hard"}
	// game2 := NewGame(p2)
	// game2.DoMoveCurrent()
	p := &EasyBot{}
	g := NewGame(p)

	g.DoMoveCurrent()
}

// 1. Инкапсуляция на уровне пакетов

// структуру можно использовать в других пакетах
type SomethingAvailableInOtherPackages struct {
}

// структуру нельзя использовать в других пакетах помимо main
type somethingNotAvailable struct {
}

// 2. Наследование
// в го есть встраивание структур, но нет полноценного наследования

type A struct {
}

func (a *A) SomeMethod() {

}

type B struct {
	A
}

type Game struct {
	player Player
}

func (g *Game) DoMoveCurrent() {
	g.player.DoMove()
}

func NewGame(p Player) *Game {
	return &Game{player: p}
}

type PlayerImpl struct {
	Name string
}

type Player interface {
	DoMove()
	Info()
}

type PlayerManual struct {
}

func (p *PlayerManual) DoMove() {
	fmt.Println("обычный игрок ходит как обычно")
}

type EasyBot struct {
}

func (p *EasyBot) DoMove() {
	fmt.Println("бот поступает не очень умно")
}

// так делать не надо
// func (p *Player) DoMove() {
// 	switch p.Name {
// 	case "Bot_Easy":
// 		fmt.Println("бот поступает не очень умно")
// 	case "Bot_Hard":
// 		fmt.Println("бот ходит как Илон Маск")
// 	default:
// 		fmt.Println("игрок подумал и сделал очень важный шаг ")
// 	}
// }

// func (p *Player) Info() {
// 	switch p.Name {
// 	case "Bot_Easy":
// 		fmt.Println("привет, я лёгкий бот и не расстроюсь при поражении")
// 	case "Bot_Hard":
// 		fmt.Println("я очень умный бот и ты будешь страдать")
// 	default:
// 		fmt.Println("обычный игрок, такой как все")
// 	}
// }
