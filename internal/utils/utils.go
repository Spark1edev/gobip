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

func ProcessIndexJS(path string) (err error) {

	f, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	document := string(f)
	document = strings.Replace(document, "{{ addr }}", fmt.Sprintf("%s://%s/write", config.Cfg.Protocol, config.Cfg.Hostname), 10)
	dst, err := os.OpenFile("./views/src/js/index.js", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	dst.Write([]byte(document))
	if err := dst.Close(); err != nil {
		log.Fatal(err)
	}
	return err
}
