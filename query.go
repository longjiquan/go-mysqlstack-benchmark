package main

func queryRow() entity {
	sqlStr := "select * from mock"
	var e entity
	err := db.QueryRow(sqlStr).Scan(&e.id, &e.name)
	if err != nil {
		panic(err)
	}
	return e
}

func queryMultiRow() {
	sqlStr := "select * from mock"
	rows, err := db.Query(sqlStr)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var e entity
		err := rows.Scan(&e.id, &e.name)
		if err != nil {
			panic(err)
		}
	}
}
