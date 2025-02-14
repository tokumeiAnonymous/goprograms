package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	header = `<!DOCTYPE html>
	<html>
		<head>
			<meta http-equiv="content-type" content="text/html; charset=utf-8">
			<title>Markdown Preview Tool</title>
		</head>
	<body> `
	footer = `
		</body>
	</html>
	`
)

func run(filename string, out io.Writer, skipPreview bool) error {
	// Read all the data from the input file and check for errors
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	temp, err := os.CreateTemp("", "mdp*.html")
	if err != nil {
		fmt.Println("Error creating temp file:", err)
		return err
	}
	htmlData := parseContent(input)
	outName := temp.Name()
	fmt.Fprintln(out, outName)

	if err := saveHTML(outName, htmlData); err != nil {
		return err
	}
	if skipPreview {
		os.Remove(outName)
		return nil
	}
	defer os.Remove(outName)
	return preview(outName)
}

func parseContent(input []byte) []byte {
	// Parse the markdown file through blackfriday and bluemonday // to generate a valid and safe HTML
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)
	// Create a buffer of bytes to write to file
	var buffer bytes.Buffer
	// Write html to bytes buffer
	buffer.WriteString(header)
	buffer.Write(body)
	buffer.WriteString(footer)
	return buffer.Bytes()
}

func saveHTML(outFname string, data []byte) error {
	// Write the bytes to the file
	return os.WriteFile(outFname, data, 0644)
}

func preview(fname string) error {
	cName := ""
	cParams := []string{}

	if runtime.GOOS == "linux" {
		cName = "xdg-open"
	} else if runtime.GOOS == "windows" {
		cName = "cmd.exe"
		cParams = []string{"/C", "start"}
	} else if runtime.GOOS == "darwin" {
		cName = "open"
	} else {
		return fmt.Errorf("OS not supported")
	}
	cParams = append(cParams, fname)
	cPath, err := exec.LookPath(cName)
	if err != nil {
		return err
	}
	err = exec.Command(cPath, cParams...).Run()
	time.Sleep(2 * time.Second)
	return err
}

func main() {
	// Parse flags
	filename := flag.String("file", "", "Markdown file to preview")
	skipPreview := flag.Bool("s", false, "Skip auto-preview")
	flag.Parse()
	// If user did not provide input file, show usage
	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}
	if err := run(*filename, os.Stdout, *skipPreview); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
