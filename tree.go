package tree

import (
	"fmt"
	"io"
	"strings"
	"unicode"

	tea "github.com/charmbracelet/bubbletea"
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

func DefaultStyles() Styles {
	return Styles{
		Shapes: lipgloss.NewStyle().Foreground(lipgloss.Color("69")),
		Root:   lipgloss.NewStyle().Bold(true).Background(lipgloss.Color("69")),
	}
}

type Item struct {
}

type ItemDelegate interface {
	Render(w io.Writer, m Model, index int, item Item) error

	Height() int

	Width() int
}

type Model struct {
	Styles   Styles

	delegate ItemDelegate
	width int
	height int
	items []Item
}

func New(items []Item, delegate ItemDelegate, width int, height int) Model {
	return Model{
		delegate: delegate,
		Styles:   DefaultStyles(),

		width:  width,
		height: height,
		items: items,
	}
}

func (m Model) Items() []Item {
	return m.items
}

func (m *Model) SetItems(items []Item) {
	m.items = items
}

func (m *Model) SetDelegate(delegate ItemDelegate) {
	m.delegate = delegate
}

func (m Model) Width() int {
	return m.width
}

func (m Model) Height() int {
	return m.height
}

func (m *Model) SetSize(width, height int) {
	m.SetSize(width, height)
}

func (m *Model) SetWidth(newWidth int) {
	m.SetSize(newWidth, m.height)
}

func (m *Model) SetHeight(newHeight int) {
	m.SetSize(m.width, newHeight)
}


func (m Model) View() string {
	var (
		availableHeight = m.height
	)

	content := lipgloss.NewStyle().Height(availableHeight).Render(m.pop)
}