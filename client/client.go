package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	com "go-exercise/libs/comunication"
	st "go-exercise/libs/structs"
)

var textToParse = flag.String("text", "", "testo da parsare")
var serverAddress = flag.String("server", st.MasterAddress, "master server")

func main() {
	flag.Parse()

	//chSend := make(chan []byte)
	//chAck := make (charn string)

	fmt.Println("File to parse: ", *textToParse)

	//connection to a server (master server)
	server := com.ConnectToHost(*serverAddress)

	//msgToSend := &st.StringMsg{textToParse}
	var msgFromServer = &st.SlaveResponse{}

	/*reader := bufio.NewReader(os.Stdin)
	fmt.Print("Insert the file path of the file to send: ")
	fileToSend, _ := reader.ReadString('\n')
	fmt.Printf("fileToSend: %s\n", fileToSend)*/
	fts := st.FileToSend{}

	replace := strings.Replace(*textToParse, "\n", "", 1)
	b, err := ioutil.ReadFile(replace)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", string(b))
	fts.File = string(b)

	//st.TextParse(fts.File)

	err = server.Call("Work.WordCount", fts, &msgFromServer)
	if err != nil {
		log.Fatal("Error in Work.WordCount: ", err)
	}
	msgFromServer.WordHashMap.PrintTable()

}
