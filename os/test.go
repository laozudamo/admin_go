package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Readfile()

	//Writefile()

	//OpenFile()

	//readDir()

	handleFile()
}

func Readfile() {
	content, err := os.ReadFile("os/test.txt")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(string(content))

}

func Writefile() {
	content := []byte("Hello jack")
	err := os.WriteFile("os/tes.txt", content, 0644)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(string(content))
}

func OpenFile() {
	file, err := os.OpenFile("os/test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer file.Close()
	//
	//contentByte, err := os.ReadFile("os/test.txt")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(contentByte))

	//defer file.Close()

	//err := file.Close()
	//if err != nil {
	//	fmt.Println("Error: ", err)
	//}
	//for {
	//	b := make([]byte, 10)
	//	n, err := file.Read(b)
	//	if err != nil {
	//		// 读到文件末尾 EOF
	//		fmt.Println("Error: ", err)
	//		return
	//	}
	//	fmt.Println(string(b), n)
	//}
	//file.WriteString

	// seek 光标移动
	//file.Seek(0, io.SeekEnd)
	//file.Seek(0, io.SeekStart)
	//file.WriteString("Hello world")
	//r := bufio.NewReader(file)
}

func readDir() {
	d, _ := os.ReadDir("os")
	//fmt.Println(d, err)
	for _, v := range d {
		fmt.Println(v.Name(), v.IsDir())
	}
}

func handleFile() {
	f, err := os.OpenFile("os/test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	reader := bufio.NewReader(f)
	n := 0
	for {
		n++
		//str, _, err := reader.ReadLine()

		//str, err := reader.ReadString('\n')

		str, err := reader.ReadString('\n')
		writer.WriteString(strconv.Itoa(n) + "" + str)
		if err != nil {
			break
		}
	}
	f.Seek(0, 0)
	writer.Flush()
}
