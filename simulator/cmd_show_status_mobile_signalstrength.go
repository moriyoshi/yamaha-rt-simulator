package simulator

import "context"

var CmdShowStatusMobileSignalStrength = &CommandSpec{
	[]Token{TShow, TStatus, TMobile, TSignalStrength}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

