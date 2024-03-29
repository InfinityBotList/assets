package popltypes

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// @ci table=changelogs
//
// Changelogs for the list
type ChangelogEntry struct {
	Version          string      `db:"version" json:"version" validate:"required" description:"The version for the changelog entry. (4.3.0 etc.)"`
	ExtraDescription string      `db:"extra_description" json:"extra_description" description:"The extra description for the version, if applicable"`
	GithubHTML       pgtype.Text `db:"github_html" json:"github_html" description:"The Github-backed HTML for the changelog entry."`
	Published        bool        `db:"published" json:"published" description:"Whether or not this is a published changelog"`
	Prerelease       bool        `db:"prerelease" json:"prerelease" description:"Whether or not this is a prerelease."`
	Added            []string    `db:"added" json:"added" validate:"required" description:"The added features for the version."`
	Updated          []string    `db:"updated" json:"updated" validate:"required" description:"The changed features for the version."`
	Removed          []string    `db:"removed" json:"removed" validate:"required" description:"The removed features for the version."`
	CreatedAt        time.Time   `db:"created_at" json:"created_at" description:"The time the changelog entry was created."`
}

type Changelog struct {
	Entries []ChangelogEntry `json:"entries" validate:"required,dive,required" description:"The changelog entries."`
}
