package simulator

import "context"

var CmdIpsecIkeLicenseKey = &CommandSpec{
	[]Token{TIpsec, TIke, TLicenseKey}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeLicenseKey = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TLicenseKey}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

