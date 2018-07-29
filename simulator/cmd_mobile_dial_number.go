package simulator

import "context"

var CmdMobileDialNumber = &CommandSpec{
	[]Token{TMobile, TDial, TNumber}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMobileDialNumberDial_String = &CommandSpec{
	[]Token{TNo, TMobile, TDial, TNumber}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

