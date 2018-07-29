package simulator

import "context"

var CmdSystemLedBrightness = &CommandSpec{
	[]Token{TSystem, TLed, TBrightness}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSystemLedBrightness = &CommandSpec{
	[]Token{TNo, TSystem, TLed, TBrightness}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

