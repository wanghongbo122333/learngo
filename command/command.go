package command

import (
	"bytes"
	"fmt"
	"strconv"

	"os/exec"

	"os"

	"strings"
)

func CommandRun() {

	fmt.Println("shell")

	var str, ip, data []byte

	var err error

	var cmd *exec.Cmd

	//

	cmd = exec.Command("whoami")

	str, err = cmd.Output()

	if err != nil {

		fmt.Println(err)

		os.Exit(1)

	}

	fmt.Println(string(str))

	fmt.Println("=======")

	//filter  line breaks

	fmt.Println(strings.Trim(string(str), "\n"))

	cmd = exec.Command("ping", "127.0.0.1", "-n", "1")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println("output is :", cmd.Stdout)

	//get IP ,Mac platform,can not implement
	pid := os.Getegid()
	cmd = exec.Command("ps -p " + strconv.Itoa(pid) + " -o %cpu,%mem | grep -v CPU")
	// cmd.Stdout = os.Stdout

	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		fmt.Print("err")
		return
	}
	fmt.Println("ps output is :", out.String())

	fmt.Println("=======")

	cmd = exec.Command("bash", "-c", `/sbin/ifconfig en0 | grep -E 'inet ' |  awk '{print $2}'`)

	ip, err = cmd.Output()

	if err != nil {

		fmt.Println(err)

		os.Exit(1)

	}

	fmt.Println(string(ip))

	fmt.Println(strings.Trim(string(ip), "\n"))

	//implement command

	fmt.Println("====================")

	cmd = exec.Command("/bin/sh", "-c", "echo wo shi shui wo zai na")

	data, err = cmd.Output()

	if err != nil {

		fmt.Println(err)

		os.Exit(1)

	}

	fmt.Println(string(data))

	fmt.Println(strings.Trim(string(data), "\n"))

}
