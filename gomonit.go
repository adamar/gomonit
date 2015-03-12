package main

import (
	// "bytes"
	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
	"encoding/xml"
	// "fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Server struct {
	Uptime        int         `xml:"uptime"`
	Poll          int         `xml:"poll"`
	StartDelay    int         `xml:"startdelay"`
	LocalHostname string      `xml:"localhostname"`
	ControlFile   string      `xml:"controlfile"`
	Httpd         Httpd       `xml:"httpd"`
	Credentials   Credentials `xml:"credentials"`
}

type Httpd struct {
	Address string `xml:"address"`
	Port    int    `xml:"port"`
	Ssl     int    `xml:"ssl"`
}

type Credentials struct {
	Username string `xml:"username"`
	Password string `xml:"password"`
}

type Platform struct {
	Name    string `xml:"name"`
	Release string `xml:"release"`
	Version string `xml:"version"`
	Machine string `xml:"machine"`
	Cpu     string `xml:"cpu"`
	Memory  string `xml:"memory"`
	Swap    string `xml:"swap"`
}

type Service struct {
	Type          int    `xml:"type"`
	CollectedSec  int    `xml:"collected_sec"`
	CollectedUsec int    `xml:"collected_usec"`
	Status        int    `xml:"status"`
	StatusHint    int    `xml:"status_hint"`
	Monitor       int    `xml:"monitor"`
	MonitorMode   int    `xml:"monitormode"`
	PendingAction int    `xml:"pendingaction"`
	Load          Load   `xml:"system>load"`
	Cpu           Cpu    `xml:"system>cpu"`
	Memory        Memory `xml:"system>memory"`
	Swap          Swap   `xml:"system>swap"`
}

type Load struct {
	Avg01 float64 `xml:"avg01"`
	Avg05 float64 `xml:"avg05"`
	Avg15 float64 `xml:"avg15"`
}

type Cpu struct {
	User   float64 `xml:"user"`
	System float64 `xml:"system"`
	Wait   float64 `xml:"wait"`
}

type Memory struct {
	Percent  float64 `xml:"percent"`
	Kilobyte int     `xml:"kilobyte"`
}

type Swap struct {
	Percent  float64 `xml:"percent"`
	Kilobyte int     `xml:"kilobyte"`
}

type Monit struct {
	Id          string    `xml:"id,attr"`
	Incarnation string    `xml:"incarnation,attr"`
	Version     string    `xml:"version,attr"`
	Server      Server    `xml:"server"`
	Platform    Platform  `xml:"platform"`
	Services    []Service `xml:"services>service"`
}

func Parse(reader io.Reader) Monit {
	var monit Monit

	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReader
	decoder.DecodeElement(&monit, nil)

	return monit
}

func MonitServer(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	monit := Parse(req.Body)

	log.Println("Got message from", monit)

	// b := new(bytes.Buffer)
	// b.ReadFrom(req.Body)
	// log.Fatal(b.String())

	io.WriteString(w, "hello, world!\n")
}

func main() {
	file, err := os.Open("stub.xml")
	if err != nil {
		log.Fatal(err)
	}

	monit := Parse(file)

	log.Println(monit)

	// http.HandleFunc("/collector", MonitServer)
	// err := http.ListenAndServe(":5001", nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
}
