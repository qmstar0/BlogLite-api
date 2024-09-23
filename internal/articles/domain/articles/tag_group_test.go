package articles_test

import (
	"github.com/qmstar0/BlogLite-api/internal/articles/domain/articles"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTagGroup(t *testing.T) {
	t.Parallel()
	var err error
	_, err = articles.NewTagGroup([]string{"1", "2", "3", "4", "5"})
	assert.Error(t, err)

	group, err := articles.NewTagGroup([]string{"1", "2", "3", "4", "4"})
	assert.NoError(t, err)
	assert.Equal(t, group.Value(), []string{"1", "2", "3", "4"})
}
