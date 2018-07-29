package simulator

import "context"

var CmdIpsecIkeEapSendCertreq = &CommandSpec{
	[]Token{TIpsec, TIke, TEap, TSend, TCertreq}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeEapSendCertreq = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TEap, TSend, TCertreq}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

