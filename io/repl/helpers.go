package repl

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/Force4760/pipes/io/colorize"
	"github.com/Force4760/pipes/src/stack"
)

// Prompt for the repl
var PROMPT = colorize.Colorize("-> ", colorize.GREEN)

// Intro message for the repl
var MSG = colorize.Colorize("Hello! This is the Pipes programming language!", colorize.BLUE)

// Repl changing functions
// clear -> clear the repl
// empty -> empty the stack
// stack -> print the stack
func Commands(line string, stk *stack.Stack) bool {
	switch line {
	case "clear":
		// Clear the repl
		ClearScreen()
		return true

	case "empty":
		// Empty the stack
		stk.Empty()
		return true

	case "stack":
		// Print the stack
		fmt.Println(stk)
		return true

	}

	return false
}

// Clear Command Map
// OS - clear cmd
var clearScreen = map[string]func(){
	"darwin": func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	},

	"linux": func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	},

	"windows": func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	},
}

// Clears the terminal window
func ClearScreen() {
	fun, valid := clearScreen[runtime.GOOS]
	if valid {
		fun()
	}
}
