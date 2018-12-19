package main

import (
	"bufio"
	"flag"
	"fmt"
	//"ops"
	"os"
	"strings"
	"unicode"
)


func getHostGroup(br *bufio.Reader, hostGroup string)([]string){
	hostScanner := bufio.NewScanner(br)
	hostSlice := make([]string, 20)
	lock := 0
	l := ""
	for hostScanner.Scan() {
		l = strings.TrimRightFunc(hostScanner.Text(), unicode.IsSpace)
		if len(l) == 0 {
			continue
		}
		if l[0] == '[' && l[1:len(l)-1] != hostGroup {
			lock = 0
		}
		if lock == 1 {
			hostSlice = append(hostSlice, l)
		}
		if l[0] == '[' && l[1:len(l)-1] == hostGroup {
			lock = 1
			continue
		}
	}

	return hostSlice
}

func main() {

	filePath := flag.String("f", "./hostname.txt", "file path")
	//cmd := flag.String("c", "uptime", "run command")
	//user := flag.String("u", "nobody", "the user run the command")
	hostGroup := flag.String("H", "", "the default hosts group")
	host := flag.String("h", "", "the hostname of the machine")

	flag.Parse()

	fi, err := os.Open(*filePath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)

	var hostSlice []string

	if len(*hostGroup) > 0 && len(*host) > 0 {
		fmt.Println("Can not use -H and -h at the same time")
		os.Exit(1)
	}
	if len(*host) > 0 {
		hostSlice = append(hostSlice, *host)
	}else if len(*hostGroup) > 0 {
		hostSlice = getHostGroup(br, *hostGroup)
	}

	for _, host := range hostSlice {
		if len(host) > 0 {
			//info := ops.GetServerInfobyHostname(string(host))
			//fmt.Println(ops.ServersyncTask(info["sn"], "root", "pkill -9  miner"))
			//fmt.Println(ops.ServersyncTask(info["sn"], "", "ps aux |grep miner"))
			fmt.Println(string(host))
			//fmt.Println(ops.ServersyncTask(info["sn"], *user, *cmd))
		}
	}
}
