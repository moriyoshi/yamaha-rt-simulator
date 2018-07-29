package simulator

import "context"

var CmdSnmpv3EngineId = &CommandSpec{
	[]Token{TSnmpv3, TEngine, TId}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpv3EngineId = &CommandSpec{
	[]Token{TNo, TSnmpv3, TEngine, TId}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
