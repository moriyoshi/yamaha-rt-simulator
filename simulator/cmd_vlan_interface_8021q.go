package simulator

import "context"

var CmdVlanInterface8021Q = &CommandSpec{
	[]Token{TVlan, TL2Interface, T8021Q, TVidArg, TNameArg}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoVlanInterface = &CommandSpec{
	[]Token{TNo, TVlan, TL2Interface, T8021Q}, 44,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
