package simulator

import "context"

var CmdIpPimSparseLog = &CommandSpec{
	[]Token{TIp, TPim, TSparse, TLog}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPimSparseLog = &CommandSpec{
	[]Token{TNo, TIp, TPim, TSparse, TLog}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

