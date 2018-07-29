package simulator

import "context"

var CmdSnmpTrapEnableSwitch = &CommandSpec{
	[]Token{TSnmp, TTrap, TEnable, TSwitch}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpTrapEnableSwitch = &CommandSpec{
	[]Token{TNo, TSnmp, TTrap, TEnable, TSwitch}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
