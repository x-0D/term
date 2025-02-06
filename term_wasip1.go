//go:build wasip1 && wasm
// +build wasip1,wasm

package term

import (
	"errors"
	"io"
	"os"
)

type terminalState struct{}

func stdStreams() (stdIn io.ReadCloser, stdOut, stdErr io.Writer) {
	return os.Stdin, os.Stdout, os.Stderr
}

func getFdInfo(in interface{}) (uintptr, bool) {
	return 0, false // No real file descriptors in Wasm
}

func getWinsize(fd uintptr) (*Winsize, error) {
	return nil, errors.New("terminal size detection not supported in WebAssembly")
}

func setWinsize(fd uintptr, ws *Winsize) error {
	return errors.New("setting terminal size not supported in WebAssembly")
}

func isTerminal(fd uintptr) bool {
	return false // No real terminals in Wasm
}

func restoreTerminal(fd uintptr, state *State) error {
	return errors.New("terminal state management not supported in WebAssembly")
}

func saveState(fd uintptr) (*State, error) {
	return nil, errors.New("terminal state management not supported in WebAssembly")
}

func disableEcho(fd uintptr, state *State) error {
	return errors.New("terminal echo control not supported in WebAssembly")
}

func setRawTerminal(fd uintptr) (*State, error) {
	return nil, errors.New("raw terminal mode not supported in WebAssembly")
}

func setRawTerminalOutput(fd uintptr) (*State, error) {
	return nil, errors.New("raw terminal output not supported in WebAssembly")
}

func makeRaw(fd uintptr) (*State, error) {
	return nil, errors.New("raw terminal mode not supported in WebAssembly")
}