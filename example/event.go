package main

import cat "../"

func main() {
	CAT := cat.Instance()
	for {
		CAT.LogEvent("Atomic", "Event")
	}
}
