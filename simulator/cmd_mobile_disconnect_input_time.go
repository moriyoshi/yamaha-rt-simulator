package simulator

import "context"

var CmdMobileDisconnectInputTime = &CommandSpec{
	[]Token{TMobile, TDisconnect, TInput, TTime}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMobileDisconnectInputTime = &CommandSpec{
	[]Token{TNo, TMobile, TDisconnect, TInput, TTime}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

