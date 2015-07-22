package cat

import (
	"errors"
	"net"
)

var (
	CONN_FACTORY = func() (conn net.Conn, err error) {
		servers := CAT_SERVERS
		size := len(servers)
		for i := 0; i < size; i++ {
			conn, err = net.Dial("tcp", servers[i])
			if err == nil {
				return conn, err
			}
		}
		return nil, errors.New("Unable to access cat servers including backups.")
	}
)
