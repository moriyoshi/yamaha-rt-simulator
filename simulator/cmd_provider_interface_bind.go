package simulator

import "context"

var CmdProviderInterfaceBind = &CommandSpec{
	[]Token{TProvider, TL2Interface, TBind}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoProviderInterfaceBind = &CommandSpec{
	[]Token{TNo, TProvider, TL2Interface, TBind}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

