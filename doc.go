/*
	Package cat works as a client for Central Application Tracking(CAT).

	Import

		import cat "/your/path/to/cat"

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
	
	//Atomic Event is not supported yet.
	Use Event

		mycat := cat.Instance()
		func bizMethod() {
			e := mycat.NewEvent("Review", "New")
			e.Add("id", 12345)
			e.Add("user", "john")
			e.SetStatus("0")
			e.Complete()
		}()
*/
package cat
