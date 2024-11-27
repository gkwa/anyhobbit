// Code generated by gen; DO NOT EDIT.

package cmd

var commands = map[string]struct {
	short string
	long  string
}{
	"bunny": {
		short: "Generate Renovate configuration using bunny preset",
		long:  "Generate Renovate configuration using bunny preset, which same as hare but with prNotPendingHours 1.",
	},
	"chimp": {
		short: "Generate Renovate configuration using chimp preset",
		long:  "Generate Renovate configuration using chimp preset, which same as gorilla but with prNotPendingHours 1.",
	},
	"eagle": {
		short: "Generate Renovate configuration using eagle preset",
		long:  "Generate Renovate configuration using eagle preset, which same as owl but with prCreation not-pending.",
	},
	"gerbil": {
		short: "Generate Renovate configuration using gerbil preset",
		long:  "Generate Renovate configuration using gerbil preset, which same as hamster but with prCreation not-pending.",
	},
	"gorilla": {
		short: "Generate Renovate configuration using gorilla preset",
		long:  "Generate Renovate configuration using gorilla preset, which same as monkey but with prCreation not-pending.",
	},
	"guinea": {
		short: "Generate Renovate configuration using guinea preset",
		long:  "Generate Renovate configuration using guinea preset, which same as gerbil but with prNotPendingHours 1.",
	},
	"hamster": {
		short: "Generate Renovate configuration using hamster preset",
		long:  "Generate Renovate configuration using hamster preset, which auto-merges all dependency types and recreates PRs without filtering update types or post-update options.",
	},
	"hare": {
		short: "Generate Renovate configuration using hare preset",
		long:  "Generate Renovate configuration using hare preset, which same as rabbit but with prCreation not-pending.",
	},
	"hawk": {
		short: "Generate Renovate configuration using hawk preset",
		long:  "Generate Renovate configuration using hawk preset, which same as eagle but with prNotPendingHours 1.",
	},
	"koala": {
		short: "Generate Renovate configuration using koala preset",
		long:  "Generate Renovate configuration using koala preset, which auto-merges dependencies with pin strategy and npm/pnpm dedupe options.",
	},
	"lion": {
		short: "Generate Renovate configuration using lion preset",
		long:  "Generate Renovate configuration using lion preset, which auto-merges all dependency types with merge type pr and ignore tests since we don't have any.",
	},
	"monkey": {
		short: "Generate Renovate configuration using monkey preset",
		long:  "Generate Renovate configuration using monkey preset, which auto-merges all updates including indirect dependencies.",
	},
	"mouse": {
		short: "Generate Renovate configuration using mouse preset",
		long:  "Generate Renovate configuration using mouse preset, which same as rat but with prCreation not-pending.",
	},
	"owl": {
		short: "Generate Renovate configuration using owl preset",
		long:  "Generate Renovate configuration using owl preset, which auto-merges and recreates PRs for all update types including replacements.",
	},
	"panda": {
		short: "Generate Renovate configuration using panda preset",
		long:  "Generate Renovate configuration using panda preset, which uses pin range strategy with auto-merging and recreation policies.",
	},
	"penguin": {
		short: "Generate Renovate configuration using penguin preset",
		long:  "Generate Renovate configuration using penguin preset, which auto-merges all dependency types with merge type branch to prevent noisey pull request emails.",
	},
	"rabbit": {
		short: "Generate Renovate configuration using rabbit preset",
		long:  "Generate Renovate configuration using rabbit preset, which auto-merges all dependency types and recreates PRs without filtering update types.",
	},
	"rat": {
		short: "Generate Renovate configuration using rat preset",
		long:  "Generate Renovate configuration using rat preset, which focuses on auto-merging standard updates with no automated testing.",
	},
	"shrew": {
		short: "Generate Renovate configuration using shrew preset",
		long:  "Generate Renovate configuration using shrew preset, which same as mouse but with prNotPendingHours 1.",
	},
	"tiger": {
		short: "Generate Renovate configuration using tiger preset",
		long:  "Generate Renovate configuration using tiger preset, which auto-merges all dependency types with merge type branch to reduce pull request noise.",
	},
}
