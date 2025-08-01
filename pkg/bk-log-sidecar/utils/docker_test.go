// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 日志平台 (BlueKing - Log) available.
// Copyright (C) 2017-2021 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.
//

package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDockerVersion(t *testing.T) {
	apiVersions := []struct {
		Match bool
		Api   string
	}{
		{
			true,
			"1.39",
		},
		{
			false,
			"",
		},
		{
			true,
			"1.30",
		},
		{
			true,
			"1.40",
		},
	}
	for _, apiVersion := range apiVersions {
		if apiVersion.Match {
			result_match, result_api := ExtractDockerApiVersion(fmt.Errorf("client version 1.40 is too new. Maximum supported API version is %s", apiVersion.Api))
			assert.Equal(t, apiVersion.Match, result_match)
			assert.Equal(t, apiVersion.Api, result_api)
			continue
		}
		result_match, result_api := ExtractDockerApiVersion(fmt.Errorf("%s", apiVersion.Api))
		assert.Equal(t, apiVersion.Match, result_match)
		assert.Equal(t, apiVersion.Api, result_api)
	}
}
