package articles_test

import (
	"github.com/qmstar0/BlogLite-api/internal/articles/domain/articles"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUri(t *testing.T) {
	t.Parallel()
	testsErr := []string{
		"test ",
		" test",
		" te st ",
		" test ",
		"te st",
		"te st",
		"*test",
		"!test",
		"te!@$@%%^@*#@&*#st",
		"'test'",
		"`test`",
		"",
		" ",
		`"UPDATE article SET uri = 'updated' WHERE uri = 'test"'`,
	}

	testsOk := []string{
		"test",
		"test-1",
		"test_1",
		"test1",
		"test123",
		"test123-test456-test789",
		"test123_test456_test789",
	}

	t.Run("correct format test", func(t *testing.T) {
		var err error
		for _, s := range testsErr {
			err = articles.NewUri(s).CheckFormat()
			assert.Error(t, err)
		}
	})
	t.Run("wrong format test", func(t *testing.T) {
		var err error
		for _, s := range testsOk {
			err = articles.NewUri(s).CheckFormat()
			assert.NoError(t, err)
		}
	})
}
