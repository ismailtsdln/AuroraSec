package audit

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
)

// AWSError wraps AWS SDK errors for consistent handling
type AWSError struct {
	Code    string
	Message string
	Err     error
}

func (e *AWSError) Error() string {
	return fmt.Sprintf("AWS Error [%s]: %s (original: %v)", e.Code, e.Message, e.Err)
}

// WrapError converts standard AWS errors into AuroraSec error types
func WrapError(err error) error {
	if err == nil {
		return nil
	}
	// Simplified wrapping for now
	return &AWSError{
		Code:    "SDK_ERROR",
		Message: "An error occurred during AWS SDK call",
		Err:     err,
	}
}

// WithRetry executes a function with standard backoff retry logic
func WithRetry(ctx context.Context, fn func() error) error {
	maxRetries := 3
	backoff := 1 * time.Second

	var err error
	for i := 0; i < maxRetries; i++ {
		err = fn()
		if err == nil {
			return nil
		}

		// Check if error is retryable (simplified)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(backoff):
			backoff *= 2
		}
	}
	return err
}

// CustomRetryer implements AWS SDK retryer interface if needed
func NewCustomRetryer() aws.Retryer {
	return retry.NewStandard(func(o *retry.StandardOptions) {
		o.MaxAttempts = 3
	})
}
