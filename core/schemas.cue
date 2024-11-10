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

rat: #RenovateConfig & bestPracticesBase & {
	lockFileMaintenance: enabled: true
	packageRules: [
		ruleBlocks.noTests & {
			rangeStrategy: "pin"
		},
	]
}

owl: #RenovateConfig & bestPracticesBase & commonPatterns.withGoPost & {
	packageRules: [commonRuleFields]
	ignorePaths: ["**/testdata/go.mod"]
}

monkey: #RenovateConfig & bestPracticesBase & commonPatterns.withGoPost & {
	packageRules: [
		ruleBlocks.indirectDeps,
		commonRuleFields & ruleBlocks.noTests,
	]
}

rabbit: #RenovateConfig & bestPracticesBase & commonPatterns.withGoPost & {
	packageRules: [commonRuleFields]
}

penguin: #RenovateConfig & bestPracticesBase & commonPatterns.withGoPost & {
	packageRules: [
		commonRuleFields & {
			automergeType: "branch"
			ignoreTests:   true
		},
	]
}

tiger: #RenovateConfig & bestPracticesBase & commonPatterns.withGoPost & {
	packageRules: [
		commonRuleFields & {
			automergeType: "branch"
		},
	]
}

panda: #RenovateConfig & bestPracticesBase & {
	lockFileMaintenance: enabled: true
	packageRules: [
		commonRuleFields & {
			rangeStrategy: "pin"
			recreateWhen:  "always"
		},
	]
}

koala: #RenovateConfig & bestPracticesBase & commonPatterns.withNpmPost & {
	packageRules: [
		commonRuleFields & {
			rangeStrategy: "pin"
			recreateWhen:  "always"
		},
	]
}
