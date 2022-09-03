package commands

import (
	"context"
	"github.com/df-mc/dragonfly/server/cmd"
	commands "github.com/saddlemc/commands/cmd"
	"github.com/saddlemc/saddle/plugin"
)

type Plugin struct{}

func (p Plugin) Name() string { return "saddle-commands" }

func (p Plugin) Setup(this *plugin.Plugin) error {
	// Create each command entry. These are passed along to a config, which will either write this data as the default
	// commands.toml file or, if it already exists, it will read the data from the file into this map. It is then used
	// to register the commands.
	// todo: commands should have an associated permission when a permission system is added to saddle
	commandConfigs := map[string]CommandInfo{
		"gamemode": {
			Runnables:   []cmd.Runnable{commands.Gamemode{}},
			Name:        "gamemode",
			Description: "change someone's gamemode",
			Aliases:     []string{"gm"},
		},
		"give": {
			Runnables:   []cmd.Runnable{commands.Give{}},
			Name:        "give",
			Description: "give an item to a player",
		},
		"kick": {
			Runnables:   []cmd.Runnable{commands.Kick{}},
			Name:        "kick",
			Description: "disconnect a player",
		},
		"kill": {
			Runnables:   []cmd.Runnable{commands.Kill{}},
			Name:        "kill",
			Description: "kill a player",
		},
		"teleport": {
			Runnables: []cmd.Runnable{
				commands.TeleportToCoords{},
				commands.TeleportToPlayer{},
				commands.TeleportPlayerToCoords{},
				commands.TeleportPlayerToPlayer{},
			},
			Name:        "teleport",
			Description: "teleport yourself or a player to a destination",
			Aliases:     []string{"tp"},
		},
		"weather": {
			Runnables:   []cmd.Runnable{commands.Weather{}},
			Name:        "weather",
			Description: "change the weather",
		},
	}
	// Saddle's built-in config system provides a simple API for loading basic configuration files. This will
	// automatically create a config.toml file in the plugins/saddle-commands directory, and handle everything for us.
	err := this.WithConfigs(plugin.Config{
		Path:  "commands.toml",
		Value: &commandConfigs,
	})
	if err != nil {
		return err
	}

	i := 0
	for _, entry := range commandConfigs {
		if entry.Disable {
			continue
		}
		cmd.Register(cmd.New(entry.Name, entry.Description, entry.Aliases, entry.Runnables...))
		i++
	}
	this.Logger().Info().Msgf("Registered %d command(s).", i)
	return nil
}

func (p Plugin) Run(ctx context.Context, this *plugin.Plugin) {
	// Make the plugin run until the server closes. We don't actually need to do anything here for this specific plugin.
	for range ctx.Done() {
	}
}
