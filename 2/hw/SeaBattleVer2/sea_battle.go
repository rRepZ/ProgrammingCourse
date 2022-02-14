package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type ShotResult int

const (
	HIT ShotResult = iota
	SINK
	MISS
)

type CellStatus int

const (
	FREE CellStatus = iota
	SHOT
	SHIP
	NEAR_SHIP
	ATTACKED
)

type Orientation int

const (
	HORIZONTAL Orientation = iota
	VERTICAL
)

type DeckStatus int

const (
	OK DeckStatus = iota
	DEAD
)

type cmdHandler func(string) string

type cell struct {
	ship   *ship
	status CellStatus
}

//func NewCell()
type ship struct {
	name        string
	x           int //крайняя левая или верхняя координата
	y           int
	decks       []int
	orientation Orientation
	health      int
}

// todo
func NewShip(name string /*x, y int,*/, shipSize int, hp int) *ship { //подумать над shipSize
	thisShip := new(ship)
	thisShip.name = name
	thisShip.health = hp
	//thisShip.x = x
	//thisShip.y = y
	for i := 0; i < shipSize; i++ {
		thisShip.decks = append(thisShip.decks, i)
	}
	return &ship{
		name: thisShip.name,
		//x:     thisShip.x,
		//y:     thisShip.y,
		decks: thisShip.decks,
	}
}

func ShipCreating(thisField *field) {
	//PlayerShipOneDeck1 := new(ship)
	//PlayerShipOneDeck2 := new(ship)
	PlayerShipOneDeck1 := NewShip("Arnold", 1, 1) //создаём корабли с одной палубой
	PlayerShipOneDeck2 := NewShip("Gregory", 1, 1)
	PlayerShipOneDeck3 := NewShip("Marston", 1, 1)
	PlayerShipOneDeck4 := NewShip("Momo", 1, 1)
	PlayerShipTwoDeck1 := NewShip("Tompson", 2, 2) //создаём корабли с двумя палубами
	PlayerShipTwoDeck2 := NewShip("Flower", 2, 2)
	PlayerShipTwoDeck3 := NewShip("Fisherman", 2, 2)
	PlayerShipThreeDeck1 := NewShip("Bob", 3, 3) // //создаём корабли с тремя палубами
	PlayerShipThreeDeck2 := NewShip("Viper", 3, 3)
	PlayerShipFourDeck := NewShip("Chrono", 4, 4) //создаём корабли с 4мя палубами
	ShipBuilder(PlayerShipOneDeck1, thisField)
	ShipBuilder(PlayerShipOneDeck2, thisField)
	ShipBuilder(PlayerShipOneDeck3, thisField)
	ShipBuilder(PlayerShipOneDeck4, thisField)
	ShipBuilder(PlayerShipTwoDeck1, thisField)
	ShipBuilder(PlayerShipTwoDeck2, thisField)
	ShipBuilder(PlayerShipTwoDeck3, thisField)
	ShipBuilder(PlayerShipThreeDeck1, thisField)
	ShipBuilder(PlayerShipThreeDeck2, thisField)
	ShipBuilder(PlayerShipFourDeck, thisField)
	fmt.Println(PlayerShipOneDeck1)
	fmt.Println(PlayerShipOneDeck2)
	fmt.Println(PlayerShipOneDeck3)
	fmt.Println(PlayerShipOneDeck4)
	fmt.Println(PlayerShipTwoDeck1)
	fmt.Println(PlayerShipTwoDeck2)
	fmt.Println(PlayerShipTwoDeck3)
	fmt.Println(PlayerShipThreeDeck1)
	fmt.Println(PlayerShipThreeDeck2)
	fmt.Println(PlayerShipFourDeck)
	fmt.Println(len(thisField.cells))

}

