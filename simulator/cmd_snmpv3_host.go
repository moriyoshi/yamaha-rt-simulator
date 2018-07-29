package simulator

import "context"

var CmdSnmpv3Host = &CommandSpec{
	[]Token{TSnmpv3, THost}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
