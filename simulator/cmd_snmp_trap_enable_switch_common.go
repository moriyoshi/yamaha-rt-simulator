package simulator

import "context"

var CmdSnmpTrapEnableSwitchCommon = &CommandSpec{
	[]Token{TSnmp, TTrap, TEnable, TSwitch, TCommon}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpTrapEnableSwitchCommon = &CommandSpec{
	[]Token{TNo, TSnmp, TTrap, TEnable, TSwitch, TCommon}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
