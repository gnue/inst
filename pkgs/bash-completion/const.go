package bash_completion

var Locals = []string{"$BASH_COMPLETION_DIR"}
var Globals = []string{"$BASH_COMPLETION_COMPAT_DIR"}

var Template = `_{{.}}() {
	args=("${COMP_WORDS[@]:1:$COMP_CWORD}")

	local IFS=$'\n'
	COMPREPLY=($(GO_FLAGS_COMPLETION=1 ${COMP_WORDS[0]} "${args[@]}"))
	return 1
}

complete -F _{{.}} {{.}}
`
