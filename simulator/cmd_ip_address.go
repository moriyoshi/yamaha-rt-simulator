package simulator

import (
	"context"
	"net"
)

func defaultBroadcastAddr(net_ net.IPNet) net.IP {
	retval := make(net.IP, len(net_.IP))
	for i := 0; i < len(net_.Mask); i++ {
		retval[i] = net_.IP[i] | ^net_.Mask[i]
	}
	return retval
}

var CmdIPIfAddress = &CommandSpec{
	[]Token{TIp, TL2Interface, TAddress, &EitherToken{TDhcp, TSomething}, TBroadcast}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		if len(tis) < 4 {
			return WrongNumberOfArgs
		}
		var v4dhcp bool
		var v4addr net.IP
		var v4net = &net.IPNet{}
		var v4broadcastAddr net.IP
		var err error
		if tis[3].Image == "dhcp" {
			if len(tis) != 4 {
				return InvalidCommandName
			}
			v4dhcp = true
		} else {
			v4addr, v4net, err = net.ParseCIDR(tis[3].Image)
			if err != nil {
				return UnrecognizedIPAddr
			}
			if len(tis) == 6 && tis[4].Image == "broadcast" {
				v4broadcastAddr = net.ParseIP(tis[5].Image)
				if v4broadcastAddr == nil {
					return UnrecognizedIPAddr
				}
			} else if len(tis) == 4 {
				v4broadcastAddr = defaultBroadcastAddr(*v4net)
			} else {
				return UnrecognizedIPAddr
			}
		}
		if_ := sess.ResolveInterface(tis[1].Image)
		if if_ == nil {
			return InvalidCommandName
		}
		sess.RunningConfig.Lock()
		defer sess.RunningConfig.Unlock()
		if sess.RunningConfig.Interfaces == nil {
			sess.RunningConfig.Interfaces = make(map[NetInterface]*InterfaceConfig)
		}
		ifConfig := sess.RunningConfig.Interfaces[if_]
		if ifConfig == nil {
			ifConfig = &InterfaceConfig{}
			sess.RunningConfig.Interfaces[if_] = ifConfig
		}
		ifConfig.Lock()
		defer ifConfig.Unlock()
		ifConfig.V4DHCP = v4dhcp
		ifConfig.V4Address = v4addr
		ifConfig.V4Network = *v4net
		ifConfig.V4BroadcastAddr = v4broadcastAddr
		return nil
	},
}
