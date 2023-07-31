package handler

import (
	"fmt"
	"strings"
)

var (
	hostEndpointColl          = "hostendpoints"
	globalNetworksetsColl     = "globalnetworksets"
	globalNetworkPoliciesColl = "globalnetworkpolicies"
)

func Event(path string) (string, string, error) {
	if strings.Contains(path, hostEndpointColl) {
		key := strings.Split(path, hostEndpointColl+"/")[1]
		return hostEndpointColl, key, nil
	}
	if strings.Contains(path, globalNetworksetsColl) {
		key := strings.Split(path, globalNetworksetsColl+"/")[1]
		return globalNetworksetsColl, key, nil
	}
	if strings.Contains(path, globalNetworkPoliciesColl+"/") {
		key := strings.Split(strings.Split(path, globalNetworkPoliciesColl)[1], "default.")[1]
		return globalNetworkPoliciesColl, key, nil
	}
	return "", "", fmt.Errorf("%s", "event have not key")
}
