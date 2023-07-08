package tree

import (
	"io"
	"testing"
)

type item string

type itemDelegate struct {}

func (d itemDelegate) Height() int {return 1}
func (d itemDelegate) Width() int {return 0}
func (d itemDelegate) Render(w io.Writer, m Model, index int, node Node) error {
	i, ok := node.(Node)
}

func TestTreeItemName (t *testing.T) {
	tree := New([]Node{"a", "b", "c"}, itemDelegate{}, 0, 0)
}