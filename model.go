package main

import ()

type KeyValue struct {
	UserName string `json:"userName"`
	Key      string `json:"key"`
	Value    string `json:"value"`
}

type UserInfo struct {
	Sub      string `json:"sub"`
	Aud      string `json:"aud"`
	Email    string `json:"email"`
	UserName string `json:"cognito:username"`
}
