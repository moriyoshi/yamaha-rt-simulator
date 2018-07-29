package simulator

import "context"

var CmdIpInterfaceSecondaryAddress = &CommandSpec{
	[]Token{TIp, TL2Interface, TSecondary, TAddress}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceSecondaryAddress = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TSecondary, TAddress}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
