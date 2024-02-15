package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Config struct {
	LogFilePath string `json:"logFilePath"`
	LogLevel    string `json:"logLevel"`
	OutputPath  string `json:"outputPath"`
}

var configPath string

func init() {
	flag.StringVar(&configPath, "config", getEnv("LOG_ANALYZER_CONFIG", "config.json"), "Path to configuration file")
}

func main() {
	flag.Parse()

	config := Config{
		LogLevel: "info",
	}

	if configPath != "" {
		readConfigFromFile(&config, configPath)
	} else {
		log.Fatal("Configuration file path not specified")
	}

	flag.StringVar(&config.LogFilePath, "file", getEnv("LOG_ANALYZER_FILE", config.LogFilePath),
		"Path to the analyzed log file")
	flag.StringVar(&config.LogLevel, "level", getEnv("LOG_ANALYZER_LEVEL", config.LogLevel),
		"Level of logs to be analyzed")
	flag.StringVar(&config.OutputPath, "output", getEnv("LOG_ANALYZER_OUTPUT", config.OutputPath),
		"Path to the file for recording statistics")
	flag.Parse()

	stats := analyzeLogFile(&config)

	if config.OutputPath != "" {
		file, err := os.OpenFile(config.OutputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			log.Fatalf("Failed to open output file: %s\n", err)
		}
		defer file.Close()
		outputStats(file, stats, config.LogLevel)
	} else {

		outputStats(os.Stdout, stats, config.LogLevel)
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func readConfigFromFile(config *Config, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Failed to open configuration file: %s\n", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		log.Printf("Error while parsing configuration file: %s\n", err)
	}
}

func analyzeLogFile(config *Config) map[string]int {
	stats := make(map[string]int)
	file, err := os.Open(config.LogFilePath)
	if err != nil {
		fmt.Printf("Error when opening a file: %s\n", err)
		return stats
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, config.LogLevel+":") {
			stats[config.LogLevel]++
		}
	}

	err = scanner.Err()
	if err != nil {
		fmt.Printf("Error when reading a file: %s\n", err)
	}

	return stats
}

func outputStats(writer io.Writer, stats map[string]int, logLevel string) {
	statsString := fmt.Sprintf("Statistics for level '%s':\n", logLevel)
	for level, count := range stats {
		statsString += fmt.Sprintf("%s: %d\n", level, count)
	}

	_, err := fmt.Fprint(writer, statsString)
	if err != nil {
		fmt.Printf("Error when writing: %s\n", err)
	}
}
