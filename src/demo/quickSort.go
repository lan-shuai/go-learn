package demo

import (
	"fmt"
)

func quickSort(arr []int, left, right int){
	if(left >= right){
		return;
	}
	i, j := left, right

	key := arr[i];

	for i < j {
		for arr[i] < key {
			i++
		}
		for arr[j] > key {
			j--
		}
		arr[i], arr[j] = arr[j], arr[i]
 	}
	arr[i] = key

	quickSort(arr, left, i - 1);
	quickSort(arr, i + 1, right)
}

func main(){
	arr := []int{1,3,5,7,9,2,4,6,8,0};

	quickSort(arr, 0, len(arr) - 1);

	fmt.Println(arr)

}