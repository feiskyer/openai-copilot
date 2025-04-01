package tools

import (
	"errors"
	"os/exec"
	"strings"
)

// Kubectl runs the given kubectl command and returns the output.
func Kubectl(command string) (string, error) {
	if strings.HasPrefix(command, "kubectl") {
		command = strings.TrimSpace(strings.TrimPrefix(command, "kubectl"))
	}

	if strings.HasPrefix(command, "edit") {
		return "", errors.New("interactive command kubectl edit is not supported")
	}

	args := parseCommandWithQuotes(command)
	cmd := exec.Command("kubectl", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return strings.TrimSpace(string(output)), err
	}

	return strings.TrimSpace(string(output)), nil
}

// parseCommandWithQuotes splits a command string into arguments, respecting quoted strings
func parseCommandWithQuotes(command string) []string {
	var args []string
	var currentArg strings.Builder
	inQuotes := false
	quoteChar := rune(0)

	for _, ch := range command {
		switch {
		case ch == '\'' || ch == '"':
			if inQuotes && ch == quoteChar {
				// End of quoted section
				inQuotes = false
				quoteChar = rune(0)
			} else if !inQuotes {
				// Start of quoted section
				inQuotes = true
				quoteChar = ch
			} else {
				// We're in quotes but found a different quote character, add it
				currentArg.WriteRune(ch)
			}
		case ch == ' ' && !inQuotes:
			// Space outside quotes - end of argument
			if currentArg.Len() > 0 {
				args = append(args, currentArg.String())
				currentArg.Reset()
			}
		default:
			// Add character to current argument
			currentArg.WriteRune(ch)
		}
	}

	// Add the last argument if there is one
	if currentArg.Len() > 0 {
		args = append(args, currentArg.String())
	}

	return args
}
