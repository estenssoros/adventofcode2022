package gen

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var day int

func init() {
	Cmd.Flags().IntVarP(&day, "day", "", 0, "day to generate")
}

var Cmd = &cobra.Command{
	Use:   "gen",
	Short: "",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if day == 0 {
			return errors.New("please provide a day to generate")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error { return gen(day) },
}

//go:embed templates/cmd.go.tpl
var cmdTemplate string

//go:embed templates/part.go.tpl
var partTemplate string

//go:embed templates/parse.go.tpl
var parseTemplate string

//go:embed templates/test.go.tpl
var testTemplate string

type Formatter struct {
	Day  int
	Part int
}

func gen(day int) error {
	dir := filepath.Join("pkg", fmt.Sprintf("day%d", day))
	_, err := os.Stat(dir)
	if err == nil {
		return errors.Errorf("%s already exists", dir)
	}
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "os.MkdirAll")
	}
	formatter := Formatter{
		Day: day,
	}
	if err := createFileFromTemplate(filepath.Join(dir, "cmd.go"), cmdTemplate, formatter); err != nil {
		return errors.Wrap(err, "createFileFromTemplate")
	}

	if err := createBlankFile(filepath.Join(dir, "input.txt")); err != nil {
		return errors.Wrap(err, "createFile")
	}
	for i := 1; i <= 2; i++ {
		formatter.Part = i
		if err := createFileFromTemplate(filepath.Join(dir, fmt.Sprintf("part%d.go", i)), partTemplate, formatter); err != nil {
			return errors.Wrap(err, "createFileFromTemplate")
		}
	}
	if err := createFileFromTemplate(filepath.Join(dir, fmt.Sprintf("day%d_test.go", day)), testTemplate, formatter); err != nil {
		return errors.Wrap(err, "crateFileFromTemplate")
	}

	if err := createFileFromTemplate(filepath.Join(dir, "parse.go"), parseTemplate, formatter); err != nil {
		return errors.Wrap(err, "crateFileFromTemplate")
	}

	return createPackageFile(filepath.Join(dir, "structs.go"), day)
}

func createBlankFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "os.Create")
	}
	return f.Close()
}

func createPackageFile(path string, day int) error {
	f, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "os.Create")
	}
	defer f.Close()
	_, err = f.Write([]byte(fmt.Sprintf("package day%d", day)))
	return err
}

func createFileFromTemplate(path, tpl string, formatter Formatter) error {
	f, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "os.Create")
	}
	defer f.Close()
	t, err := template.New("").Parse(tpl)
	if err != nil {
		return errors.Wrap(err, "template.New.Parse")
	}
	return t.Execute(f, formatter)
}
