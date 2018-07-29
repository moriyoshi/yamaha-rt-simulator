package simulator

import "context"

var CmdDhcpDuplicateCheck = &CommandSpec{
	[]Token{TDhcp, TDuplicate, TCheck}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDhcpDuplicateCheck = &CommandSpec{
	[]Token{TNo, TDhcp, TDuplicate, TCheck}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

