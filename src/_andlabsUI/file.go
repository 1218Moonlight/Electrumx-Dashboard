package _andlabsUI

import (
	"os"
	"bufio"
)

func readFile(fileName string) ([]string, []string, error){
	return revertLine(fileName)
}

func revertLine(fileName string) ([]string, []string, error){
	file, err := os.Open(fileName)
	if checkError(err, true) {return nil, nil, err}
	defer file.Close()
	origin := []string{}
	revert := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		origin = append(origin, scanner.Text())
	}
	for i := len(origin)-1; i >=0; i-- {
		revert = append(revert, origin[i])
	}
	return origin, revert, nil
}