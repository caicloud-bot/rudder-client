/*
Copyright 2019 caicloud authors. All rights reserved.
*/

// Code generated by listerfactory-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "github.com/caicloud/clientset/listerfactory/internalinterfaces"
	storagev1beta1 "k8s.io/api/storage/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	kubernetes "k8s.io/client-go/kubernetes"
	v1beta1 "k8s.io/client-go/listers/storage/v1beta1"
)

var _ v1beta1.CSIDriverLister = &cSIDriverLister{}

// cSIDriverLister implements the CSIDriverLister interface.
type cSIDriverLister struct {
	client           kubernetes.Interface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCSIDriverLister returns a new CSIDriverLister.
func NewCSIDriverLister(client kubernetes.Interface) v1beta1.CSIDriverLister {
	return NewFilteredCSIDriverLister(client, nil)
}

func NewFilteredCSIDriverLister(client kubernetes.Interface, tweakListOptions internalinterfaces.TweakListOptionsFunc) v1beta1.CSIDriverLister {
	return &cSIDriverLister{
		client:           client,
		tweakListOptions: tweakListOptions,
	}
}

// List lists all CSIDrivers in the indexer.
func (s *cSIDriverLister) List(selector labels.Selector) (ret []*storagev1beta1.CSIDriver, err error) {
	listopt := v1.ListOptions{
		LabelSelector: selector.String(),
	}
	if s.tweakListOptions != nil {
		s.tweakListOptions(&listopt)
	}
	list, err := s.client.StorageV1beta1().CSIDrivers().List(listopt)
	if err != nil {
		return nil, err
	}
	for i := range list.Items {
		ret = append(ret, &list.Items[i])
	}
	return ret, nil
}

// Get retrieves the CSIDriver from the index for a given name.
func (s *cSIDriverLister) Get(name string) (*storagev1beta1.CSIDriver, error) {
	return s.client.StorageV1beta1().CSIDrivers().Get(name, v1.GetOptions{})
}
