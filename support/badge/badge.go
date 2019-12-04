package badge

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	session     string
	leaderboard string
)

func init() {
	session = os.Getenv("AOC_SESSION")
	if "" == session {
		log.Fatal("AOC_SESSION not set or empty")
	}

	leaderboard = os.Getenv("LEADERBOARD")
	if "" == leaderboard {
		log.Fatal("LEADERBOARD not set or empty")
	}
}

// Badge returns text needed for
func Badge(w http.ResponseWriter, r *http.Request) {

	userid := r.URL.Query().Get("user_id")
	if userid == "" {
		http.Error(w, "user_id parameter is required", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("https://adventofcode.com/2019/leaderboard/private/view/%s.json", leaderboard)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	req.Header.Add("COOKIE", fmt.Sprintf("session=%s", session))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	dec := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var lb leaderboardResponse

	err = dec.Decode(&lb)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	shield := shieldsIOResponse{1, "AoC Stars", "", "308413", "D90A16"}

	for i, m := range lb.Members {
		if i == userid {
			shield.Message = strconv.Itoa(m.Stars)
			break
		}
	}

	if shield.Message == "" {
		http.Error(w, fmt.Sprintf("No entry found for user %q", userid), http.StatusNotFound)
		return
	}

	enc := json.NewEncoder(w)
	err = enc.Encode(shield)
	if err != nil {
		log.Println(err)
	}
}

type shieldsIOResponse struct {
	SchemaVersion int    `json:"schemaVersion"`
	Label         string `json:"label"`
	Message       string `json:"message"`
	Color         string `json:"color"`
	LabelColor    string `json:"labelColor"`
}

type leaderboardResponse struct {
	Members map[string]member `json:"members"`
}
type member struct {
	Stars int `json:"stars"`
}
