package cat

var cat_lock chan int = make(chan int, 1)
var cat_initialized bool = false

//Cat_init_if initialize cat.go,
//which must be down before any other operations, 
//for which Instance called it automatically.
func Cat_init_if() {
	cat_lock <- 0
	if !cat_initialized {
		/*read config
		start sender*/
		cat_initialized = true
	}
	<-cat_lock
}

//As it's not recommended to apply thread local in go, 
//apps with cat.go have to call Instance, 
//keep and manage the instance returned properly.
func Instance() *cat {
	Cat_init_if()
	return &cat{M.Newtree()}
}
