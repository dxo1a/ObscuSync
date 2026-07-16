package config

var registry = map[string]Game{
	"factorio": {
		ID:   "factorio",
		Name: "Factorio",
		ScanFolders: []string{
			"mods",
			"config",
		},
	},

	"minecraft": {
		ID:   "minecraft",
		Name: "Minecraft Java",
		ScanFolders: []string{
			"mods",
			"config",
			"resourcepacks",
			"shaderpacks",
		},
	},
}

func GetGame(id string) (Game, bool) {
	game, ok := registry[id]
	return game, ok
}
