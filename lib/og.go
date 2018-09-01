package og

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

type OgConfig struct {
	Blocks  bool
	Dirty   bool
	Print   bool
	Verbose bool
	Paths   []string
	OutPath string
}

var config OgConfig

func Compile(config_ OgConfig) {
	config = config_
	for _, p := range config.Paths {
		filepath.Walk(p, walker)
	}
}
func walker(filePath string, info os.FileInfo, err error) error {
	if info.IsDir() == true {
		return nil
	}
	if path.Ext(filePath) != ".og" {
		return nil
	}
	source, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	res := ProcessFile(filePath, string(source))
	finalizeFile(filePath, res)
	return nil
}
func ProcessFile(filePath string, data string) string {
	if config.Verbose == true {
		fmt.Println(filePath)
	}
	preprocessed := Preproc(string(data))
	if config.Blocks == true {
		return preprocessed
	}
	res := Parse(string(preprocessed))
	if config.Dirty == true {
		return res
	}
	final := format(res)
	return final
}
func finalizeFile(filePath string, data string) {
	if config.Print == true {
		fmt.Println(data)
	} else {
		writeFile(filePath, data)
	}
}
func writeFile(filePath string, data string) {
	if config.OutPath != "./" {
		splited := strings.SplitN(filePath, "/", 2)
		filePath = splited[1]
	}
	newPath := strings.Replace(path.Join(config.OutPath, filePath), ".og", ".go", 1)
	os.MkdirAll(filepath.Dir(newPath), os.ModePerm)
	ioutil.WriteFile(newPath, []byte(data), os.ModePerm)
	if config.Verbose == true {
		fmt.Println("->", newPath)
	}
}
func format(str string) string {
	cmd := exec.Command("gofmt")
	stdin, _ := cmd.StdinPipe()
	go func() {
		defer stdin.Close()
		stdin.Write([]byte(str))
	}()
	final, _ := cmd.CombinedOutput()
	return string(final)
}
