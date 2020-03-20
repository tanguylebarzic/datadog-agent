// +build linux_bpf

package network

// EnabledKProbes returns a map of kprobes that are enabled per config settings.
// This map does not include the probes used exclusively in the offset guessing process.
func EnabledKProbes(c *Config, pre410Kernel bool) map[KProbeName]struct{} {
	enabled := make(map[KProbeName]struct{}, 0)

	if c.CollectTCPConns {
		if pre410Kernel {
			enabled[TCPSendMsgPre410] = struct{}{}
		} else {
			enabled[TCPSendMsg] = struct{}{}
		}
		enabled[TCPCleanupRBuf] = struct{}{}
		enabled[TCPClose] = struct{}{}
		enabled[TCPRetransmit] = struct{}{}
		enabled[InetCskAcceptReturn] = struct{}{}
		enabled[TCPv4DestroySock] = struct{}{}

		if c.BPFDebug {
			enabled[TCPSendMsgReturn] = struct{}{}
		}
	}

	if c.CollectUDPConns {
		enabled[UDPRecvMsgReturn] = struct{}{}
		if pre410Kernel {
			enabled[UDPSendMsgPre410] = struct{}{}
			enabled[UDPRecvMsgPre410] = struct{}{}
		} else {
			enabled[UDPRecvMsg] = struct{}{}
			enabled[UDPSendMsg] = struct{}{}
		}

	}

	return enabled
}
