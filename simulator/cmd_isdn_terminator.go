package simulator

import "context"

var CmdIsdnTerminator = &CommandSpec{
	[]Token{TIsdn, TTerminator}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIsdnTerminator = &CommandSpec{
	[]Token{TNo, TIsdn, TTerminator}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

