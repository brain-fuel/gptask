package task

import (
	"errors"
	"testing"

	stdworkflow "goforge.dev/goplus/std/workflow"
)

// go-task classifies each task run into a std/workflow.Outcome; these pin the
// mapping and the error projection that RunTask relies on.

func TestClassifyRunOutcome(t *testing.T) {
	boom := errors.New("boom")
	cases := []struct {
		name             string
		preCondMet       bool
		upToDate         bool
		runErr           error
		wantSkip         bool
		wantSkipReason   string
		wantErr          error
	}{
		{"completed", true, false, nil, false, "", nil},
		{"up-to-date", true, true, nil, true, "up to date", nil},
		{"precondition", false, false, nil, true, "precondition not met", nil},
		{"failed", true, false, boom, false, "", boom},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			o := classifyRunOutcome(c.preCondMet, c.upToDate, c.runErr)
			reason, isSkip := stdworkflow.SkipReason(o)
			if isSkip != c.wantSkip || reason != c.wantSkipReason {
				t.Fatalf("skip = (%q,%v), want (%q,%v)", reason, isSkip, c.wantSkipReason, c.wantSkip)
			}
			if got := stdworkflow.OutcomeError(o); got != c.wantErr {
				t.Fatalf("OutcomeError = %v, want %v", got, c.wantErr)
			}
			// a non-failing outcome never surfaces an error to the runner.
			if c.wantErr == nil && !stdworkflow.Succeeded(o) {
				t.Fatal("non-error outcome must Succeed")
			}
		})
	}
}
