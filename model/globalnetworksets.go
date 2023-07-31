package models

type GlobalNetworkSet struct {
	ID         string                   `bson:"_id" json:"_id"`
	Kind       string                   `bson:"subject" json:"kind"`
	ApiVersion string                   `bson:"subject_id" json:"apiVersion"`
	Metadata   GlobalNetworkSetMetaData `bson:"metadata" json:"metadata"`
	Spec       GlobalNetworkSetSpec     `bson:"spec" json:"spec"`
}

type GlobalNetworkSetMetaData struct {
	Name              string                 `bson:"name" json:"name"`
	UID               string                 `bson:"uid" json:"uid"`
	CreationTimestamp string                 `bson:"creationTimestamp" json:"creationTimestamp"`
	Labels            GlobalNetworkSetLabels `bson:"labels" json:"labels"`
}
type GlobalNetworkSetLabels struct {
	Name string `bson:"name,omitempty" json:"name,omitempty"`
	Zone string `bson:"zone,omitempty" json:"zone,omitempty"`
}
type GlobalNetworkSetSpec struct {
	Nets []string `bson:"nets,omitempty" json:"nets,omitempty"`
}
