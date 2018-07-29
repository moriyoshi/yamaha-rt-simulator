package simulator

import "context"

var CmdIpsecIkeLocalName = &CommandSpec{
	[]Token{TIpsec, TIke, TLocal, TName}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeLocalName = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TLocal, TName}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

