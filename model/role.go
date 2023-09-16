package model

type Role int

const (
	MENBER Role = iota
	ADMIN   
)

func (r Role) String() string{
	return []string{"MENBER", "ADMIN",}[r]
}