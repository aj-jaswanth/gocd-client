package gocd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GoCd struct {
	Url string
}

type auth struct {
	username, password string
}

var gocdAuth auth

type Agent struct {
	Uuid, Agent_name, Ip_address, Os, Free_space, Status, Sandbox, Hostname string
	Resources, Environments                                                 []string
}

func (gocd *GoCd) SetAuth(username, password string) {
	gocdAuth = auth{username, password}
}

func (gocd *GoCd) GetAgents() (agents []Agent) {
	request, ok := http.NewRequest(http.MethodGet, gocd.Url+"/agents", nil)
	request.SetBasicAuth(gocdAuth.username, gocdAuth.password)

	client := http.Client{}
	response, ok := client.Do(request)
	if ok != nil {
		return
	}

	body, ok := ioutil.ReadAll(response.Body)
	json.Unmarshal(body, &agents)
	return
}

func (gocd *GoCd) UpdateAgentHostName(agent Agent, hostname string) (updatedAgent Agent) {
	request, ok := http.NewRequest(http.MethodPatch, gocd.Url+"/agents/"+agent.Uuid, bytes.NewReader([]byte(`{"hostname":"`+hostname+`"}`)))
	request.SetBasicAuth(gocdAuth.username, gocdAuth.password)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/vnd.go.cd.v2+json")

	client := http.Client{}
	response, ok := client.Do(request)
	if ok != nil {
		return
	}
	body, ok := ioutil.ReadAll(response.Body)
	json.Unmarshal(body, &updatedAgent)
	return
}
