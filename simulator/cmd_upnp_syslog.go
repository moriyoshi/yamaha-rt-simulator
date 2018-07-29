package simulator

import "context"

var CmdUpnpSyslog = &CommandSpec{
	[]Token{TUpnp, TSyslog}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoUpnpSyslog = &CommandSpec{
	[]Token{TNo, TUpnp, TSyslog}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

