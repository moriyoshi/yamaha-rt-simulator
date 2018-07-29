package simulator

import "context"

var CmdSshdListen = &CommandSpec{
	[]Token{TSshd, TListen}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSshdListenPort = &CommandSpec{
	[]Token{TNo, TSshd, TListen}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

