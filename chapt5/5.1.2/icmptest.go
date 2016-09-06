package main

import (
	"os"
	"fmt"
	"net"
	"bytes"
	"io"
)

func main() {
	fmt.Println("hello --------")
	if len(os.Args) != 2 {
		fmt.Println("参数数量不对")
		os.Exit(1)
	}
	service := os.Args[1]

	conn, err := net.Dial("ip4:icmp", service)
	fmt.Println("---------1--------")

	checkError(err)

	fmt.Println("---------2--------")

	var msg [512]byte
	msg[0] = 8
	msg[1] = 0
	msg[2] = 0
	msg[3] = 0
	msg[4] = 0
	msg[5] = 13
	msg[6] = 0
	msg[7] = 37

	len := 8
	check := checkSum(msg[0:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	fmt.Println("---------2--------")

	_, err = conn.Write(msg[0:len])
	checkError(err)

	_, err = readFully(conn);
	checkError(err)

	fmt.Println("Got response")
	if msg[5] == 13 {
		fmt.Println("Indentifier matchs")
	}
	if msg[7] == 37 {
		fmt.Println("sqquence matchs")
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println("errororrrrrrrrrrrrrrrrrrrrrrrr")
		fmt.Fprintf(os.Stderr, "Fatal error : %s", err.Error())
		os.Exit(1)
	}
}

func checkSum(msg []byte) uint16 {
	sum := 0

	for n := 1; n < len(msg) - 1; n += 2 {
		sum += int(msg[n]) * 256 + int(msg[n + 1])
	}
	sum = (sum >> 16) + (sum & 0xffff)
	// sum += (sum >> 16)
	var answer uint16 = uint16(^sum)

	fmt.Printf("校验和为:  %d", answer)

	return answer
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}