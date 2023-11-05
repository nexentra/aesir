package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
	"github.com/nexentra/aesir/importer"
)

var (
	term   = termenv.EnvColorProfile()
	subtle = makeFgStyle("241")
	dot    = colorFg(" • ", "236")
)

func makeFgStyle(color string) func(string) string {
	return termenv.Style{}.Foreground(term.Color(color)).Styled
}

func colorFg(val, color string) string {
	return termenv.String(val).Foreground(term.Color(color)).String()
}

func checkbox(label string, checked bool) string {
	if checked {
		return colorFg("[x] "+label, "212")
	}
	return fmt.Sprintf("[ ] %s", label)
}

type Model struct {
	page         string
	filepicker   filepicker.Model
	selectedFile string
	Choice       int
	Chosen       bool
	quitting     bool
	err          error
}

type clearErrorMsg struct{}

func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return clearErrorMsg{}
	})
}

func (m Model) Init() tea.Cmd {
	return m.filepicker.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		// case "enter":
			// m.page = "fs"
		}
	case tea.WindowSizeMsg:
		m.filepicker, cmd = m.filepicker.Update(msg)
	case clearErrorMsg:
		m.err = nil
	default:
		m.filepicker, cmd = m.filepicker.Update(msg)
	}

	switch m.page {
	case "fs":
		var fcmd tea.Cmd
		m.filepicker, fcmd = m.filepicker.Update(msg)
		cmd = tea.Batch(cmd, fcmd)
		if didSelect, path := m.filepicker.DidSelectFile(msg); didSelect {
			m.selectedFile = path
		}

		if didSelect, path := m.filepicker.DidSelectDisabledFile(msg); didSelect {
			m.err = errors.New(path + " is not valid.")
			m.selectedFile = ""
			return m, tea.Batch(cmd, clearErrorAfter(2*time.Second))
		}

		return m, cmd
	case "repl":
		return m, tea.Quit
	default:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "j", "down":
				m.Choice++
				if m.Choice > 1 {
					m.Choice = 0
				}
			case "k", "up":
				m.Choice--
				if m.Choice < 0 {
					m.Choice = 1
				}
			case "enter":
				if m.Choice == 0 {
					m.page = "fs"
					// m.filepicker.CurrentDirectory = "./"
				}
				m.Chosen = true
				return m, cmd
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	if m.quitting {
		return ""
	}

	if m.selectedFile != "" {
		res, err := importer.Importer(string(m.selectedFile))
		if err != nil {
			print(err)
		}
		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		res.Inspect()
		w.Close()
		out, _ := io.ReadAll(r)
		os.Stdout = rescueStdout

		return string(out)
	}

	switch m.page {
	case "fs":
		var s strings.Builder
		s.WriteString("\n  ")
		if m.err != nil {
			s.WriteString(m.filepicker.Styles.DisabledFile.Render(m.err.Error()))
		} else if m.selectedFile == "" {
			s.WriteString("Pick a file:")
		} else {
			s.WriteString("Selected file: " + m.filepicker.Styles.Selected.Render(m.selectedFile))
		}
		s.WriteString("\n\n" + m.filepicker.View() + "\n")
		return s.String()
	case "repl":
		return "repl view"
	default:

		c := m.Choice

		tpl := "What to do today?\n\n"
		tpl += "%s\n\n"
		tpl += "Program quits in %s seconds\n\n"
		tpl += subtle("j/k, up/down: select") + dot + subtle("enter: choose") + dot + subtle("q, esc: quit")

		choices := fmt.Sprintf(
			"%s\n%s\n",
			checkbox("Choose an aesir file", c == 0),
			checkbox("Fire up REPL", c == 1),
		)
		return fmt.Sprintf(tpl, choices, colorFg(strconv.Itoa(m.Choice), "79"))

	}
}

func main() {
	fp := filepicker.New()
	fp.AllowedTypes = []string{".ae"}
	fp.CurrentDirectory, _ = os.UserHomeDir()

	m := Model{
		filepicker: fp,
	}

	p := tea.NewProgram(
		&m,
		tea.WithOutput(os.Stdout),
	)
	// var tm tea.Model
	var err error
	if _, err = p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	// mm := tm.(Model)
	// fmt.Println("\n  You selected: " + m.filepicker.Styles.Selected.Render(mm.selectedFile) + "\n")
}
