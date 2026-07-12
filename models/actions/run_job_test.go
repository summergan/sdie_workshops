// Copyright 2026 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package actions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAggregateJobStatus(t *testing.T) {
	tests := []struct {
		name     string
		statuses []Status
		want     Status
	}{
		{
			name:     "all skipped jobs aggregate to skipped",
			statuses: []Status{StatusSkipped, StatusSkipped},
			want:     StatusSkipped,
		},
		{
			name:     "success and skipped jobs aggregate to success",
			statuses: []Status{StatusSuccess, StatusSkipped},
			want:     StatusSuccess,
		},
		{
			name:     "cancelled terminal job aggregates to cancelled",
			statuses: []Status{StatusSuccess, StatusCancelled, StatusFailure},
			want:     StatusCancelled,
		},
		{
			name:     "running job keeps run running",
			statuses: []Status{StatusBlocked, StatusRunning, StatusCancelled},
			want:     StatusRunning,
		},
		{
			name:     "waiting job keeps run waiting when no job is running",
			statuses: []Status{StatusSuccess, StatusWaiting, StatusBlocked},
			want:     StatusWaiting,
		},
		{
			name:     "blocked job is visible when no job is waiting or running",
			statuses: []Status{StatusSuccess, StatusBlocked},
			want:     StatusBlocked,
		},
		{
			name:     "failure terminal job aggregates to failure",
			statuses: []Status{StatusSuccess, StatusFailure, StatusSkipped},
			want:     StatusFailure,
		},
		{
			name:     "empty job list falls back to unknown",
			statuses: nil,
			want:     StatusUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jobs := make([]*ActionRunJob, 0, len(tt.statuses))
			for _, status := range tt.statuses {
				jobs = append(jobs, &ActionRunJob{Status: status})
			}

			assert.Equal(t, tt.want, aggregateJobStatus(jobs))
		})
	}
}
