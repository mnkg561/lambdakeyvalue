package main

import ()

//KeyValue This will be used for KeyValue model
type KeyValue struct {
	UserName string `json:"userName"`
	Key      string `json:"key"`
	Value    string `json:"value"`
}

//UserInfo This will be used for UserInfo
type UserInfo struct {
	Sub      string `json:"sub"`
	Aud      string `json:"aud"`
	Email    string `json:"email"`
	UserName string `json:"cognito:username"`
}
