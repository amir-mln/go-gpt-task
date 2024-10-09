package main

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"go-gpt-task/configs"
	"go-gpt-task/prompting"
	"go-gpt-task/repositories"
	"go-gpt-task/usecases"
	"os"
)

type flags struct {
	inputFile string
	batchMode bool
}

func loadFlags() (flags, error) {
	btch := flag.Bool("batch-mode", false, "wether to send the request in batch mode or not")
	file := flag.String("input-file", "", "the path to the file that contains the laptops data")

	flag.Parse()

	if *file == "" {
		return flags{}, errors.New("the input file cli argument is required")
	}

	return flags{inputFile: *file, batchMode: *btch}, nil
}

func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	cfg, err := configs.Load()
	if err != nil {
		panic(err)
	}

	flgs, err := loadFlags()
	if err != nil {
		panic(err)
	}

	data, err := readFile(flgs.inputFile)
	if err != nil {
		panic(err)
	}

	gpt := prompting.NewGPTPromptParser(cfg.APIKey)
	db := repositories.NewDatabase()
	cache := repositories.NewCache()
	uc := usecases.NewUsecases(&db, cache, gpt)

	for _, d := range data {
		fmt.Printf("Received A Prompt, Content=%q\n", d)

		laptop, err := uc.ParsePrompt(context.Background(), d)
		if err != nil {
			fmt.Printf("Parsing Resulted in an Error, Message=%q\n\n", err)
			continue
		}

		bytes, _ := json.MarshalIndent(laptop, "", "\t")
		fmt.Printf("Parsed Laptop Information, Laptop=%v\n\n", string(bytes))
	}
}
