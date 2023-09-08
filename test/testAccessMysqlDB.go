package test

import "gocode.ethan/ethereum-dev/utils"

func TestAccessMysqlDB() {
	db := utils.GetMysqlDB()
	// 关闭整个程序之前执行db.Close()
	defer db.Close()

	utils.Query(db, "SELECT id, `key`, value FROM aggregator_ethan.`system`;")

	utils.Insert(db, "INSERT INTO `system` (`key`, value) VALUES(?, ?);", "ens", "ether.eth")

	utils.Query(db, "SELECT id, `key`, value FROM aggregator_ethan.`system`;")
}
