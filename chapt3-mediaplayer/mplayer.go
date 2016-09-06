package main

import (
	"chapt3-mediaplayer/mlib"
	"fmt"
	"strconv"
	"chapt3-mediaplayer/mp"
	"bufio"
	"os"
	"strings"
)

var lib *library.MusicManager
var id int = 1
var ctr1, signal chan int

func handleLibCommands(tokens []string)  {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e,_ := lib.Get(i)
			fmt.Println(i+1,":",e.Name,e.Artist,e.Source,e.Type)
		}
	case "add":
		if len(tokens) == 6 {
			id++
			lib.Add(&library.MusicEntry{strconv.Itoa(id),
			tokens[2],tokens[3],tokens[4],tokens[5]})
		} else {
			fmt.Println("这个真的加不进去...")
		}
	case "remove":
		if len(tokens) == 3 {
			temp,_ := strconv.Atoi(tokens[1])
			lib.Remove(temp)
		}
	default:
		fmt.Println("不知道你在说什么...")
	}
}

func handlePlayCommand(tokens []string)  {
	if len(tokens) != 2 {
		fmt.Println("输的不对,再来一次")
		return
	}
	e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println(tokens[1],"这个音乐不存在")
	}
	mp.Play(e.Source,e.Type)
}

func main() {
	fmt.Println(`可输入一下指令:
	lib list -- 看现存的音乐
	lib add <name><artist><resource><type>
	lib remove <id>
	play <name>
	`)
	lib = library.NewMusicManager()
	r := bufio.NewReader(os.Stdin)

	for  {
		fmt.Print("输入指令-> ")
		rawLine,_,_ := r.ReadLine()
		line := string(rawLine)

		if line =="q" || line== "e" {
			break
		}

		tokens := strings.Split(line," ")
		if tokens[0]=="lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
			fmt.Println("命令输的有问题")
		}
	}


}











