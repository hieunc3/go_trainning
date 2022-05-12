package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
)

func main() {
	//readFromFile()
	//writeFromFile()
	data := make(chan int)
	done := make(chan bool)
	wg := sync.WaitGroup{}

	//create 100 producer goroutines
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce(data, &wg)
	}

	go consume(data, done)

	go func() {
		wg.Wait()
		close(data)
	}()
	d := <-done

	if d {
		fmt.Println("File written successfully")
	} else {
		fmt.Println("File writing failed")
	}
}

func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}

func consume(data chan int, done chan bool) {
	f, err := os.Create("concurrent")
	if err != nil {
		fmt.Println(err)
		f.Close()
		done <- false
		return
	}

	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}

func readFromFile() {
	//data, err := ioutil.ReadFile("text1.txt")

	// data, err := ioutil.ReadFile("F:\\go_workspace\\go_step_1\\go-basic\\go_tranning_day1\\file_handling\\text1.txt")
	// if err != nil {
	// 	fmt.Println("File reading error", err)
	// 	return
	// }

	//the ReadFile func returns a byte slice -> we need to convert byte slice to string in order to display
	//fmt.Println("Contents of file: ", string(data))

	//read by line
	fptr := flag.String("fpath", "text1.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}

	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func writeFromFile() {
	//Creates a file named text2.txt. If a file with that name already exists, then the create func truncates the file
	// f, err := os.Create("text2.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// //use write method to write a slice of byte
	// l, err := f.WriteString("Write text to file1")
	// if err != nil {
	// 	fmt.Println(err)
	// 	f.Close()
	// 	return
	// }
	// fmt.Println(l, "bytes written successfully")
	// err = f.Close()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	//write strings line by line to a file
	// f, err := os.Create("lines")
	// if err != nil {
	// 	fmt.Println(err)
	// 	f.Close()
	// 	return
	// }
	// d := []string{"Welcome to the world of Go1.", "Go is a compiled language.", "It is easy to learn Go."}

	// for _, v := range d {
	// 	fmt.Fprintln(f, v)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// }
	// err = f.Close()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("file written successfully")

	//apepending to a file
	//we open the file in append & write only mode
	// f, err := os.OpenFile("lines", os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// newLine := "Append newline to file."
	// _, err = fmt.Fprintln(f, newLine)
	// if err != nil {
	// 	fmt.Println(err)
	// 	f.Close()
	// 	return
	// }
	// err = f.Close()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("file appended successfully")
}
