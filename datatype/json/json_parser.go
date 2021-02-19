package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type ClientParticipated struct {
	ClientId       string     `json:"client_id,omitempty"`
	ExperimentId   string     `json:"experiment_id,omitempty"`
	VariantId      string     `json:"variant_id,omitempty"`
	ParticipatedAt *time.Time `json:"participated_at,string,omitempty"`
}

func main() {
	json_string := "{\"client_id\":\"cli_1\",\"experiment_id\":\"exp_1\",\"variant_id\":\"variant_id\",\"participated_at\":\"2021-01-27T12:09:20.100216+07:00\"}"
	var cp ClientParticipated
	if err := json.Unmarshal([]byte(json_string), &cp); err != nil {
		fmt.Printf("err: %v", err)
	}
	fmt.Println(cp)
}
