package cmdline

import "testing"

func TestConfig(t *testing.T) {
	if len(commands) != len(HELP_OPTION.commands) {
		t.Error("Some command does not have a help option.")
	}
}

func TestOptionsComplete(t *testing.T) {
	for _, opt := range options {
		if opt.short == "" || opt.long == "" || opt.description == "" || opt.commands == nil || (opt.hasParam && opt.isParamValid == nil) {
			t.Error("Incomplete option", opt.short, opt.long)
		}
	}
}

func TestCommandsComplete(t *testing.T) {
	for _, cmd := range commands {
		if cmd.name == "" || cmd.description == "" || cmd.usage == "" {
			t.Error("Incomplete command", cmd.name)
		}
	}
}

func TestDuplicateOption(t *testing.T) {
	for _, cmd := range commands {
		foundLongOpts := make(map[string]string)
		foundShortOpts := make(map[string]string)
		for _, opt := range options {
			if opt.supportsCommand(cmd) {
				_, exists := foundLongOpts[opt.long]
				if exists {
					t.Error("Duplicate option", opt.long, "for", cmd.name, "command")
				}
				_, exists = foundShortOpts[opt.short]
				if exists {
					t.Error("Duplicate option", opt.short, "for", cmd.name, "command")
				}
				foundLongOpts[opt.long] = "found"
				foundShortOpts[opt.short] = "found"
			}
		}
	}
}
