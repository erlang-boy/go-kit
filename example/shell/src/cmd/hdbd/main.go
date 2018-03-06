package main

import (
	. "hive_shell/config"
	"hive_shell/hsmain"
)

func main() {
	hsmain.Run(HDBD)
}
