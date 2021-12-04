package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

func main() {
	readFromStdInAndWriteTosStdOut()
}

func readFromStdInAndWriteTosStdOut() {
	r, w := io.Pipe()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer r.Close()

		buf, err := io.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.WriteString(os.Stdout, "<<< processed "+string(buf))
		if err != nil {
			log.Fatal(err)
		}

	}()

	fi, _ := os.Stdin.Stat()

	if (fi.Mode() & os.ModeCharDevice) == 0 {
		fmt.Println("data is from pipe")

		bytes, _ := io.ReadAll(os.Stdin)
		_, err := w.Write(bytes)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		fmt.Println(">>> data is from terminal")

		ConsoleReader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter input data : ")

		input, err := ConsoleReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		_, err = w.Write([]byte(input))
		if err != nil {
			log.Fatal(err)
		}

	}

	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()
}
