// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/openshift/cluster-kube-apiserver-operator/pkg/apis/kubeapiserver/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KubeAPIServerOperatorConfigLister helps list KubeAPIServerOperatorConfigs.
type KubeAPIServerOperatorConfigLister interface {
	// List lists all KubeAPIServerOperatorConfigs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.KubeAPIServerOperatorConfig, err error)
	// Get retrieves the KubeAPIServerOperatorConfig from the index for a given name.
	Get(name string) (*v1alpha1.KubeAPIServerOperatorConfig, error)
	KubeAPIServerOperatorConfigListerExpansion
}

// kubeAPIServerOperatorConfigLister implements the KubeAPIServerOperatorConfigLister interface.
type kubeAPIServerOperatorConfigLister struct {
	indexer cache.Indexer
}

// NewKubeAPIServerOperatorConfigLister returns a new KubeAPIServerOperatorConfigLister.
func NewKubeAPIServerOperatorConfigLister(indexer cache.Indexer) KubeAPIServerOperatorConfigLister {
	return &kubeAPIServerOperatorConfigLister{indexer: indexer}
}

// List lists all KubeAPIServerOperatorConfigs in the indexer.
func (s *kubeAPIServerOperatorConfigLister) List(selector labels.Selector) (ret []*v1alpha1.KubeAPIServerOperatorConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KubeAPIServerOperatorConfig))
	})
	return ret, err
}

// Get retrieves the KubeAPIServerOperatorConfig from the index for a given name.
func (s *kubeAPIServerOperatorConfigLister) Get(name string) (*v1alpha1.KubeAPIServerOperatorConfig, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("kubeapiserveroperatorconfig"), name)
	}
	return obj.(*v1alpha1.KubeAPIServerOperatorConfig), nil
}
