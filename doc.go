/*
	Package cat works as a client for Central Application Tracking(CAT).

	Import

		import cat "/your/path/to/cat"

	Config

		cat.DOMAIN   = "your appid"
		cat.HOSTNAME = "your hostname" //optional
		cat.IP       = "your hostip"   //optional
		cat.TEMPFILE = ".cat"          //optional, "your/path/to/.cat"
		cat.CAT_HOST = cat.UAT         //or "http://cat.uat.qa.nt.ctripcorp.com"

	Use Transaction

		mycat := cat.Instance()
		func() {
			t := mycat.NewTransaction("URL", "Page")
			defer func() {
				p := recover()
				mycat.LogPanic(p)
				t.SetStatus(p)
				t.Complete()
			}()
			// do your bussiness here
			// perhaps panic
			t.AddData("k0", "v0")
			t.AddData("k1", "v1")
		}()

	Use Event

		mycat := cat.Instance()
		func() {
			e := mycat.NewEvent("Review", "New")
			e.AddData("k0", "v0")
			e.AddData("k1", "v1")
			e.SetStatus("0")
			e.Complete()
		}()


	Use Heartbeat

		mycat := cat.Instance()
		func() {
			h := mycat.NewHeartbeat("Heartbeat", "192.168.141.131")
			h.Set("System", "CPU", "0.3")
			h.Set("System", "DISK", "0.9")
			h.SetStatus("0")
			h.Complete()
		}()

	Log Error As Event

		mycat := cat.Instance()
		func() {
			err, ret := someMethod()
			mycat.LogError(err)
		}()

*/
package cat
