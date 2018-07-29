package simulator

import "context"

var CmdIpsecIkeEapRequest = &CommandSpec{
	[]Token{TIpsec, TIke, TEap, TRequest}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeEapRequest = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TEap, TRequest}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

