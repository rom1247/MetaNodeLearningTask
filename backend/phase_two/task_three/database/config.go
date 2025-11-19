package database

type DBType string

const (
	SQLite = "sqlite"
	MySQL  = "mysql"
)

type Config struct {
	DBType     DBType
	MySqlDsn   string
	SqlitePath string
}

func Load(dbType DBType) *Config {
	switch dbType {
	case SQLite:
		return &Config{
			DBType:     SQLite,
			SqlitePath: "file:test.db?cache=shared&_temp_store=1",
		}

	case MySQL:
		return &Config{
			DBType:   MySQL,
			MySqlDsn: "webmail_user:xhtest2018@tcp(192.168.80.199:3306)/yyl_go_center?charset=utf8mb4&parseTime=True&loc=Asia/Shanghai",
		}
	}
	return nil

}
