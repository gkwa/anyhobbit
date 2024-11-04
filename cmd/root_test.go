package cmd

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gkwa/anyhobbit/internal/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCustomLogger(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	cmd := rootCmd
	cmd.SetArgs([]string{"version"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		t.Fatalf("Failed to copy output: %v", err)
	}
	output := buf.String()

	if !strings.Contains(output, "Version:") {
		t.Errorf("Expected output to contain version information, but got: %s", output)
	}

	t.Logf("Command output: %s", output)
}

func TestJSONLogger(t *testing.T) {
	oldVerbose, oldLogFormat := verbose, logFormat
	verbose, logFormat = 1, "json"
	defer func() {
		verbose, logFormat = oldVerbose, oldLogFormat
	}()

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	customLogger := logger.NewConsoleLogger(verbose, logFormat == "json")
	cliLogger = customLogger

	cmd := rootCmd
	cmd.SetArgs([]string{"version"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		t.Fatalf("Failed to copy output: %v", err)
	}
	output := buf.String()

	if !strings.Contains(output, "Version:") {
		t.Errorf("Expected output to contain version information, but got: %s", output)
	}

	t.Logf("Command output: %s", output)
}

func TestCommandExecution(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name           string
		args           []string
		expectedFile   string
		expectedInFile []string // strings that should be present in output
		notExpected    []string // strings that should not be present
	}{
		{
			name:         "owl command creates config",
			args:         []string{"owl", "-o", filepath.Join(tmpDir, "owl.json")},
			expectedFile: filepath.Join(tmpDir, "owl.json"),
			expectedInFile: []string{
				"config:best-practices",
				"recreateWhen",
			},
			notExpected: []string{
				"config:recommended",
				"indirect",
				"replacement",
			},
		},
		{
			name:         "monkey command includes indirect deps",
			args:         []string{"monkey", "-o", filepath.Join(tmpDir, "monkey.json")},
			expectedFile: filepath.Join(tmpDir, "monkey.json"),
			expectedInFile: []string{
				"matchDepTypes",
				"indirect",
				"enabled",
			},
		},
		{
			name:         "hamster command uses recommended base",
			args:         []string{"hamster", "-o", filepath.Join(tmpDir, "hamster.json")},
			expectedFile: filepath.Join(tmpDir, "hamster.json"),
			expectedInFile: []string{
				"config:recommended",
				"gomodTidyE",
			},
			notExpected: []string{
				"config:best-practices",
				"replacement",
			},
		},
		{
			name:         "default output file",
			args:         []string{"cat", "-o", filepath.Join(tmpDir, "default.json")},
			expectedFile: filepath.Join(tmpDir, "default.json"),
			expectedInFile: []string{
				"rangeStrategy",
				"pin",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Execute command
			rootCmd.SetArgs(tt.args)
			err := rootCmd.Execute()
			require.NoError(t, err)

			// Check file exists and content
			content, err := os.ReadFile(tt.expectedFile)
			require.NoError(t, err, "should be able to read output file")

			// Check expected contents
			for _, expected := range tt.expectedInFile {
				assert.Contains(t, string(content), expected)
			}

			// Check things that shouldn't be there
			for _, notExpected := range tt.notExpected {
				assert.NotContains(t, string(content), notExpected)
			}
		})
	}
}

func TestInvalidCommand(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name        string
		args        []string
		expectedErr string
	}{
		{
			name:        "non-existent animal",
			args:        []string{"giraffe"},
			expectedErr: "unknown command",
		},
		{
			name:        "invalid output dir",
			args:        []string{"owl", "-o", filepath.Join(tmpDir, "nonexistent", "config.json")},
			expectedErr: "error writing config file",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootCmd.SetArgs(tt.args)
			err := rootCmd.Execute()
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.expectedErr)
		})
	}
}
