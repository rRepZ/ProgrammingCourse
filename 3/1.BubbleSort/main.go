package bubblesort

func BubbleSort(arr []int) {
	wasSwap := true // когда не будет ни одной перестановки -> значит отсортирована
	for wasSwap {
		wasSwap = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				wasSwap = true
			}
		}
	}
}

func main() {

	arr := []int{1, 2, 3, 4, 5}
	BubbleSort(arr)
}
