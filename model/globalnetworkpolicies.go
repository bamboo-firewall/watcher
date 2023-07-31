package models

import (
	"github.com/projectcalico/api/pkg/lib/numorstring"
)

type Action string

const (
	Allow Action = "Allow"
	Deny         = "Deny"
	Log          = "Log"
	Pass         = "Pass"
)

type PolicyType string

const (
	PolicyTypeIngress PolicyType = "Ingress"
	PolicyTypeEgress  PolicyType = "Egress"
)

type Port map[int]int

type GlobalNetworkPolicies struct {
	ID         string                        `bson:"_id" json:"_id"`
	Kind       string                        `bson:"subject" json:"kind"`
	ApiVersion string                        `bson:"subject_id" json:"apiVersion"`
	Metadata   GlobalNetworkPoliciesMetadata `bson:"metadata,omitempty" json:"metadata,omitempty"`
	Spec       GlobalNetworkPoliciesSpec     `bson:"spec,omitempty" json:"spec,omitempty"`
}

type GlobalNetworkPoliciesMetadata struct {
	Name              string `bson:"name" json:"name"`
	UID               string `bson:"uid" json:"uid"`
	ResourceVersion   string `bson:"resourceVersion,omitempty" json:"resourceVersion,omitempty"`
	CreationTimestamp string `bson:"creationTimestamp" json:"creationTimestamp"`
}
type GlobalNetworkPoliciesSpec struct {
	Order             *float64     `bson:"order,omitempty" json:"order,omitempty"`
	Ingress           []Rule       `bson:"ingress,omitempty" json:"ingress,omitempty" validate:"omitempty,dive"`
	Egress            []Rule       `bson:"egress,omitempty" json:"egress,omitempty" validate:"omitempty,dive"`
	Selector          string       `bson:"selector,omitempty" json:"selector,omitempty" validate:"selector"`
	Types             []PolicyType `bson:"types,omitempty" json:"types,omitempty" validate:"omitempty,dive,policyType"`
	NamespaceSelector string       `bson:"namespaceSelector, omitempty" json:"namespaceSelector,omitempty" validate:"selector"`
}

type Rule struct {
	Action   Action      `bson:"action" json:"action" validate:"action"`
	Protocol string      `bson:"protocol,omitempty" json:"protocol,omitempty" validate:"omitempty"`
	ICMP     *ICMPFields `bson:"icmp,omitempty" json:"icmp,omitempty" validate:"omitempty"`
	Source   EntityRule  `bson:"source,omitempty" json:"source,omitempty" validate:"omitempty"`
	// Destination contains the match criteria that apply to destination entity.
	Destination EntityRule `bson:"destination, omitempty" json:"destination,omitempty" validate:"omitempty"`
}

type ICMPFields struct {
	Type *int `bson:"type,omitempty" json:"type,omitempty" validate:"omitempty,gte=0,lte=254"`
	Code *int `bson:"code,omitempty" json:"code,omitempty" validate:"omitempty,gte=0,lte=255"`
}

type EntityRule struct {
	Nets              []string             `bson:"nets,omitempty" json:"nets,omitempty" validate:"omitempty,dive,net"`
	Selector          string               `bson:"selector,omitempty" json:"selector,omitempty" validate:"omitempty,selector"`
	NamespaceSelector string               `bson:"namespaceSelector,omitempty" json:"namespaceSelector,omitempty" validate:"omitempty,selector"`
	Services          *ServiceMatch        `bson:"services,omitempty" json:"services,omitempty" validate:"omitempty"`
	Ports             []numorstring.Port   `bson:"ports,omitempty" json:"ports,omitempty" validate:"omitempty,dive"`
	NotNets           []string             `bson:"notNets,omitempty" json:"notNets,omitempty" validate:"omitempty,dive,net"`
	NotSelector       string               `bson:"notSelector,omitempty" json:"notSelector,omitempty" validate:"omitempty,selector"`
	NotPorts          []numorstring.Port   `bson:"notPorts,omitempty" json:"notPorts,omitempty" validate:"omitempty,dive"`
	ServiceAccounts   *ServiceAccountMatch `bson:"serviceAccounts,omitempty" json:"serviceAccounts,omitempty" validate:"omitempty"`
}

type ServiceMatch struct {
	Name      string `bson:"name,omitempty" json:"name,omitempty" validate:"omitempty,name"`
	Namespace string `bson:"namespace,omitempty" json:"namespace,omitempty" validate:"omitempty,name"`
}

type ServiceAccountMatch struct {
	Names    []string `bson:"names,omitempty" json:"names,omitempty" validate:"omitempty"`
	Selector string   `bson:"selector,omitempty" json:"selector,omitempty" validate:"omitempty,selector"`
}
type RuleMetadata struct {
	Annotations map[string]string `bson:"annotations,omitempty" json:"annotations,omitempty"`
}
