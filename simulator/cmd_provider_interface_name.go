package simulator

import "context"

var CmdProviderInterfaceName = &CommandSpec{
	[]Token{TProvider, TL2Interface, TName}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoProviderInterfaceName = &CommandSpec{
	[]Token{TNo, TProvider, TL2Interface, TName}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

