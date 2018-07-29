package simulator

import "context"

var CmdDhcpManualLease = &CommandSpec{
	[]Token{TDhcp, TManual, TLease}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
