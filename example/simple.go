package main

import "time"
import cat "../"

func main() {
	cat := cat.Instance()
	tr := cat.NewTransaction("ImgSvr", "Resize")

	err := recover()
	tr.SetStatus(err)
	tr.Complete()

	time.Sleep(3*time.Second)
}
