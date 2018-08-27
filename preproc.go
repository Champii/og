package main

import "strings"

func indentCount(str string) int {
	for i := range str {
		if str[i] != ' ' {
			return i
		}
	}

	return 0
}

func Preproc(str string) string {
	rawLines := strings.Split(str, "\n")

	var lines []string
	for _, v := range rawLines {
		if len(v) > 0 {
			lines = append(lines, v)
		}
	}

	var res []string

	lastIndent := 0
	indentSize := 0

	for i := range lines {
		indent := indentCount(lines[i])

		if indentSize == 0 && indent != lastIndent {
			indentSize = indent
		}

		if indent > lastIndent {
			res[len(res)-1] = res[len(res)-1] + "{"
		} else if indent < lastIndent {
			indentBuff := lastIndent - indent

			for indentBuff > 0 {
				res = append(res, "}")
				indentBuff -= indentSize
			}
		}

		lastIndent = indent

		res = append(res, lines[i])
	}

	for lastIndent > 0 {
		res = append(res, "}")
		lastIndent -= indentSize
	}

	joined := strings.Join(res, "\n") + "\n"

	return joined
}
