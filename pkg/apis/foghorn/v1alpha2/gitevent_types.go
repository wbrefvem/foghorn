package v1alpha2

import (
	"github.com/wbrefvem/go-scm/scm"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GitEventSpec defines the desired state of GitEvent
type GitEventSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	EventType     string                `json:"eventType,omitempty"`
	ParsedWebhook scm.WebhookSerializer `json:"parsedWebhook"`
}

// GitEventStatus defines the observed state of GitEvent
type GitEventStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GitEvent is the Schema for the gitevents API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=gitevents,scope=Namespaced
type GitEvent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GitEventSpec   `json:"spec,omitempty"`
	Status GitEventStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GitEventList contains a list of GitEvent
type GitEventList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GitEvent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GitEvent{}, &GitEventList{})
}
