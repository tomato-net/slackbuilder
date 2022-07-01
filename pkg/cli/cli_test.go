package cli

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		give    []Option
		wantCLI *CLI
	}{
		{
			name:    "setting options",
			give:    []Option{WithCommandName("test")},
			wantCLI: &CLI{commandName: "test"},
		},
		{
			name:    "no options",
			wantCLI: &CLI{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCLI, err := New(tt.give...)

			assert.Nil(t, err)
			assert.Equal(t, tt.wantCLI.commandName, gotCLI.commandName)
			assert.Equal(t, tt.wantCLI.description, gotCLI.description)
			assert.True(t, gotCLI.cmd.HasSubCommands())
		})
	}
}

func TestCLI_Run(t *testing.T) {
	t.Run("does not error", func(t *testing.T) {
		c, err := New()

		assert.Nil(t, err)
		assert.Nil(t, c.Run())
	})
}
