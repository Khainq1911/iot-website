package model

type Information struct {
	Node_id string  `json: "node_id"`
	Power   float32 `json: "power"`
	Current float32 `json: "current"`
	Voltage float32 `json: "voltage"`
	Energy  float32 `json: "energy"`
	Status  bool    `json: "status"`
}

type UpdateStatus struct {
	Status bool "json: status"
}
