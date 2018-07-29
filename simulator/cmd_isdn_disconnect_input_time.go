package simulator

import "context"

var CmdIsdnDisconnectInputTime = &CommandSpec{
	[]Token{TIsdn, TDisconnect, TInput, TTime}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIsdnDisconnectInputTime = &CommandSpec{
	[]Token{TNo, TIsdn, TDisconnect, TInput, TTime}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

