package simulator

import "context"

var CmdIpsecIkeLocalAddress = &CommandSpec{
	[]Token{TIpsec, TIke, TLocal, TAddress}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeLocalAddress = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TLocal, TAddress}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
