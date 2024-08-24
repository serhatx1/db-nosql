package DB

type Table struct {
	Name      string
	Fields    map[string]string
	Variables map[string]interface{}
}
