package api

import (
	"fmt"
	"net/url"
	"strings"
)

// MustParams ...
func MustParams(params []string, values url.Values) error {
	missing := make([]string, 0)
	for _, param := range params {
		if len(values.Get(param)) == 0 {
			missing = append(missing, param)
		}
	}

	if len(missing) > 0 {
		return fmt.Errorf("Mandatory parameters missing: %s", strings.Join(missing, ", "))
	}

	return nil
}
