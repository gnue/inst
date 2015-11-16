package bash_completion

const Name = "bash"
const InstallPath = "$BASH_COMPLETION_DIR"

var Template = `_{{.}}() {
	args=("${COMP_WORDS[@]:1:$COMP_CWORD}")

	local IFS=$'\n'
	COMPREPLY=($(GO_FLAGS_COMPLETION=1 ${COMP_WORDS[0]} "${args[@]}"))
	return 1
}

complete -F _{{.}} {{.}}
`
