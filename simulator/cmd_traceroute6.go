package simulator

import "context"

var CmdTraceroute6 = &CommandSpec{
	[]Token{TTraceroute6}, 1,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

