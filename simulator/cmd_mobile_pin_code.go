package simulator

import "context"

var CmdMobilePinCode = &CommandSpec{
	[]Token{TMobile, TPin, TCode}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMobilePinCode = &CommandSpec{
	[]Token{TNo, TMobile, TPin, TCode}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

