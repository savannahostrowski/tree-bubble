package tree

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	bottomLeft string = " └──"
	
	white = lipgloss.Color("#ffffff")
	black = lipgloss.Color("#000000")
	purple  = lipgloss.Color("#bd93f9")
)

type Styles struct {
	Shapes     lipgloss.Style
	RootValue  lipgloss.Style
	RootDesc   lipgloss.Style
	ChildValue lipgloss.Style
	ChildDesc  lipgloss.Style
}

func defaultStyles() Styles {
	return Styles{
		Shapes:     lipgloss.NewStyle().Margin(0, 0, 0, 0).Foreground(purple),
		RootValue:  lipgloss.NewStyle().Margin(0, 0, 0, 0).Background(purple),
		RootDesc:   lipgloss.NewStyle().Margin(0, 0, 0, 0).Foreground(purple),
		ChildDesc:  lipgloss.NewStyle().Margin(0, 0, 0, 0).Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#ffffff"}), 
		ChildValue: lipgloss.NewStyle().Margin(0, 0, 0, 0).Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#ffffff"}),
	}
}

type Node struct {
	Value    string
	Desc     string
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
		m.renderTree(&b, m.nodes, 0)
	}
	return b.String()
}

func (m *Model) renderTree(b *strings.Builder, remainingNodes []Node, indent int) string {
	// Root Value - Root Description
	// 	└── Child Value - Child Description
	// 	└── Child Value - Child Description
	//  └── Child Value -  Child Description
	// 	└── Child Value - Child Description
	// 	└── Child Value - Child Description

	for _, node := range remainingNodes {

		if indent == 0 {
			str := m.Styles.RootValue.Render(node.Value) + "\t\t" + m.Styles.RootDesc.Render(node.Desc) + "\n"
			b.WriteString(str)
		}
		if node.Children != nil {
			m.renderTree(b, node.Children, indent+1)
		} else {
			str := strings.Repeat(" ", indent*2) + m.Styles.Shapes.Render(bottomLeft) + m.Styles.ChildValue.Render(node.Value) + "\t\t" + m.Styles.ChildDesc.Render(node.Desc) + "\n"
			b.WriteString(str)
		}

	}
	return b.String()
}
