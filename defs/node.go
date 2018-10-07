package defs

type Node struct {
	ID       int    `json:"id, val"`
	Icon     string `json:"icon, val"`
	Text     string `json:"text, val"`
	Url      string `json:"url, val"`
	Level    string `json:"level, val"`
	ParentId int    `json:"parentId, val"`
	Children []Node `json:"children, val"`
}
