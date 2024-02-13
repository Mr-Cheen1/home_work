package main

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestLogAnalyzerIntegration(t *testing.T) {
	os.Setenv("LOG_ANALYZER_FILE", "logfile.log")
	os.Setenv("LOG_ANALYZER_LEVEL", "info")
	os.Setenv("LOG_ANALYZER_OUTPUT", "output.txt")
	defer func() {
		os.Unsetenv("LOG_ANALYZER_FILE")
		os.Unsetenv("LOG_ANALYZER_LEVEL")
		os.Unsetenv("LOG_ANALYZER_OUTPUT")
		os.Remove("output.txt")
	}()

	main()

	output, err := os.ReadFile("output.txt")
	if err != nil {
		t.Fatalf("Error reading output file: %s", err)
	}

	expected := "info: 3"
	if !strings.Contains(string(output), expected) {
		t.Errorf("Expected %q, received %q", expected, string(output))
	}
}

func TestGetEnv(t *testing.T) {
	const envKey = "TEST_ENV_VAR"
	const defaultValue = "default"
	expectedValue := "testValue"

	os.Setenv(envKey, expectedValue)
	defer os.Unsetenv(envKey)

	value := getEnv(envKey, defaultValue)
	if value != expectedValue {
		t.Errorf("getEnv() = %v, want %v", value, expectedValue)
	}

	os.Unsetenv(envKey)
	value = getEnv(envKey, defaultValue)
	if value != defaultValue {
		t.Errorf("getEnv() = %v, want %v", value, defaultValue)
	}
}

func TestReadConfigFromFile_Success(t *testing.T) {
	expectedConfig := Config{
		LogFilePath: "./test.log",
		LogLevel:    "info",
		OutputPath:  "./output.txt",
	}

	configFilePath := "test_config.json"
	configFile, err := json.Marshal(expectedConfig)
	if err != nil {
		t.Fatalf("Failed to serialize the expected configuration: %s", err)
	}
	err = os.WriteFile(configFilePath, configFile, 0o644)
	if err != nil {
		t.Fatalf("Failed to create a temporary configuration file: %s", err)
	}
	defer func() {
		err = os.Remove(configFilePath)
		if err != nil {
			t.Logf("Failed to delete the temporary configuration file: %s", err)
		}
	}()

	var config Config

	readConfigFromFile(&config, configFilePath)

	if !reflect.DeepEqual(config, expectedConfig) {
		t.Errorf("readConfigFromFile() = %+v, want %+v", config, expectedConfig)
	}
}

func TestReadConfigFromFile_Error(t *testing.T) {
	var config Config

	readConfigFromFile(&config, "non_existent_config.json")

	if config.LogFilePath != "" || config.LogLevel != "" || config.OutputPath != "" {
		t.Errorf("Expected empty config, got %v", config)
	}
}

func TestAnalyzeLogFile_Success(t *testing.T) {
	logFilePath := "test.log"
	logContent := []byte("info: first info message\nerror: first error message\ninfo: second info message")
	os.WriteFile(logFilePath, logContent, 0o600)
	defer os.Remove(logFilePath)

	config := Config{
		LogFilePath: logFilePath,
		LogLevel:    "info",
	}

	expectedStats := map[string]int{
		"info": 2,
	}

	stats := analyzeLogFile(&config)
	if !reflect.DeepEqual(stats, expectedStats) {
		t.Errorf("analyzeLogFile() = %v, want %v", stats, expectedStats)
	}
}

func TestAnalyzeLogFile_Error(t *testing.T) {
	config := Config{
		LogFilePath: "non_existent.log",
		LogLevel:    "info",
	}

	stats := analyzeLogFile(&config)
	if len(stats) != 0 {
		t.Errorf("Expected empty stats for non-existent file, got %v", stats)
	}
}

func TestOutputStats_ToFile(t *testing.T) {
	stats := map[string]int{
		"info": 2,
	}
	config := Config{
		LogLevel:   "info",
		OutputPath: "output_test.txt",
	}
	defer os.Remove(config.OutputPath)

	outputStats(&config, stats)

	output, err := os.ReadFile(config.OutputPath)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	expectedOutput := "Statistics for level 'info':\ninfo: 2\n"
	if string(output) != expectedOutput {
		t.Errorf("outputStats() = %q, want %q", string(output), expectedOutput)
	}
}
