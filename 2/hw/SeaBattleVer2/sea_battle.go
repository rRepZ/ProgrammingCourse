package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type ShotResult int

const (
	HIT ShotResult = iota
	SINK
	MISS
)

type CellStatus int

const (
	SHOT CellStatus = iota
	NEAR_SHIP
)

type Orientation int

const (
	HORIZONTAL = iota
	VERTICAL
)

type cmdHandler func(string) string

type tempCelsius int

func (t *tempCelsius) toFarenheit() int {
	// todo
	return 123
}

type cell struct {
	ship   *ship
	status *CellStatus
}

type ship struct {
	name        string
	x           int
	y           int
	decks       []int
	orientation Orientation
}

// todo
func newShip(name string, x, y int) *ship {

}

func (s *ship) shot() {

}

type field struct {
	cells [][]*cell
}

func (f *field) shot(x, y int) ShotResult {
	// здесь обработка выстрела
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	p1 := &player{}
	p2 := &player{}
	p1.enemy = p2
	p2.enemy = p1

	game := newGame(p1, p2, p1)

	// старт сервера
	for {
		s.Scan()
		cmd := s.Text()

		handler, err := ValidateAndParse(cmd, game)
		if err != nil {
			fmt.Printf("invalid input: %s", err.Error())
			continue
		}

		output := handler(cmd)
		fmt.Println(output)
	}
}

func ValidateAndParse(input string, game *game) (cmdHandler, error) {
	if len(input) < 2 {
		return nil, fmt.Errorf("string length should be > 2")
	}

	switch input {
	case "status":
		// todo
		// return
	}

	// todo make extensible
	if err := ValidateShoot(input); err != nil {
		return nil, err
	} else {
		return game.HandleShoot, nil
	}

	return nil, fmt.Errorf("unknown command")
}

// ValidateShoot валидация команды выстрела
func ValidateShoot(input string) error {
	x := rune(input[0])
	// todo заглавные буквы тоже принимать на вход
	if x < rune('a') || x > rune('z') {
		return fmt.Errorf("invalid letter, should be from range [a-z]")
	}

	_, err := strconv.Atoi(input[1:])
	if err != nil {
		return err
	}

	return nil
}

type player struct {
	name       string
	enemy      *player
	stepsCount int

	field field
}

func (p *player) doMove(x, y int) (result ShotResult, fieldAfterShot [][]int8) {
	result, fieldAfterShot = p.enemy.getShot(x, y)
	return
}

func (p *player) getShot(x, y int) (result ShotResult, fieldAfterShot [][]int8) {
	res := p.field.shot(x, y)
	return res, p.field
}

type game struct {
	player1 *player
	player2 *player

	currentPlayer *player
}

func (g *game) HandleShoot(input string) string {
	x := rune(input[0])
	y, _ := strconv.Atoi(input[1:])

	res, field := g.currentPlayer.doMove(x, y)
	// todo преобразовать в строковое представление и вернуть
}

func newGame(p1, p2, curr *player) *game {
	// todo create fields
	return &game{
		player1:       p1,
		player2:       p2,
		currentPlayer: curr,
	}
}

func newGameWithFirstMove(firstMove int) *game {
	// todo create fields
	return &game{
		fields:            fields,
		currentPlayerMove: firstMove,
	}
}

func mapResult(result ShotResult) string {
	// switch по результатам и их строковое представление
}
