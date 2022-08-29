package main

import "goTestProject/consulDemo/base"

func main() {
	//base.Register("127.0.0.1", 8021, "darren", []string{"11"}, "darren")
	//base.AllServices()
	base.ServicesWithFilter()
}
