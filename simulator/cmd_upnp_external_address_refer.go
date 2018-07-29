package simulator

import "context"

var CmdUpnpExternalAddressRefer = &CommandSpec{
	[]Token{TUpnp, TExternal, TAddress, TRefer, TPp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoUpnpExternalAddressRefer = &CommandSpec{
	[]Token{TNo, TUpnp, TExternal, TAddress, TRefer, TPp}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
