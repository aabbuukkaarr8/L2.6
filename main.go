package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"} // len - 3 cap - 3
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3" // s = ["3", "2", "3"] len -3 cap -3
	// в итоге на наш массив повлияет только первая строчка, так как после нее, мы увеличили cap - создали новый массив
	// который существует только внутри функции modifySlice, и будет удален из стека после конца функции.
	i = append(i, "4") // new massive s = ["3", "2", "3", "4"] len - 4 cap - 6
	i[1] = "5"
	i = append(i, "6")
}