func ShipBuilder(thisShip *ship, thisField *field) {
	var rand_i, rand_j int
	isFind := false
	shipSize := len(thisShip.decks) - 1 // для прохождения по нужному количеству клеток
	//for i := 0; i < 1; i++ {
	for isFind == false {
		rand_i = rand.Intn(10)
		rand_j = rand.Intn(10)
		fmt.Println("вошли")
		for CheckField(thisField, rand_i, rand_j) == false {
			fmt.Println(CheckField(thisField, rand_i, rand_j))

			rand_i = rand.Intn(10)
			rand_j = rand.Intn(10)

		}
		fmt.Println(CheckField(thisField, rand_i, rand_j))

		for i := rand_i - shipSize; i <= rand_i+shipSize; i += shipSize { // делаем проверку для второго отсека корабля (сверху вниз)
			if i >= len(thisField.cells) { //i стало равно 10 (выход за пределы поля)

				break //выход, дальше некуда итерировать
			}
			if i <= -1 { //проверка выхода за область
				i = rand_i //чтобы проверить с правого конца (для больших кораблей)
			}

			switch {
			case i != rand_i:
				//j := rarand_j //когда просмтариваем сверху или снизу, то j статична
				if CheckField(thisField, i, rand_j) == true {
					isFind = true
					switch {
					case i < rand_i:
						for ; i < rand_i; i++ {
							thisField.cells[i][rand_j].status = SHIP
							thisField.cells[i][rand_j].ship = thisShip
							thisShip.orientation = VERTICAL
						}
						thisShip.x = rand_j
						thisShip.y = i

					case i > rand_i:
						thisShip.x = rand_j
						thisShip.y = i
						for ; i > rand_i; i-- {
							thisField.cells[i][rand_j].status = SHIP
							thisField.cells[i][rand_j].ship = thisShip
							thisShip.orientation = VERTICAL
						}
					}

					break
				}

			case i == rand_i:
				for j := rand_j - shipSize; j <= rand_j+shipSize; j += 2 * shipSize {

					if j <= -1 {
						j = rand_j + shipSize
					}
					if j >= len(thisField.cells) { //когда проверяет границы
						break
					}

					if CheckField(thisField, i, j) == true {
						fmt.Println(i, j)
						fmt.Println(rand_i, rand_j)
						isFind = true
						switch {
						case j < rand_j:

							for ; j < rand_j; j++ {
								thisField.cells[i][j].status = SHIP
								thisField.cells[i][j].ship = thisShip
								thisShip.orientation = HORIZONTAL
							}
							thisShip.x = j
							thisShip.y = i
							fmt.Println("координаты (идём вправо)", i, j)
							//fmt.Println("координаты", i, j)
						case j > rand_j:
							thisShip.x = j
							thisShip.y = i
							fmt.Println("координаты идём влево", i, j)
							for ; j > rand_j; j-- {
								thisField.cells[i][j].status = SHIP
								thisField.cells[i][j].ship = thisShip
								thisShip.orientation = HORIZONTAL
							}

							//здесь записывать координаты
						}

						break
					}

				}
			}

			if isFind {

				break
			}

		}

	}
	if len(thisShip.decks) == 1 {
		thisShip.x = rand_j
		thisShip.y = rand_i
	}
	thisField.cells[rand_i][rand_j].ship = thisShip
	thisField.cells[rand_i][rand_j].status = SHIP
	PointAround(thisField, thisShip)

}

func PointAround(thisField *field, thisShip *ship) {
	this_i := thisShip.y
	j := thisShip.x
	fmt.Println(thisShip)
	switch {
	case thisShip.orientation == VERTICAL:
		if this_i+1 != len(thisField.cells) {
			thisField.cells[this_i+1][j].status = NEAR_SHIP
			fmt.Println("вертикальный")
			if j-1 != -1 { //ставим на диагональные квадраты
				thisField.cells[this_i+1][j-1].status = NEAR_SHIP
			}
			if j+1 != len(thisField.cells) {
				thisField.cells[this_i+1][j+1].status = NEAR_SHIP
			}

		}
		for lenShip := len(thisShip.decks); lenShip > 0; lenShip-- {

			if j+1 != len(thisField.cells) {
				thisField.cells[this_i][j+1].status = NEAR_SHIP
			}
			if j-1 != -1 {
				thisField.cells[this_i][j-1].status = NEAR_SHIP
			}
			if this_i-1 != -1 {
				this_i--
				if thisField.cells[this_i][j].status != SHIP {
					thisField.cells[this_i][j].status = NEAR_SHIP
					if j-1 != -1 { //ставим на диагональные квадраты
						thisField.cells[this_i][j-1].status = NEAR_SHIP
					}
					if j+1 != len(thisField.cells) {
						thisField.cells[this_i][j+1].status = NEAR_SHIP
					}
				}
			}
		}

		this_i = thisShip.y
		j = thisShip.x
	case thisShip.orientation == HORIZONTAL:
		fmt.Println("горизонтальный")
		//this_j := j
		if j+1 != len(thisField.cells) {
			thisField.cells[this_i][j+1].status = NEAR_SHIP

			if this_i-1 != -1 { //ставим на диагональные квадраты
				thisField.cells[this_i-1][j+1].status = NEAR_SHIP
			}
			if this_i+1 != len(thisField.cells) {
				thisField.cells[this_i+1][j+1].status = NEAR_SHIP
			}

		}
		for lenShip := len(thisShip.decks); lenShip > 0; lenShip-- {

			if this_i+1 != len(thisField.cells) {
				thisField.cells[this_i+1][j].status = NEAR_SHIP
			}
			if this_i-1 != -1 {
				thisField.cells[this_i-1][j].status = NEAR_SHIP
			}
			if j-1 != -1 {
				j--
				if thisField.cells[this_i][j].status != SHIP {
					thisField.cells[this_i][j].status = NEAR_SHIP
					if this_i-1 != -1 { //ставим на диагональные квадраты
						thisField.cells[this_i-1][j].status = NEAR_SHIP
					}
					if this_i+1 != len(thisField.cells) {
						thisField.cells[this_i+1][j].status = NEAR_SHIP
					}
				}
			}
		}
	}
	//return thisField

}

