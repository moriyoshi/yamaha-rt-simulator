package simulator

import "context"

var CmdMobileSignalStrengthGo = &CommandSpec{
	[]Token{TMobile, TSignalStrength, TGo}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

