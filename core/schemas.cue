package renovate

import "list"

#UpdateType: "minor" | "patch" | "pin" | "digest" | "replacement"

#DepType: "*" | "indirect"

#Manager: "gomod"

#PackageRule: {
	rangeStrategy?: string
	ignoreTests?:   bool
	matchDepTypes?: [...#DepType]
	matchManagers?: [...#Manager]
	matchUpdateTypes?: [...#UpdateType]
	enabled?:           bool
	automerge?:         bool | *false
	automergeType?:     string
	automergeStrategy?: string
	recreateWhen?:      string
}

#RenovateConfig: {
	$schema: "https://docs.renovatebot.com/renovate-schema.json"
	extends?: [...string]
	lockFileMaintenance?: enabled: bool
	prHourlyLimit:     0
	prConcurrentLimit: 0
	packageRules: [...#PackageRule]
	ignorePaths?: [...string]
	postUpdateOptions?: [...string]
	platformAutomerge: true
	prCreation:        string | *"immediate"
}

let bestPracticesBase = {
	extends: [
		"config:best-practices",
		":dependencyDashboard",
	]
}

let commonRuleFields = {
	matchDepTypes: ["*"]
	automerge:         true
	automergeStrategy: "merge-commit"
	recreateWhen:      "always"
}

ruleBlocks: {
	indirectDeps: #PackageRule & {
		matchDepTypes: ["indirect"]
		enabled: true
		matchManagers: ["gomod"]
	}

	noTests: #PackageRule & commonRuleFields & {
		ignoreTests:   true
		automergeType: "branch"
	}
}

updateTypes: {
	standard: ["minor", "patch", "pin", "digest"]
	withReplacement: list.Concat([standard, ["replacement"]])
}

let goPostUpdateOptions = [
	"gomodTidyE",
	"gomodMassage",
	"gomodUpdateImportPaths",
]

let npmPostUpdateOptions = [
	"npmDedupe",
	"pnpmDedupe",
]

let commonPatterns = {
	withGoPost: postUpdateOptions:  goPostUpdateOptions
	withNpmPost: postUpdateOptions: npmPostUpdateOptions
}

// @animal
// preset: focuses on auto-merging standard updates with no automated testing
rat: #RenovateConfig & bestPracticesBase & {
	lockFileMaintenance: enabled: true
	packageRules: [
		ruleBlocks.noTests & {
			rangeStrategy: "pin"
		},
	]
}

// @animal
// preset: auto-merges and recreates PRs for all update types including replacements
owl: #RenovateConfig & bestPracticesBase & commonPatterns.withGoPost & {
	packageRules: [commonRuleFields]
	ignorePaths: ["**/testdata/go.mod"]
}

// @animal
// preset: auto-merges all updates including indirect dependencies
monkey: #RenovateConfig & bestPracticesBase & commonPatterns.withGoPost & {
	packageRules: [
		ruleBlocks.indirectDeps,
		commonRuleFields & ruleBlocks.noTests,
	]
}

// @animal
// preset: auto-merges all dependency types and recreates PRs without filtering update types
rabbit: #RenovateConfig & bestPracticesBase & commonPatterns.withGoPost & {
	packageRules: [commonRuleFields]
}

// @animal
// preset: auto-merges all dependency types and recreates PRs without filtering update types or post-update options
hamster: #RenovateConfig & bestPracticesBase & {
	packageRules: [commonRuleFields]
}

// @animal
// preset: auto-merges all dependency types with merge type branch to prevent noisey pull request emails
penguin: #RenovateConfig & bestPracticesBase & commonPatterns.withGoPost & {
	packageRules: [
		commonRuleFields & {
			automergeType: "branch"
			ignoreTests:   true
		},
	]
}

// @animal
// preset: auto-merges all dependency types with merge type branch to reduce pull request noise
tiger: #RenovateConfig & bestPracticesBase & commonPatterns.withGoPost & {
	packageRules: [
		commonRuleFields & {
			automergeType: "branch"
		},
	]
}

// @animal
// preset: uses pin range strategy with auto-merging and recreation policies
panda: #RenovateConfig & bestPracticesBase & {
	lockFileMaintenance: enabled: true
	packageRules: [
		commonRuleFields & {
			rangeStrategy: "pin"
			recreateWhen:  "always"
		},
	]
}

// @animal
// preset: auto-merges dependencies with pin strategy and npm/pnpm dedupe options
koala: #RenovateConfig & bestPracticesBase & commonPatterns.withNpmPost & {
	packageRules: [
		commonRuleFields & {
			rangeStrategy: "pin"
			recreateWhen:  "always"
		},
	]
}

// @animal
// preset: auto-merges all dependency types with merge type pr and ignore tests since we don't have any
lion: #RenovateConfig & bestPracticesBase & commonPatterns.withGoPost & {
	packageRules: [
		commonRuleFields & {
			ignoreTests: true
		},
	]
}

// @animal
// preset: same as rat but with prCreation not-pending
mouse: rat & {
	prCreation: "not-pending"
}

// @animal
// preset: same as owl but with prCreation not-pending
eagle: owl & {
	prCreation: "not-pending"
}

// @animal
// preset: same as monkey but with prCreation not-pending
gorilla: monkey & {
	prCreation: "not-pending"
}

// @animal
// preset: same as rabbit but with prCreation not-pending
hare: rabbit & {
	prCreation: "not-pending"
}

// @animal
// preset: same as hamster but with prCreation not-pending
gerbil: hamster & {
	prCreation: "not-pending"
}
