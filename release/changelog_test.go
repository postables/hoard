package release

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChangelogForReleases(t *testing.T) {
	changelog, err := changelogForReleases(
		[]Release{
			{
				Version: "0.1.0",
				Notes:   `Everything broken`,
			},
			{
				Version: "0.0.2",
				Notes:   `Wonderful things were achieved`,
			},
			{
				Version: "0.0.1",
				Notes:   `Marvelous advances were made`,
			},
		})
	assert.NoError(t, err)

	assert.Equal(t, `# Hoard Changelog
## Version 0.1.0
Everything broken

## Version 0.0.2
Wonderful things were achieved

## Version 0.0.1
Marvelous advances were made
`, changelog)
}
