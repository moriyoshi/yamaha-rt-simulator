package simulator

import "context"

var CmdClearSwitchingHubMacaddress = &CommandSpec{
	[]Token{TClear, TSwitchingHub, TMacaddress}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

