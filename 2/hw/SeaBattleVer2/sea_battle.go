package main

// TODO все инициализируется в конструкторах, зависимости  передаются в параметры конструктора
// TODO не мутировать состояния по ходу действия программы, инициализировать новые объекты
// TODO отрефакторить main, чтобы соответствовать тому как я написал
// TODO убрать принты по ходу программы
// TODO имена переменных использовать краткие, ёмкие, не использовать snake_case(jRand) => randJ, f => field или f
// TODO быть уверенным, что используются все конструкторы(NewField, NewCell)
// TODO избавиться от повисших в воздухе функций, всё сделать методами
// TODO начать переписывать на интерфейсы

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	fSize int = 10
)

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
	GetShot(x, y int)
	GetX() int
	GetY() int
	// и так далее...
}
*/

type ShipNames struct {
	Ships []string `json:"ships"`
}

type shipImpl struct {
	name        string
	x           int //крайняя левая или верхняя координата
	y           int
	decks       []int //изменить в boolgi
	orientation Orientation
	health      int
}

func NewShip(name string, shipSize int, hp int) *shipImpl { //подумать над shipSize
	decks := make([]int, 0, shipSize)
	for i := 0; i < shipSize; i++ {
		decks = append(decks, i)
	}

	return &shipImpl{
		name:   name,
		decks:  decks,
		health: hp,
	}
}

/*
func (*shipImpl) GetShot(x, y int) {

}
*/

func (f *field) FillWithRandomShips() { //CreateShips
	dataShip, err := ioutil.ReadFile("./input/input1.json")
	if err != nil {
		log.Fatalf("Can't read file: %s", err)
		return
	}
	var ships ShipNames
	err = json.Unmarshal(dataShip, &ships)
	if err != nil {
		log.Fatalf("Marshal error: %s", err)
		return
	}
	for i := 3; i >= 0; i-- {
		for j := i; j < 4; j++ {
			for {
				s := NewShip(ships.Ships[f.shipsOnField], i+1, i+1)
				f.AddShipIfFits(s)
				f.shipsOnField++
				break

			}
		}
	}

}

func (f *field) AddShipIfFits(s *shipImpl) { //rename to
	var iRand, jRand int
	freeSpaceForShip := false
	shipSize := len(s.decks) - 1 // для прохождения по нужному количеству клеток

	for !freeSpaceForShip {
		iRand = rand.Intn(fSize)
		jRand = rand.Intn(fSize)

		for !f.CheckField(iRand, jRand) {

			iRand = rand.Intn(fSize)
			jRand = rand.Intn(fSize)
		}

		for i := iRand - shipSize; i <= iRand+shipSize; i += shipSize { // делаем проверку для второго отсека корабля (сверху вниз)
			if i >= len(f.cells) { //i стало равно 10 (выход за пределы поля)
				break //выход, дальше некуда итерировать
			}

			if i <= -1 { //проверка выхода за область
				i = iRand //чтобы проверить с правого конца (для больших кораблей)
			}

			switch {
			case i != iRand:
				//j := rajRand //когда просмтариваем сверху или снизу, то j статична
				if f.CheckField(i, jRand) {
					freeSpaceForShip = true
					switch {
					case i < iRand:
						for ; i < iRand; i++ {
							// todo не мутировать состояние ячейки, создавать всегда новую
							s.orientation = VERTICAL
							f.cells[i][jRand] = NewCell(s, SHIP)
							//f.cells[i][jRand].status = SHIP
							//f.cells[i][jRand].ship = s
							//s.orientation = VERTICAL
						}
						s.x = jRand
						s.y = i

					case i > iRand:
						s.x = jRand
						s.y = i
						for ; i > iRand; i-- {
							s.orientation = VERTICAL
							f.cells[i][jRand] = NewCell(s, SHIP)
							/*
								f.cells[i][jRand].status = SHIP
								f.cells[i][jRand].ship = s
							*/
						}
					}

					break
				}

			case i == iRand:
				for j := jRand - shipSize; j <= jRand+shipSize; j += 2 * shipSize {

					if j <= -1 {
						j = jRand + shipSize
					}
					if j >= len(f.cells) { //когда проверяет границы
						break
					}

					if f.CheckField(i, j) {

						freeSpaceForShip = true
						switch {
						case j < jRand:
							for ; j < jRand; j++ {
								s.orientation = HORIZONTAL
								f.cells[i][j] = NewCell(s, SHIP)
							}
							s.x = j
							s.y = i
						case j > jRand:
							s.x = j
							s.y = i
							for ; j > jRand; j-- {
								s.orientation = HORIZONTAL
								f.cells[i][j] = NewCell(s, SHIP)
							}
						}

						break
					}

				}
			}

			if freeSpaceForShip {
				break
			}
		}

	}
	if len(s.decks) == 1 {
		s.x = jRand
		s.y = iRand
	}

	f.cells[iRand][jRand] = NewCell(s, SHIP)
	/*
		f.cells[iRand][jRand].ship = s
		f.cells[iRand][jRand].status = SHIP
	*/
	//PointAround(f, s)

}

