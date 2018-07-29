package simulator

import "context"

var CmdIpsecIkePkiFile = &CommandSpec{
	[]Token{TIpsec, TIke, TPki, TFile, TSomething /* gateway_id */, TCertificate, TSomething /* cert_id */, TCrl}, 8,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkePkiFile = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TPki, TFile}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

