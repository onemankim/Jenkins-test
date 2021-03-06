// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "stonelb/pkg/apis/network/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FloatingIPLister helps list FloatingIPs.
// All objects returned here must be treated as read-only.
type FloatingIPLister interface {
	// List lists all FloatingIPs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FloatingIP, err error)
	// FloatingIPs returns an object that can list and get FloatingIPs.
	FloatingIPs(namespace string) FloatingIPNamespaceLister
	FloatingIPListerExpansion
}

// floatingIPLister implements the FloatingIPLister interface.
type floatingIPLister struct {
	indexer cache.Indexer
}

// NewFloatingIPLister returns a new FloatingIPLister.
func NewFloatingIPLister(indexer cache.Indexer) FloatingIPLister {
	return &floatingIPLister{indexer: indexer}
}

// List lists all FloatingIPs in the indexer.
func (s *floatingIPLister) List(selector labels.Selector) (ret []*v1alpha1.FloatingIP, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FloatingIP))
	})
	return ret, err
}

// FloatingIPs returns an object that can list and get FloatingIPs.
func (s *floatingIPLister) FloatingIPs(namespace string) FloatingIPNamespaceLister {
	return floatingIPNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FloatingIPNamespaceLister helps list and get FloatingIPs.
// All objects returned here must be treated as read-only.
type FloatingIPNamespaceLister interface {
	// List lists all FloatingIPs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FloatingIP, err error)
	// Get retrieves the FloatingIP from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FloatingIP, error)
	FloatingIPNamespaceListerExpansion
}

// floatingIPNamespaceLister implements the FloatingIPNamespaceLister
// interface.
type floatingIPNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FloatingIPs in the indexer for a given namespace.
func (s floatingIPNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FloatingIP, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FloatingIP))
	})
	return ret, err
}

// Get retrieves the FloatingIP from the indexer for a given namespace and name.
func (s floatingIPNamespaceLister) Get(name string) (*v1alpha1.FloatingIP, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("floatingip"), name)
	}
	return obj.(*v1alpha1.FloatingIP), nil
}
