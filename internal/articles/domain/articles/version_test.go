package articles_test

import (
	"github.com/qmstar0/BlogLite-api/internal/articles/domain/articles"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewVersion(t *testing.T) {
	t.Parallel()
	var err error
	var hash = "a94a8fe5ccb19ba61c4c0873d391e987982fbbd"

	_, err = articles.NewVersion("", "", "", "", "", "")
	assert.Error(t, err)

	_, err = articles.NewVersion("title", "description", "text", hash, "source", "note")
	assert.NoError(t, err)
}
