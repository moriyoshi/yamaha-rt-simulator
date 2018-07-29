package simulator

import "context"

var CmdSnmpv3VacmView = &CommandSpec{
	[]Token{TSnmpv3, TVacm, TView}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpv3VacmView = &CommandSpec{
	[]Token{TNo, TSnmpv3, TVacm, TView}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

