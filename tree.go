package tree

import (
	"io"

	"github.com/charmbracelet/lipgloss"
)

type Shape string

const (
	topLeft     Shape = "╭"
	bottomLeft  Shape = "╰"
	topRight    Shape = "╮"
	bottomRight Shape = "╯"
	bar         Shape = "|"
	left        Shape = "├"
	right       Shape = "┤"
)

type Styles struct {
	Shapes lipgloss.Style
	Root   lipgloss.Style
}

func defaultStyles() Styles {
	return Styles{
		Shapes: lipgloss.NewStyle().Foreground(lipgloss.Color("69")),
		Root:   lipgloss.NewStyle().Padding(0, 0, 1, 2).Bold(true).Background(lipgloss.Color("62")).Foreground(lipgloss.Color("230")),
	}
}


type Node interface {
	Children() []Node
}


type NodeDelegate interface {
	Render(w io.Writer, m Model, index int, node Node) error

	Height() int

	Width() int
}

type Model struct {
	Styles   Styles

	delegate NodeDelegate
	width int
	height int
	nodes []Node
}

func New(nodes []Node, delegate NodeDelegate, width int, height int) Model {
	return Model{
		delegate: delegate,
		Styles:   defaultStyles(),

		width:  width,
		height: height,
		nodes: nodes,
	}
}

func (m Model) Nodes() []Node {
	return m.nodes
}

func (m *Model) SetNodes(nodes []Node) {
	m.nodes = nodes
}

func (m *Model) SetDelegate(delegate NodeDelegate) {
	m.delegate = delegate
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


