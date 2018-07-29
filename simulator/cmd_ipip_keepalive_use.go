package simulator

import "context"

var CmdIpipKeepaliveUse = &CommandSpec{
	[]Token{TIpip, TKeepalive, TUse}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpipKeepaliveUse = &CommandSpec{
	[]Token{TNo, TIpip, TKeepalive, TUse}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

