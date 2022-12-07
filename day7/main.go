package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := getData()

	directories, files := parseData(data)

	fmt.Printf("Part1: %d\n", part1(directories, files))
	fmt.Printf("Part2: %d\n", part2(directories, files))
}
func part1(directories []string, files map[string]int) int {
	directoriesSize := getDirectoriesSize(directories, files)

	result := 0
	for _, s := range directoriesSize {
		if s <= 100000 {
			result += s
		}
	}

	return result
}
func part2(directories []string, files map[string]int) int {
	fileSystemSize := 70000000
	requiredUnusedSpace := 30000000
	spaceTaken := getDirectorySize("/", files)
	unusedSpace := fileSystemSize - spaceTaken
	directoriesSize := getDirectoriesSize(directories, files)
	toDelete := requiredUnusedSpace - unusedSpace

	result := spaceTaken
	for _, s := range directoriesSize {
		if s < toDelete {
			continue
		}
		if s < result {
			result = s
		}
	}

	return result
}
func getDirectoriesSize(directories []string, fileList map[string]int) map[string]int {
	result := make(map[string]int)

	for _, d := range directories {
		result[d] = getDirectorySize(d, fileList)
	}

	return result
}
func getDirectorySize(directory string, fileList map[string]int) int {
	result := 0

	for f, s := range fileList {
		if strings.HasPrefix(f, directory) {
			result += s
		}
	}

	return result
}
func parseData(data []string) ([]string, map[string]int) {
	directories := []string{}
	files := make(map[string]int)

	currentDirectory := []string{""}
	currentLine := 1

	for {
		if data[currentLine] == "$ ls" {
			_, lineCount := parseDirectory(data, currentLine+1, strings.Join(currentDirectory, "/"), files)
			currentLine += lineCount
		} else if data[currentLine] == "$ cd .." {
			currentDirectory = currentDirectory[:len(currentDirectory)-1]
			currentLine++
		} else if strings.HasPrefix(data[currentLine], "$ cd ") {
			currentDirectory = append(currentDirectory, data[currentLine][5:])
			directories = append(directories, strings.Join(currentDirectory, "/"))
			currentLine++
		}

		if currentLine >= len(data) {
			break
		}
	}

	return directories, files
}
func parseDirectory(lines []string, lineNumber int, currentDirectory string, targetFileList map[string]int) ([]string, int) {

	directories := []string{}

	currentLineNumber := lineNumber
	for {
		if currentLineNumber >= len(lines) || lines[currentLineNumber][0] == '$' {
			return directories, currentLineNumber - lineNumber + 1
		}

		if strings.HasPrefix(lines[currentLineNumber], "dir") {
			directories = append(directories, lines[currentLineNumber][4:])
			currentLineNumber++
		} else {
			parts := strings.Split(lines[currentLineNumber], " ")
			size, _ := strconv.Atoi(parts[0])
			targetFileList[currentDirectory+"/"+parts[1]] = size
			currentLineNumber++
		}
	}
}
func getData() []string {
	result := []string{}

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
func getTestData() []string {
	return []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"14848514 b.txt",
		"8504156 c.dat",
		"dir d",
		"$ cd a",
		"$ ls",
		"dir e",
		"29116 f",
		"2557 g",
		"62596 h.lst",
		"$ cd e",
		"$ ls",
		"584 i",
		"$ cd ..",
		"$ cd ..",
		"$ cd d",
		"$ ls",
		"4060174 j",
		"8033020 d.log",
		"5626152 d.ext",
		"7214296 k",
	}
}
