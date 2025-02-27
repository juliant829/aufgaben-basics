package strings

// Erwartet zwei Strings s und t und prüft, ob t in s als Teilstring vorkommt.
func ContainsSubstring(s, t string) bool {

	for i := 0; i <= len(s)-len(t); i++ {
		if s[i:i+len(t)] == t {
			return true
		}
	}
	return false
}
