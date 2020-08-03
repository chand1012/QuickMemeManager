package main

func stringInSlice(s string, a []string) bool {
	for _, thing := range a {
		if thing == s {
			return true
		}
	}
	return false
}
