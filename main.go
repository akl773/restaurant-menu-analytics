package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Entry struct {
	EaterID    int
	FoodMenuID int
}

func main() {

}

func readLogFile(filePath string) ([]Entry, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var entries []Entry
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid log entry: %s", line)
		}

		eaterID, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return nil, fmt.Errorf("invalid eater ID: %s", parts[0])
		}

		foodMenuID, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return nil, fmt.Errorf("invalid foodmenu ID: %s", parts[1])
		}

		entries = append(entries, Entry{EaterID: eaterID, FoodMenuID: foodMenuID})
	}

	return entries, nil
}
