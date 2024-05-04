package client

import "time"

type Client struct {
	ClientID   int
	ClientName string
	Files      *[]File
	StartTimer time.Time
}

type File struct {
	FileID int
	Size   float64
}

func CreateEmptyClintsArr() *[]Client {
	clients := make([]Client, 0)
	return &clients
}

func NewClient(clientID int, clientName string) Client {
	emptyFilesList := make([]File, 0)
	return Client{ClientID: clientID, ClientName: clientName, Files: &emptyFilesList, StartTimer: time.Now()}
}

func NewFile(fileID int, size float64) File {
	return File{FileID: fileID, Size: size}
}

func (c *Client) ElapsedTime() float64 {
	return time.Since(c.StartTimer).Seconds()
}

func IDCounter() func() int {
	var counter int = 0
	return func() int {
		counter++
		return counter
	}
}
