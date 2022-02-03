package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func PlacingShip(thisField [10][11]string) [10][11]string {
	var rand_i, rand_j int
	isFind := false          //для проверки подходящей точки на поле
	for i := 0; i < 4; i++ { //ставим одинарные корабли
		rand_i = rand.Intn(10)
		rand_j = rand.Intn(10)

		for rand_j == 0 {
			rand_j = rand.Intn(10)
		}

		for CheckField(thisField, rand_i, rand_j) == false {
			rand_i = rand.Intn(10)
			rand_j = rand.Intn(10)

			fmt.Println("i", rand_i)
			for rand_j == 0 {
				rand_j = rand.Intn(10)
			}
			fmt.Println("j", rand_j)
		}

		if CheckField(thisField, rand_i, rand_j) {
			fmt.Println("ставим i", rand_i+1)
			fmt.Println("ставим j", rand_j)
			thisField[rand_i][rand_j] = "□"
		}
	}
	for i := 0; i < 3; i++ { // ставим 2-хпалубники
		rand_i = rand.Intn(10)
		rand_j = rand.Intn(10)
		for rand_j == 0 {
			rand_j = rand.Intn(10)
		}
		for CheckField(thisField, rand_i, rand_j) == false { //неправильный цикл

			rand_i = rand.Intn(10)
			rand_j = rand.Intn(10)

			fmt.Println("i", rand_i)
			for rand_j == 0 {
				rand_j = rand.Intn(10)
			}
			fmt.Println("j", rand_j)

		}
		for isFind == false {
			for i := rand_i; rand_i-1 <= i+1; rand_i += 2 { // делаем проверку для второго отсека корабля || что делать с диагональю?
				fmt.Println("Итерируем i")
				for j := rand_j; rand_j <= j; rand_j++ {
					if CheckField(thisField, rand_i-1, rand_j-1) == true {
						isFind = true
						thisField[rand_i-1][rand_j] = "□"
						thisField[i][j] = "□"
						fmt.Println("Выходим, нашли")
						break
					}

				}
				if isFind {
					fmt.Println("Выходим, нашли")
					break
				}

			}
		}
		isFind = false
		/*
			if CheckField(thisField, rand_i, rand_j) {
				fmt.Println("ставим i", rand_i+1)
				fmt.Println("ставим j", rand_j)
				thisField[rand_i][rand_j] = "□"
			}
		*/

	}
	return thisField
}

func CheckField(thisField [10][11]string, i int, j int) bool { //проверяем поля для правильности размещения кораблей
	checkPoint := true
	check_i := i - 1
	check_j := j - 1
	if check_i == -1 { //проверяем края полей, || убедиться, что логика верная
		check_i++
	} else if i == 9 {
		println("изменили i")
		i--
	}
	if check_j == 0 {
		check_j++
	} else if j == 10 {
		println("изменили j")
		j--
	}
	println("Проверяем ", i+1, " ", j+1)
	println("Вошли")
	for min_i := check_i; min_i <= i+1; min_i++ {
		for min_j := check_j; min_j <= j+1; min_j++ {
			if thisField[min_i][min_j] != "#" {
				checkPoint = false
				println("вернули false")
				return checkPoint
			}
		}

	}
	//fmt.Println(thisField)
	return checkPoint
}

func PlayersField() [10][11]string {
	var field [10][11]string
	for i := 0; i < 10; i++ {
		for j := 1; j < 11; j++ {
			field[i][j] = "#"
		}
	}
	for i := 0; i < 10; i++ {
		field[i][0] = strconv.Itoa(i + 1)
	}
	return field

}

func FieldDraw(thisField [10][11]string) {
	fmt.Println("   1 2 3 4 5 6 7 8 9 10")
	for i := 0; i < 10; i++ {

		fmt.Println(thisField[i])
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	/*
		var n int
		fmt.Scanf("%d\n", &n)

		scanner := bufio.NewScanner(os.Stdin)

		rights := make([]int, 0)
		lefts := make([]int, 0)

		scanner.Scan()

		numsStr := strings.Split(scanner.Text(), " ")
		nums := make([]int, 0, len(numsStr))
	*/
	MyField := PlayersField()
	FieldDraw(MyField)
	MyField = PlacingShip(MyField)
	FieldDraw(MyField)

	//fmt.Println(rand.Intn(10))
}
