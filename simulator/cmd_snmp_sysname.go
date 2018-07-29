package simulator

import "context"

var CmdSnmpSysname = &CommandSpec{
	[]Token{TSnmp, TSysname}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpSysname = &CommandSpec{
	[]Token{TNo, TSnmp, TSysname}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

