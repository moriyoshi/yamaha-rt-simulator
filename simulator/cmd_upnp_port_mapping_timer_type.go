package simulator

import "context"

var CmdUpnpPortMappingTimerType = &CommandSpec{
	[]Token{TUpnp, TPort, TMapping, TTimer, TType}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoUpnpMappingTimerType = &CommandSpec{
	[]Token{TNo, TUpnp, TMapping, TTimer, TType}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

