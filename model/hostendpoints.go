package models

type HostEndPoint struct {
	ID         string               `bson:"_id" json:"_id"`
	Kind       string               `bson:"subject" json:"kind"`
	ApiVersion string               `bson:"subject_id" json:"apiVersion"`
	Metadata   HostEndPointMetadata `bson:"metadata" json:"metadata"`
	Spec       HostEndPointSpec     `bson:"spec" json:"spec"`
}

type HostEndPointMetadata struct {
	Name              string             `bson:"name" json:"name"`
	UID               string             `bson:"uid" json:"uid"`
	CreationTimestamp string             `bson:"creationTimestamp" json:"creationTimestamp"`
	Labels            HostEndPointLabels `bson:"labels" json:"labels"`
}
type HostEndPointLabels struct {
	IP        string `bson:"ip,omitempty" json:"ip,omitempty"`
	NameSpace string `bson:"namespace,omitempty" json:"namespace,omitempty"`
	Project   string `bson:"project,omitempty" json:"project,omitempty"`
	Role      string `bson:"role,omitempty" json:"role,omitempty"`
	Zone      string `bson:"zone,omitempty" json:"zone,omitempty"`
}
type HostEndPointSpec struct {
	Node          string   `bson:"node,omitempty" json:"node,omitempty"`
	InterfaceName string   `bson:"interfaceName,omitempty" json:"interfaceName,omitempty"`
	ExpectedIPs   []string `bson:"expectedIPs,omitempty" json:"expectedIPs,omitempty"`
}
