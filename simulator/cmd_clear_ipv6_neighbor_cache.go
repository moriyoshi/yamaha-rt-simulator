package simulator

import "context"

var CmdClearIpv6NeighborCache = &CommandSpec{
	[]Token{TClear, TIpv6, TNeighbor, TCache}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

