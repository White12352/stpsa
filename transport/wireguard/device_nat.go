package wireguard

import (
	"github.com/sagernet/sing-tun"
	"github.com/sagernet/sing/common/buf"
)

var _ Device = (*natDeviceWrapper)(nil)

type natDeviceWrapper struct {
	Device
	outbound chan *buf.Buffer
	mapping  *tun.NatMapping
	writer   *tun.NatWriter
}

func NewNATDevice(upstream Device, ipRewrite bool) NatDevice {
	wrapper := &natDeviceWrapper{
		Device:   upstream,
		outbound: make(chan *buf.Buffer, 256),
		mapping:  tun.NewNatMapping(ipRewrite),
	}
	if ipRewrite {
		wrapper.writer = tun.NewNatWriter(upstream.Inet4Address(), upstream.Inet6Address())
	}
	return wrapper
}

func (d *natDeviceWrapper) Read(bufs [][]byte, sizes []int, offset int) (count int, err error) {
	select {
	case packet := <-d.outbound:
		defer packet.Release()
		sizes[0] = copy(bufs[0][offset:], packet.Bytes())
		return 1, nil
	default:
	}
	return d.Device.Read(bufs, sizes, offset)
}

func (d *natDeviceWrapper) Write(bufs [][]byte, offset int) (count int, err error) {
	packet := bufs[0][offset:]
	handled, err := d.mapping.WritePacket(packet)
	if handled {
		return 1, err
	}
	return d.Device.Write(bufs, offset)
}

func (d *natDeviceWrapper) CreateDestination(session tun.RouteSession, conn tun.RouteContext) tun.DirectDestination {
	d.mapping.CreateSession(session, conn)
	return &natDestinationWrapper{d, session}
}

var _ tun.DirectDestination = (*natDestinationWrapper)(nil)

type natDestinationWrapper struct {
	device  *natDeviceWrapper
	session tun.RouteSession
}

func (d *natDestinationWrapper) WritePacket(buffer *buf.Buffer) error {
	if d.device.writer != nil {
		d.device.writer.RewritePacket(buffer.Bytes())
	}
	d.device.outbound <- buffer
	return nil
}

func (d *natDestinationWrapper) Close() error {
	d.device.mapping.DeleteSession(d.session)
	return nil
}

func (d *natDestinationWrapper) Timeout() bool {
	return false
}
