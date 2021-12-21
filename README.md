# JNDILOG
use JNDILOG instead of DNSLOG

1.start JNDILOG server

2.send JNDI POC (eg: ldap://127.0.0.1:33333/randomstr or rmi://127.0.0.1:33333/randomstr)

3.receive JNDILOG "randomstr"

此项目是从vscan脱离出来的JNDILOG服务，旨在不使用DNSLOG的情况下检测JNDI漏洞


```
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
			fmt.Println("get randomstr JNDI connect")
			break
		}
	}
}
```