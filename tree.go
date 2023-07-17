package tree

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	bottomLeft string = " └──"

	white  = lipgloss.Color("#ffffff")
	black  = lipgloss.Color("#000000")
	purple = lipgloss.Color("#bd93f9")
)

type Styles struct {
	Shapes     lipgloss.Style
	RootValue  lipgloss.Style
	RootDesc   lipgloss.Style
	ChildValue lipgloss.Style
	ChildDesc  lipgloss.Style
	Help  lipgloss.Style
}

func defaultStyles() Styles {
	return Styles{
		Shapes:     lipgloss.NewStyle().Margin(0, 0, 0, 0).Foreground(purple),
		RootValue:  lipgloss.NewStyle().Margin(0, 0, 0, 0).Background(purple),
		RootDesc:   lipgloss.NewStyle().Margin(0, 0, 0, 0).Foreground(purple),
		ChildDesc:  lipgloss.NewStyle().Margin(0, 0, 0, 0).Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#ffffff"}),
		ChildValue: lipgloss.NewStyle().Margin(0, 0, 0, 0).Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#ffffff"}),
		Help: 	 lipgloss.NewStyle().Margin(0, 0, 0, 0).Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#ffffff"}),
	}
}

type Node struct {
	Value    string
	Desc     string
	Children []Node
}

type Model struct {
	KeyMap KeyMap
	Styles Styles

	width  int
	height int
	nodes  []Node
	cursor int

	Help     help.Model
	showHelp bool

	AdditionalShortHelpKeys func() []key.Binding
}

func New(nodes []Node, width int, height int) Model {
	return Model{
		KeyMap: DefaultKeyMap(),
		Styles: defaultStyles(),

		width:  width,
		height: height,
		nodes:  nodes,

		showHelp: true,
		Help:     help.New(),
	}
}

// KeyMap holds the key bindings for the table.
type KeyMap struct {
	Bottom      key.Binding
	Top         key.Binding
	SectionDown key.Binding
	SectionUp   key.Binding
	Down        key.Binding
	Up          key.Binding
	Quit        key.Binding

	ShowFullHelp  key.Binding
	CloseFullHelp key.Binding

}

// DefaultKeyMap is the default key bindings for the table.
func DefaultKeyMap() KeyMap {
	return KeyMap{
		Bottom: key.NewBinding(
			key.WithKeys("bottom"),
			key.WithHelp("end", "bottom"),
		),
		Top: key.NewBinding(
			key.WithKeys("top"),
			key.WithHelp("home", "top"),
		),
		SectionDown: key.NewBinding(
			key.WithKeys("secdown"),
			key.WithHelp("secdown", "section down"),
		),
		SectionUp: key.NewBinding(
			key.WithKeys("secup"),
			key.WithHelp("secup", "section up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down"),
			key.WithHelp("↓", "down"),
		),
		Up: key.NewBinding(
			key.WithKeys("up"),
			key.WithHelp("↑", "up"),
		),

		ShowFullHelp: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "more"),
		),
		CloseFullHelp: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "close help"),
		),

		Quit: key.NewBinding(
			key.WithKeys("q", "esc"),
			key.WithHelp("q", "quit"),
		),
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

func (m Model) Cursor() int {
	return m.cursor
}

func (m *Model) SetCursor(cursor int) {
	m.cursor = cursor
}

func (m Model) isCursorAtRoot() bool {
	return m.cursor == 0
}

func (m Model) isCursorAtBottom() bool {
	return m.cursor == len(m.nodes)-1
}

func (m *Model) setShowHelp() bool {
	return m.showHelp
}

func (m *Model) navUp() {
	if m.isCursorAtRoot() {
		return
	}
	m.cursor--
}

func (m *Model) navDown() {
	if m.isCursorAtBottom() {
		return
	}
	m.cursor++
}

// func (m *Model) navTop() {
// 	m.cursor = 0
// }

// func (m *Model) navBottom() {
// 	m.cursor = len(m.nodes) - 1
// }

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.Up):
			fmt.Println(msg)
			m.navUp()
		case key.Matches(msg, m.KeyMap.Down):
			fmt.Println(msg)
			m.navDown()
			// case key.Matches(msg, m.KeyMap.Top):
			// 	m.navTop()
			// case key.Matches(msg, m.KeyMap.Bottom):
			// 	m.navBottom()
		case key.Matches(msg, m.KeyMap.ShowFullHelp):
			fallthrough
		case key.Matches(msg, m.KeyMap.CloseFullHelp):
			m.Help.ShowAll = !m.Help.ShowAll
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	availableHeight := m.height
	var b strings.Builder


	nodes := m.Nodes()

	var help string
	if m.showHelp {
		help = m.helpView()
		availableHeight -= lipgloss.Height(help)
	}

	b.WriteString(lipgloss.NewStyle().Height(availableHeight).Render(m.renderTree(&b, m.nodes, 0)))
	b.WriteString(help)

	if len(nodes) == 0 {
		return "No data"
	}
	return b.String()
}

func (m *Model) renderTree(b *strings.Builder, remainingNodes []Node, indent int) string {
	// Root Value - Root Description
	// 	└── Child Value - Child Description
	// 	└── Child Value - Child Description
	//  └── Child Value -  Child Description
	// 		└── Child Value - Child Description
	// 		└── Child Value - Child Description

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

func (m Model) helpView() string {
	return m.Styles.Help.Render(m.Help.View(m))
}

func (m Model) ShortHelp() []key.Binding {
	kb := []key.Binding{
		m.KeyMap.Up,
		m.KeyMap.Down,
	}

	if m.AdditionalShortHelpKeys != nil {
		kb = append(kb, m.AdditionalShortHelpKeys()...)
	}

	return append(kb,
		m.KeyMap.Quit,
	)
}

func (m Model) FullHelp() [][]key.Binding {
	kb := [][]key.Binding{{
		m.KeyMap.Up,
		m.KeyMap.Down,
	}}


	return append(kb,
		[]key.Binding{
			m.KeyMap.Quit,
			m.KeyMap.CloseFullHelp,
		})
}