// Copyright 2016-2018 Authors of Cilium
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

package kvstore

import (
	"time"
)

var (
	// LockLeaseTTL is the time-to-live of the lease dedicated for locks
	LockLeaseTTL = 25 * time.Second

	// RetryInterval is the interval in which retries occur in the case of
	// errors in communication with the KVstore. This should be set to a
	// small value to account for temporary errors while communicating with
	// the KVstore.
	RetryInterval = 1 * time.Minute
)
