package simulator

import "context"

var CmdIpPimSparseJoinPruneSend = &CommandSpec{
	[]Token{TIp, TPim, TSparse, TJoinPrune, TSend}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPimSparseJoinPruneSend = &CommandSpec{
	[]Token{TNo, TIp, TPim, TSparse, TJoinPrune, TSend}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

