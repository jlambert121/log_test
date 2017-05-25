package main

import (
	"fmt"
	"strconv"
	"time"
)

// Send number (of messages) over duration (s)
func send_msg_per_time(number int, duration int) {
	sleep := time.Duration(int64(float64(duration) / float64(number) * 1000)) * time.Millisecond

	var count = 0

	for count < number {
		count++
		fmt.Printf("Message number %s of %s (%s)\n", strconv.Itoa(count), strconv.Itoa(number), time.Now().Format("15:04:05.00000"))
		time.Sleep(sleep)
	}
}

func main() {
	// 30 messages over 30 seconds
	send_msg_per_time(30, 30)
	fmt.Print("Sleeping 60s\n")
	time.Sleep(60 * time.Second)

	// 300 messages over 30 seconds
	send_msg_per_time(300, 30)
	fmt.Print("Sleeping 60s\n")
	time.Sleep(60 * time.Second)

	// 1000 messages over 30 seconds
	send_msg_per_time(1000, 30)
	fmt.Print("Sleeping 60s\n")
	time.Sleep(60 * time.Second)

	// 2000 messages over 30 seconds
	send_msg_per_time(2000, 30)
	fmt.Print("Sleeping 60s\n")
	time.Sleep(60 * time.Second)

	// 3000 messages over 30 seconds
	send_msg_per_time(3000, 30)
	fmt.Print("Sleeping 60s\n")
	time.Sleep(60 * time.Second)

	// 5000 messages over 30 seconds
	send_msg_per_time(5000, 30)
	fmt.Print("Sleeping 60s\n")
	time.Sleep(60 * time.Second)

	// 7000 messages over 30 seconds
	send_msg_per_time(7000, 30)
	fmt.Print("Sleeping 60s\n")
	time.Sleep(60 * time.Second)

}
