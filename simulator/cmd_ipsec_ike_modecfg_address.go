package simulator

import "context"

var CmdIpsecIkeModeCfgAddress = &CommandSpec{
	[]Token{TIpsec, TIke, TModeCfg, TAddress}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeModeCfgAddress = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TModeCfg, TAddress}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

