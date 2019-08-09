package slasterisk

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
)

type VmInfo struct {
	Timestamp int64
	CallerID  string
	Duration  int
}

func GetLastVM(dir string, mailbox string) {
	inboxDir := dir + mailbox + "INBOX/"
	files, err := ioutil.ReadDir(inboxDir)
	if err != nil {
		log.Fatalf("Could not get last VM for mailbox %v\n%v", mailbox, err)
	}

	if len(files) == 0 {
		fmt.Println("Mailbox inbox is empty.")
	}

	var newestFile string
	var newestTime int64

	for _, f := range files {
		fi, err := os.Stat(dir + f.Name())
		if err != nil {
			fmt.Println(err)
		}
		currTime := fi.ModTime().Unix()
		if currTime > newestTime {
			newestTime = currTime
			newestFile = f.Name()
		}
	}
	fmt.Println(newestFile)
}

func ParseVMInfo(path string) (info *VmInfo) {
	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("Could not open %v\n%v", file, err)
	}

	scanner := bufio.NewScanner(file)
	var data [][]string

	for scanner.Scan() {
		line := scanner.Text()
		reg := regexp.MustCompile(`(.*)=(.*)`)
		match := reg.FindStringSubmatch(line)
		data = append(data, match)
	}

	parsedInfo := make(map[string]string)

	for _, v := range data {
		if v != nil {
			parsedInfo[v[1]] = v[2]
		}
	}

	info = new(VmInfo)
	info.CallerID = parsedInfo["callerid"]
	info.Duration, _ = strconv.Atoi(parsedInfo["duration"])
	info.Timestamp, _ = strconv.ParseInt(parsedInfo["origtime"], 10, 64)
	return info
}
