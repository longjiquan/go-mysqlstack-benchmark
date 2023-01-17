package main

func prepareInsert() {
	sqlStr := "insert into user(name, age) values (?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec("小王子", 18)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec("沙河娜扎", 18)
	if err != nil {
		panic(err)
	}
}
