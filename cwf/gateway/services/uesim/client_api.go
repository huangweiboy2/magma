/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

// Package hss provides a thin client for using the UESim service.
// This can be used by apps to discover and contact the service, without knowing about
// the RPC implementation.
package uesim

import (
	"context"
	"errors"
	"fmt"

	cwfprotos "magma/cwf/cloud/go/protos"
	"magma/cwf/gateway/registry"

	"github.com/golang/glog"
	"google.golang.org/grpc"
)

// Wrapper for GRPC Client functionality
type ueSimClient struct {
	cwfprotos.UESimClient
	cc *grpc.ClientConn
}

// getUESimClient is a utility function to get a RPC connection to the
// UESim service
func getUESimClient() (*ueSimClient, error) {
	conn, err := registry.GetConnection(registry.UeSim)
	if err != nil {
		errMsg := fmt.Sprintf("UESim client initialization error: %s", err)
		glog.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return &ueSimClient{
		cwfprotos.NewUESimClient(conn),
		conn,
	}, err
}

// AddUE tries to add this UE to the server.
// Input: The UE data which will be added.
func AddUE(ue *cwfprotos.UEConfig) error {
	cli, err := getUESimClient()
	if err != nil {
		return err
	}
	_, err = cli.AddUE(context.Background(), ue)
	return err
}

// Authenticate triggers an authentication for the UE with the specified IMSI.
// Input: The IMSI of the UE to try to authenticate.
// Output: The resulting Radius packet returned by the Radius server.
func Authenticate(id *cwfprotos.AuthenticateRequest) (*cwfprotos.AuthenticateResponse, error) {
	cli, err := getUESimClient()
	if err != nil {
		return nil, err
	}
	return cli.Authenticate(context.Background(), id)
}

// GenTraffic triggers traffic generation for the UE with the specified IMSI.
// Input: The IMSI of the UE to simulate traffic for
func GenTraffic(req *cwfprotos.GenTrafficRequest) error {
	cli, err := getUESimClient()
	if err != nil {
		return err
	}
	_, err = cli.GenTraffic(context.Background(), req)
	return err
}
