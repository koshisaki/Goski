package scan

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Splitline() (int, []int) {
	scanner := bufio.NewScanner(os.Stdin)

	// 最初の1文字目はn なので先に受け取り、integerに変換しておく
	scanner.Scan()
	n := scanner.Text()
	N, _ := strconv.Atoi(n)
	fmt.Printf("N: %d \n", N)

	var val []int
	for i := 0; i < N; i++ {
		scanner.Scan()
		vals := scanner.Text()
		vali, _ := strconv.Atoi(vals)
		val = append(val, vali)
	}

	fmt.Println(val)
	fmt.Println("scan end!")

	return N, val
}
