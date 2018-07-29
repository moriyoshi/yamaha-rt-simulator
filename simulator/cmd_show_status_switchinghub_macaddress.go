package simulator

import "context"

var CmdShowStatusSwitchingHubMacaddress = &CommandSpec{
	[]Token{TShow, TStatus, TSwitchingHub, TMacaddress}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
