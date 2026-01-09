package audit

import (
	"context"
	"time"
)

// Severity represents the impact of a security finding
type Severity string

const (
	SeverityCritical Severity = "CRITICAL"
	SeverityHigh     Severity = "HIGH"
	SeverityMedium   Severity = "MEDIUM"
	SeverityLow      Severity = "LOW"
	SeverityInfo     Severity = "INFO"
)

// Finding represents a single security issue discovered during audit
type Finding struct {
	Module      string    `json:"module"`
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Severity    Severity  `json:"severity"`
	Resource    string    `json:"resource"`
	Status      string    `json:"status"` // PASS, FAIL, WARN
	Remediation string    `json:"remediation"`
	Timestamp   time.Time `json:"timestamp"`
}

// Result contains all findings from an audit run
type Result struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Findings  []Finding `json:"findings"`
	Summary   struct {
		Total    int `json:"total"`
		Critical int `json:"critical"`
		High     int `json:"high"`
		Medium   int `json:"medium"`
		Low      int `json:"low"`
		Passed   int `json:"passed"`
		Failed   int `json:"failed"`
	} `json:"summary"`
}

// Module interface that all security modules must implement
type Module interface {
	Name() string
	Description() string
	Audit(ctx context.Context) ([]Finding, error)
}

// Engine coordinates the execution of audit modules
type Engine struct {
	modules []Module
}

func NewEngine() *Engine {
	return &Engine{
		modules: make([]Module, 0),
	}
}

func (e *Engine) RegisterModule(m Module) {
	e.modules = append(e.modules, m)
}

func (e *Engine) Run(ctx context.Context) (*Result, error) {
	result := &Result{
		StartTime: time.Now(),
		Findings:  make([]Finding, 0),
	}

	for _, m := range e.modules {
		findings, err := m.Audit(ctx)
		if err != nil {
			return nil, err
		}
		result.Findings = append(result.Findings, findings...)
	}

	result.EndTime = time.Now()
	e.calculateSummary(result)
	return result, nil
}

func (e *Engine) calculateSummary(r *Result) {
	r.Summary.Total = len(r.Findings)
	for _, f := range r.Findings {
		if f.Status == "PASS" {
			r.Summary.Passed++
			continue
		}
		r.Summary.Failed++
		switch f.Severity {
		case SeverityCritical:
			r.Summary.Critical++
		case SeverityHigh:
			r.Summary.High++
		case SeverityMedium:
			r.Summary.Medium++
		case SeverityLow:
			r.Summary.Low++
		}
	}
}
