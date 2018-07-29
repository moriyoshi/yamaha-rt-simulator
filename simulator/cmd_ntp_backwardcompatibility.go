package simulator

import "context"

var CmdNtpBackwardCompatibility = &CommandSpec{
	[]Token{TNtp, TBackwardCompatibility}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNtpBackwardCompatibility = &CommandSpec{
	[]Token{TNo, TNtp, TBackwardCompatibility}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

