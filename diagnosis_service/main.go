package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

type diagnosisInfo struct {
	Hostname    string   `json:"Hostname"`
	IpAddresses []string `json:"IP Addresses"` //readable format, like x.x.x.x or ::1
}

func fetchDiagnosisInfo() *diagnosisInfo {
	var di diagnosisInfo

	//hostname
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("get hostname failed, err %v\n", err)
	} else {
		fmt.Printf("Hostname: %s\n", hostname)
		di.Hostname = hostname
	}

	//ip addresses
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("lookup network interface addrs failed, err %v\n", err)
	} else {
		for _, addr := range addrs {
			fmt.Printf("%s ip address-->%s\n", addr.Network(), addr.String())
			di.IpAddresses = append(di.IpAddresses, addr.String())
		}
	}

	return &di
}

func (d diagnosisInfo) String() string {
	jsonstr, err := json.Marshal(d)
	if err != nil {
		fmt.Printf("to json failed, err %v\n", err)
		return err.Error()
	}
	return string(jsonstr)
}

func (d diagnosisInfo) diagnosis(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("diagnosis")
	if item == "ping" {
		fmt.Fprintf(w, "%s", d)
	} else {
		w.WriteHeader(http.StatusNotFound) //404
		fmt.Fprintf(w, "no diagnosis command %s", req.URL)
	}
}

func main() {

	//Sample Request: http://localhost:8000/diagnosis?diagnosis=ping
	di := fetchDiagnosisInfo()
	mux := http.NewServeMux()
	mux.HandleFunc("/diagnosis", di.diagnosis)
	log.Fatal(http.ListenAndServe(":8000", mux))
}
