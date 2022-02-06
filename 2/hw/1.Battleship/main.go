package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
) //⧆

func GetShot(thisField [10][10]string, i int, j int) [10][10]string {

	if thisField[i][j] == "#" {
		thisField[i][j] = "❋"
	}
	if thisField[i][j] == "□" {
		thisField[i][j] = "⧆"
		if CheckDefeat(thisField, i, j) {
			thisField = PointAround(thisField, i, j)
		}

	}
	return thisField
}

func PointAround(thisField [10][10]string, i int, j int) [10][10]string {
	this_i := i
	CheckOrientHorizontal := true
	if j-1 != -1 {
		if thisField[this_i][j-1] == "⧆" {
			CheckOrientHorizontal = false
		}
	}
	if j+1 != 10 {
		if thisField[this_i][j+1] == "⧆" {
			CheckOrientHorizontal = false
		}
	}
	fmt.Println("Check = ", CheckOrientHorizontal)
	switch {
	case CheckOrientHorizontal:
		for thisField[this_i][j] == "⧆" {
			if j+1 != 10 {
				thisField[this_i][j+1] = "X"
			}
			if j-1 != -1 {
				thisField[this_i][j-1] = "X"
			}

			if this_i+1 != 10 {
				this_i++
				if thisField[this_i][j] != "⧆" {
					thisField[this_i][j] = "X"
					if j-1 != -1 { //ставим на диагональные квадраты
						thisField[this_i][j-1] = "X"
					}
					if j+1 != 10 {
						thisField[this_i][j+1] = "X"
					}
				}

			} else {
				break
			}
		}

		this_i = i

		for thisField[this_i][j] == "⧆" {
			if j+1 != 10 {
				thisField[this_i][j+1] = "X"
			}
			if j-1 != -1 {
				thisField[this_i][j-1] = "X"
			}
			if this_i-1 != -1 {
				this_i--
				if thisField[this_i][j] != "⧆" { //верхний край
					thisField[this_i][j] = "X"
					if j-1 != -1 { //ставим на диаогнальные квадраты
						thisField[this_i][j-1] = "X"
					}
					if j+1 != 10 {
						thisField[this_i][j+1] = "X"
					}

				}
			} else {
				break
			}
		}
	default:

		this_j := j
		for thisField[i][this_j] == "⧆" { //а если слева направо?
			if i+1 != 10 {
				thisField[i+1][this_j] = "X"
			}
			if i-1 != -1 {
				thisField[i-1][this_j] = "X"
			}

			if this_j+1 != 10 {
				this_j++
				if thisField[i][this_j] != "⧆" {
					thisField[i][this_j] = "X"
					if i-1 != -1 { //ставим на диагональные квадраты
						thisField[i-1][this_j] = "X"
					}
					if i+1 != 10 {
						thisField[i+1][this_j] = "X"
					}
				}

			} else {
				break
			}
		}

		this_j = j

		for thisField[i][this_j] == "⧆" { //а если слева направо?
			if i+1 != 10 {
				thisField[i+1][this_j] = "X"
			}
			if i-1 != -1 {
				thisField[i-1][this_j] = "X"
			}
			if this_j-1 != -1 {
				this_j--
				if thisField[i][this_j] != "⧆" { //верхний край
					thisField[i][this_j] = "X"
					if i-1 != -1 { //ставим на диаогнальные квадраты
						thisField[i-1][this_j] = "X"
					}
					if i+1 != 10 {
						thisField[i+1][this_j] = "X"
					}

				}
			} else {
				break
			}
		}
	}
	return thisField

}

