package yamlfile

import (
	"fmt"
	"os"
	"path/filepath"

	yaml3 "gopkg.in/yaml.v3"
)

// UnmarshalAllYAMLFiles reads and unmarshals all YAML files in the specified directory with the specified suffix (including dot) as a map of the absolute file path to the target YAML container.
func UnmarshalAllYAMLFiles[TTarget any](parentDirPath string, targetFileSuffix string) (map[string]TTarget, error) {
	dirEntries, readErr := os.ReadDir(parentDirPath)
	if readErr != nil {
		return nil, fmt.Errorf("error reading directory '%s': %w", parentDirPath, readErr)
	}

	allTargets := make(map[string]TTarget, 0)
	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			continue
		}

		targetFileName := dirEntry.Name()
		if filepath.Ext(targetFileName) != targetFileSuffix {
			continue
		}

		var target TTarget
		filePathAbs := filepath.Join(parentDirPath, targetFileName)
		fileLoadErr := UnmarshalYAMLFile(filePathAbs, &target)
		if fileLoadErr != nil {
			return nil, fileLoadErr
		}

		allTargets[filePathAbs] = target
	}
	return allTargets, nil
}

// UnmarshalYAMLFile reads and unmarshals the specified YAML file into the target YAML container.
func UnmarshalYAMLFile[TTarget any](filePathAbs string, target *TTarget) error {
	fileContents, readFileErr := os.ReadFile(filePathAbs)
	if readFileErr != nil {
		return fmt.Errorf("error reading file contents '%s': %w", filePathAbs, readFileErr)
	}

	unmarshalErr := yaml3.Unmarshal(fileContents, target)
	if unmarshalErr != nil {
		return fmt.Errorf("error unmarshalling yaml file contents '%s': %w", filePathAbs, unmarshalErr)
	}

	return nil
}
