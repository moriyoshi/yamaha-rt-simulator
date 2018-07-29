package simulator

import "context"

var CmdIpFilterSet = &CommandSpec{
	[]Token{TIp, TFilter, TSet}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpFilterSetNameDirection = &CommandSpec{
	[]Token{TNo, TIp, TFilter, TSet}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

