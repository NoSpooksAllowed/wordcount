package input

func ReadInputFromCmd(args []string) string {
	if len(args) < 2 {
		return ""
	}

	return args[1]
}
