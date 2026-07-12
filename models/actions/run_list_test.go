// Copyright 2026 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package actions

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStatusInfoListIncludesAggregateStatuses(t *testing.T) {
	statusInfos := GetStatusInfoList(context.Background())
	statuses := make([]Status, 0, len(statusInfos))
	displayedStatuses := make([]string, 0, len(statusInfos))
	for _, statusInfo := range statusInfos {
		statuses = append(statuses, Status(statusInfo.Status))
		displayedStatuses = append(displayedStatuses, statusInfo.DisplayedStatus)
	}

	assert.Equal(t, []Status{
		StatusSuccess,
		StatusFailure,
		StatusCancelled,
		StatusSkipped,
		StatusWaiting,
		StatusRunning,
		StatusBlocked,
	}, statuses)
	assert.Equal(t, []string{
		"success",
		"failure",
		"cancelled",
		"skipped",
		"waiting",
		"running",
		"blocked",
	}, displayedStatuses)
}
