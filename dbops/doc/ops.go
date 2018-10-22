package doc

import (
	"admin-go/db"
	"admin-go/defs"
	"log"
)

func ListData() []defs.Doc {
	rows, err := db.MySQL.Query("select id, title, type, cdate, edate, status, content from i_doc")

	if err != nil {
		log.Printf("query docs error, error message are :%v", err.Error())
		return nil
	}
	defer rows.Close()

	var docs = make([]defs.Doc, 0)

	for rows.Next() {
		var doc defs.Doc

		if err := rows.Scan(&doc.ID, &doc.Title, &doc.Type, &doc.Cdate, &doc.Edate, &doc.Status, &doc.Content); err != nil {
			log.Printf("scan docs val error %v", err.Error())
			return nil
		}
		docs = append(docs, doc)
	}

	return docs
}
