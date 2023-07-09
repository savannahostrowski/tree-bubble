package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	tree "github.com/savannahostrowski/tree-bubble"
	"golang.org/x/term"
)

var (
	styleDoc = lipgloss.NewStyle().Padding(1)
)

type node struct {
	value    string
	desc     string
	children []tree.Node
}

func (n node) Value() string         { return n.value }
func (n node) Description() string   { return n.desc }
func (n node) Children() []tree.Node { return n.children }

func main() {
	err := tea.NewProgram(initialModel()).Start()
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}

func initialModel() model {
	w, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		w = 80
		h = 24
	}
	top, right, bottom, left := styleDoc.GetPadding()
	w = w - left - right
	h = h - top - bottom

	t := tree.New([]tree.Node{
		node{value: "history | grep docker", desc: "Used in a Unix-like operating system to search through the " +
			"command history for any entries that contain the word 'docker.'", children: []tree.Node{
			node{value: "history", desc: "Shows the history of all commands in the terminal", children: nil},
			node{value: "|", desc: "Used to combine two or more commands", children: nil},
			node{value: "grep", desc: "Short for 'global regular expression print'; A command used in searching and matching text files contained in the regular expressions.", children: nil},
			node{value: "docker", desc: "Used to interact with Docker", children: nil},
		},
		}}, w, h)

	return model{tree: t}
}

type model struct {
	tree tree.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.tree, cmd = m.tree.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return styleDoc.Render(m.tree.View())
}
