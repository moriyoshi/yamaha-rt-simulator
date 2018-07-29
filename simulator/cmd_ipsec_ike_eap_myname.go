package simulator

import "context"

var CmdIpsecIkeEapMyname = &CommandSpec{
	[]Token{TIpsec, TIke, TEap, TMyname}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeEapMyname = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TEap, TMyname}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

