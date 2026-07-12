// Copyright 2023 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package repo

import (
	"testing"

	"code.gitea.io/gitea/models/db"
	"code.gitea.io/gitea/models/unittest"
	"code.gitea.io/gitea/modules/optional"

	"github.com/stretchr/testify/assert"
)

func TestMigrate_InsertReleases(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	a := &Attachment{
		UUID: "a0eebc91-9c0c-4ef7-bb6e-6bb9bd380a12",
	}
	r := &Release{
		Attachments: []*Attachment{a},
	}

	err := InsertReleases(db.DefaultContext, r)
	assert.NoError(t, err)
}

func TestFindReleasesOptionsKeyword(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	releases, err := db.Find[Release](db.DefaultContext, FindReleasesOptions{
		ListOptions:   db.ListOptions{ListAll: true},
		RepoID:        1,
		IncludeDrafts: true,
		IncludeTags:   true,
		HasSha1:       optional.Some(true),
		Keyword:       "V1",
	})
	assert.NoError(t, err)

	tagNames := make([]string, 0, len(releases))
	for _, release := range releases {
		tagNames = append(tagNames, release.TagName)
	}
	assert.EqualValues(t, []string{"v1.0", "v1.1"}, tagNames)

	count, err := db.Count[Release](db.DefaultContext, FindReleasesOptions{
		ListOptions:   db.ListOptions{ListAll: true},
		RepoID:        1,
		IncludeDrafts: true,
		IncludeTags:   true,
		HasSha1:       optional.Some(true),
		Keyword:       "V1",
	})
	assert.NoError(t, err)
	assert.EqualValues(t, 2, count)
}
