package core

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"github.com/spf13/viper"
	"github.com/tidwall/pretty"
)

//go:embed schemas.cue cue.mod/module.cue
var schemaFS embed.FS

const maxLineWidth = 999999

// loadCueValues loads and compiles both schema and module files
func loadCueValues() (*cue.Value, error) {
	schemaBytes, err := schemaFS.ReadFile("schemas.cue")
	if err != nil {
		return nil, fmt.Errorf("error reading schema: %w", err)
	}

	modBytes, err := schemaFS.ReadFile("cue.mod/module.cue")
	if err != nil {
		return nil, fmt.Errorf("error reading module: %w", err)
	}

	ctx := cuecontext.New()

	modVal := ctx.CompileBytes(modBytes)
	if err := modVal.Err(); err != nil {
		return nil, fmt.Errorf("error compiling module: %w", err)
	}

	schemaVal := ctx.CompileBytes(schemaBytes)
	if err := schemaVal.Err(); err != nil {
		return nil, fmt.Errorf("error compiling schema: %w", err)
	}

	unified := modVal.Unify(schemaVal)
	if err := unified.Validate(); err != nil {
		return nil, fmt.Errorf("error validating unified schema: %w", err)
	}

	return &unified, nil
}

// GenerateConfig handles the common flow for all config generators
func GenerateConfig(configName string) error {
	unified, err := loadCueValues()
	if err != nil {
		return err
	}

	config := unified.LookupPath(cue.ParsePath(configName))
	if err := config.Err(); err != nil {
		return fmt.Errorf("error looking up %s config: %w", configName, err)
	}

	var renovateConfig map[string]interface{}
	if err := config.Decode(&renovateConfig); err != nil {
		return fmt.Errorf("error decoding %s config: %w", configName, err)
	}

	jsonBytes, err := json.Marshal(renovateConfig)
	if err != nil {
		return fmt.Errorf("error marshaling to JSON: %w", err)
	}

	// Configure pretty options with sorted keys
	opts := pretty.Options{
		Indent:   "  ",
		SortKeys: true,
		Width:    maxLineWidth,
	}
	prettyJSON := pretty.PrettyOptions(jsonBytes, &opts)

	outFile := viper.GetString("outfile")
	if err := os.WriteFile(outFile, prettyJSON, 0o644); err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	if !viper.GetBool("quiet") {
		fmt.Printf("Generated %s using %s preset\n", outFile, configName)
		fmt.Print(string(prettyJSON))
	}

	return nil
}
