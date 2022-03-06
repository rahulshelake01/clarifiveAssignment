package models

import (
	"time"
)

type Member struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Father    uint   `json:"father"`
	Mother    uint   `json:"mother"`
	Brother   uint   `json:"brother"`
	Sister    uint   `json:"sister"`
	Uncle     uint   `json:"uncle"`
	Aunt      uint   `json:"aunt"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AddMemberRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Father    uint   `json:"father"`
	Mother    uint   `json:"mother"`
	Brother   uint   `json:"brother"`
	Sister    uint   `json:"sister"`
	Uncle     uint   `json:"uncle"`
	Aunt      uint   `json:"aunt"`
}

type GetMemberRequest struct {
	Name string `schema:"name"`
}

type AddMemberResponse struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type GetMemberResponse struct {
	Success    bool     `json:"success"`
	StatusCode int      `json:"statusCode"`
	Message    string   `json:"message"`
	Data       []Member `json:"data"`
}
