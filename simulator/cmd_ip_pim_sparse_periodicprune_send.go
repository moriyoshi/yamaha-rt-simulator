package simulator

import "context"

var CmdIpPimSparsePeriodicPruneSend = &CommandSpec{
	[]Token{TIp, TPim, TSparse, TPeriodicPrune, TSend}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPimSparsePeriodicPruneSend = &CommandSpec{
	[]Token{TNo, TIp, TPim, TSparse, TPeriodicPrune, TSend}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

