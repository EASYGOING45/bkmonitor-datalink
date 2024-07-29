// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package daemon

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Demo struct {
	RootCtx context.Context //
}

func NewDemoTask(ctx context.Context) *Demo {
	return &Demo{RootCtx: ctx}
}

type Payload struct {
	Content string `json:"content"`
}

func (d *Demo) Start(runInstanceCtx context.Context, errorReceiveChan chan<- error, payload []byte) {
	var p Payload
	var temp = json.Unmarshal(payload, &p)
	if temp != nil {
		errorReceiveChan <- errors.New("failed to parse payload")
		return
	}

	//content := string(payload)
	fmt.Printf("start task instance, payload: %s \n", p.Content)
	if p.Content != "ok" {
		errorReceiveChan <- errors.New("receive abnormal payload")
		return
	}

	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case currentTime := <-ticker.C:
			fmt.Println("当前时间: ", currentTime.Format("2006-01-02 15:04:05"))
		case <-runInstanceCtx.Done():
			fmt.Println("receive task instance done, return")
			return
		}
	}
}

func (t *Demo) GetTaskDimension(payload []byte) string {
	return "demo_task"
}
