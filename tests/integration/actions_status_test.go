// Copyright 2026 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package integration

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	actions_model "code.gitea.io/gitea/models/actions"
	"code.gitea.io/gitea/models/db"
	repo_model "code.gitea.io/gitea/models/repo"
	unit_model "code.gitea.io/gitea/models/unit"
	"code.gitea.io/gitea/models/unittest"
	repo_service "code.gitea.io/gitea/services/repository"
	"code.gitea.io/gitea/tests"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestActionsRunStatusFilterOptions(t *testing.T) {
	defer tests.PrepareTestEnv(t)()

	repo := unittest.AssertExistsAndLoadBean(t, &repo_model.Repository{ID: 4})
	require.NoError(t, repo_service.UpdateRepositoryUnits(db.DefaultContext, repo, []repo_model.RepoUnit{
		{
			RepoID: repo.ID,
			Type:   unit_model.TypeActions,
		},
	}, nil))

	session := loginUser(t, "user5")
	req := NewRequest(t, "GET", "/user5/repo4/actions")
	resp := session.MakeRequest(t, req, http.StatusOK)

	htmlDoc := NewHTMLParser(t, resp.Body)
	statusDropdown := htmlDoc.doc.Find(".ui.secondary.filter.menu .ui.dropdown.jump.item").Last()
	require.NotEmpty(t, statusDropdown.Text())

	expectedStatusLabels := map[actions_model.Status]string{
		actions_model.StatusSuccess:   "Success",
		actions_model.StatusFailure:   "Failure",
		actions_model.StatusCancelled: "Canceled",
		actions_model.StatusSkipped:   "Skipped",
		actions_model.StatusWaiting:   "Waiting",
		actions_model.StatusRunning:   "Running",
		actions_model.StatusBlocked:   "Blocked",
	}
	for _, status := range []actions_model.Status{
		actions_model.StatusSuccess,
		actions_model.StatusFailure,
		actions_model.StatusCancelled,
		actions_model.StatusSkipped,
		actions_model.StatusWaiting,
		actions_model.StatusRunning,
		actions_model.StatusBlocked,
	} {
		item := statusDropdown.Find(fmt.Sprintf(`a.item[href*="status=%d"]`, status))
		assert.Equal(t, 1, item.Length())
		assert.Equal(t, expectedStatusLabels[status], strings.TrimSpace(item.Text()))
		assert.Equal(t, 1, item.Find("span[data-tooltip-content] svg").Length())
	}
}
