package pkg

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"
)

var (
	members  = map[string]net.Conn{}
	Massages string
	Mutex    sync.Mutex
	Count    int = 0
)

func ToChat(member net.Conn) {
	var name, message string
	file, err := os.Open("./pkg/penguin.txt")
	Error(err)
	r := bufio.NewReader(file)
	buf := make([]byte, 500)
	io.ReadFull(r, buf)
	member.Write(buf)

	Mutex.Lock()
	scan := bufio.NewScanner(member)
	for scan.Scan() {
		name = scan.Text()
		name = strings.TrimSpace(name)
		_, check := members[name]
		if len(name) == 0 || check {
			member.Write([]byte("Enter correct name: "))
		} else {
			members[name] = member
			break
		}
	}
	member.Write([]byte(Messages))
	DistributionMessage(name+" has joined our chats...\n", name)
	Mutex.Unlock()

	for {
		check := scan.Scan()
		if !check {
			break
		}
		text := scan.Text()
		str := strings.Trim(text, "\n\t\r")
		str1 := strings.TrimSpace(str)
		if len(str1) != 0 {
			Mutex.Lock()
			message = string(fmt.Sprintf("[%s][%s]: ", TimeValue(), name)) + text + "\n"
			DistributionMessage(message, name)
			Mutex.Unlock()
		}

	}
	Mutex.Lock()
	DistributionMessage(name+" has left our chat...\n", name)
	delete(members, name)
	Count--
	Mutex.Unlock()
}
