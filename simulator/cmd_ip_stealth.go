package simulator

import "context"

var CmdIpStealth = &CommandSpec{
	[]Token{TIp, TStealth}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpStealth = &CommandSpec{
	[]Token{TNo, TIp, TStealth}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
