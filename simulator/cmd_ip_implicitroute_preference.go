package simulator

import "context"

var CmdIpImplicitRoutePreference = &CommandSpec{
	[]Token{TIp, TImplicitRoute, TPreference}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpImplicitRoutePreferencePreference = &CommandSpec{
	[]Token{TNo, TIp, TImplicitRoute, TPreference}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

