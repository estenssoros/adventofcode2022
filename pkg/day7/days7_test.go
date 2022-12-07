package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateWd(t *testing.T) {
	d := NewDirectory("/")
	assert.Equal(t, "/", d.Name)
	d = updateWd(d, "asdf")
	assert.Equal(t, "asdf", d.Name)
}

var testInput = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func TestPart1(t *testing.T) {
	dir, err := parseInput(testInput)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 48381165, dir.Size())
	assert.Equal(t, 584, dir.Children["a"].Children["e"].Size())
	assert.Equal(t, 94853, dir.Children["a"].Size())
	assert.Equal(t, 24933642, dir.Children["d"].Size())
	// ju, err := json.MarshalIndent(dir, "", " ")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(ju))
}
