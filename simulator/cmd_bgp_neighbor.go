package simulator

import "context"

var CmdBgpNeighbor = &CommandSpec{
	[]Token{TBgp, TNeighbor}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoBgpNeighbor = &CommandSpec{
	[]Token{TNo, TBgp, TNeighbor}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

