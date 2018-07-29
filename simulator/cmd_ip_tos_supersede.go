package simulator

import "context"

var CmdIpTosSupersede = &CommandSpec{
	[]Token{TIp, TTos, TSupersede}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTosSupersede = &CommandSpec{
	[]Token{TNo, TIp, TTos, TSupersede}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

