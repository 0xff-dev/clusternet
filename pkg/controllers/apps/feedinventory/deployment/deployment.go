/*
Copyright 2022 The Clusternet Authors.

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

package deployment

import (
	"encoding/json"

	k8sappsv1 "k8s.io/api/apps/v1"

	appsapi "github.com/clusternet/clusternet/pkg/apis/apps/v1alpha1"
	"github.com/clusternet/clusternet/pkg/controllers/apps/feedinventory/utils"
)

type Plugin struct {
	name string
}

func NewPlugin() *Plugin {
	return &Plugin{
		name: "deployment",
	}
}

func (pl *Plugin) Parser(rawData []byte) (*int32, appsapi.ReplicaRequirements, string, error) {
	// use apps/v1 for deserialization
	// all the fields that we needed are backward compatible with apps/v1beta1, apps/v1beta2, extensions/v1beta1
	deploy := &k8sappsv1.Deployment{}

	if err := json.Unmarshal(rawData, deploy); err != nil {
		return nil, appsapi.ReplicaRequirements{}, "", err
	}

	return deploy.Spec.Replicas, utils.GetReplicaRequirements(deploy.Spec.Template.Spec), "/spec/replicas", nil
}

func (pl *Plugin) Name() string {
	return pl.name
}

func (pl *Plugin) Kind() string {
	return "Deployment"
}