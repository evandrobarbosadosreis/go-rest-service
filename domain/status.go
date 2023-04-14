package domain

type Status string

const (
	active   = "active"
	inactive = "inactive"
)

func (s Status) GetValue() string {
	if s == active {
		return "1"
	}
	if s == inactive {
		return "0"
	}
	return ""
}

func (s Status) IsValid() bool {
	return string(s) == active || string(s) == inactive
}
