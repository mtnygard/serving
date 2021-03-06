/*
Copyright 2018 The Knative Authors

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

package names

import (
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/knative/serving/pkg/apis/serving/v1alpha1"
)

func TestNamer(t *testing.T) {
	tests := []struct {
		name  string
		route *v1alpha1.Route
		f     func(*v1alpha1.Route) string
		want  string
	}{{
		name: "K8sService",
		route: &v1alpha1.Route{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "blah",
				Namespace: "default",
			},
		},
		f:    K8sService,
		want: "blah",
	}, {
		name: "K8sServiceFullname",
		route: &v1alpha1.Route{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "bar",
				Namespace: "default",
			},
		},
		f:    K8sServiceFullname,
		want: "bar.default.svc.cluster.local",
	}, {
		name: "ClusterIngressPrefix",
		route: &v1alpha1.Route{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "bar",
				Namespace: "default",
				UID:       "1234-5678-910",
			},
		},
		f:    ClusterIngress,
		want: "route-1234-5678-910",
	}, {
		name: "Certificate",
		route: &v1alpha1.Route{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "bar",
				Namespace: "default",
				UID:       "1234-5678-910",
			},
		},
		f:    Certificate,
		want: "route-1234-5678-910",
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.f(test.route)
			if got != test.want {
				t.Errorf("%s() = %v, wanted %v", test.name, got, test.want)
			}
		})
	}
}
