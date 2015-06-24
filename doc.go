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
			t.AddTag("k0", "v0")
			t.AddTag("k1", "v1")
		}
	
	Use Event

		//Atomic Event is not supported yet.

		mycat := cat.Instance()
		func bizMethod() {
			e := mycat.NewEvent("Review", "New")
			e.AddTag("k0", "v0")
			e.AddTag("k1", "v1")
			e.SetStatus("0")
			e.Complete()
		}()

	Log Error As Event
	
		mycat := cat.Instance()
		func bizMethod() {
			err, ret := someMethod()
			mycat.LogError(err)
		}()

*/
package cat
