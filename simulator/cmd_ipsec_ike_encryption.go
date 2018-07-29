package simulator

import "context"

var CmdIpsecIkeEncryption = &CommandSpec{
	[]Token{TIpsec, TIke, TEncryption}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeEncryption = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TEncryption}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

