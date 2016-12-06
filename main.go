package main

import (
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"os"
	"io/ioutil"
	"fmt"
	"flag"
	"text/template"
	"path/filepath"
	"bytes"
)

type data struct {
	Filename string
	Fullname string
	Exif     *exif.Exif
}

func (b data) String() string {
	return fmt.Sprintf("Filename: %s\nFullname: %s\nExif: {\n%s\n}", b.Filename, b.Fullname, b.Exif.String())
}

var Renames map[string]string = make(map[string]string)
var UnittestMode = false

func (v data) Format(fmt string) string {
	dt,_ := v.Exif.DateTime()
	return dt.Format(fmt)
}

func main() {
	testMode := flag.Bool("test", true, "Do not execute changes")
	templateString := flag.String("template", `{{.Format ("2006/01")}}/{{.Filename}}`, "Template to use, e.g. " + `{{.Format ("2006-01-02")}}-{{.Format ("2006-01-02 03:04:05"}}`)
	flag.Parse()
	exif.RegisterParsers(mknote.All...)

	t, err := template.New("filename").Parse(*templateString)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Test mode is %t\n", *testMode)

	files := scanDir(".")
	for index := range (files) {
		fileName := files[index]
		processFile(t, *testMode, fileName)
	}
}

func processFile(template *template.Template, testMode bool, fileName string) {

	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	x, err := exif.Decode(f)
	if err != nil {
		fmt.Printf("SKIPPED: %s: %s\n", fileName, err)
	} else {
		context := data{
			Exif: x,
			Filename: filepath.Base(fileName),
			Fullname: fileName,
		}
		buf := bytes.Buffer{}
		err := (*template).Execute(&buf, context)
		if err != nil {
			panic(err)
		}
		targetName := buf.String()
		fmt.Printf("%s -> %s\n", fileName, targetName)
		if !testMode {
			os.MkdirAll(filepath.Dir(targetName), os.ModeDir | 0777)
			os.Rename(targetName, targetName)
		} else if UnittestMode {
			Renames[fileName] = targetName
		}
	}

}

func scanDir(dir string) []string {
	listing, _ := ioutil.ReadDir(dir)
	result := []string{}
	for index := range (listing) {
		if (listing[index].IsDir()) {
			add := scanDir(dir + string(os.PathSeparator) + listing[index].Name())
			result = append(result, add...)
		} else if (listing[index].Name()[0] != '.') {
			result = append(result, dir + string(os.PathSeparator) + listing[index].Name())
		}
	}
	return result
}
