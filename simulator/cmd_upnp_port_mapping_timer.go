package simulator

import "context"

var CmdUpnpPortMappingTimer = &CommandSpec{
	[]Token{TUpnp, TPort, TMapping, TTimer}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoUpnpPortMappingTimer = &CommandSpec{
	[]Token{TNo, TUpnp, TPort, TMapping, TTimer}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
