package models

type LabMember struct {
	Id          string `db:"id"`
	Name        string `db:"name"`
	Picture     string `db:"picture"`
	Description string `db:"description"`
}

type LabAlumni struct {
	id       string `db:"id"`
	Name     string `db:"name"`
	Tenure   string `db:"tenure"`
	Position string `db:"position"`
}
