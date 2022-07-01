package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()

	t.Run("creates command", func(t *testing.T) {
		c := New("test-name", "test-description")

		assert.Equal(t, "test-name", c.Use)
		assert.Equal(t, "test-description", c.Long)
	})

	t.Run("has subcommands", func(t *testing.T) {
		c := New("test-name", "test-description")

		assert.True(t, c.HasSubCommands())
	})
}
