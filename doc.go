/*
	Package cat works as a client for Central Application Tracking(CAT).

	Import

		import cat "/your/path/to/cat"

	Use Transaction

		mycat := cat.Instance()
		func bizMethod() {
			t := mycat.NewTransaction("URL", "Page")
			defer func {
				p := recover()
				mycat.LogPanic(p)
				t.SetStatus(p)
				t.Complete()
			}()
			// do your bussiness here
			// perhaps panic
			t.Add("k0", "v0")
			t.Add("k1", "v1")
		}
	
	//Atomic Event is not supported yet.
	Use Event

		mycat := cat.Instance()
		func bizMethod() {
			e := mycat.NewEvent("Review", "New")
			e.Add("k0", "v0")
			e.Add("k1", "v1")
			e.SetStatus("0")
			e.Complete()
		}()
*/
package cat
