package cmd

var commands = map[string]struct {
	short string
	long  string
}{
	"rat": {
		short: "Generate Renovate configuration using rat preset",
		long:  "Generate Renovate configuration using rat preset, which focuses on auto-merging standard updates with no automated testing.",
	},
	"monkey": {
		short: "Generate Renovate configuration using monkey preset",
		long:  "Generate Renovate configuration using monkey preset, which auto-merges all updates including indirect dependencies.",
	},
	"owl": {
		short: "Generate Renovate configuration using owl preset",
		long:  "Generate Renovate configuration using owl preset, which auto-merges and recreates PRs for all update types including replacements.",
	},
	"rabbit": {
		short: "Generate Renovate configuration using rabbit preset",
		long:  "Generate Renovate configuration using rabbit preset, which auto-merges all dependency types and recreates PRs without filtering update types.",
	},
	"penguin": {
		short: "Generate Renovate configuration using penguin preset",
		long:  "Generate Renovate configuration using penguin preset, which auto-merges all dependency types with merge type pr to notify us that merge has happend.",
	},
	"tiger": {
		short: "Generate Renovate configuration using tiger preset",
		long:  "Generate Renovate configuration using tiger preset, which auto-merges all dependency types with merge type branch to reduce pull request noise.",
	},
	"panda": {
		short: "Generate Renovate configuration using panda preset",
		long:  "Generate Renovate configuration using panda preset, which uses pin range strategy with auto-merging and recreation policies.",
	},
}
