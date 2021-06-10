package discord

import (
	"errors"
	"fmt"
	"strings"
)

// NewCommand creates a new command and returns it!
func NewCommand(sig string, desc string, handler newCommandHandlerFunc) *Command {
	cmd := &Command{}
	cmd.Signature = sig
	cmd.Description = desc
	cmd.Handler = handler
	return cmd
}

// FormatParamString removes certain characters like "" from the param response string.
func FormatParamString(pr string) string {
	f := func(r rune) bool {
		switch {
		case r == '"':
			return true
		default:
			return false
		}
	}

	return strings.TrimFunc(pr, f)
}

// ParseParams gets a slice of arguments in a string.
// 		Goes through each rune and determines what to do with it.
//		If the rune is in Quotes "something something", it will return the whole string in quotes.
func ParseParams(content string) []string {
	inQuotes := false
	f := func(r rune) bool {
		switch {
		case r == '"':
			inQuotes = !inQuotes
			return false
		case inQuotes:
			return false
		default:
			return r == ' '
		}
	}

	return strings.FieldsFunc(content, f)
}

// SignatureWithPrefix returns a concatenated string with prefix and signature
func (c *Command) SignatureWithPrefix() string {
	return fmt.Sprintf("%s%s", c.Prefix, c.Signature)
}

// AddParam adds a param to a command.
func (c *Command) AddParam(param *CommandParameter) {
	c.Parameters = append(c.Parameters, param)
}

// AddParameter adds a new parameter to the command.
// Parameters should be added in order.
func (c *Command) AddParameter(name string, description string, required bool) {
	param := &CommandParameter{
		Name:        name,
		Description: description,
		Position:    len(c.Parameters) + 1,
		Required:    required,
	}
	c.Parameters = append(c.Parameters, param)
}

func (c *Command) HelpUsage() string {
	return fmt.Sprintf("%shelp %s", c.Prefix, c.Signature)
}

// HelpString returns a string of how this Command is used.
func (c *Command) HelpString() string {
	parts := make([]string, 0)
	for _, p := range c.Parameters {
		if !p.Required {
			parts = append(parts, fmt.Sprintf("<%s:optional>", p.Name))
		} else {
			parts = append(parts, fmt.Sprintf("<%s:required>", p.Name))
		}
	}
	joined := strings.Join(parts, " ")
	return fmt.Sprintf("%s %s", c.SignatureWithPrefix(), joined)
}

// GetParam : parses the params and gets the correct value for the param with the name given.
// TODO: change this so it can be used like:
//		p := c.GetParam(content, "paramName")
//		if !p {
//			did not find the param.
//		}
func (c *Command) GetParam(content string, name string) (pr string, err error) {
	trimmedContent := strings.TrimPrefix(content, c.SignatureWithPrefix())

	pr = ""
	param := &CommandParameter{}
	foundParam := false
	foundMatchingParam := false
	for _, p := range c.Parameters {
		if p.Name == name {
			param = p
			foundParam = true
			break
		}
	}

	if len(c.Parameters) <= 0 {
		fmt.Println("command has no parameters")
		err = errors.New("command has no parameters set")
		return
	}

	if !foundParam {
		fmt.Printf("couldn't find the param named '%s' in the command '%s'", name, c.Signature)
		err = errors.New("could not find param")
		return
	}

	parsed := ParseParams(trimmedContent)

	if len(parsed) > len(c.Parameters) {
		err = errors.New("too many parameters passed")
		return
	}

	for pos, s := range parsed {
		if pos == param.Position-1 {
			pr = FormatParamString(s)
			err = nil
			foundMatchingParam = true
			break
		}
	}

	if param.Required && !foundMatchingParam {
		helpUsage := c.HelpUsage()
		// this error message really shouldnt be directed towards the user, but eh.. this is the only way I could really figure out how for now.
		err = fmt.Errorf("could not find required parameter `%s`. To see required parameters, check `%s`", param.Name, helpUsage)
		return
	}

	return
}

// GetRequiredParam
// TODO actually not really sure where I was going with this. Something about sending a response back to the user..
func (c *Command) GetRequiredParam(content string, name string) (pr string, err error) {
	pr, err = c.GetParam(content, name)
	if err != nil {

	}

	return
}
