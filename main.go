//go:generate fileb0x b0x.yaml
package main

import "github.com/timo-reymann/fritzbox-based-presence/cmd"

func main() {
	cmd.Run()
}
