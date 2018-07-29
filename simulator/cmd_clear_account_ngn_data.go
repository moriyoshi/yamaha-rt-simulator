package simulator

import "context"

var CmdClearAccountNgnData = &CommandSpec{
	[]Token{TClear, TAccount, TNgn, TData}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

