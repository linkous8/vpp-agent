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

package vpp1904

import (
	interfaces "github.com/ligato/vpp-agent/api/models/vpp/interfaces"
	"github.com/ligato/vpp-agent/plugins/vpp/binapi/vpp1904/vhost_user"
)

// AddVhostUserInterface implements interface handler.
func (h *InterfaceVppHandler) AddVhostUserInterface(ifName string, vhstUsrIface *interfaces.VhostUserLink) (swIfIdx uint32, err error) {
	req := &vhost_user.CreateVhostUserIf{
		SockFilename: []byte(vhstUsrIface.SockFilename),
	}

	if vhstUsrIface.IsServer {
		req.IsServer = 1
	} else {
		req.IsServer = 0
	}

	if vhstUsrIface.DisableMrgRxbuf {
		req.DisableMrgRxbuf = 1
	} else {
		req.DisableMrgRxbuf = 0
	}

	if vhstUsrIface.DisableIndirectDesc {
		req.DisableIndirectDesc = 1
	} else {
		req.DisableIndirectDesc = 0
	}

	reply := &vhost_user.CreateVhostUserIfReply{}

	if err := h.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return 0, err
	}

	return reply.SwIfIndex, h.SetInterfaceTag(ifName, reply.SwIfIndex)
}

// DeleteVhostUserInterface implements interface handler.
func (h *InterfaceVppHandler) DeleteVhostUserInterface(ifName string, idx uint32) error {
	req := &vhost_user.DeleteVhostUserIf{
		SwIfIndex: idx,
	}
	reply := &vhost_user.DeleteVhostUserIfReply{}

	if err := h.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return err
	}

	return h.RemoveInterfaceTag(ifName, idx)
}
