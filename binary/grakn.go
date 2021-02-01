package main

import "fmt"
// import "flag"
import "syscall"
import "os/exec"
import "os"
import "log"	

func main() {

	current_dir, err := os.Getwd()

	server_classpath := fmt.Sprintf("%s/server/conf/:%s/server/lib/common/*:%s/server/lib/prod/*", current_dir, current_dir, current_dir)

	fmt.Println(server_classpath)

	

	path, err := exec.LookPath("java")

	if os.Args[1] == "server" {
		x := []string{path, "-classpath", server_classpath, "grakn.core.server.GraknServer"}
		x = append(x, os.Args[2:]...)

		fmt.Println(x)

		fmt.Println(path)

		if err != nil {
			log.Fatal("installing java is in your future")
		}


		if err := syscall.Exec(path, x, []string{}); err != nil {
			fmt.Println(err)
		}
	}
}