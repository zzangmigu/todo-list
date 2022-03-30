package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	inputToDo string
	toDoSlice []string
)

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
func addToDo() {
	fmt.Println("새로 추가할 계획을 작성해 주세요")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	inputToDo = stdin.Text()
	toDoSlice = append(toDoSlice, inputToDo)

	fmt.Printf("입력한 계획 [%s]이/가 저장되었습니다\n", inputToDo)
}

func displayToDo() {
	// tempLen := len(toDoSlice)
	// for i := 0; i < tempLen-1; i++ {
	// }
	fmt.Println(toDoSlice)
}

func deleteToDo() {
	fmt.Println("삭제할 계획을 작성해 주세요")
	stdin2 := bufio.NewScanner(os.Stdin)
	stdin2.Scan()
	inputToDo = stdin2.Text()
	var foundNum int
	for i := range toDoSlice {
		if toDoSlice[i] == inputToDo {
			toDoSlice = append(toDoSlice[:i], toDoSlice[i+1:]...)
			fmt.Printf("입력한 계획 [%s]이/가 삭제되었습니다\n", inputToDo)
			foundNum = 1
			break
		}
	}
	if foundNum != 1 {
		fmt.Println("해당 계획이 없습니다. 메인으로 돌아갑니다")
	}
}

func editToDo() {

}

func clearToDo() {

}

var choosedNum int

func main() {
T:
	for {
		fmt.Println("0000000000000000000000")
		fmt.Println("0     Choose Menu    0")
		fmt.Println("0  1. Write ToDo     0")
		fmt.Println("0  2. Display ToDo   0")
		fmt.Println("0  3. Delete ToDo    0")
		fmt.Println("0  4. Edit ToDO      0")
		fmt.Println("0  5. Clear ToDo     0")
		fmt.Println("0  6. Close Menu     0")
		fmt.Println("0000000000000000000000")

		fmt.Scanln(&choosedNum)
		switch choosedNum {
		case 1:
			addToDo()
		case 2:
			displayToDo()
		case 3:
			deleteToDo()
		case 4:
			editToDo()
		case 5:
			clearToDo()
		case 6:
			break T
		default:
			fmt.Println("올바른 메뉴 번호를 입력해 주세요 ")
		}
	}

}
