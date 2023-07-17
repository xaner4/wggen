package wggen

func stringInSlice(s []string, v string) bool {
	for _, i := range s {
		if i == v {
			return true
		}
	}
	return false
}