func ShipBuilder(thisField [10][10]string, numberOfShips int, shipSize int) [10][10]string {
	var rand_i, rand_j int
	isFind := false
	shipSize = shipSize - 1 // для прохождения по нужному количеству клеток
	for i := 0; i < numberOfShips; i++ {
		for isFind == false {
			rand_i = rand.Intn(10)
			rand_j = rand.Intn(10)
			for CheckField(thisField, rand_i, rand_j) == false {

				rand_i = rand.Intn(10)
				rand_j = rand.Intn(10)

			}

			for i := rand_i - shipSize; i <= rand_i+shipSize; i += shipSize { // делаем проверку для второго отсека корабля (сверху вниз)
				if i >= 10 { //i стало равно 10 (выход за пределы поля)

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
								thisField[i][rand_j] = "□"
							}
						case i > rand_i:
							for ; i > rand_i; i-- {
								thisField[i][rand_j] = "□"
							}
						}

						break
					}

				case i == rand_i:
					for j := rand_j - shipSize; j <= rand_j+shipSize; j += 2 * shipSize {

						if j <= -1 {
							j = rand_j + shipSize
						}
						if j >= 10 { //когда проверяет границы
							break
						}

						if CheckField(thisField, i, j) == true {
							isFind = true
							switch {
							case j < rand_j:
								for ; j < rand_j; j++ {
									thisField[i][j] = "□"
								}
							case j > rand_j:
								for ; j > rand_j; j-- {
									thisField[i][j] = "□"
								}
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
		isFind = false
		thisField[rand_i][rand_j] = "□"

	}
	return thisField
}

func PlacingShip(thisField [10][10]string) [10][10]string {

	thisField = ShipBuilder(thisField, 4, 1)
	thisField = ShipBuilder(thisField, 3, 2)
	thisField = ShipBuilder(thisField, 2, 3)
	thisField = ShipBuilder(thisField, 1, 4)
	return thisField
}

func CheckDefeat(thisField [10][10]string, i int, j int) bool { //делаем проверку, полностью ли побеждён корабль
	checkPoint := true

	this_i := i
	for thisField[this_i][j] == "⧆" {
		if this_i+1 != 10 {
			this_i++
		} else {
			break
		}

		if thisField[this_i][j] == "□" {

			checkPoint = false
			return checkPoint
		}

	}
	this_i = i
	for thisField[this_i][j] == "⧆" {
		if this_i-1 != -1 {
			this_i--
		} else {
			break
		}

		if thisField[this_i][j] == "□" {

			checkPoint = false
			return checkPoint
		}

	}

	this_j := j
	for thisField[i][this_j] == "⧆" {
		if this_j-1 != -1 {
			this_j--
		} else {
			break
		}

		if thisField[i][this_j] == "□" {

			checkPoint = false
			return checkPoint
		}

	}
	this_j = j
	for thisField[i][this_j] == "⧆" {
		if this_j+1 != 10 {
			this_j++
		} else {
			break
		}

		if thisField[i][this_j] == "□" {

			checkPoint = false
			return checkPoint
		}
	}

	return checkPoint
}

func CheckField(thisField [10][10]string, i int, j int) bool { //проверяем поля для правильности размещения кораблей
	checkPoint := true

	check_i := i - 1
	check_j := j - 1
	if check_i == -1 { //проверяем края полей
		check_i++
	} else if i == 9 {

		i--
	}
	if check_j == -1 {
		check_j++
	} else if j == 9 {

		j--
	}

	for min_i := check_i; min_i <= i+1; min_i++ {
		for min_j := check_j; min_j <= j+1; min_j++ {
			if thisField[min_i][min_j] != "#" {
				checkPoint = false

				return checkPoint
			}
		}

	}

	return checkPoint
}

func PlayersField() [10][10]string {

	var field [10][10]string
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			field[i][j] = "#"
		}
	}

	return field
}

func FillHidden(thisField [10][10]string, hiddenField [10][10]string) [10][10]string {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			switch {
			case thisField[i][j] == "⧆":
				hiddenField[i][j] = "⧆"
			case thisField[i][j] == "X":
				hiddenField[i][j] = "X"
			case thisField[i][j] == "❋":
				hiddenField[i][j] = "❋"
			}
		}
	}
	return hiddenField
}

func IsWin(thisField [10][10]string) bool {
	defShips := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if thisField[i][j] == "⧆" {
				defShips++
			}
		}
	}
	if defShips == 20 {
		return true
	}
	return false
}

