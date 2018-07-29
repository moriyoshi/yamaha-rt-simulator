package simulator

import "context"

var CmdSnmpSyslocation = &CommandSpec{
	[]Token{TSnmp, TSyslocation}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpSyslocation = &CommandSpec{
	[]Token{TNo, TSnmp, TSyslocation}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

