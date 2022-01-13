package accessors

import "fmt"

type DBReader interface {
	readData(ipAddress string) IpDataRecord
	close()
}

type IpDataRecord struct {
	Country string
	City    string
}

type DBReaderMockData struct {
	login, password, url string
}

func (d *DBReaderMockData) readData(ipAddress string) IpDataRecord {
	var data = make(map[string]IpDataRecord)
	data["192.168.1.1"] = IpDataRecord{"Isreal", "Jeruslaem"}
	data["192.168.1.0"] = IpDataRecord{"Isreal", "Tel Aviv"}
	data["0.0.0.0"] = IpDataRecord{"Spain", "Madrid"}
	data["128.0.0.0"] = IpDataRecord{"Isreal", "Eilat"}
	return data[ipAddress]
}

func (d *DBReaderMockData) close() {
	fmt.Println("Mock DB, nothing to close")
}

func ReadData(db DBReader, ipAddress string) IpDataRecord {
	ipValue := db.readData(ipAddress)
	db.close()
	return ipValue
}
