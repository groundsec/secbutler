package utils

func BooleanColorCode(boolValue bool) string {
	if boolValue {
		return ANSICodes["Green"]
	} else {
		return ANSICodes["Red"]
	}
}
