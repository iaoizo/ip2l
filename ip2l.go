package ip2lo

import (
	"fmt"
)

const DEFAULT_FILE = "IPV6-COUNTRY-REGION-CITY-LATITUDE-LONGITUDE-ISP.BIN"

type Ip2L struct {
	db     *DB
	dbFile string
}

type Result struct {
	Country_long  string
	Country_short string
	City          string
	Latitude      float32
	Longitude     float32
	Isp           string
}

var ip2LObj *Ip2L

func init() {
	ip2LObj = &Ip2L{
		dbFile: DEFAULT_FILE,
	}
	db, err := ip2LObj.openDb()
	if err != nil {
		fmt.Println(err.Error())
	}
	ip2LObj.setDb(db)
}

// 获取解析实例
func NewIp2L() *Ip2L {
	return ip2LObj
}

// bin 文件加载
func (i *Ip2L) LoadDbFile(dbF string) error {
	oldFile := i.dbFile
	i.dbFile = dbF
	db, err := i.openDb()
	if err != nil {
		i.dbFile = oldFile
		return err
	}
	i.Close()
	i.setDb(db)
	return nil
}

// db 文件打开
func (i *Ip2L) openDb() (*DB, error) {
	db, err := OpenDB(i.dbFile)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// 设置db
func (i *Ip2L) setDb(db *DB) {
	i.db = db
}

// 解析ip
func (i *Ip2L) Parse(ip string) (*Result, error) {
	results, err := i.db.Get_all(ip)
	if err != nil {
		return nil, err
	}
	return &Result{
		results.Country_long,
		results.Country_short,
		results.City,
		results.Latitude,
		results.Longitude,
		results.Isp,
	}, nil
}

// 关闭db文件
func (i *Ip2L) Close() {
	i.db.Close()
}
