package og

import (
	"bufio"
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

var (
	config OgConfig
)

func Compile(config_ OgConfig) error {
	config = config_
	for _, p := range config.Paths {
		if err := filepath.Walk(p, walker); err != nil {
			fmt.Println("Error", err)
			return err
		}
	}
	return nil
}
func walker(filePath string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
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
	res, err := ProcessFile(filePath, string(source), false)
	if err != nil {
		return err
	}
	finalizeFile(filePath, res)
	return nil
}
func ProcessFile(filePath string, data string, isInterpret bool) (string, error) {
	if config.Verbose == true {
		fmt.Print(filePath)
	}
	preprocessed := Preproc(string(data))
	if config.Blocks == true {
		return preprocessed, nil
	}
	res := ""
	if !isInterpret {
		res = Parse(filePath, string(preprocessed))
	} else {
		res = ParseInterpret(filePath, string(preprocessed))
	}
	if config.Dirty == true {
		return res, nil
	}
	return format(res)
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
func format(str string) (string, error) {
	cmd := exec.Command("gofmt")
	stdin, _ := cmd.StdinPipe()
	stdin.Write([]byte(str))
	stdin.Close()
	final, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(final))
		return "", err
	}
	return string(final), nil
}
func RunInterpreter() {
	running := true
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for running {
		fmt.Print("> ")
		if !scanner.Scan() {
			return
		}
		ln := scanner.Text()
		if len(ln) == 0 {
			continue
		}
		str, err := ProcessFile("STDIN", ln, true)
		if err != nil {
			fmt.Println(err)
		} else {
			execCode(str)
		}
	}
}
func execCode(str string) {
	skelton := `package main
  import "fmt"
  func main() {
  fmt.Print(` + str[:len(str)-1] + `)
  }`
	ioutil.WriteFile("/tmp/main.go", []byte(skelton), os.ModePerm)
	cmd := exec.Command("go", "run", "/tmp/main.go")
	stdin, _ := cmd.StdinPipe()
	stdin.Write([]byte(str))
	stdin.Close()
	final, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(final))
}
