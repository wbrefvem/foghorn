package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ActionName defines what type of action should be taken
type ActionName string

const (
	// CreatePipelineRun indicates that a Tekton PipelineRun should be created
	// in response to the webhook that triggered this action.
	CreatePipelineRun ActionName = "CreatePipelineRun"
	// CommentOnIssue indicates that we should leave a comment on a specified
	// issue on specified repo on a configured git provider.
	CommentOnIssue ActionName = "CommentOnIssue"
	// CommentOnPR indicates that we should comment on a pull request (or equivalent
	// contstruct) for a repo on a configured git provider.
	CommentOnPR ActionName = "CommentOnPR"
	// LabelIssue indicates that we should label an issue on a repo on a configured
	// git provider.
	LabelIssue ActionName = "LabelIssue"
	// LabelPR indicates that we should label a PR on a repo on a configured git
	// provider
	LabelPR ActionName = "LabelPR"
	// MergePR indicates that a PR is ready to merge in response to an approve
	// or lgtm event.
	MergePR ActionName = "MergePR"
)

// ActionSpec defines the desired state of Action
type ActionSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Type        ActionName `json:"type"`
	ParentEvent *GitEvent  `json:"parentEvent"`
}

// ActionStatus defines the observed state of Action
type ActionStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Action is the Schema for the actions API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=actions,scope=Namespaced
type Action struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ActionSpec   `json:"spec,omitempty"`
	Status ActionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ActionList contains a list of Action
type ActionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Action `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Action{}, &ActionList{})
}
