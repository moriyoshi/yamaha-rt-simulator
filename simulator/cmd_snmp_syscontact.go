package simulator

import "context"

var CmdSnmpSyscontact = &CommandSpec{
	[]Token{TSnmp, TSyscontact}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpSyscontact = &CommandSpec{
	[]Token{TNo, TSnmp, TSyscontact}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

