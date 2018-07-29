package simulator

import "context"

var CmdSnmpTrapLinkUpdownSeparateL2SwitchPort = &CommandSpec{
	[]Token{TSnmp, TTrap, TLinkUpdown, TSeparateL2SwitchPort}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpTrapLinkUpdownSeparateL2SwitchPort = &CommandSpec{
	[]Token{TNo, TSnmp, TTrap, TLinkUpdown, TSeparateL2SwitchPort}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

