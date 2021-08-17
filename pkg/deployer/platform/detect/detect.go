/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2021 Red Hat, Inc.
 */

package detect

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/api/errors"

	"github.com/k8stopologyawareschedwg/deployer/pkg/clientutil"
	"github.com/k8stopologyawareschedwg/deployer/pkg/deployer/platform"
)

func Detect() (platform.Platform, error) {
	ocpCli, err := clientutil.NewOCPClientSet()
	if err != nil {
		return platform.Unknown, err
	}
	sccs, err := ocpCli.SecurityV1.SecurityContextConstraints().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return platform.Kubernetes, nil
		}
		return platform.Unknown, err
	}
	if len(sccs.Items) > 0 {
		return platform.OpenShift, nil
	}
	return platform.Kubernetes, nil
}
