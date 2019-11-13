//  Copyright (c) 2019 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package vpp1904_test

import (
	"testing"

	ifModel "github.com/ligato/vpp-agent/api/models/vpp/interfaces"
	"github.com/ligato/vpp-agent/plugins/vpp/binapi/vpp1904/interfaces"
	"github.com/ligato/vpp-agent/plugins/vpp/binapi/vpp1904/vhost_user"
	. "github.com/onsi/gomega"
)

func TestAddVhostUserInterface(t *testing.T) {
	ctx, ifHandler := ifTestSetup(t)
	defer ctx.TeardownTestCtx()

	ctx.MockVpp.MockReply(&vhost_user.CreateVhostUserIfReply{
		SwIfIndex: 1,
	})
	ctx.MockVpp.MockReply(&interfaces.SwInterfaceTagAddDelReply{})

	swIfIdx, err := ifHandler.AddVhostUserInterface("vhost_user", &ifModel.VhostUserLink{
		IsServer:     true,
		SockFilename: "filename",
	})
	Expect(err).To(BeNil())
	Expect(swIfIdx).To(BeEquivalentTo(1))
	var msgCheck bool
	for _, msg := range ctx.MockChannel.Msgs {
		vppMsg, ok := msg.(*vhost_user.CreateVhostUserIf)
		if ok {
			Expect(vppMsg.IsServer).To(BeEquivalentTo(1))
			Expect(vppMsg.SockFilename).To(BeEquivalentTo([]byte("filename")))
			Expect(vppMsg.Renumber).To(BeEquivalentTo(0))
			Expect(vppMsg.DisableMrgRxbuf).To(BeEquivalentTo(0))
			Expect(vppMsg.DisableIndirectDesc).To(BeEquivalentTo(0))
			Expect(vppMsg.CustomDevInstance).To(BeEquivalentTo(0))
			Expect(vppMsg.UseCustomMac).To(BeEquivalentTo(0))
			msgCheck = true
		}
	}
	Expect(msgCheck).To(BeTrue())
}

func TestDeleteVhostUserInterface(t *testing.T) {
	ctx, ifHandler := ifTestSetup(t)
	defer ctx.TeardownTestCtx()

	ctx.MockVpp.MockReply(&vhost_user.DeleteVhostUserIfReply{})
	ctx.MockVpp.MockReply(&interfaces.SwInterfaceTagAddDelReply{})

	err := ifHandler.DeleteTapInterface("vhost_user", 1)
	Expect(err).To(BeNil())
	var msgCheck bool
	for _, msg := range ctx.MockChannel.Msgs {
		vppMsg, ok := msg.(*vhost_user.DeleteVhostUserIf)
		if ok {
			Expect(vppMsg.SwIfIndex).To(BeEquivalentTo(1))
			msgCheck = true
		}
	}
	Expect(msgCheck).To(BeTrue())
}
