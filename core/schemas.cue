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

ruleBlocks: {
	indirectDeps: #PackageRule & {
		matchDepTypes: ["indirect"]
		enabled: true
		matchManagers: ["gomod"]
	}
	automerge: #PackageRule & {
		automerge: true
	}
	mergeCommit: #PackageRule & {
		automergeStrategy: "merge-commit"
	}
	noTests: #PackageRule & {
		ignoreTests:   true
		automergeType: "branch"
	}
	allDeps: #PackageRule & {
		matchDepTypes: ["*"]
	}
	recreate: #PackageRule & {
		recreateWhen: "always"
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

let commonPatterns = {
	withGoPost: postUpdateOptions: goPostUpdateOptions
}

cat: #RenovateConfig & bestPracticesBase & {
	packageRules: [
		ruleBlocks.automerge &
		ruleBlocks.mergeCommit &
		ruleBlocks.noTests &
		ruleBlocks.allDeps & {
			rangeStrategy: "pin"
		},
	]
}

owl: #RenovateConfig & bestPracticesBase & commonPatterns.withGoPost & {
	packageRules: [
		ruleBlocks.automerge &
		ruleBlocks.mergeCommit &
		ruleBlocks.allDeps &
		ruleBlocks.recreate,
	]
	ignorePaths: ["**/testdata/go.mod"]
}

monkey: #RenovateConfig & bestPracticesBase & commonPatterns.withGoPost & {
	packageRules: [
		ruleBlocks.indirectDeps,
		ruleBlocks.automerge &
		ruleBlocks.mergeCommit &
		ruleBlocks.noTests &
		ruleBlocks.allDeps,
	]
}

hamster: #RenovateConfig & bestPracticesBase & commonPatterns.withGoPost & {
	packageRules: [
		ruleBlocks.automerge &
		ruleBlocks.mergeCommit &
		ruleBlocks.allDeps &
		ruleBlocks.recreate,
	]
}

rabbit: #RenovateConfig & bestPracticesBase & commonPatterns.withGoPost & {
	packageRules: [
		ruleBlocks.automerge &
		ruleBlocks.mergeCommit &
		ruleBlocks.allDeps &
		ruleBlocks.recreate,
	]
}
