package DB

type DB struct {
	Name   string
	Tables map[string]*Table
}
