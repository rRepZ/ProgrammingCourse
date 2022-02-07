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

type cmdHandler func(string) string

type field [][]int

func main() {
	s := bufio.NewScanner(os.Stdin)
	game := newGame()

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

// HandleShoot обработка выстрела

type game struct {
	fields            []field
	currentPlayerMove int
}

func (g *game) HandleShoot(input string) string {
	x := rune(input[0])
	y, _ := strconv.Atoi(input[1:])

	// взять все игровые поля, кроме g.fields[g.currentPlayerMove]
	// написать сюда логику обработки выстрела по этим полям

	var output string
	for _, f := range fieldsToShoot {
		// попал не попал и всё такое
		// shotResult := g.handleShootInternal()
		// output += mapResult(shotResult)
	}

	// отмаппить результат выполнения команды на строковое представление
}

func newGame() *game {
	// todo create fields
	return &game{
		fields: fields,
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
