package main

import (
	"fmt"
	"github.com/veo/jndilog"
)

func main() {
	go jndilog.JndiServer("127.0.0.1:33333")
	// send ldap://127.0.0.1:33333/randomstr  or rmi://127.0.0.1:33333/randomstr POC
	for {
		if jndilog.Jndilogchek("randomstr") {
			fmt.Println("randomstr JNDI RCE")
			break
		}
	}
}
