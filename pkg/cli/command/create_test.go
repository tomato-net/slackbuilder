package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCreate(t *testing.T) {
	t.Run("creates command", func(t *testing.T) {
		c := NewCreate()

		assert.Equal(t, "create", c.Use)
		assert.Equal(t, []string{"new"}, c.SuggestFor)
	})

	t.Run("registers subcommands", func(t *testing.T) {
		c := NewCreate()

		assert.True(t, c.HasSubCommands())
	})
}

func TestNewCreateSlashCMD(t *testing.T) {
	t.Run("creates command", func(t *testing.T) {
		c := NewCreateSlashCMD()

		assert.Equal(t, "slashcmd", c.Use)
	})
}
