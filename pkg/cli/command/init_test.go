package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInit(t *testing.T) {
	t.Run("creates command", func(t *testing.T) {
		c := NewInit()

		assert.Equal(t, "init", c.Use)
	})
}