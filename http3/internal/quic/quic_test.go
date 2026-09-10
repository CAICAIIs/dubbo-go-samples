/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package quic

import (
	"strconv"
	"testing"
)

import (
	"dubbo.apache.org/dubbo-go/v3/protocol/triple"
)

func TestOptions(t *testing.T) {
	h3 := triple.NewOptions(Options()...).Triple.Http3
	if h3 == nil {
		t.Fatal("Http3 config is nil")
	}

	if !h3.Enable {
		t.Errorf("Enable = false, want true")
	}
	if h3.KeepAlivePeriod != keepAlivePeriod.String() {
		t.Errorf("KeepAlivePeriod = %q, want %q", h3.KeepAlivePeriod, keepAlivePeriod.String())
	}
	if h3.MaxIdleTimeout != maxIdleTimeout.String() {
		t.Errorf("MaxIdleTimeout = %q, want %q", h3.MaxIdleTimeout, maxIdleTimeout.String())
	}
	if h3.MaxIncomingStreams != maxIncomingStreams {
		t.Errorf("MaxIncomingStreams = %d, want %d", h3.MaxIncomingStreams, maxIncomingStreams)
	}
	if h3.MaxIncomingUniStreams != maxIncomingUniStreams {
		t.Errorf("MaxIncomingUniStreams = %d, want %d", h3.MaxIncomingUniStreams, maxIncomingUniStreams)
	}
	if h3.InitialStreamReceiveWindow != strconv.FormatUint(uint64(initialStreamReceiveWindow), 10) {
		t.Errorf("InitialStreamReceiveWindow = %q, want %q",
			h3.InitialStreamReceiveWindow, strconv.FormatUint(uint64(initialStreamReceiveWindow), 10))
	}
	if h3.MaxStreamReceiveWindow != strconv.FormatUint(uint64(maxStreamReceiveWindow), 10) {
		t.Errorf("MaxStreamReceiveWindow = %q, want %q",
			h3.MaxStreamReceiveWindow, strconv.FormatUint(uint64(maxStreamReceiveWindow), 10))
	}
	if h3.InitialConnectionReceiveWindow != strconv.FormatUint(uint64(initialConnectionReceiveWindow), 10) {
		t.Errorf("InitialConnectionReceiveWindow = %q, want %q",
			h3.InitialConnectionReceiveWindow, strconv.FormatUint(uint64(initialConnectionReceiveWindow), 10))
	}
	if h3.MaxConnectionReceiveWindow != strconv.FormatUint(uint64(maxConnectionReceiveWindow), 10) {
		t.Errorf("MaxConnectionReceiveWindow = %q, want %q",
			h3.MaxConnectionReceiveWindow, strconv.FormatUint(uint64(maxConnectionReceiveWindow), 10))
	}
}
