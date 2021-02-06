package api

type Pairs struct {
	Success bool     `json:"success"`
	Data    []string `json:"data"`
}

type VolumeData struct {
	VolumeUsd float64 `json:"volumeUsd"`
	Volume    float64 `json:"volume"`
}

type Volume struct {
	Success bool         `json:"success"`
	Data    []VolumeData `json:"data"`
}
