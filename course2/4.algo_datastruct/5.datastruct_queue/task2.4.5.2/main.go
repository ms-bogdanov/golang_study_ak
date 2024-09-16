package main

import "fmt"

type BrowserHistory struct {
	stack []string
}

func (h *BrowserHistory) Visit(url string) {
	fmt.Println("Visiting ", url)
	h.stack = append(h.stack, url)
}

func (h *BrowserHistory) Back() {
	if len(h.stack) < 2 {
		fmt.Println("No more adresses to return")
		return
	}

	lastStackElemIndex := len(h.stack) - 1
	fmt.Println("Back to ", h.stack[lastStackElemIndex-1])
	h.stack = h.stack[:lastStackElemIndex]
}

func (h *BrowserHistory) PrintHistory() {
	if len(h.stack) == 0 {
		fmt.Println("No history")
		return
	}

	fmt.Println("Browser history")

	for lastStackElemIndex := len(h.stack) - 1; lastStackElemIndex >= 0; lastStackElemIndex-- {
		fmt.Println((lastStackElemIndex + 1), ": ", h.stack[lastStackElemIndex])
	}
}

func main() {
	history := &BrowserHistory{}
	history.Visit("www.google.com")
	history.Visit("www.github.com")
	history.Visit("www.openai.com")
	history.Back()
	history.PrintHistory()
}
