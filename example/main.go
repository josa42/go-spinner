package main

import (
	"fmt"
	"time"

	spinner "github.com/josa42/go-spinner"
)

func main() {
	fmt.Print("\u001b[?25l")
	time.Sleep(1 * time.Second)
	fmt.Println("")

	spinner4()

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	// fmt.Println("\nSpinner 1:")
	// spinner1()

	// time.Sleep(1 * time.Second)
	// fmt.Println("\nSpinner 2:")
	// spinner2()

	// time.Sleep(1 * time.Second)
	// fmt.Println("\nSpinner 3:")
	// spinner3()

	// time.Sleep(1 * time.Second)
	// fmt.Println("\nSpinner 4:")
	// spinner4()

	// time.Sleep(1 * time.Second)
}

func spinner1() {
	s := spinner.New("Loading something")

	time.Sleep(2 * time.Second)
	s.Done()
}

func spinner2() {
	s := spinner.New("Loading something")

	time.Sleep(2 * time.Second)
	s.Fail()
}

func spinner3() {
	s := spinner.New("Loading...")

	time.Sleep(1 * time.Second)
	s.Message("Still loading...")

	time.Sleep(1 * time.Second)
	s.Message("Almost done")

	time.Sleep(1 * time.Second)
	s.Message("Done")
	s.Done()
}

func spinner4() {
	s := spinner.New("One")

	time.Sleep(1 * time.Second)
	s.Next("Two")
	time.Sleep(1 * time.Second)

	s.Next("Three")
	time.Sleep(1 * time.Second)
	s.Done()
}
