package task

import (
	"cmp"
	"fmt"

	"goforge.dev/gptask/internal/hash"
	"goforge.dev/gptask/taskfile/ast"
)

func (e *Executor) GetHash(t *ast.Task) (string, error) {
	r := cmp.Or(t.Run, e.Taskfile.Run)
	var h hash.HashFunc
	switch r {
	case "always":
		h = hash.Empty
	case "once":
		h = hash.Name
	case "when_changed":
		h = hash.Hash
	default:
		return "", fmt.Errorf(`task: invalid run "%s"`, r)
	}
	return h(t)
}