func FieldDraw(thisField [10][10]string) {
	fmt.Println("   a b c d e f g h i j")
	for i := 0; i < 10; i++ {
		fmt.Printf(" %d", i)
		fmt.Println(thisField[i])
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var n int
	isExit := false
	//fmt.Scanf("%d", &n)

	isLoop := false
	MyField := PlayersField()
	BotField := PlayersField()
	HiddenField := PlayersField()
	fmt.Println("menu ")
	fmt.Println("1 -- to start battleship game; 2 -- to shot; 3 -- to clear ur field and restart; 5 -- to see help info")
	fmt.Println("6 -- to exit")
	for isExit == false {
		if isLoop {
			fmt.Println("choose 2 to continue, 5 for help, 7 to see ur field")
		}
		fmt.Scanf("%d", &n)
		switch {

		case n == 1:

			if isLoop {
				fmt.Println("U want restart game? Try to select option № 3")
				break
			}
			MyField = PlacingShip(MyField)
			BotField = PlacingShip(BotField)
			fmt.Println("My field ")
			FieldDraw(MyField)
			//fmt.Println("Bot field ")

			FieldDraw(HiddenField)
			isLoop = true
		case n == 2:

			if isLoop {
				NextMove := true
				for NextMove {
					NextMove = false
					if IsWin(BotField) {
						fmt.Println("Вы Победили!!!")
						break
					} else if IsWin(MyField) {
						fmt.Println("Вы Проиграли!!!")
						break
					}
					//==================================
					scanner := bufio.NewScanner(os.Stdin)
					fmt.Println("Введите значение поля в формате 'a1'...")
					scanner.Scan()
					StrToParse := strings.Split(scanner.Text(), "")

					y, _ := strconv.ParseInt(StrToParse[0], 20, 32)
					if y > 20 || y < 10 {
						fmt.Println("Ошибка ввода, попробуйте ещё раз")
						break
					}
					x, _ := strconv.ParseInt(StrToParse[1], 20, 32)
					if x > 10 || x < 0 {
						fmt.Println("Такого поля не существует, попробуйте ещё раз")
						break
					}
					//fmt.Println(x)
					//fmt.Println(y)
					//===================================
					BotField = GetShot(BotField, int(x), int(y-10))
					HiddenField = FillHidden(BotField, HiddenField)
					if BotField[int(x)][int(y-10)] == "⧆" {
						fmt.Println("Nice shot!!!")
						if CheckDefeat(BotField, int(x), int(y-10)) {
							fmt.Println("Bot ship was destroyed!!!!")
						}

						NextMove = true
					}
					FieldDraw(HiddenField)

				}
				NextMove = true
				var rand_i, rand_j int
				for NextMove {
					NextMove = false
					if IsWin(BotField) {
						//fmt.Println("Вы Победили!!!")
						break
					} else if IsWin(MyField) {
						fmt.Println("Вы Проиграли!!!")
						break
					}
					rand_i = rand.Intn(10)
					rand_j = rand.Intn(10)
					for MyField[rand_i][rand_j] != "□" && MyField[rand_i][rand_j] != "#" {
						rand_i = rand.Intn(10)
						rand_j = rand.Intn(10)
					}
					MyField = GetShot(MyField, rand_i, rand_j)
					if MyField[rand_i][rand_j] == "⧆" {
						fmt.Println("Ur ship was attecked!!!")
						if CheckDefeat(MyField, rand_i, rand_j) {
							fmt.Println("your ship was destroyed!!!!")
						}
						NextMove = true
					}

				}

				//FieldDraw(MyField)

			} else {
				fmt.Println("At first u must start a game, choose 1")
			}
		case n == 3:
			MyField = PlayersField()
			BotField = PlayersField()
			HiddenField = PlayersField()
			isLoop = false
		case n == 4:
			fmt.Println("Cheats on =)")
			FieldDraw(BotField)
		case n == 5:
			fmt.Println("1 -- to start battleship game; 2 -- to shot; 3 -- to clear ur field and restart; 5 -- to see help info")
			fmt.Println("6 -- to exit")
		case n == 6:
			isExit = true
		case n == 7:
			FieldDraw(MyField)
		}
	}

	//fmt.Println(rand.Intn(10))
}
