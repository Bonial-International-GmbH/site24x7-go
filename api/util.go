package api

// String returns a pointer to string s.
func String(s string) *string {
	return &s
}

// Bool returns a pointer to bool b.
func Bool(b bool) *bool {
	return &b
}