func CheckField(thisField *field, i int, j int) bool { //проверяем поля для правильности размещения кораблей
	checkPoint := true

	check_i := i - 1
	check_j := j - 1
	if check_i == -1 { //проверяем края полей
		check_i++
	} else if i == len(thisField.cells)-1 {

		i--
	}
	if check_j == -1 {
		check_j++
	} else if j == len(thisField.cells)-1 {

		j--
	}

	for min_i := check_i; min_i <= i+1; min_i++ {
		for min_j := check_j; min_j <= j+1; min_j++ {
			if thisField.cells[min_i][min_j].status == SHIP {
				checkPoint = false

				return checkPoint
			}
		}

	}

	return checkPoint
}

func (s *ship) shot() { //метод для игрока??

}

type field struct {
	cells [][]*cell
}

func NewField(fieldSize int) *field {

	thisField := new(field)
	thisField.cells = make([][]*cell, 0, fieldSize)
	for i := 0; i < fieldSize; i++ {
		arr := make([]*cell, 0, fieldSize)
		thisField.cells = append(thisField.cells, arr)
		for j := 0; j < fieldSize; j++ {
			thisField.cells[i] = append(thisField.cells[i], new(cell))
		}
	}
	return &field{cells: thisField.cells}
}

func DrawField(f *field) {

	for i := 0; i < len(f.cells); i++ {
		for j := 0; j < len(f.cells); j++ {
			fmt.Print(f.cells[i][j].status)
		}
		fmt.Println()
	}
}

func (thisField *field) shot(i, j int) ShotResult { //return shotResult?
	var resultOfShot ShotResult
	// здесь обработка выстрела
	if thisField.cells[i][j].status == FREE {
		thisField.cells[i][j].status = SHOT
		resultOfShot = MISS
		return resultOfShot
	}
	if thisField.cells[i][j].status == SHIP {
		thisField.cells[i][j].status = ATTACKED
		if thisField.cells[i][j].ship.health-1 > 0 {
			thisField.cells[i][j].ship.health--
			resultOfShot = HIT
			return resultOfShot
		} else {
			thisField.cells[i][j].ship.health--
			PointAround(thisField, thisField.cells[i][j].ship)
			resultOfShot = SINK
			return resultOfShot
		}

	}

	return MISS
}

/*
func (thisField *field) IsDead() {
	thisField.cells[i][j].ship
}
*/

func main() {
	rand.Seed(time.Now().UnixNano())
	s := bufio.NewScanner(os.Stdin)

	p1 := &player{}
	p2 := &player{}
	p1.enemy = p2
	p2.enemy = p1

	game := NewGame(p1, p2, p1)
	p1Field := new(field)
	p1Field = NewField(10)
	ShipCreating(p1Field)
	DrawField(p1Field)

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
	/*
		if err := ValidateShoot(input); err != nil {
			return nil, err
		} else {
			return game.HandleShoot, nil
		}
	*/
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

	playerField field
}

/*
func (p *player) doMove(x, y int) (result ShotResult, fieldAfterShot [][]int8) {
	result, fieldAfterShot = p.enemy.getShot(x, y)
	return
}

func (p *player) getShot(x, y int) (result ShotResult, fieldAfterShot [][]int8) {
	res := p.playerField.shot(x, y)
	return res, p.playerField
}
*/
type game struct {
	player1 *player
	player2 *player

	currentPlayer *player
}

/*
func (g *game) HandleShoot(input string) string {
	x := rune(input[0])
	y, _ := strconv.Atoi(input[1:])

	res, field := g.currentPlayer.doMove(x, y)
	// todo преобразовать в строковое представление и вернуть
}
*/
func NewGame(p1, p2, curr *player) *game {
	// todo create fields
	return &game{
		player1:       p1,
		player2:       p2,
		currentPlayer: curr,
	}
}

/*
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
*/
