package simulator

import "context"

var CmdSystemTemperatureThreshold = &CommandSpec{
	[]Token{TSystem, TTemperature, TThreshold}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSystemTemperatureThreshold = &CommandSpec{
	[]Token{TNo, TSystem, TTemperature, TThreshold}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

