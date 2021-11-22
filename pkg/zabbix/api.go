package zabbix

import (
	"fmt"

	z "github.com/cavaliercoder/go-zabbix"
	"github.com/spf13/viper"
)

type Zabbix struct {
	url      string
	username string
	password string
	session  *z.Session
}

type Interface struct {
	Name  string `json:"name"`
	Rx    int    `json:"rx"`
	Tx    int    `json:"tx"`
	Speed int    `json:"speed"`
}

type Interfaces []Interface

func New() *Zabbix {
	zbx := Zabbix{
		url:      viper.GetString("zabbix.url"),
		username: viper.GetString("zabbix.username"),
		password: viper.GetString("zabbix.password"),
	}

	session, err := z.NewSession(
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

func (zbx *Zabbix) GetHosts() ([]z.Host, error) {
	// Get hosts (id, name) and the groups they're in (id, name).
	hosts, err := zbx.session.GetHosts(z.HostGetParams{
		GetParameters: z.GetParameters{
			OutputFields: z.SelectFields{"hostid", "name"},
		},
		SelectGroups: z.SelectFields{"groupid", "name"},
	})

	if err != nil {
		return nil, err
	}

	return hosts, nil
}

func (zbx *Zabbix) GetValues(hostids []string) ([]z.Item, error) {
	// Get a list of relevant items for the host with id `hostid`.
	values, err := zbx.session.GetItems(z.ItemGetParams{
		GetParameters: z.GetParameters{
			OutputFields: "extend",
		},
	})

	if err != nil {
		return nil, err
	}

	return values, nil
}
