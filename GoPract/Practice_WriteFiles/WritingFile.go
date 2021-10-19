package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("demo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString("Rab Ne Bana Di Jodi (transl. A Match Made By God), also known as RNBDJ, is a 2008 Indian Hindi-language romantic comedy film written and directed by Aditya Chopra and produced by Yash Chopra and Aditya Chopra under the banner Yash Raj Films. The film stars Shah Rukh Khan and Anushka Sharma, who makes her Bollywood debut with this film.)")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
