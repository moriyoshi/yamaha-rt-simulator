package simulator

import "context"

var CmdIpsecIkeEspEncapsulation = &CommandSpec{
	[]Token{TIpsec, TIke, TEspEncapsulation}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeEspEncapsulation = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TEspEncapsulation}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

