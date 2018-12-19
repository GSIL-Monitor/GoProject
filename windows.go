package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"strconv"
)


func listAllProcess(){
	cmd := exec.Command("tasklist")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}

func listProcess(text string) string{
	cmd := exec.Command("tasklist", "/fi", text)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func killExE(exename string) bool {
	log.Fatal("kill exe")
	exename = exename + '.exe'
	arg := []string{"/im", exename}
	cmd :=exec.Command("taskkill", arg...)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	return true
}

func main(){
	c := exec.Command("cmd", "/C", "del", "C:\\a.txt")
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}