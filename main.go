package main

import "warden/warden"

func main() {
	proxy := warden.New(9090)

	proxy.Add("myservice.mine", 30001)
	proxy.Add("myservice.yours", 30002)

	err := proxy.Start()
	if err != nil {
		panic(err)
	}
}
