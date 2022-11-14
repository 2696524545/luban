package types

import (
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterOptions struct {
	Cluster string `uri:"cluster" binding:"required"`
}

type NameOptions struct {
	ClusterOptions `json:",inline"`
	Name           string `uri:"name" binding:"required"`
}

type NamespaceOptions struct {
	ClusterOptions `json:",inline"`
	Namespace      string `uri:"namespace" binding:"required"`
}

type RequestOptions struct {
	NamespaceOptions `json:",inline"`
	Name             string `uri:"name" binding:"required"`
}

type Node struct {
	Name       string             `json:"name"`
	Role       string             `json:"role"`
	Version    string             `json:"version"`
	Ready      v1.ConditionStatus `json:"ready"`
	NodeIP     string             `json:"nodeIP"`
	NodeInfo   v1.NodeSystemInfo  `json:"nodeInfo"`
	CreateAt   string             `json:"createAt"`
	ObjectMeta metaV1.ObjectMeta  `json:"objectMeta"`
}

type NodeDetail struct {
	// Extends list item structure.
	Node `json:",inline"`

	// NodePhase is the current lifecycle phase of the node.
	Phase v1.NodePhase `json:"phase"`

	// PodCIDR represents the pod IP range assigned to the node.
	PodCIDR string `json:"podCIDR"`

	// ID of the node assigned by the cloud provider.
	ProviderID string `json:"providerID"`

	// Unschedulable controls node schedulability of new pods. By default node is schedulable.
	Unschedulable bool `json:"unschedulable"`

	// Set of ids/uuids to uniquely identify the node.
	NodeInfo v1.NodeSystemInfo `json:"nodeInfo"`

	//// Conditions is an array of current node conditions.
	Conditions []Condition `json:"conditions"`

	// Container images of the node.
	ContainerImages []string `json:"containerImages"`

	// PodListComponent contains information about pods belonging to this node.
	PodList v1.PodList `json:"podList"`

	// Taints
	Taints []v1.Taint `json:"taints,omitempty"`

	// Addresses is a list of addresses reachable to the node. Queried from cloud provider, if available.
	Addresses []v1.NodeAddress `json:"addresses,omitempty"`

	Ready  v1.ConditionStatus `json:"ready"`
	NodeIP string             `json:"nodeIP"`
	UID    string             `json:"uid"`
}

type Condition struct {
	// Type of a condition.
	Type string `json:"type"`
	// Status of a condition.
	Status v1.ConditionStatus `json:"status"`
	// Last probe time of a condition.
	LastProbeTime metaV1.Time `json:"lastProbeTime"`
	// Last transition time of a condition.
	LastTransitionTime metaV1.Time `json:"lastTransitionTime"`
	// Reason of a condition.
	Reason string `json:"reason"`
	// Message of a condition.
	Message string `json:"message"`
}

type NodeList struct {
	ListMeta ListMeta `json:"listMeta"`
	Nodes    []Node   `json:"nodes"`
}

type ListMeta struct {
	TotalItems int `json:"totalItems"`
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
