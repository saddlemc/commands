package commands

import "github.com/df-mc/dragonfly/server/cmd"

// CommandInfo is a config entry for a command. Here it can be customized in order to fit the rest of the server.
type CommandInfo struct {
	// Runnables is not an actual configuration entry. It is used for the plugin to associate a certain command with a
	// CommandInfo instance.
	Runnables []cmd.Runnable `toml:"-"`

	// Disable will make the plugin not register the command on startup if set to true.
	Disable bool `toml:"disable"`
	// Name is what the player types in to execute the command, without the slash. For example, /give is named "give".
	Name string `toml:"name"`
	// Description is a short message that explains what the command does. Will show up to a user when typing in a
	// command name.
	Description string `toml:"description"`
	// Aliases are different ways to type a command. They are alternative names for the command. For example, /gamemode
	// has a /gm alias. Once again, the slash should not be part of the string.
	Aliases []string `toml:"aliases"`
}
