package simulator

import "context"

var CmdIpRouting = &CommandSpec{
	[]Token{TIp, TRouting}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpRouting = &CommandSpec{
	[]Token{TNo, TIp, TRouting}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

