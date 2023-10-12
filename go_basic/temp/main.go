package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("begin")
	para := `"curl -XGET -u elastic:bytepower -H 'Content-type: application/json' http://es-cluster.bytepower.app:9200/_cat/indices\?format\=json\&bytes='mb'\&h='index,store.size,pri'"`
	fmt.Println(exec.Command("ssh", "bp-global", para).Output())

	v := string(Runcmd(para, false))
	// v := string(Runcmd("./script.sh", true))
	fmt.Println("tag:", v)
	fmt.Println("end")
}

func Runcmd(cmd string, shell bool) []byte {
	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			fmt.Println("error:", err)
			log.Fatal(err)
			//panic("some error found")
		}
		return out
	}
	out, err := exec.Command("ssh", cmd).Output()
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	return out
}
