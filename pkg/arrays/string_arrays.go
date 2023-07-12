package arrays

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type StringArray struct {
	array []string
}

func (a StringArray) IndexOfPrefix(prefix string) int {
	for index, value := range a.array {
		if strings.HasPrefix(value, prefix) {
			return index
		}
	}
	return -1
}

func (a *StringArray) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range a.array {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func LoadFromFile(filename string) (array *StringArray, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	ar := make([]string, 0, 200)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ar = append(ar, scanner.Text())
	}
	return &StringArray{ar}, scanner.Err()
}

func (a *StringArray) AppendItem(line string) *StringArray {
	a.array = append(a.array, line)
	return a
}

func (a *StringArray) UpdateItem(lineNo int, line string) *StringArray {
	a.array[lineNo] = line
	return a
}

func (a *StringArray) RemoveItem(index int) *StringArray {
	a.array = append(a.array[:index], a.array[index+1:]...)
	return a
}

func (a *StringArray) FindByLine(finder func(line string) bool) (lineNumber int, lineContent string, err error) {
	for i, line := range a.array {
		if finder(line) {
			lineNumber = i
			lineContent = line
			return
		}
	}

	err = errors.New("line not found")
	return
}
