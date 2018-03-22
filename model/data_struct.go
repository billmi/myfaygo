package model

type DataInfo struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}
