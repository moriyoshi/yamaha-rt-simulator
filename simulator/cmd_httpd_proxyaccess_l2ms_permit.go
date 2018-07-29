package simulator

import "context"

var CmdHttpdProxyAccessL2MsPermit = &CommandSpec{
	[]Token{THttpd, TProxyAccess, TL2Ms, TPermit}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHttpdProxyAccessL2MsPermit = &CommandSpec{
	[]Token{TNo, THttpd, TProxyAccess, TL2Ms, TPermit}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

