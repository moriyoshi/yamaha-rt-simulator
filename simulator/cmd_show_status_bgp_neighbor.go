package simulator

import "context"

var CmdShowStatusBgpNeighbor = &CommandSpec{
	[]Token{TShow, TStatus, TBgp, TNeighbor}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
