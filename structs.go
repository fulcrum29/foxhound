package main

import "log/slog"

type api struct {
	logger *slog.Logger
}

type UHC struct {
	Name         string `json:"name"`
	ReplicaCount int    `json:"replicaCount"`
	Resources    struct {
		Limits struct {
			CPU    string `json:"cpu"`
			Memory string `json:"memory"`
		} `json:"limits"`
		Requests struct {
			CPU    string `json:"cpu"`
			Memory string `json:"memory"`
		} `json:"requests"`
	} `json:"resources"`
}