func (f *field) pointAround(s *shipImpl) {
	this_i := s.y
	j := s.x
	fmt.Println(s)
	switch {
	case s.orientation == VERTICAL:
		if this_i+1 != len(f.cells) {
			f.cells[this_i+1][j].status = NEAR_SHIP

			if j-1 != -1 { //ставим на диагональные квадраты
				f.cells[this_i+1][j-1].status = NEAR_SHIP
			}
			if j+1 != len(f.cells) {
				f.cells[this_i+1][j+1].status = NEAR_SHIP
			}

		}
		for lenShip := len(s.decks); lenShip > 0; lenShip-- {

			if j+1 != len(f.cells) {
				f.cells[this_i][j+1].status = NEAR_SHIP
			}
			if j-1 != -1 {
				f.cells[this_i][j-1].status = NEAR_SHIP
			}
			if this_i-1 != -1 {
				this_i--
				if f.cells[this_i][j].status != ATTACKED {
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

		this_i = s.y
		j = s.x
	case s.orientation == HORIZONTAL:
		if j+1 != len(f.cells) {
			f.cells[this_i][j+1].status = NEAR_SHIP

			if this_i-1 != -1 { //ставим на диагональные квадраты
				f.cells[this_i-1][j+1].status = NEAR_SHIP
			}
			if this_i+1 != len(f.cells) {
				f.cells[this_i+1][j+1].status = NEAR_SHIP
			}

		}
		for lenShip := len(s.decks); lenShip > 0; lenShip-- {

			if this_i+1 != len(f.cells) {
				f.cells[this_i+1][j].status = NEAR_SHIP
			}
			if this_i-1 != -1 {
				f.cells[this_i-1][j].status = NEAR_SHIP
			}
			if j-1 != -1 {
				j--
				if f.cells[this_i][j].status != ATTACKED {
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
	//return f

}

func (f *field) CheckField(i int, j int) bool { //проверяем поля для правильности размещения кораблей
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
	cells        [][]*cell
	shipsOnField int
}

func NewField(fieldSize int) *field {
	f := new(field)
	f.cells = make([][]*cell, 0, fieldSize)
	for i := 0; i < fieldSize; i++ {
		arr := make([]*cell, 0, fieldSize)
		f.cells = append(f.cells, arr)
		for j := 0; j < fieldSize; j++ {
			f.cells[i] = append(f.cells[i], new(cell))
		}
	}
	return &field{
		cells:        f.cells,
		shipsOnField: 0,
	}
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

// todo возвращать string, вместо вывода в консоль
func (f *field) DrawPlayerField(playerField [][]string, isHidden bool) /* string */ {
	switch {
	case !isHidden:
		fmt.Println("  a b c d e f g h i j")
		for i := 0; i < len(f.cells); i++ {
			fmt.Printf("%d", i)
			// todo то же самое, только результат получается в string
			// fmt.Sprintf("%d", i)
			for j := 0; j < len(f.cells); j++ {
				switch f.cells[i][j].status {
				case FREE:
					playerField[i][j] = "#"
				case NEAR_SHIP:
					fallthrough
				case SHOT:
					playerField[i][j] = "X"
				case SHIP:
					playerField[i][j] = "□"
				case ATTACKED:
					playerField[i][j] = "⧆"
				}

			}
			fmt.Println(playerField[i])
		}
	case isHidden:
		// todo аналогично см.выше
		fmt.Println("  a b c d e f g h i j")
		for i := 0; i < len(f.cells); i++ {
			fmt.Printf("%d", i)
			for j := 0; j < len(f.cells); j++ {
				if f.cells[i][j].status == FREE {
					playerField[i][j] = "#"
				}
				if f.cells[i][j].status == NEAR_SHIP || f.cells[i][j].status == SHOT {
					playerField[i][j] = "X"
				}
				if f.cells[i][j].status == SHIP {
					playerField[i][j] = "#"
				}
				if f.cells[i][j].status == ATTACKED {
					playerField[i][j] = "⧆"
				}

			}
			fmt.Println(playerField[i])
		}
	}

}

func (f *field) shot(i, j int) ShotResult { //return shotResult?
	var resultOfShot ShotResult

	// fmt.Println("f.cells[i][j].status", f.cells)

	// здесь обработка выстрела

	// TODO делегировать изменения состояния корабля - кораблю
	if f.cells[i][j].status == FREE || f.cells[i][j].status == NEAR_SHIP {

		f.cells[i][j].status = SHOT
		resultOfShot = MISS

		fmt.Println(resultOfShot)
		return resultOfShot
	}
	if f.cells[i][j].status == SHIP {
		// TODO
		// result := f.cells[i][j].ship.Shot()
		// return result
		f.cells[i][j].status = ATTACKED

		if f.cells[i][j].ship.health-1 > 0 {
			f.cells[i][j].ship.health--
			resultOfShot = HIT
			return resultOfShot
		} else {

			f.cells[i][j].ship.health--
			f.pointAround(f.cells[i][j].ship)
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

	f1 := NewField(fSize)
	f2 := NewField(fSize)
	f1.FillWithRandomShips()
	f2.FillWithRandomShips()

	p1 := NewPlayerNoEnemy("Player", f1)
	p2 := NewPlayer("Bot", p1, f2)
	p1.enemy = p2

	game := NewGame(p1, p2, p1)
	enemyF := FieldToDraw(fSize)

	var cmd, output string

	// старт сервера
	for {
		isContinue := true
		for isContinue {
			//if game.currentPlayer.name != "Bot" {
			fmt.Println("")
			game.player2.playerField.DrawPlayerField(enemyF, true)

			// TODO принты только вот в этом месте программы, больше нигде не нужно жёстко привязываться к консоли
			// вывод чужого поля
			//fmt.Println(game.GetCurrentPlayerEnemyFields())

			s.Scan()
			cmd = s.Text()
			handler, err := ValidateAndParse(cmd, game)
			if err != nil {
				fmt.Printf("invalid input: %s \n", err.Error())
				continue
			}

			output = handler(cmd)
			//} else {
			//	output = game.HandleShoot("")
			//}

			// TODO здесь весь вывод должен происходить
			fmt.Println(output)

			if output == "Промах!" {
				isContinue = false
			}

			if output == "Победа!" {
				return
			}
		}
		game.SwitchPlayer(p1, p2)
	}
}

func ValidateAndParse(input string, game *game) (cmdHandler, error) {
	if len(input) < 2 {
		return nil, fmt.Errorf("string length should be > 2\n")
	}

	switch input {
	case "status":
		return game.HandleStatus, nil
	}

	// todo make extensible

	if err := ValidateShoot(input); err != nil {
		return nil, err
	} else {
		shoot := game.HandleShoot
		return shoot, nil
	}

	//return nil, fmt.Errorf("unknown command")
}

// ValidateShoot валидация команды выстрела
func ValidateShoot(input string) error {
	x := rune(input[0])
	// todo заглавные буквы тоже принимать на вход
	if x < rune('a') || x > rune('j') {
		return fmt.Errorf("invalid letter, should be from range [a-z]\n")
	}

	num, err := strconv.Atoi(input[1:])
	if err != nil {
		return err
	}
	if num > 9 || num < 0 {
		return fmt.Errorf("invalid number, should be from range [0-9]\n")
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

func NewPlayerNoEnemy(name string, f *field) *player {
	return &player{
		name:        name,
		stepsCount:  0,
		playerField: f,
	}
}

func (p *player) doMove(x, y int) (result ShotResult, fieldAfterShot *field) {

	result, fieldAfterShot = p.enemy.getShot(x, y)

	return result, fieldAfterShot
}

func (p *player) getShot(x, y int) (result ShotResult, fieldAfterShot *field) {
	res := p.playerField.shot(x, y)

	return res, p.playerField
}

type game struct {
	player1 *player
	player2 *player

	currentPlayer *player
}

func (g *game) HandleShoot(input string) string {
	var x, y int
	//if g.currentPlayer.name != "Bot" {
	x = int([]rune(input)[0] - []rune("a")[0])
	//x, _ := strconv.Atoi(input[:1])
	fmt.Println("X", x)
	y, _ = strconv.Atoi(input[1:])
	fmt.Println("Y", y)

	/*	} else {
		fmt.Println("Ход бота ", y)
		x = rand.Intn(fSize)
		y = rand.Intn(fSize)
		for g.currentPlayer.enemy.playerField.cells[y][x].status != FREE && g.currentPlayer.enemy.playerField.cells[y][x].status != SHIP {
			x = rand.Intn(fSize)
			y = rand.Intn(fSize)
		}
	}*/
	res, _ := g.currentPlayer.doMove(y, x) //res, fiels
	//.Println(field)
	// todo преобразовать в строковое представление и вернуть
	if res == SINK {
		//g.currentPlayer.enemy.playerField.pointAround(g.currentPlayer.enemy.playerField.cells[y][x].ship)
		g.currentPlayer.enemy.playerField.shipsOnField--

		if g.currentPlayer.enemy.playerField.shipsOnField == 0 {
			return "Победа!"
		} else {
			return "Корабль убит!"
		}
	}
	if res == HIT {
		return "Попадание!"
	}
	if res == MISS {
		return "Промах!"
	}

	return input
}

func (g *game) HandleStatus(input string) string {
	pfield := FieldToDraw(fSize)
	g.player1.playerField.DrawPlayerField(pfield, false)

	// todo преобразовать в string
	return "status"
}

func (g *game) SwitchPlayer(p1 *player, p2 *player) {
	switch {
	case g.currentPlayer == p1:
		g.currentPlayer = p2
		break
	case g.currentPlayer == p2:
		g.currentPlayer = p1
		break
	}
}
func NewGame(p1, p2, curr *player) *game {

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
