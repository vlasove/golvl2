package command

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vlasove/materials/tasks_2/utils/shell/internal/app/shell"
)

func TestCommand_BuildPrefix(t *testing.T) {
	s := shell.New()
	writer, reader := bytes.Buffer{}, bytes.Buffer{}
	st := New(s, &reader, &writer)
	res, err := st.buildPrefix()
	assert.NoError(t, err)
	assert.True(t, len(res) > 0)
}
