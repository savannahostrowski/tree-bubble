package tree

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	topLeft     string = "╭"
	bottomLeft  string = "╰"
	topRight    string = "╮"
	bottomRight string = "╯"
	bar         string = "|"
	left        string = "├"
	right       string = "┤"
)

type Styles struct {
	Shapes lipgloss.Style
	Root   lipgloss.Style
	Child  lipgloss.Style
}

func defaultStyles() Styles {
	return Styles{
		Shapes: lipgloss.NewStyle().Foreground(lipgloss.Color("69")),
		Root:   lipgloss.NewStyle().Padding(0, 0, 1, 2).Bold(true).Background(lipgloss.Color("62")).Foreground(lipgloss.Color("230")),
		Child:  lipgloss.NewStyle().Padding(0, 0, 1, 2).Foreground(lipgloss.Color("230")),
	}
}

type Node struct{
	Value string
	Desc string
	Children []Node
}

type Model struct {
	Styles Styles

	width  int
	height int
	nodes  []Node
}

func New(nodes []Node, width int, height int) Model {
	return Model{
		Styles: defaultStyles(),

		width:  width,
		height: height,
		nodes:  nodes,
	}
}

func (m Model) Nodes() []Node {
	return m.nodes
}

func (m *Model) SetNodes(nodes []Node) {
	m.nodes = nodes
}

func (m Model) Width() int {
	return m.width
}

func (m Model) Height() int {
	return m.height
}

func (m *Model) SetSize(width, height int) {
	m.width = width
	m.height = height
}

func (m *Model) SetWidth(newWidth int) {
	m.SetSize(newWidth, m.height)
}

func (m *Model) SetHeight(newHeight int) {
	m.SetSize(m.width, newHeight)
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	nodes := m.Nodes()

	var b strings.Builder

	if len(nodes) == 0 {
		return "No data"
	}

	if len(nodes) > 0 {
		m.renderTree(&b, m.nodes)
	}
	return b.String()
}

func (m *Model) renderTree(b *strings.Builder, remainingNodes []Node) string {
	if len(m.nodes) == 0 {
		return b.String()
	}

	for i, node := range m.nodes {
		if len(node.Children) > 0 {
			m.renderTree(b, node.Children)
			
		} else {
			if i == 0 {
				fmt.Println(node.Value)
				b.WriteString(m.Styles.Shapes.Render(topLeft) + node.Value)
				b.WriteString(m.Styles.Root.Render("\n\t" + node.Desc))
				b.WriteString(m.Styles.Root.Render(string(bar)))
			} else {
				b.WriteString(m.Styles.Shapes.Render(left))
				b.WriteString(m.Styles.Root.Render("\n\t" + node.Desc))
			}
		}
	}
	return b.String()
}
