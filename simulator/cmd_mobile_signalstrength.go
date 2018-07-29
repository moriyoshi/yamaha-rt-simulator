package simulator

import "context"

var CmdMobileSignalStrength = &CommandSpec{
	[]Token{TMobile, TSignalStrength}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMobileSignalStrength = &CommandSpec{
	[]Token{TNo, TMobile, TSignalStrength}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

