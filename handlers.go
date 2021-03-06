package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const notifiURL = "https://notifi.it/api"

type GrafanaHookBody struct {
	DashboardID int `json:"dashboardId"`
	EvalMatches []struct {
		Value  int    `json:"value"`
		Metric string `json:"metric"`
		Tags   struct {
		} `json:"tags"`
	} `json:"evalMatches"`
	ImageURL string `json:"imageUrl"`
	Message  string `json:"message"`
	OrgID    int    `json:"orgId"`
	PanelID  int    `json:"panelId"`
	RuleID   int    `json:"ruleId"`
	RuleName string `json:"ruleName"`
	RuleURL  string `json:"ruleUrl"`
	State    string `json:"state"`
	Tags     struct {
		TagName string `json:"tag name"`
	} `json:"tags"`
	Title string `json:"title"`
}

func ApiProxyHandler(w http.ResponseWriter, r *http.Request) {
	var g GrafanaHookBody

	log.SetOutput(f)

	credentials, ok := r.URL.Query()["credentials"]
	if !ok {
		http.Error(w, "No credentials", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		log.Println("error decoding body")
		http.Error(w, "json error: "+err.Error(), http.StatusNotAcceptable)
		return
	}

	// post to notifi
	notifiArgs := fmt.Sprintf("credentials=%s&title=%s&message=%s&link=%s&image=%s", credentials[0], g.RuleName, g.Message, g.RuleURL, g.ImageURL)
	notifiArgs = strings.Replace(notifiArgs, " ", "%20", -1)
	notifiURL := fmt.Sprintf("%s?%s", notifiURL, notifiArgs)
	log.Println(notifiURL)

	client := http.Client{
		Timeout: 1 * time.Second,
	}
	resp, err := client.Get(notifiURL)
	defer resp.Body.Close()
	if err != nil {
		http.Error(w, "get error: "+err.Error(), http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		log.Println(body)
		http.Error(w, string(body), http.StatusBadRequest)
		return
	}
}
