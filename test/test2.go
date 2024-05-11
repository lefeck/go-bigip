package main

// A PoolList holds a list of Pool.
type PoolList struct {
	Items    []Pool `json:"items"`
	Kind     string `json:"kind"`
	SelfLink string `json:"selfLink"`
}

type Pool struct {
	Kind                   string `json:"kind"`
	Name                   string `json:"name"`
	Partition              string `json:"partition"`
	FullPath               string `json:"fullPath"`
	Generation             int    `json:"generation"`
	SelfLink               string `json:"selfLink"`
	AllowNat               string `json:"allowNat"`
	AllowSnat              string `json:"allowSnat"`
	Description            string `json:"description,omitempty"`
	IgnorePersistedWeight  string `json:"ignorePersistedWeight"`
	IPTosToClient          string `json:"ipTosToClient"`
	IPTosToServer          string `json:"ipTosToServer"`
	LinkQosToClient        string `json:"linkQosToClient"`
	LinkQosToServer        string `json:"linkQosToServer"`
	LoadBalancingMode      string `json:"loadBalancingMode"`
	MinActiveMembers       int    `json:"minActiveMembers"`
	MinUpMembers           int    `json:"minUpMembers"`
	MinUpMembersAction     string `json:"minUpMembersAction"`
	MinUpMembersChecking   string `json:"minUpMembersChecking"`
	Monitor                string `json:"monitor,omitempty"`
	QueueDepthLimit        int    `json:"queueDepthLimit"`
	QueueOnConnectionLimit string `json:"queueOnConnectionLimit"`
	QueueTimeLimit         int    `json:"queueTimeLimit"`
	ReselectTries          int    `json:"reselectTries"`
	ServiceDownAction      string `json:"serviceDownAction"`
	SlowRampTime           int    `json:"slowRampTime"`
	MembersReference       struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"membersReference"`
}
