package report

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ismailtsdln/AuroraSec/internal/pkg/audit"
)

func SaveJSON(result *audit.Result, filepath string) error {
	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal results to JSON: %v", err)
	}

	if filepath != "" {
		err := os.WriteFile(filepath, data, 0644)
		if err != nil {
			return fmt.Errorf("failed to save JSON report to %s: %v", filepath, err)
		}
		fmt.Printf("✅ JSON report saved to %s\n", filepath)
	} else {
		fmt.Println(string(data))
	}

	return nil
}
