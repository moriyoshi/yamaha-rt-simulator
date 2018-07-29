package simulator

import "context"

var CmdSipServerCallRemoteDomain = &CommandSpec{
	[]Token{TSip, TServer, TCall, TRemote, TDomain}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSipServerCallRemoteDomain = &CommandSpec{
	[]Token{TNo, TSip, TServer, TCall, TRemote, TDomain}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
