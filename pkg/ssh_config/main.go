package main

import (
	"fmt"

	"github.com/kyl2016/Play-With-Golang/utility"
)

func main() {
	// f, _ := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "config"))
	// cfg, _ := ssh_config.Decode(f)
	// for _, host := range cfg.Hosts {
	// 	for _, pattern := range host.Patterns {
	// 		fmt.Println(pattern.String())
	// 	}

	// 	// fmt.Println("patterns:", host.Patterns)
	// 	for _, node := range host.Nodes {
	// 		// Manipulate the nodes as you see fit, or use a type switch to
	// 		// distinguish between Empty, KV, and Include nodes.
	// 		// fmt.Println(node.String())
	// 		if kv, ok := node.(*ssh_config.KV); ok {
	// 			fmt.Println(kv.Key, kv.Value)
	// 		}
	// 	}
	// }

	// Print the config to stdout:
	// fmt.Println(cfg.String())
	client, err := NewSshClient("ubuntu", "69.234.208.147", 22, "/Users/kyl/.ssh/key-087-cn.pem", "")
	utility.PanicIfNotNil(err)

	output, err := client.RunCommand("sudo cat /etc/telegraf/telegraf.conf | grep 2bp-carbon.tick.vpc")
	fmt.Println(output, err)
}
