package enumerations

type Roles int

const (
	NORMAL Roles = iota
	ADMIN
	SUPERVISOR
)

func (e *Roles) Scan() error {
	return nil
}
