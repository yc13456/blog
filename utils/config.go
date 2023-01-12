package utils

type SqlParam struct {
	UserName string
	PassWord string
	Host     string
	Port     int
	Dbname   string
	Timeout  string
}

type RedisParam struct {
	Addr 		string
	PassWord  	string
	DB 			int
	PoolSize 	int
}