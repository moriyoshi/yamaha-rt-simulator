package simulator

import "context"

var CmdBgpConfigureRefresh = &CommandSpec{
	[]Token{TBgp, TConfigure, TRefresh}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

