package testx

import (
	"bytes"
	"fmt"
	"github.com/carmark/pseudo-terminal-go/terminal"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	term := terminal.NewTerminal(os.Stdout, "test")
	term.GetHistory()
}

func TestEval(t *testing.T) {
	output := &bytes.Buffer{}

	// Desired Command to get current cursor position in the terminal
	c1 := exec.Command("echo", "-e", "\033[6n")
	c2 := exec.Command("read", "-d", "R")

	// Example commands that work (un-comment 19 and 20 to show that this method should work)
	//c1 = exec.Command("ls")
	//c2 = exec.Command("wc", "-l")

	c2.Stdin, _ = c1.StdoutPipe()
	c2.Stdout = output
	_ = c2.Start()
	_ = c1.Run()
	_ = c2.Wait()
	fmt.Println("Output: " + strings.TrimSpace(output.String()))
}
