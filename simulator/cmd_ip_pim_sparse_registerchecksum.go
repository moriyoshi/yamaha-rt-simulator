package simulator

import "context"

var CmdIpPimSparseRegisterChecksum = &CommandSpec{
	[]Token{TIp, TPim, TSparse, TRegisterChecksum}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPimSparseRegisterChecksum = &CommandSpec{
	[]Token{TNo, TIp, TPim, TSparse, TRegisterChecksum}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

