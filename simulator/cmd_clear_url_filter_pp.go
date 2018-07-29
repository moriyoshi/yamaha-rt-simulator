package simulator

import "context"

var CmdClearUrlFilterPp = &CommandSpec{
	[]Token{TClear, TUrl, TFilter, TPp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
