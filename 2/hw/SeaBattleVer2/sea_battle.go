package main

// TODO все инициализируется в конструкторах, зависимости  передаются в параметры конструктора
// TODO не мутировать состояния по ходу действия программы, инициализировать новые объекты
// TODO отрефакторить main, чтобы соответствовать тому как я написал
// TODO убрать принты по ходу программы
// TODO имена переменных использовать краткие, ёмкие, не использовать snake_case(rand_j) => randJ, thisField => field или f
// TODO быть уверенным, что используются все конструкторы(NewField, NewCell)
// TODO избавиться от повисших в воздухе функций, всё сделать методами
// TODO начать переписывать на интерфейсы

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
	// todo чем shot отличается от attacked?
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
	ship   *shipImpl
	status CellStatus
}

func NewCell(ship *shipImpl, status CellStatus) *cell {
	return &cell{
		ship:   ship,
		status: status,
	}
}

/*
type Ship interface {
	IsFit(f *field) bool
	GetShot(x, y int)
	GetX() int
	GetY() int
	// и так далее...
}
*/

type shipImpl struct {
	name        string
	x           int //крайняя левая или верхняя координата
	y           int
	decks       []int
	orientation Orientation
	health      int
}

// todo
func NewShip(name string /*x, y int,*/, shipSize int, hp int) *shipImpl { //подумать над shipSize
	// some code
	thisShip := new(shipImpl)
	thisShip.name = name
	thisShip.health = hp
	//thisShip.x = x
	//thisShip.y = y
	for i := 0; i < shipSize; i++ {
		thisShip.decks = append(thisShip.decks, i)
	}
	return &shipImpl{
		name:   thisShip.name,
		decks:  thisShip.decks,
		health: thisShip.health,
	}
}

//IsFit проверка возможности размещения
/*
func (*shipImpl) IsFit(f *field) bool {

}
*/
func (*shipImpl) GetShot(x, y int) {

}

//массив имён //перемешать в случайном порядке
// TODO создавать в цикле согласно указанному флоу
func ShipCreating(f *field) { //CreateShips
	for i := 0; i < 4; i++ {
		for j := i; j < 4; j++ {
			for {
				s := NewShip("Arnold", i+1, i+1)

				f.AddShip(s)

				break

			}
		}
	}
	//PlayerShipOneDeck1 := new(shipImpl)
	//PlayerShipOneDeck2 := new(shipImpl)
	/*
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
		ShipBuilder(PlayerShipOneDeck1, f)
		ShipBuilder(PlayerShipOneDeck2, f)
		ShipBuilder(PlayerShipOneDeck3, f)
		ShipBuilder(PlayerShipOneDeck4, f)
		ShipBuilder(PlayerShipTwoDeck1, f)
		ShipBuilder(PlayerShipTwoDeck2, f)
		ShipBuilder(PlayerShipTwoDeck3, f)
		ShipBuilder(PlayerShipThreeDeck1, f)
		ShipBuilder(PlayerShipThreeDeck2, f)
		ShipBuilder(PlayerShipFourDeck, f)
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
	*/

}

