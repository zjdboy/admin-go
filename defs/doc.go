package defs

type Doc struct {
	ID      int    `json:"id, val"`
	Title   string `json:"title, val"`
	Type    string `json:"type, val"`
	Cdate   string `json:"cdate, val"`
	Edate   string `json:"edate, val"`
	Status  int    `json:"status, val"`
	Content string `json:"content, val"`
	Imgs    string `json:"imgs, val"`
}
