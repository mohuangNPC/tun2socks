//go:build unix

package engine

import (
	"net/url"

	"github.com/mohuangNPC/tun2socks/v2/core/device"
	"github.com/mohuangNPC/tun2socks/v2/core/device/tun"
)

func parseTUN(u *url.URL, mtu uint32) (device.Device, error) {
	return tun.Open(u.Host, mtu)
}
