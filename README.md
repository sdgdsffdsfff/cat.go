PACKAGE DOCUMENTATION

package cat
    import "./"

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

    Use Event

	mycat := cat.Instance()
	func bizMethod() {
		e := mycat.NewEvent("Review", "New")
		e.Add("id", 12345)
		e.Add("user", "john")
		e.SetStatus("0")
		e.Complete()
	}()

VARIABLES

var LF = "\n"

var TAB = "\t"

var Tchan chan Transaction = make(chan Transaction)

FUNCTIONS

func Cat_init_if()
    Cat_init_if initialize cat.go, which must be down before any other
    operations, for which Instance called it automatically.

func Instance() interface{}
    As it's not recommended to apply thread local in go, apps with cat.go
    have to call Instance, keep and manage the instance returned properly.

func Invoke(f Function, values ...interface{}) ([]reflect.Value, error)
    Invoke panics if f's Kind is not Func. As accurate validation is skipped
    for performance concern, don't call Invoke unless you know what you're
    doing.

TYPES

type Function interface{}

type Message interface {
    SetStatus(Status)
    Add(string, string)
    GetType() string
    GetName() string
    GetStatus() string
    GetTimestamp() time.Time
    GetData() []byte
}

func NewMessage(t string, n string) Message

type Status interface{}

type Transaction interface {
    Message
    AddChild(Message) Transaction
    Complete()
}

func NewTransaction(t string, n string, f Function) Transaction

type Tree interface {
    NewTransaction(string, string) Transaction
}

func NewTree() Tree


