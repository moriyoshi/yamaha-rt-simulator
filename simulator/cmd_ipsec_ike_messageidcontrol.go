package simulator

import "context"

var CmdIpsecIkeMessageIdControl = &CommandSpec{
	[]Token{TIpsec, TIke, TMessageIdControl}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeMessageIdControl = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TMessageIdControl}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

