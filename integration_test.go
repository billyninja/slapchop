package main

import (
	"crypto/tls"
	"encoding/json"
	"github.com/billyninja/slapchop/actions"
	"github.com/billyninja/slapchop/chopper"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

var testPort = "3002"
var testHost = "http://localhost:" + testPort
var client = http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
}

func TestChain(t *testing.T) {

	server_control := InitServer(testPort)

	// CREATE
	req, err := actions.NewfileUploadRequest(testHost+"/chopit/testuser", "uploadfile", "test/gopher.jpg")
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	cr := chopper.CreateResponse{}
	err = json.Unmarshal(body, &cr)
	if err != nil {
		t.Error(err)
		return
	}
	resp.Body.Close()

	// READ-ALL
	req, err = http.NewRequest("GET", testHost+"/chopit/testuser/", nil)
	resp, err = client.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	body, _ = ioutil.ReadAll(resp.Body)
	rar := chopper.ReadAllResponse{}
	err = json.Unmarshal(body, &rar)
	if err != nil {
		t.Error(err)
		return
	}

	if len(rar.Slapchops) == 0 || rar.User != "testuser" {
		t.Error("Bad Response")
		return
	}
	resp.Body.Close()

	// READ-ONE
	req, err = http.NewRequest("GET", testHost+"/chopit/testuser/"+cr.ChopId, nil)
	resp, err = client.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	body, _ = ioutil.ReadAll(resp.Body)
	ror := chopper.ReadAllResponse{}
	err = json.Unmarshal(body, &ror)
	if err != nil {
		t.Error(err)
		return
	}
	resp.Body.Close()

	// DELETE
	req, err = http.NewRequest("DELETE", testHost+"/chopit/testuser/"+cr.ChopId, nil)
	resp, err = client.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	// Closing the server
	server_control <- os.Interrupt
}
