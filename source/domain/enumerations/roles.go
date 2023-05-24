package enumerations

type Roles int

const (
	NORMAL     = 10
	SUPERVISOR = 20
	ADMIN      = 30
	GOD        = 90
)

func (e *Roles) Scan() error {
	return nil
}
