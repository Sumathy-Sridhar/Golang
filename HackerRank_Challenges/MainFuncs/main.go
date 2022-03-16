package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//var str string
	//fmt.Scanf("%s", &str)
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	fmt.Printf("Hello, World.\n%s", str)

}
