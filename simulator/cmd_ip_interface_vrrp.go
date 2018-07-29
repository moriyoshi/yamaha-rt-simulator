package simulator

import "context"

var CmdIpInterfaceVrrp = &CommandSpec{
	[]Token{TIp, TL2Interface, TVrrp}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceVrrp = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TVrrp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

