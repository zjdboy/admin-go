package defs

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Username       string `json:"username"`
	jwt.StandardClaims
}


type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

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

type Menu struct {
	ID    int    `json:"id, val"`
	Icon  string `json:"icon, val"`
	Menu  string `json:"menu, val"`
	Url   string `json:"url, val"`
	Pid   int    `json:"pid, val"`
	Level int    `json:"level, val"`
}

type Type struct {
	ID      int    `json:"id, val"`
	Type    string `json:"type, val"`
	Content string `json:"content, val"`
}

type Node struct {
	ID       int    `json:"id, val"`
	Icon     string `json:"icon, val"`
	Text     string `json:"text, val"`
	Url      string `json:"url, val"`
	Level    string `json:"level, val"`
	ParentId int    `json:"parentId, val"`
	Children []Node `json:"children, val"`
}

