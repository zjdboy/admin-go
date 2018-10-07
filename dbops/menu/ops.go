package menu

import (
	"admin-go/db"
	"admin-go/defs"
	"log"
)

func ListData() []*defs.Menu {

	rows, err := db.MySQL.Query("select id, icon, menu, url, pid, level from s_menu")

	if err != nil {
		log.Printf("query menus error, error message are :%v", err)
		return nil
	}

	var menus []*defs.Menu

	for rows.Next() {
		var id, level, pid int
		var icon, menu, url string

		if err := rows.Scan(&id, &icon, &menu, &url, &pid, &level); err != nil {
			log.Printf("scan menus val error %v", err)
			return nil
		}
		one := &defs.Menu{ID: id, Icon: icon, Menu: menu, Url: url, Pid: pid, Level: level}
		menus = append(menus, one)
	}
	defer rows.Close()
	return menus
}
