package gocd

import (
	"encoding/json"
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
	Uuid, Agentname, Ip_address, Os, Free_space, Status, Sandbox string
	Resources, Environments                                      []string
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
	return agents
}
