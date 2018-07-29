package simulator

import "context"

var CmdIpInterfaceProxyarp = &CommandSpec{
	[]Token{TIp, TL2Interface, TProxyarp, TVrrp}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpInterfaceProxyarpVrrp = &CommandSpec{
	[]Token{TIp, TL2Interface, TProxyarp, TVrrp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceProxyarp = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TProxyarp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

