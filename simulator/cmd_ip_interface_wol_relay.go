package simulator

import "context"

var CmdIpInterfaceWolRelay = &CommandSpec{
	[]Token{TIp, TL2Interface, TWol, TRelay}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceWolRelay = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TWol, TRelay}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

