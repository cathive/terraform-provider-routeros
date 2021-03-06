package routeros

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func mustAtoi(s string) int {
	if i, err := strconv.Atoi(s); err != nil {
		panic(fmt.Errorf(`unable to convert "%s" to numeric value: %v`, s, err))
	} else {
		return i
	}
}

func mustParseBool(s string) bool {
	if b, err := strconv.ParseBool(s); err != nil {
		panic(fmt.Errorf(`unable to convert "%s" to boolean value: %v`, s, err))
	} else {
		return b
	}
}

func stringToBoolMap(ps string) map[string]bool {
	policies := make(map[string]bool)
	s := strings.Split(ps, ",")
	for _, policy := range s {
		if policy[:1] == "!" {
			policies[policy[1:]] = false
		} else {
			policies[policy] = true
		}
	}
	return policies
}

func boolMapToString(m map[string] bool) string {
	lst := make([]string, len(m))
	idx := 0
	for key, val := range m {
		if val {
			lst[idx] = key
		} else {
			lst[idx] = fmt.Sprintf("!%s", key)
		}
		idx++
	}

	// We sort the list to make sure that the output is predictable/reproducible.
	// TODO The sort order should match the RouterOS API sort order.
	sort.Strings(lst)

	return strings.Join(lst, ",")
}
