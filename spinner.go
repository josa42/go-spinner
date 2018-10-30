package spinner

import (
	"fmt"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/buger/goterm"
)

var Chars = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

func New(message string) *Spinner {
	s := &Spinner{
		message:  message,
		chars:    Chars,
		lock:     &sync.RWMutex{},
		stopChan: make(chan struct{}, 1),
	}
	s.start()
	return s
}

type Spinner struct {
	message     string
	delay       time.Duration
	isRunning   bool
	chars       []string
	lock        *sync.RWMutex
	stopChan    chan struct{}
	deleteCount int
}

func (s *Spinner) Message(message string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.message = message
}

func (s *Spinner) Chars(chars []string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.chars = chars
}

func (s *Spinner) start() *Spinner {
	if s.isRunning {
		return s
	}
	s.isRunning = true

	s.hideCursor()

	go func() {
		for {
			for i := 0; i < len(s.chars); i++ {
				select {
				case <-s.stopChan:
					return
				default:
					s.lock.Lock()

					s.erase()
					s.print(goterm.Color(s.chars[i], goterm.BLUE))
					s.resetCursor()

					s.lock.Unlock()

					time.Sleep(100 * time.Millisecond)
				}
			}
		}
	}()

	return s
}

func (s *Spinner) stop(c string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.isRunning {
		s.isRunning = false
		s.stopChan <- struct{}{}

		s.erase()
		s.print(c)
		fmt.Println("")
		s.showCursor()
	}
}

func (s *Spinner) Done() {
	s.stop(goterm.Color("√", goterm.GREEN))
}

func (s *Spinner) Fail() {
	s.stop(goterm.Color("✖", goterm.RED))
}

func (s *Spinner) Next(message string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.isRunning {
		s.print(goterm.Color("√", goterm.GREEN))
		fmt.Println("")

		s.message = message
	}
}

func (s *Spinner) print(c string) {
	fmt.Printf("%s %s", goterm.Bold(c), s.message)
}

func (s *Spinner) erase() {
	fmt.Printf("\033[K")
}

func (s *Spinner) resetCursor() {
	fmt.Printf("\033[%dD", utf8.RuneCountInString(s.message)+2)
}

func (s *Spinner) hideCursor() {
	fmt.Print("\u001b[?25l")
}

func (s *Spinner) showCursor() {
	fmt.Print("\u001b[?25h")
}
