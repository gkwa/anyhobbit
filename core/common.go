package core

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"github.com/spf13/viper"
	"github.com/tidwall/pretty"
)

//go:embed schemas.cue cue.mod/module.cue
var schemaFS embed.FS

const MaxLineWidth = 999999

// ConfigsByName contains a config and its animal name
type ConfigsByName struct {
	Name   string
	Config []string
}

// ListAllConfigs returns all animal configs sorted by name
func ListAllConfigs() ([]ConfigsByName, error) {
	unified, err := loadCueValues()
	if err != nil {
		return nil, err
	}

	// Get all fields with @animal attribute
	var animals []string
	iter, _ := unified.Fields()
	for iter.Next() {
		sel := iter.Selector()
		docs := iter.Value().Doc()
		for _, doc := range docs {
			for _, c := range doc.List {
				if strings.Contains(c.Text, "@animal") {
					animals = append(animals, sel.String())
					break
				}
			}
		}
	}
	sort.Strings(animals)

	// Get formatted config for each animal
	var configs []ConfigsByName
	for _, animal := range animals {
		config, err := LoadAnimalConfig(animal)
		if err != nil {
			return nil, fmt.Errorf("error loading %s config: %w", animal, err)
		}

		jsonBytes, err := json.Marshal(config)
		if err != nil {
			return nil, fmt.Errorf("error marshaling %s config: %w", animal, err)
		}

		opts := pretty.Options{
			Indent:   "  ",
			SortKeys: true,
			Width:    MaxLineWidth,
		}
		prettyJSON := pretty.PrettyOptions(jsonBytes, &opts)
		configs = append(configs, ConfigsByName{
			Name:   animal,
			Config: strings.Split(string(prettyJSON), "\n"),
		})
	}

	return configs, nil
}

// LoadAnimalConfig loads and parses a single animal config
func LoadAnimalConfig(configName string) (map[string]interface{}, error) {
	unified, err := loadCueValues()
	if err != nil {
		return nil, err
	}

	config := unified.LookupPath(cue.ParsePath(configName))
	if err := config.Err(); err != nil {
		return nil, fmt.Errorf("error looking up %s config: %w", configName, err)
	}

	var renovateConfig map[string]interface{}
	if err := config.Decode(&renovateConfig); err != nil {
		return nil, fmt.Errorf("error decoding %s config: %w", configName, err)
	}

	return renovateConfig, nil
}

// GenerateConfig handles the common flow for all config generators
func GenerateConfig(configName string) error {
	renovateConfig, err := LoadAnimalConfig(configName)
	if err != nil {
		return err
	}

	jsonBytes, err := json.Marshal(renovateConfig)
	if err != nil {
		return fmt.Errorf("error marshaling to JSON: %w", err)
	}

	// Configure pretty options with sorted keys
	opts := pretty.Options{
		Indent:   "  ",
		SortKeys: true,
		Width:    MaxLineWidth,
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
