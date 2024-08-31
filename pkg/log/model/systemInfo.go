package model

type SystemInfo struct {
	Hostname string `json:"hostname"`
	Version  string `json:"version"`
	Build    string `json:"build"`
	Env      string `json:"env"`
}
