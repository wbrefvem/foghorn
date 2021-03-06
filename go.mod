module github.com/foghornci/foghorn

go 1.12

require (
	github.com/go-logr/logr v0.1.0
	github.com/jenkins-x/go-scm v1.5.33
	github.com/json-iterator/go v1.1.7
	github.com/onsi/ginkgo v1.8.0
	github.com/onsi/gomega v1.5.0
	github.com/sirupsen/logrus v1.4.2
	google.golang.org/appengine v1.5.0 // indirect
	gopkg.in/src-d/go-git.v4 v4.13.1
	gopkg.in/yaml.v2 v2.2.2
	k8s.io/api v0.0.0-20190620084959-7cf5895f2711
	k8s.io/apimachinery v0.0.0-20190816221834-a9f1d8a9c101
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/code-generator v0.0.0-20190912042602-ebc0eb3a5c23 // indirect
	sigs.k8s.io/controller-runtime v0.2.1
)

replace github.com/jenkins-x/go-scm => github.com/wbrefvem/go-scm v1.5.1-0.20190925164711-8398556133f9
