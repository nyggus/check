package check

// Any checks if any of the conditions is met. For no conditions, it returns false.
func Any(Conditions []bool) bool {
	if len(Conditions) == 0 {
		return false
	}
	for _, condition := range Conditions {
		if condition {
			return true
		}
	}
	return false
}

// AnyInMapInt checks if any of the map's values is true. For no conditions, it returns false.
func AnyInMapInt(Conditions map[int]bool) bool {
	if len(Conditions) == 0 {
		return false
	}
	for _, condition := range Conditions {
		if condition {
			return true
		}
	}
	return false
}

// WhichInMapInt checks which of the map's values is true.
// Returns a tuple of true/false, keys, where keys are the keys with true.
func WhichInMapInt(Conditions map[int]bool) ([]int, bool) {
	var exists bool
	if len(Conditions) == 0 {
		return []int{}, false
	}
	keys := make([]int, 0)
	for key, condition := range Conditions {
		if condition {
			exists = true
			keys = append(keys, key)
		}
	}
	return keys, exists
}

// WhichInMapString checks which of the map's values is true.
// Returns a tuple of true/false, keys, where keys are the keys with true.
func WhichInMapString(Conditions map[string]bool) ([]string, bool) {
	var exists bool
	if len(Conditions) == 0 {
		return []string{}, false
	}
	keys := make([]string, 0)
	for key, condition := range Conditions {
		if condition {
			exists = true
			keys = append(keys, key)
		}
	}
	return keys, exists
}

// AnyInMapString checks if any of the map's values is true.
// Returns a tuple of true/false, keys, where keys are the keys with true.
func AnyInMapString(Conditions map[string]bool) bool {
	if len(Conditions) == 0 {
		return false
	}
	for _, condition := range Conditions {
		if condition {
			return true
		}
	}
	return false
}

// All checks if all conditions are met. For no conditions, it returns false.
func All(Conditions []bool) bool {
	if len(Conditions) == 0 {
		return false
	}
	for _, condition := range Conditions {
		if !condition {
			return false
		}
	}
	return true
}

// AllMapInt checks if all values in the map are true. For no conditions, it returns false.
func AllInMapInt(Conditions map[int]bool) bool {
	if len(Conditions) == 0 {
		return false
	}
	for _, condition := range Conditions {
		if !condition {
			return false
		}
	}
	return true
}

// AllInMapString checks if all values in the map are true. For no conditions, it returns false.
func AllInMapString(Conditions map[string]bool) bool {
	if len(Conditions) == 0 {
		return false
	}
	for _, condition := range Conditions {
		if !condition {
			return false
		}
	}
	return true
}
