package arrays

import (
	"bufio"
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

func (a StringArray) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range a {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func LoadFromFile(filename string) (array StringArray, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	array = make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		array = append(array, scanner.Text())
	}
	return array, scanner.Err()
}

func (a StringArray) RemoveItem(index int) (array StringArray){

}