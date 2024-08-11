package zabbix

// https://www.zabbix.com/documentation/current/en/manual/api/reference/host/object#host-tag
type Tag struct {
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type Tags []Tag
