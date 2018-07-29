package simulator

import "context"

var CmdIpsecIkeNatTraversal = &CommandSpec{
	[]Token{TIpsec, TIke, TNatTraversal}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeNatTraversal = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TNatTraversal}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

