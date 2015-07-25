package cat

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	. "os"
	"strconv"
	"strings"
)

//Configs about cat is required to be correct.
//The client is able to complement some of the configs,
//but is incapable of validation.
//Fortunately invalid configs causes failure of cat's init,
//thus users can be immediately aware of config fault.
var (
	PROD        string = "http://cat.ctripcorp.com"
	FAT         string = "http://cat.fws.qa.nt.ctripcorp.com"
	UAT         string = "http://cat.uat.qa.nt.ctripcorp.com"
	CAT_HOST    string = FAT
	CAT_SERVERS []string
	DOMAIN      string = "900407"
	HOSTNAME    string = ""
	IP          string = ""
	TEMPFILE    string = ".cat"
)

func cat_config_init() {
	var err error
	var resp *http.Response
	var metaServer string

	metaServer = fmt.Sprintf(CAT_HOST+"/cat/s/router?domain=%s", DOMAIN)
	resp, err = http.Get(metaServer)
	if err != nil {
		panic("Fail to access the cat meta server.")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	CAT_SERVERS = strings.Split(string(body), ";")

	if DOMAIN == "" {
		panic("DOMAIN is required, set your appid to it.")
	}

	if HOSTNAME == "" {
		HOSTNAME, err = Hostname()
		if err != nil {
			panic("Fail to auto-get HOSTNAME, try config it manually.")
		}
	}

	if IP == "" {
		IP = cat_get_ip()
	}
}

func cat_get_ip() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, addr := range addrs {
		add := strings.Split(addr.String(), "/")[0]
		if add == "127.0.0.1" || add == "::1" {
			continue
		}
		first := strings.Split(add, ".")[0]
		if _, err := strconv.Atoi(first); err == nil {
			return add
		}
	}
	return ""
}
