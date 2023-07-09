package tree

import (
	"fmt"
	"io"
	"testing"
)

type node string

type nodeDelegate struct{}

func (d nodeDelegate) Height() int {return 1}
func (d nodeDelegate) Width() int {return 0}
func (d nodeDelegate) Render(w io.Writer, m Model, index int, node Node) {
	i, ok := node.(Node)

	if !ok {
		return 
	}

	str := fmt.Sprintf("%d. %s", index + 1, i)
	fmt.Fprintln(w, m.Styles.Root.Render(str))
}

func TestFlatTree(t *testing.T) {
	tree := New([]Node{
		node("root node"),
		node("one"),
		node("two"),
		node("three"),
		node("four"),
		node("five"),
		node("six"),
		node("seven"),
		node("eight"),
		node("nine"),
		node("ten"),
		node("eleven"),
	}, nodeDelegate{}, 0, 0)

	if len(tree.Nodes()) != 12 {
		t.Errorf("expected 12 nodes, got %d", len(tree.Nodes()))
	}
}

func TestNestedTree(t *testing.T) {
	nestedTree := New([]Node{
		node("nested root"),
		node("nested two"),
		node("nested three"),
		node("nested four"),
		node("nested five"),
	}, nodeDelegate{}, 0, 0)

	tree := New([]Node{
		node("root node"),
		node("one"),
		nestedTree,
		node("three"),
		node("four"),
	}, nodeDelegate{}, 0, 0)

	numOfNodes := 0
	for i, n := range tree.Nodes() {
		// If the node is a nested tree, we need to count the nodes in the nested tree
		if _, ok := n.(Tree); ok {
			numOfNodes += len(n.(Tree).Nodes())
			continue
		}
		

	}

	if numOfNodes != 9 {
		t.Errorf("expected 9 nodes, got %d", numOfNodes)
	}
}