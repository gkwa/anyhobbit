package renovate

#UpdateType: "minor" | "patch" | "pin" | "digest" | "replacement"
#DepType:    "*" | "indirect"
#Manager:    "gomod"
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
	withReplacement: standard + ["replacement"]
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
// preset: auto-merges all dependency types with merge type pr to notify us that merge has happened
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
