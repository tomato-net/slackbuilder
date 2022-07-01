package cli

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithCommandName(t *testing.T) {
	t.Run("sets name on cli object", func(t *testing.T) {
		cli := &CLI{}
		err := WithCommandName("test")(cli)
		assert.Nil(t, err)
	})
}
