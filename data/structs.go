package data

/*Field : field biodata*/
type Field struct {
	ID          int    `json:ID`
	Name        string `json:"Name"`
	Job         string `json:"Job"`
	Description string `json:"Description"`
	Motto       string `json:"Motto"`
}

type biodata []Field
