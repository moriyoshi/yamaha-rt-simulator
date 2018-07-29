package simulator

import "context"

var CmdSnmpv3ContextName = &CommandSpec{
	[]Token{TSnmpv3, TContext, TName}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpv3ContextName = &CommandSpec{
	[]Token{TNo, TSnmpv3, TContext, TName}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

