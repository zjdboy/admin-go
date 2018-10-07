package defs

type Menu struct {
	ID    int    `json:"id, val"`
	Icon  string `json:"icon, val"`
	Menu  string `json:"menu, val"`
	Url   string `json:"url, val"`
	Pid   int    `json:"pid, val"`
	Level int    `json:"level, val"`
}
