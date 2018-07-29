package simulator

import "context"

var CmdNetvolanteDnsHostnameHost = &CommandSpec{
	[]Token{TNetvolanteDns, THostname, THost}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNetvolanteDnsHostnameHostPp = &CommandSpec{
	[]Token{TNetvolanteDns, THostname, THost, TPp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNetvolanteDnsHostnameHostInterfaceHost = &CommandSpec{
	[]Token{TNo, TNetvolanteDns, THostname, THost}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNetvolanteDnsHostnameHostPpHost = &CommandSpec{
	[]Token{TNo, TNetvolanteDns, THostname, THost, TPp}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
