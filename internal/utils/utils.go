package utils

import (
	"fmt"
	"gobip/internal/config"
	"log"
	"os"
	"strings"
)

func AppendText(path string, text string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(text)); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatalln(err)
	}
	return err
}

func WriteText(path string, text string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(text)); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatalln(err)
	}
	return err
}

func ProcessIndexJS(path string) (err error) {
	var (
		content  []byte
		document string
	)
	content, err = os.ReadFile(path)
	if err != nil {
		return err
	}
	document = string(content)
	document = strings.Replace(document, "{% addr %}", fmt.Sprintf("%s://%s/write", config.Cfg.Hostname, config.Cfg.Protocol), 2)
	return WriteText("./views/src/js/index.js", document)
}
