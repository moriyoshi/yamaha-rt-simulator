package simulator

import "context"

var CmdSnmpTrapEnableSnmpTrap = &CommandSpec{
	[]Token{TSnmp, TTrap, TEnable, TSnmp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdSnmpTrapEnableSnmp = &CommandSpec{
	[]Token{TSnmp, TTrap, TEnable, TSnmp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpTrapEnableSnmp = &CommandSpec{
	[]Token{TNo, TSnmp, TTrap, TEnable, TSnmp}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

