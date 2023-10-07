package os

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteSimpleString(filename string, message string) error {

	// To start, here’s how to dump a string (or just bytes) into a file.
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile(filename, d1, 0644)
	check(err)
	return err
}

// You can Write byte slices as you’d expect.
func WriteByteSlices(filename string, message string) {
	f, err := os.Create(filename)
	check(err)
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// Issue a Sync to flush writes to stable storage.
	f.Sync()

	// It’s idiomatic to defer a Close immediately after opening a file.
	defer f.Close()

}

// A WriteString is also available.
func WriteString(filename string, message string) {

	f, err := os.Create(filename)
	check(err)
	n3, err := f.WriteString(message)
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// Issue a Sync to flush writes to stable storage.
	f.Sync()
}

// bufio provides buffered writers in addition to the buffered readers we saw earlier.
func WriteBuffer(filename string, message string) {
	f, err := os.Create(filename)
	check(err)
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	//Use Flush to ensure all buffered operations have been applied to the underlying writer.
	w.Flush()
}
