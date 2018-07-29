package simulator

import "context"

var CmdSnmpv3VacmAccess = &CommandSpec{
	[]Token{TSnmpv3, TVacm, TAccess, TSomething /* group_id */, TRead, TSomething /* read_view */, TWrite}, 7,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpv3VacmAccess = &CommandSpec{
	[]Token{TNo, TSnmpv3, TVacm, TAccess}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

