/*
Copyright © 2024 Enrique Marín

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package form

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/emarifer/go-cli-bubbletea-todoapp/internal/models"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	Bold   = "\033[1m"
	Maroon = "\x1b[38:5:1m"
	Reset  = "\033[0m"
)

type errMsg error

type model struct {
	inputs  []textinput.Model
	focused int
	err     error
}

const (
	name = iota
	description
)

const (
	hotPink  = lipgloss.Color("#FF06B7")
	darkGray = lipgloss.Color("#767676")
)

var (
	inputStyle    = lipgloss.NewStyle().Foreground(hotPink)
	continueStyle = lipgloss.NewStyle().Foreground(darkGray)
	task          = models.Task{}
)

func initialModel() model {
	var inputs []textinput.Model = make([]textinput.Model, 2)

	inputs[name] = textinput.New()
	inputs[name].Placeholder = "task name (max. 20 chars)…"
	inputs[name].Focus()
	inputs[name].CharLimit = 20 // maximum number of characters for the `Name` field
	inputs[name].Width = 30
	inputs[name].Prompt = ""
	inputs[name].Validate = nameValidator

	inputs[description] = textinput.New()
	inputs[description].Placeholder = "task description…"
	inputs[description].Width = 50
	inputs[description].Prompt = ""
	inputs[description].Validate = descriptionValidator

	return model{
		inputs:  inputs,
		focused: 0,
		err:     nil,
	}
}

func nameValidator(value string) error {
	task.Name = value

	return nil
}

func descriptionValidator(value string) error {
	task.Description = value

	return nil
}

func (m model) valueTrimmer() {
	trimmedValue := strings.Trim(m.inputs[m.focused].Value(), " ")
	m.inputs[m.focused].SetValue(trimmedValue)
}

func (m model) checkMinLen() error {
	var err error

	c := m.inputs[m.focused]
	if len(c.Value()) <= 1 {
		if m.focused == name {
			err = fmt.Errorf(
				Maroon + Bold + "the `name` field must be at least 2 characters long" + Reset,
			)
		} else {
			err = fmt.Errorf(
				Maroon + Bold + "the `description` field must be at least 2 characters long" + Reset,
			)
		}
	}

	return err
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.valueTrimmer()
			if m.checkMinLen() != nil {
				m.err = m.checkMinLen()
				break
			}

			if m.focused == len(m.inputs)-1 {
				return m, tea.Quit
			}
			m.nextInput()

		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyShiftTab, tea.KeyCtrlP:
			m.prevInput()

		case tea.KeyTab, tea.KeyCtrlN:
			m.nextInput()
		}

		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf(
			`Add a task:
	
  %s
  %s
  
  %s
  %s
  
  %s
`,
			inputStyle.Width(30).Render("Task Name:"),
			m.inputs[name].View(),
			inputStyle.Width(30).Render("Task Description:"),
			m.inputs[description].View(),
			continueStyle.Render("Continue ->"),
		) + "\n" +
			m.err.Error()
	}

	return fmt.Sprintf(
		`Add a task:

  %s
  %s
  
  %s
  %s
  
  %s
`,
		inputStyle.Width(30).Render("Task Name:"),
		m.inputs[name].View(),
		inputStyle.Width(30).Render("Task Description:"),
		m.inputs[description].View(),
		continueStyle.Render("Continue ->"),
	) + "\n"
}

// nextInput focuses the next input field
func (m *model) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

// prevInput focuses the previous input field
func (m *model) prevInput() {
	m.focused--

	// wrap around
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}

func Create() models.Task {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		log.Printf("💥 error initializing form interface: %s", err)

		os.Exit(1)
	}

	return task
}

/*
golang bubble textinput validation example:

https://www.google.com/search?q=golang+bubble+textinput+validation+example&oq=golang+bubble+textinput+validati&aqs=chrome.2.69i57j33i10i160l4.16983j0j7&sourceid=chrome&ie=UTF-8

https://github.com/charmbracelet/bubbles/discussions/289
https://gist.github.com/bashbunni/fed91563900a9f6e20cde881fe68ac31

https://github.com/charmbracelet/bubbletea/tree/main/examples/credit-card-form
*/
