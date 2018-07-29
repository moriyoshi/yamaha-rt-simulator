package simulator

import "context"

var CmdNslookup = &CommandSpec{
	[]Token{TNslookup}, 1,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

