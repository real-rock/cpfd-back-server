package core

import "time"

var Location *time.Location
var Members = []string{"KooJunMo", "SonJiIn", "ParkSeungBum", "LeeGwangJo", "HeoJinSeok", "KimWanJoo", "LeeKiSeok", "GUEST1", "GUEST2", "GUEST3"}
var Objects = []string{"KooJunMo", "SonJiIn", "ParkSeungBum", "LeeGwangJo", "HeoJinSeok", "KimWanJoo", "LeeKiSeok", "AP", "DOOR", "WINDOW", "AC", "GUEST1", "GUEST2", "GUEST3", "AC2"}
var CsvHeaders = []string{"NAME", "TIME", "ACTION", "TYPE"}

const MysqlFilePath = "/var/lib/mysql-files"
const FileDir = "/mysql-files"

const StartDate = "2022-05-06 09:00:00"

func init() {
	loc, _ := time.LoadLocation("Asia/Seoul")
	Location = loc
}
