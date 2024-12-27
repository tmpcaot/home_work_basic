package main

import (
	"os"
	"testing"
)

// TestAppConfig_ConfigFile tests the ConfigFile method of appConfig.
func TestAppConfig_ConfigFile(t *testing.T) {
	// Set the environment variable for testing.
	os.Setenv("LOG_ANALYZER_FILE", "env_file.log")
	defer os.Unsetenv("LOG_ANALYZER_FILE")
	cfg := &appConfig{}

	// Test with an empty string.
	cfg.ConfigFile("")
	if cfg.File != "env_file.log" {
		t.Errorf("Expected File to be set from environment variable LOG_ANALYZER_FILE, got '%s'", cfg.File)
	}

	// Set an explicit file path.
	cfg.ConfigFile("test.log")
	if cfg.File != "test.log" {
		t.Errorf("Expected File to be 'test.log', but got '%s'", cfg.File)
	}
}

// TestAppConfig_ConfigLevel tests the ConfigLevel method of appConfig.
func TestAppConfig_ConfigLevel(t *testing.T) {
	os.Setenv("LOG_ANALYZER_LEVEL", "DEBUG")
	defer os.Unsetenv("LOG_ANALYZER_LEVEL") // Clean up.

	cfg := &appConfig{}

	cfg.ConfigLevel("")
	if cfg.Level != "DEBUG" {
		t.Errorf("Expected Level to be set from environment variable LOG_ANALYZER_LEVEL, got '%s'", cfg.Level)
	}

	cfg.ConfigLevel("INFO")
	if cfg.Level != "INFO" {
		t.Errorf("Expected Level to be 'INFO', but got '%s'", cfg.Level)
	}
}

// TestAppConfig_ConfigOutput tests the ConfigOutput method of appConfig.
func TestAppConfig_ConfigOutput(t *testing.T) {
	os.Setenv("LOG_ANALYZER_OUTPUT", "output.log")
	defer os.Unsetenv("LOG_ANALYZER_OUTPUT")

	cfg := &appConfig{}

	// Test with an empty string.
	cfg.ConfigOutput("")
	if cfg.Output != "output.log" {
		t.Errorf("Expected Output to be set from environment variable LOG_ANALYZER_OUTPUT, got '%s'", cfg.Output)
	}

	// Set an explicit output path.
	cfg.ConfigOutput("output.txt")
	if cfg.Output != "output.txt" {
		t.Errorf("Expected Output to be 'output.txt', but got '%s'", cfg.Output)
	}
}

// TestReadFile tests the ReadFile function.
func TestReadFile(t *testing.T) {
	// Create a temporary file
	tmpfile, err := os.CreateTemp("", "testfile.log")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Write test data to the file.
	_, err = tmpfile.WriteString("192.168.1.1 INFO GET 200 search-engine\n")
	if err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	words, err := ReadFile(tmpfile.Name())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(words) == 0 {
		t.Errorf("Expected to read data from file, got empty slice")
	}
}

// TestSort tests the sort function.
func TestSort(t *testing.T) {
	sl := []status{
		{"192.168.1.1", "INFO", "GET", "200", "search-engine"},
		{"192.168.1.1", "ERROR", "POST", "500", "search-engine"},
		{"10.0.0.1", "INFO", "GET", "200", "search-engine"},
	}

	result := sort(sl, "INFO")

	// Check counts for IPs, methods, and codes.
	if result["ip"]["192.168.1.1"] != 1 {
		t.Errorf("Expected 1 count for '192.168.1.1', got %d", result["ip"]["192.168.1.1"])
	}
	if result["ip"]["10.0.0.1"] != 1 { // Correcting expected count.
		t.Errorf("Expected 1 count for '10.0.0.1', got %d", result["ip"]["10.0.0.1"])
	}

	if result["method"]["GET"] != 2 { // Both entries are 'GET'.
		t.Errorf("Expected 2 count for 'GET', got %d", result["method"]["GET"])
	}
	if result["method"]["POST"] != 0 {
		t.Errorf("Expected 0 count for 'POST', got %d", result["method"]["POST"])
	}

	if result["code"]["200"] != 2 { // Both entries are '200'.
		t.Errorf("Expected 2 count for '200', got %d", result["code"]["200"])
	}
	if result["code"]["500"] != 0 {
		t.Errorf("Expected 0 count for '500', got %d", result["code"]["500"])
	}
}

// TestWriteFile tests the writeFile function.
func TestWriteFile(t *testing.T) {
	outputFile := "test_output.log"
	defer os.Remove(outputFile)

	input := map[string]map[string]int64{
		"ip": {
			"192.168.1.1": 1,
		},
	}

	err := writeFile(input, outputFile)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check if the output file has been created and has data.
	data, err := os.ReadFile(outputFile)
	if err != nil {
		t.Errorf("Expected to read output file, got error: %v", err)
	}
	if string(data) == "" {
		t.Error("Expected output file not to be empty")
	}
}

func TestMain(m *testing.M) {
	exitVal := m.Run()
	os.Exit(exitVal)
}