func (f *field) AddShip(thisShip *shipImpl) { //rename to
	var rand_i, rand_j int
	isFind := false
	shipSize := len(thisShip.decks) - 1 // для прохождения по нужному количеству клеток
	//for i := 0; i < 1; i++ {
	for isFind == false {
		rand_i = rand.Intn(10)
		rand_j = rand.Intn(10)
		fmt.Println("вошли")
		for !CheckField(f, rand_i, rand_j) {
			fmt.Println(CheckField(f, rand_i, rand_j))

			rand_i = rand.Intn(10)
			rand_j = rand.Intn(10)

		}
		fmt.Println(CheckField(f, rand_i, rand_j))

		for i := rand_i - shipSize; i <= rand_i+shipSize; i += shipSize { // делаем проверку для второго отсека корабля (сверху вниз)
			if i >= len(f.cells) { //i стало равно 10 (выход за пределы поля)

				break //выход, дальше некуда итерировать
			}
			if i <= -1 { //проверка выхода за область
				i = rand_i //чтобы проверить с правого конца (для больших кораблей)
			}

			switch {
			case i != rand_i:
				//j := rarand_j //когда просмтариваем сверху или снизу, то j статична
				if CheckField(f, i, rand_j) == true {
					isFind = true
					switch {
					case i < rand_i:
						for ; i < rand_i; i++ {
							// todo не мутировать состояние ячейки, создавать всегда новую
							thisShip.orientation = VERTICAL
							f.cells[i][rand_j] = NewCell(thisShip, SHIP)
							//thisField.cells[i][rand_j].status = SHIP
							//thisField.cells[i][rand_j].ship = thisShip
							//thisShip.orientation = VERTICAL
						}
						thisShip.x = rand_j
						thisShip.y = i

					case i > rand_i:
						thisShip.x = rand_j
						thisShip.y = i
						for ; i > rand_i; i-- {
							thisShip.orientation = VERTICAL
							f.cells[i][rand_j] = NewCell(thisShip, SHIP)
							/*
								thisField.cells[i][rand_j].status = SHIP
								thisField.cells[i][rand_j].ship = thisShip
							*/
						}
					}

					break
				}

			case i == rand_i:
				for j := rand_j - shipSize; j <= rand_j+shipSize; j += 2 * shipSize {

					if j <= -1 {
						j = rand_j + shipSize
					}
					if j >= len(f.cells) { //когда проверяет границы
						break
					}

					if CheckField(f, i, j) == true {
						fmt.Println(i, j)
						fmt.Println(rand_i, rand_j)
						isFind = true
						switch {
						case j < rand_j:

							for ; j < rand_j; j++ {
								thisShip.orientation = HORIZONTAL
								f.cells[i][j] = NewCell(thisShip, SHIP)
								/*
									f.cells[i][j].status = SHIP
									f.cells[i][j].ship = thisShip
								*/
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
								thisShip.orientation = HORIZONTAL
								f.cells[i][j] = NewCell(thisShip, SHIP)
								/*
									thisField.cells[i][j].status = SHIP
									thisField.cells[i][j].ship = thisShip
								*/
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

	f.cells[rand_i][rand_j] = NewCell(thisShip, SHIP)
	/*
		thisField.cells[rand_i][rand_j].ship = thisShip
		thisField.cells[rand_i][rand_j].status = SHIP
	*/
	PointAround(f, thisShip)

}

func PointAround(f *field, thisShip *shipImpl) {
	this_i := thisShip.y
	j := thisShip.x
	fmt.Println(thisShip)
	switch {
	case thisShip.orientation == VERTICAL:
		if this_i+1 != len(f.cells) {
			f.cells[this_i+1][j].status = NEAR_SHIP
			fmt.Println("вертикальный")
			if j-1 != -1 { //ставим на диагональные квадраты
				f.cells[this_i+1][j-1].status = NEAR_SHIP
			}
			if j+1 != len(f.cells) {
				f.cells[this_i+1][j+1].status = NEAR_SHIP
			}

		}
		for lenShip := len(thisShip.decks); lenShip > 0; lenShip-- {

			if j+1 != len(f.cells) {
				f.cells[this_i][j+1].status = NEAR_SHIP
			}
			if j-1 != -1 {
				f.cells[this_i][j-1].status = NEAR_SHIP
			}
			if this_i-1 != -1 {
				this_i--
				if f.cells[this_i][j].status != SHIP {
					f.cells[this_i][j].status = NEAR_SHIP
					if j-1 != -1 { //ставим на диагональные квадраты
						f.cells[this_i][j-1].status = NEAR_SHIP
					}
					if j+1 != len(f.cells) {
						f.cells[this_i][j+1].status = NEAR_SHIP
					}
				}
			}
		}

		this_i = thisShip.y
		j = thisShip.x
	case thisShip.orientation == HORIZONTAL:
		fmt.Println("горизонтальный")
		//this_j := j
		if j+1 != len(f.cells) {
			f.cells[this_i][j+1].status = NEAR_SHIP

			if this_i-1 != -1 { //ставим на диагональные квадраты
				f.cells[this_i-1][j+1].status = NEAR_SHIP
			}
			if this_i+1 != len(f.cells) {
				f.cells[this_i+1][j+1].status = NEAR_SHIP
			}

		}
		for lenShip := len(thisShip.decks); lenShip > 0; lenShip-- {

			if this_i+1 != len(f.cells) {
				f.cells[this_i+1][j].status = NEAR_SHIP
			}
			if this_i-1 != -1 {
				f.cells[this_i-1][j].status = NEAR_SHIP
			}
			if j-1 != -1 {
				j--
				if f.cells[this_i][j].status != SHIP {
					f.cells[this_i][j].status = NEAR_SHIP
					if this_i-1 != -1 { //ставим на диагональные квадраты
						f.cells[this_i-1][j].status = NEAR_SHIP
					}
					if this_i+1 != len(f.cells) {
						f.cells[this_i+1][j].status = NEAR_SHIP
					}
				}
			}
		}
	}
	//return thisField

}

func CheckField(f *field, i int, j int) bool { //проверяем поля для правильности размещения кораблей
	checkPoint := true

	check_i := i - 1
	check_j := j - 1
	if check_i == -1 { //проверяем края полей
		check_i++
	} else if i == len(f.cells)-1 {

		i--
	}
	if check_j == -1 {
		check_j++
	} else if j == len(f.cells)-1 {

		j--
	}

	for min_i := check_i; min_i <= i+1; min_i++ {
		for min_j := check_j; min_j <= j+1; min_j++ {
			if f.cells[min_i][min_j].status == SHIP {
				checkPoint = false

				return checkPoint
			}
		}

	}

	return checkPoint
}

func (s *shipImpl) shot() { //метод для игрока??

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

func FieldToDraw(fieldSize int) [][]string {
	playerField := make([][]string, 0, fieldSize)
	for i := 0; i < fieldSize; i++ {
		arr := make([]string, 0, fieldSize)
		playerField = append(playerField, arr)
		for j := 0; j < fieldSize; j++ {
			playerField[i] = append(playerField[i], "#")
		}
	}
	return playerField
}

func (f *field) DrawPlayerField(playerField [][]string) {

	for i := 0; i < len(f.cells); i++ {
		for j := 0; j < len(f.cells); j++ {
			if f.cells[i][j].status == FREE {
				playerField[i][j] = "#"
			}
			if f.cells[i][j].status == NEAR_SHIP || f.cells[i][j].status == SHOT {
				playerField[i][j] = "X"
			}
			if f.cells[i][j].status == SHIP {
				playerField[i][j] = "□"
			}
			if f.cells[i][j].status == ATTACKED {
				playerField[i][j] = "⧆"
			}

		}
		fmt.Println(playerField[i])
	}
}

func (f *field) shot(i, j int) ShotResult { //return shotResult?
	var resultOfShot ShotResult
	fmt.Println("проверка")
	fmt.Println("f.cells[i][j].status", f.cells)

	// здесь обработка выстрела
	if f.cells[i][j].status == FREE || f.cells[i][j].status == NEAR_SHIP {
		fmt.Println("проверка 1")
		f.cells[i][j].status = SHOT
		resultOfShot = MISS
		fmt.Println("проверка 1")
		fmt.Println(resultOfShot)
		return resultOfShot
	}
	if f.cells[i][j].status == SHIP {
		f.cells[i][j].status = ATTACKED
		fmt.Println("проверка2")
		if f.cells[i][j].ship.health-1 > 0 {
			f.cells[i][j].ship.health--
			resultOfShot = HIT
			return resultOfShot
		} else {
			fmt.Println("проверка3")
			f.cells[i][j].ship.health--
			PointAround(f, f.cells[i][j].ship)
			resultOfShot = SINK
			return resultOfShot
		}

	}
	fmt.Println("выход")
	return MISS
}

func main() {
	rand.Seed(time.Now().UnixNano())
	s := bufio.NewScanner(os.Stdin)

	p1 := &player{}
	p2 := &player{}
	p1Field := new(field)
	p1Field = NewField(10)
	p2Field := new(field)
	p2Field = NewField(10)
	ShipCreating(p1Field)
	ShipCreating(p2Field)
	p1 = NewPlayer("Player", p2, p1Field)
	p2 = NewPlayer("Bot", p1, p2Field)

	game := NewGame(p1, p2, p1)
	DrawField(p1Field)
	DrawField(p2Field)

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
	pfield := FieldToDraw(10)
	if len(input) < 2 {
		return nil, fmt.Errorf("string length should be > 2")
	}

	switch input {
	case "status":
		game.player1.playerField.DrawPlayerField(pfield)

		return nil, fmt.Errorf("status was checked") //???
	}

	// todo make extensible

	if err := ValidateShoot(input); err != nil {
		return nil, err
	} else {
		return game.HandleShoot, nil
	}

	//return nil, fmt.Errorf("unknown command")
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

	playerField *field
}

func NewPlayer(name string, enemy *player, f *field) *player {
	return &player{
		name:        name,
		enemy:       enemy,
		stepsCount:  0,
		playerField: f,
	}

}

func (p *player) doMove(x, y int) (result ShotResult, fieldAfterShot *field) {
	fmt.Println("nil? ", fieldAfterShot)
	result, fieldAfterShot = p.enemy.getShot(x, y)
	fmt.Println(result, fieldAfterShot)
	return result, fieldAfterShot
}

func (p *player) getShot(x, y int) (result ShotResult, fieldAfterShot *field) {
	res := p.playerField.shot(x, y)
	fmt.Println("nil2? ", res, p.playerField)
	return res, p.playerField
}

type game struct {
	player1 *player
	player2 *player

	currentPlayer *player
}

func (g *game) HandleShoot(input string) string {
	//x := rune(input[0])
	x := int([]rune(input)[0] - []rune("a")[0])
	//x, _ := strconv.Atoi(input[:1])
	fmt.Println("X", x)
	y, _ := strconv.Atoi(input[1:])
	fmt.Println("y", y)
	res, field := g.currentPlayer.doMove(y, x)
	fmt.Println(field)
	// todo преобразовать в строковое представление и вернуть
	if res == SINK {
		return "ранен"
	}
	fmt.Println("return", y)
	return input
}

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
