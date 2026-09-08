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

package main

import (
	"context"
	"time"
)

import (
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"dubbo.apache.org/dubbo-go/v3/protocol"
	"dubbo.apache.org/dubbo-go/v3/protocol/triple"
	"dubbo.apache.org/dubbo-go/v3/server"
	"dubbo.apache.org/dubbo-go/v3/tls"

	"github.com/dubbogo/gost/log/logger"
)

import (
	greet "github.com/apache/dubbo-go-samples/http3/proto"
)

// QUIC transport tuning values; unset options fall back to quic-go defaults.
const (
	quicKeepAlivePeriod                = 30 * time.Second
	quicMaxIdleTimeout                 = 90 * time.Second
	quicMaxIncomingStreams             = 1024
	quicMaxIncomingUniStreams          = 1024
	quicInitialStreamReceiveWindow     = 512 * 1024
	quicMaxStreamReceiveWindow         = 2 * 1024 * 1024
	quicInitialConnectionReceiveWindow = 2 * 1024 * 1024
	quicMaxConnectionReceiveWindow     = 8 * 1024 * 1024
)

type GreetTripleServer struct {
}

func (srv *GreetTripleServer) Greet(ctx context.Context, req *greet.GreetRequest) (*greet.GreetResponse, error) {
	resp := &greet.GreetResponse{Greeting: "Hello " + req.Name + ", from Go HTTP/3!"}
	return resp, nil
}

func main() {
	logger.SetLoggerLevel("debug")

	srv, err := server.NewServer(
		server.WithServerProtocol(
			protocol.WithPort(20000),
			protocol.WithTriple(
				triple.WithHttp3Enable(),
				triple.WithHttp3KeepAlivePeriod(quicKeepAlivePeriod),
				triple.WithHttp3MaxIdleTimeout(quicMaxIdleTimeout),
				triple.WithHttp3MaxIncomingStreams(quicMaxIncomingStreams),
				triple.WithHttp3MaxIncomingUniStreams(quicMaxIncomingUniStreams),
				triple.WithHttp3InitialStreamReceiveWindow(quicInitialStreamReceiveWindow),
				triple.WithHttp3MaxStreamReceiveWindow(quicMaxStreamReceiveWindow),
				triple.WithHttp3InitialConnectionReceiveWindow(quicInitialConnectionReceiveWindow),
				triple.WithHttp3MaxConnectionReceiveWindow(quicMaxConnectionReceiveWindow),
			),
		),
		server.WithServerTLSOption(
			tls.WithCertFile("../../x509/server2_cert.pem"),
			tls.WithKeyFile("../../x509/server2_key_pkcs8.pem"),
			tls.WithServerName("dubbogo.test.example.com"),
		),
	)
	if err != nil {
		panic(err)
	}

	if err := greet.RegisterGreetServiceHandler(srv, &GreetTripleServer{}); err != nil {
		panic(err)
	}

	logger.Info("Starting HTTP/3 enabled Dubbo-go server on port 20000...")
	if err := srv.Serve(); err != nil {
		panic(err)
	}
}
