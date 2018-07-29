package simulator

import "context"

var CmdTraceroute = &CommandSpec{
	[]Token{TTraceroute}, 1,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

