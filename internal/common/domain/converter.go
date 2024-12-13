package domain

func ToSafeStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func ToSafeFloat64(i *float64) float64 {
	if i == nil {
		return 0
	}
	return *i
}
