package zabbix

import (
	"fmt"

	zbxapi "github.com/cavaliercoder/go-zabbix"
	"github.com/spf13/viper"
)

type Zabbix struct {
	url      string
	username string
	password string
	session  *zbxapi.Session
}

type Interface struct {
	name string
	rx   int
	tx   int
}

func New() *Zabbix {
	zbx := Zabbix{
		url:      viper.GetString("zabbix.url"),
		username: viper.GetString("zabbix.username"),
		password: viper.GetString("zabbix.password"),
	}

	session, err := zbxapi.NewSession(
		zbx.url,
		zbx.password,
		zbx.password,
	)

	// TODO: Proper error handling.
	if err != nil {
		fmt.Printf("Unable to connect to Zabbix: %v\n", err)
	}

	zbx.session = session

	return &zbx
}

func (zbx *Zabbix) ListHosts() {}
