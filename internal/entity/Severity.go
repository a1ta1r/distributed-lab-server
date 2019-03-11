package entity

const (
	SEVERITY_MINOR = 1 + iota
	SEVERITY_MAJOR
	SEVERITY_CRITICAL
)

type Severity int

var Severities = [...]string {
	"Minor",
	"Major",
	"Critical",
}

func (s Severity) String() string {
	return Severities[s - 1]
}