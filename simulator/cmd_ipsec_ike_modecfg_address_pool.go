package simulator

import "context"

var CmdIpsecIkeModeCfgAddressPool = &CommandSpec{
	[]Token{TIpsec, TIke, TModeCfg, TAddress, TPool}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeModeCfgAddressPool = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TModeCfg, TAddress, TPool}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
