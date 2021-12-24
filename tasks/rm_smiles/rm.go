package rm

import "strings"

func RmSmiles(in string) string {
	if len(in) < 3 {
		return in
	}
	sb := strings.Builder{}

	left, startSmile := 0, len(in)

	for i := 0; i < len(in); i++ {
		if in[i] == '(' || in[i] == ')' {
			if i < 2 {
				continue
			}
			if in[i-1] == '-' && in[i-2] == ':' {
				startSmile = i - 2
				sb.WriteString(in[left:startSmile])
				for i < len(in) && in[i] == in[startSmile+2] {
					i++
				}

				left = i
				i--
			}
		}
	}
	sb.WriteString(in[left:])

	return sb.String()
}

func RmSmiles_splits(in string) string {
	if len(in) < 3 {
		return in
	}
	sb := strings.Builder{}

	splitted := strings.Split(in, ":-")
	sb.WriteString(splitted[0])

	for i := 1; i < len(splitted); i++ {
		if len(splitted[i]) == 0 {
			continue
		}
		if splitted[i][0] == ')' || splitted[i][0] == '(' {
			start := 1
			for ; start < len(splitted[i]); start++ {
				if splitted[i][start] != splitted[i][0] {
					break
				}
			}
			sb.WriteString(splitted[i][start:])
		} else {
			sb.WriteString(":-")
			sb.WriteString(splitted[i])
		}
	}
	return sb.String()
}
