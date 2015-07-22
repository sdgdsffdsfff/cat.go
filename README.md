PACKAGE DOCUMENTATION

package cat

    Package cat works as a client for Central Application Tracking(CAT).

    Import

	import cat "/your/path/to/cat"

    Config

        cat.DOMAIN   = "your appid"  
        cat.HOSTNAME = "your hostname" //optional  
        cat.IP       = "your hostip"   //optional  
        cat.CAT_HOST = cat.UAT         // or "http://cat.uat.qa.nt.ctripcorp.com"  


    Use Transaction

	mycat := cat.Instance()
	func bizMethod() {
		t := mycat.NewTransaction("URL", "Page")
		defer func {
			err := recover()
			t.SetStatus(err)
			t.Complete()
		}()
		// do your bussiness here
		t.Add("k1", "v1")
		t.Add("k2", "v2")
		t.Add("k3", "v3")
	}

    Use Event

	mycat := cat.Instance()
	func bizMethod() {
		e := mycat.NewEvent("Review", "New")
		e.Add("id", 12345)
		e.Add("user", "john")
		e.SetStatus("0")
		e.Complete()
	}()

FUNCTIONS

    func Cat_init_if()  
        Cat_init_if initialize cat.go, which must be down before any other  
        operations, for which Instance called it automatically.  
    
    func Instance() interface{}  
        As it's not recommended to apply thread local in go, apps with cat.go  
        have to call Instance, keep and manage the instance returned properly.  
