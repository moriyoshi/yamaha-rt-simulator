package simulator

import "context"

var CmdBgpNeighborPreSharedKey = &CommandSpec{
	[]Token{TBgp, TNeighbor, TPreSharedKey}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoBgpNeighborPreSharedKey = &CommandSpec{
	[]Token{TNo, TBgp, TNeighbor, TPreSharedKey}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

