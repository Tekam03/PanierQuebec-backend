package util

func FromPtr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func ToPtr(s string) *string {
	return &s
}