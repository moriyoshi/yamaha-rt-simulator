package simulator

import "context"

var CmdIpInterfaceArpStatic = &CommandSpec{
	[]Token{TIp, TL2Interface, TArp, TStatic}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceArpStatic = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TArp, TStatic}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

