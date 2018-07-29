package simulator

import "context"

var CmdDhcpManualRelease = &CommandSpec{
	[]Token{TDhcp, TManual, TRelease}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

