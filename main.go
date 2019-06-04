package main

import (
	"fmt"
	"github.com/ananagame/rich-go/client"
	"github.com/fhs/gompd/mpd"
	"log"
	"time"
)

func main() {

	err := client.Login("585569957098553357")
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 5)
	var conn *mpd.Client
	conn, err = mpd.Dial("tcp", "localhost:6600")

	line := ""
	line1 := ""
	title := "MPD"
	// Loop printing the current status of MPD.
	for {
		status, err := conn.Status()
		if err != nil {
			log.Fatalln(err)
		}
		song, err := conn.CurrentSong()
		if err != nil {
			log.Fatalln(err)
		}
		if status["state"] == "play" {
			title = "Now playing"
			line1 = fmt.Sprintf("%s - %s", song["Artist"], song["Title"])
		} else {
			if song["Artist"] == "" {
				title = "Nothing's playing"
				line1 = ""
			} else {
				title = "Paused"
				line1 = fmt.Sprintf("%s - %s", song["Artist"], song["Title"])
			}
		}
		if line != line1 {
			line = line1
			fmt.Println(line)
			err = client.SetActivity(client.Activity{
				State:      title,
				Details:    line1,
				LargeImage: "default",   // TODO: Add image
				LargeText:  "MPD",       // TODO: Add image alt
				SmallImage: "Unknown",   // TODO: Add image
				SmallText:  "NoneSmall", // TODO: Add image alt
			})
		}
		time.Sleep(1e9)
	}
}
