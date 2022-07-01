package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRoot(t *testing.T) {
	t.Parallel()

	t.Run("creates command", func(t *testing.T) {
		c := NewRoot("test-name", "test-description")

		assert.Equal(t, "test-name", c.Use)
		assert.Equal(t, "test-description", c.Long)
	})
}
