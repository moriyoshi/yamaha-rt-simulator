package simulator

import "context"

var CmdShowIpv6NeighborCache = &CommandSpec{
	[]Token{TShow, TIpv6, TNeighbor, TCache}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

