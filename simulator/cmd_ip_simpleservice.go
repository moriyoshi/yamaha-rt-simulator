package simulator

import "context"

var CmdIpSimpleService = &CommandSpec{
	[]Token{TIp, TSimpleService}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpSimpleService = &CommandSpec{
	[]Token{TNo, TIp, TSimpleService}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

