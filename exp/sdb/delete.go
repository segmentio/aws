// Copyright 2012 Aaron Jacobs. All Rights Reserved.
// Author: aaronjjacobs@gmail.com (Aaron Jacobs)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sdb

import (
	"fmt"
)

func (d *domain) DeleteAttributes(
	item ItemName,
	deletes []DeleteUpdate,
	preconditions []Precondition) error {
	return fmt.Errorf("TODO")
}

func (d *domain) BatchDeleteAttributes(deletes map[ItemName][]DeleteUpdate) error {
	return fmt.Errorf("TODO")
}