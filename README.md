# Laptop Prompts with ChatGPT API 

A simple application that reads raw laptop information from an input file
and uses ChatGPT API to extracts details about the laptop's hardware.

## How to Run

1. After cloning the repo, run `go mod download` command.
2. Add a `.env` file in the root of the project with an `API_KEY` variable
```env
API_KEY=your_chatgpt_api_key
```
3. Add your input data in a textual file where each line holds the value for an input laptop.
```text
Dell Inspiron; Processor i7-10510U ; RAM 16GB; 512GB SSD Missing battery
MacBook Pro with M1 chip, 8GB RAM, 256 GB SSD storage Battery removed
ThinkPad, i5 CPU, 8GB memory, storage: 1TB HDD
Asus ROG, Processor: AMD Ryzen 7; RAM 16 GB; 1TB SSD; Damaged battrey
Dell Inspiron; Processor: i5-1135G7; RAM 8GB; Storage: 256.123548 SSD; Missing charger
Invalid Data that doesn't do anything about a laptop
```
4. Run the application using `go run main.go --input-file path_to_your_input_file`.