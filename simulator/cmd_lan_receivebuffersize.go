package simulator

import "context"

var CmdLanReceiveBufferSize = &CommandSpec{
	[]Token{TLan, TReceiveBufferSize}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoLanReceiveBufferSize = &CommandSpec{
	[]Token{TNo, TLan, TReceiveBufferSize}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
