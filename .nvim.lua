local project_config = {
	lint = {
	},
	branches = {
		main = "develop",
		prod = "main",
	},
	debug = {
		launch_file = ".nvimlaunch.json",
	},
	autocmds = {
		{
			event = "BufWritePre",
			pattern = "*.go",
			command = "!procmux signal-start --name run-day",
		},
	},
	commands = {
		{
			command = "procmux signal-start --name run-day",
			description = "procmux signal start run day",
		},
	},
}

local setup = function() end

_G.EXRC_M = {
	project_config = project_config,
	setup = setup,
}
