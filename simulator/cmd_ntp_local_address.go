package simulator

import "context"

var CmdNtpLocalAddress = &CommandSpec{
	[]Token{TNtp, TLocal, TAddress}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNtpLocalAddress = &CommandSpec{
	[]Token{TNo, TNtp, TLocal, TAddress}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

