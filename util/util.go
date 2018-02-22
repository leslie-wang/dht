package util

import "net"

// CompactPeer is to wrapper for anacrolix/torrent/util
type CompactPeer struct {
	IP   net.IP
	Port int
}
