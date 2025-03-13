package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

type History struct {
	History []string `json:"history"`
}

func copyToClipboard(data string) {
	cmd := exec.Command("wl-copy")
	stdin, _ := cmd.StdinPipe()
	cmd.Start()
	stdin.Write([]byte(data))
	stdin.Close()
	cmd.Wait()
}

func pasteFromClipboard() string {
	out, _ := exec.Command("wl-paste").Output()
	return string(out)
}

func readFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(content)
}

func writeFile(filename string, data string) {
	ioutil.WriteFile(filename, []byte(data), 0644)
}

func getHistory() []string {
	home, _ := os.UserHomeDir()
	data := readFile(home + "/cclip.json")

	var history History
	if err := json.Unmarshal([]byte(data), &history); err != nil {
		return []string{"No history"}
	}

	return history.History
}

func writeHistory(history []string) {
	home, _ := os.UserHomeDir()
	data, _ := json.Marshal(History{History: history})
	writeFile(home+"/cclip.json", string(data))
}

func trim(str string) string {
	return strings.TrimSpace(str)
}

func showInRofi(options []string) string {
	cmd := exec.Command("rofi", "-dmenu")
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	go func() {
		for _, item := range options {
			modified := strings.ReplaceAll(item, "\n", " ") // Prevent multiline issue
			stdin.Write([]byte(modified + "\n"))
		}
		stdin.Close()
	}()

	scanner := bufio.NewScanner(stdout)
	var result string
	if scanner.Scan() {
		result = scanner.Text()
	}
	cmd.Wait()

	return result
}

func finder(target string, list []string) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No args passed, quitting...")
		return
	} else if len(os.Args) > 2 {
		fmt.Println("Too many args, quitting...")
		return
	}

	arg := os.Args[1]

	switch arg {
	case "--listen":
		fmt.Println("listening")
		for {
			data := pasteFromClipboard()
			history := getHistory()
			if !finder(data, history) {
				history = append(history, data)
				writeHistory(history)
			}
			time.Sleep(500 * time.Millisecond)
		}
	case "--show":
		history := getHistory()
		selected := showInRofi(history)
		for _, item := range history {
			if strings.ReplaceAll(item, "\n", " ") == selected {
				copyToClipboard(item)
				break
			}
		}
	default:
		fmt.Println("Unknown argument")
	}
}
