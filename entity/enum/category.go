package enum

type Category string

const (
	OTHER    Category = "OTHER"
	SALE     Category = "SALE"
	PURCHASE Category = "PURCHASE"
)

func ValidateCategory(value string) bool {
	switch Category(value) {
	case OTHER, SALE, PURCHASE:
		return true
	}
	return false
}
