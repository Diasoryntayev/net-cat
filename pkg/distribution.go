package pkg

import "fmt"

func DistributionMessage(message, username string) {
	for i := range members {
		if i != username {
			members[i].Write([]byte(ClearLine(message) + message))
		}
		members[i].Write([]byte(fmt.Sprintf("[#{TimeValue()}][#{i}]: ")))
	}
	Messages += message
}
