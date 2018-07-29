package simulator

import "context"

var CmdShowStatusL2Tp = &CommandSpec{
	[]Token{TShow, TStatus, TL2Tp, TTunnel}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
