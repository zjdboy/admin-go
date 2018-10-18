package menu

import (
	"admin-go/db"
	"admin-go/defs"
	"log"
)

func ListData() []defs.Menu {

	rows, err := db.MySQL.Query("select id, icon, menu, url, pid, level from s_menu")

	if err != nil {
		log.Fatalf("query menus error, error message are :%v", err.Error())
		return nil
	}
	defer rows.Close()

	var menus = make([]defs.Menu, 0)

	for rows.Next() {
		var menu defs.Menu

		if err := rows.Scan(&menu.ID, &menu.Icon, &menu.Menu, &menu.Url, &menu.Pid, &menu.Level); err != nil {
			log.Fatalf("scan menus val error %v", err.Error())
			return nil
		}
		menus = append(menus, menu)
	}

	return menus
}
