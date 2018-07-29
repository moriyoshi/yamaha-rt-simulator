package simulator

import "context"

var CmdIpPimSparseRendezvousPointStatic = &CommandSpec{
	[]Token{TIp, TPim, TSparse, TRendezvousPoint, TStatic}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPimSparseRendezvousPointStatic = &CommandSpec{
	[]Token{TNo, TIp, TPim, TSparse, TRendezvousPoint, TStatic}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

