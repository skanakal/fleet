/*
Copyright (c) 2020 - 2025 SUSE LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/rancher/fleet/pkg/apis/fleet.cattle.io/v1alpha1"
	"github.com/rancher/wrangler/v3/pkg/generic"
)

// ContentController interface for managing Content resources.
type ContentController interface {
	generic.NonNamespacedControllerInterface[*v1alpha1.Content, *v1alpha1.ContentList]
}

// ContentClient interface for managing Content resources in Kubernetes.
type ContentClient interface {
	generic.NonNamespacedClientInterface[*v1alpha1.Content, *v1alpha1.ContentList]
}

// ContentCache interface for retrieving Content resources in memory.
type ContentCache interface {
	generic.NonNamespacedCacheInterface[*v1alpha1.Content]
}
