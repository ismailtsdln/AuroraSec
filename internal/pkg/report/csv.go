package report

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/ismailtsdln/AuroraSec/internal/pkg/audit"
)

func SaveCSV(result *audit.Result, filepath string) error {
	var f *os.File
	var err error

	if filepath != "" {
		f, err = os.Create(filepath)
		if err != nil {
			return fmt.Errorf("failed to create CSV file: %v", err)
		}
		defer f.Close()
	} else {
		f = os.Stdout
	}

	writer := csv.NewWriter(f)
	defer writer.Flush()

	// Header
	header := []string{"Severity", "Status", "Module", "ID", "Title", "Resource", "Remediation", "Timestamp"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("failed to write CSV header: %v", err)
	}

	// Data
	for _, finding := range result.Findings {
		row := []string{
			string(finding.Severity),
			finding.Status,
			finding.Module,
			finding.ID,
			finding.Title,
			finding.Resource,
			finding.Remediation,
			finding.Timestamp.Format("2006-01-02 15:04:05"),
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("failed to write CSV row: %v", err)
		}
	}

	if filepath != "" {
		fmt.Printf("✅ CSV report saved to %s\n", filepath)
	}

	return nil
}
