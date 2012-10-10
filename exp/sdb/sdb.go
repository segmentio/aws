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
)

// SimpleDB represents an authenticated connection to a particular endpoint the
// SimpleDB service.
type SimpleDB interface {
	// Return the domain with the specified name. The domain must have previously
	// been created on the service.
	OpenDomain(name string) (Domain, error)

	// Create a domain with the supplied name on the service.
	CreateDomain(name string) error

	// Delete the domain with the supplied name from the service.
	DeleteDomain(name string) error

	// Retrieve a set of items and their attributes based on a query string.
	//
	// constistentRead specifies whether completely fresh data is needed or not.
	//
	// If the selected result set is too large, a "next token" will be returned.
	// It can be passed to the Select method to resume where the previous result
	// set left off. For the initial query, use nil.
	//
	// For more info:
	//
	//     http://goo.gl/GTsSZ
	//
	Select(
		query string,
		constistentRead bool,
		nextToken []byte) (res map[ItemName][]Attribute, tok []byte, err error)
}

// Return a SimpleDB connection tied to the given region, using the sipplied
// access key to authenticate requests.
func NewSimpleDB(region Region, key aws.AccessKey) (db SimpleDB, err error)
