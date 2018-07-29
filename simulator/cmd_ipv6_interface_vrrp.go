package simulator

import "context"

var CmdIpv6InterfaceVrrp = &CommandSpec{
	[]Token{TIpv6, TL2Interface, TVrrp}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6InterfaceVrrp = &CommandSpec{
	[]Token{TNo, TIpv6, TL2Interface, TVrrp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

