package enumerations

type Roles int

const (
	GOD Roles = iota
	ADMIN
	SUPERVISOR
	NORMAL
)

func (e *Roles) Scan() error {
	return nil
}
