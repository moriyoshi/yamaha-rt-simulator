package simulator

import "context"

var CmdIpsecIkeLicenseKeyUse = &CommandSpec{
	[]Token{TIpsec, TIke, TLicenseKey, TUse}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeLicenseKeyUse = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TLicenseKey, TUse}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

