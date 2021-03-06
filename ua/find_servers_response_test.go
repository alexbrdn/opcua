// Copyright 2018-2019 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ua

import (
	"testing"
	"time"
)

func TestFindServersResponse(t *testing.T) {
	cases := []CodecTestCase{
		{
			Name: "normal",
			Struct: &FindServersResponse{
				ResponseHeader: &ResponseHeader{
					Timestamp:          time.Date(2018, time.August, 10, 23, 0, 0, 0, time.UTC),
					RequestHandle:      1,
					ServiceDiagnostics: &DiagnosticInfo{},
					StringTable:        []string{},
					AdditionalHeader:   NewExtensionObject(nil),
				},
				Servers: []*ApplicationDescription{
					&ApplicationDescription{
						ApplicationURI: "app-uri",
						ProductURI:     "prod-uri",
						ApplicationName: &LocalizedText{
							EncodingMask: LocalizedTextText,
							Text:         "app-name",
						},
						ApplicationType:     ApplicationTypeServer,
						GatewayServerURI:    "gw-uri",
						DiscoveryProfileURI: "prof-uri",
						DiscoveryURLs:       []string{"discov-uri-1", "discov-uri-2"},
					},
					&ApplicationDescription{
						ApplicationURI: "app-uri",
						ProductURI:     "prod-uri",
						ApplicationName: &LocalizedText{
							EncodingMask: LocalizedTextText,
							Text:         "app-name",
						},
						ApplicationType:     ApplicationTypeServer,
						GatewayServerURI:    "gw-uri",
						DiscoveryProfileURI: "prof-uri",
						DiscoveryURLs:       []string{"discov-uri-1", "discov-uri-2"},
					},
				},
			},
			Bytes: []byte{
				// Timestamp
				0x00, 0x98, 0x67, 0xdd, 0xfd, 0x30, 0xd4, 0x01,
				// RequestHandle
				0x01, 0x00, 0x00, 0x00,
				// ServiceResult
				0x00, 0x00, 0x00, 0x00,
				// ServiceDiagnostics
				0x00,
				// StringTable
				0x00, 0x00, 0x00, 0x00,
				// AdditionalHeader
				0x00, 0x00, 0x00,
				// Servers
				// ArraySize
				0x02, 0x00, 0x00, 0x00,
				// ApplicationURI
				0x07, 0x00, 0x00, 0x00, 0x61, 0x70, 0x70, 0x2d, 0x75, 0x72, 0x69,
				// ProductURI
				0x08, 0x00, 0x00, 0x00, 0x70, 0x72, 0x6f, 0x64, 0x2d, 0x75, 0x72, 0x69,
				// ApplicationName
				0x02, 0x08, 0x00, 0x00, 0x00, 0x61, 0x70, 0x70, 0x2d,
				0x6e, 0x61, 0x6d, 0x65,
				// ApplicationType
				0x00, 0x00, 0x00, 0x00,
				// GatewayServerURI
				0x06, 0x00, 0x00, 0x00, 0x67, 0x77, 0x2d, 0x75, 0x72, 0x69,
				// DiscoveryProfileURI
				0x08, 0x00, 0x00, 0x00, 0x70, 0x72, 0x6f, 0x66, 0x2d, 0x75, 0x72, 0x69,
				// DiscoveryURIs
				0x02, 0x00, 0x00, 0x00,
				0x0c, 0x00, 0x00, 0x00, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x2d, 0x75, 0x72, 0x69, 0x2d, 0x31,
				0x0c, 0x00, 0x00, 0x00, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x2d, 0x75, 0x72, 0x69, 0x2d, 0x32,
				// ApplicationURI
				0x07, 0x00, 0x00, 0x00, 0x61, 0x70, 0x70, 0x2d, 0x75, 0x72, 0x69,
				// ProductURI
				0x08, 0x00, 0x00, 0x00, 0x70, 0x72, 0x6f, 0x64, 0x2d, 0x75, 0x72, 0x69,
				// ApplicationName
				0x02, 0x08, 0x00, 0x00, 0x00, 0x61, 0x70, 0x70, 0x2d,
				0x6e, 0x61, 0x6d, 0x65,
				// ApplicationType
				0x00, 0x00, 0x00, 0x00,
				// GatewayServerURI
				0x06, 0x00, 0x00, 0x00, 0x67, 0x77, 0x2d, 0x75, 0x72, 0x69,
				// DiscoveryProfileURI
				0x08, 0x00, 0x00, 0x00, 0x70, 0x72, 0x6f, 0x66, 0x2d, 0x75, 0x72, 0x69,
				// DiscoveryURIs
				0x02, 0x00, 0x00, 0x00,
				0x0c, 0x00, 0x00, 0x00, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x2d, 0x75, 0x72, 0x69, 0x2d, 0x31,
				0x0c, 0x00, 0x00, 0x00, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x2d, 0x75, 0x72, 0x69, 0x2d, 0x32,
			},
		},
	}
	RunCodecTest(t, cases)
}
