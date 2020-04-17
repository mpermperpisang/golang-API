package data

/*Field : field biodata*/
type Field struct {
	ID          string `json:"ID"`
	Name        string `json:"Name"`
	Job         string `json:"Job"`
	Description string `json:"Description"`
	Motto       string `json:"Motto"`
}

type biodata []Field

/*Datas : biodata content*/
var Datas = biodata{
	{
		ID:          "1",
		Name:        "Ferawati Hartanti Pratiwi (MperMperPisang)",
		Job:         "Test Engineer",
		Description: "mpermperpisang pisangku belum matang",
		Motto:       "Pelembut pakaian",
	},
}
