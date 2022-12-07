package day7

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func parseInput(input string) (*Directory, error) {
	dir := NewDirectory("/")
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		line := s.Text()
		if isCommand(line) {
			switch commandType(line) {
			case CommandChangeDir:
				dir = updateWd(dir, line)
				continue
			case CommandList:
				continue
			}
		}
		fileInfo, err := parseFile(line)
		if err != nil {
			return nil, errors.Wrap(err, "parseFile")
		}
		dir.AddFileInfo(fileInfo)
	}
	return dir.Root(), nil
}

func isDir(line string) bool {
	return strings.HasPrefix(line, "dir")
}

func updateWd(dir *Directory, cmd string) *Directory {
	arg := strings.TrimPrefix(cmd, "$ cd ")
	switch arg {
	case "/":
		return dir.Root()
	case "..":
		return dir.Parent
	default:
		if d, ok := dir.Children[arg]; ok {
			return d
		}
		dir.AddFileInfo(&FileInfo{Name: arg, IsDir: true})
		return dir.Children[arg]
	}
}

func isCommand(line string) bool {
	return strings.HasPrefix(line, "$ ")
}

var (
	CommandChangeDir = "cd"
	CommandList      = "ls"
)

func commandType(line string) string {
	if strings.HasPrefix(line, "$ cd") {
		return CommandChangeDir
	}
	return CommandList
}

func parseFile(line string) (*FileInfo, error) {
	if isDir(line) {
		return &FileInfo{
			Name:  strings.TrimPrefix(line, "dir "),
			IsDir: true,
		}, nil
	}
	fields := strings.Fields(line)
	if len(fields) != 2 {
		return nil, errors.Errorf("could not parse: %s", line)
	}
	size, err := strconv.Atoi(fields[0])
	if err != nil {
		return nil, errors.Wrap(err, "strconv.Atoi")
	}
	return &FileInfo{
		Name: fields[1],
		Size: size,
	}, nil
}
