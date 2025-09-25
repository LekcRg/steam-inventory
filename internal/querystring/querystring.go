package querystring

import (
	"fmt"
	"net/url"
	"strings"
)

func BuildQuery(m map[string]string) string {
	strSeparator := "&"
	i := 1

	var builder strings.Builder

	for key, val := range m {
		if i == len(m) {
			strSeparator = ""
		}

		i++

		fmt.Fprintf(
			&builder, "%s=%s%s",
			key, url.QueryEscape(val),
			strSeparator,
		)
	}

	return builder.String()
}
