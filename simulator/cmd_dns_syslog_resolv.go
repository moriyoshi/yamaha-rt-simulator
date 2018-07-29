package simulator

import "context"

var CmdDnsSyslogResolv = &CommandSpec{
	[]Token{TDns, TSyslog, TResolv}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDnsSyslogResolv = &CommandSpec{
	[]Token{TNo, TDns, TSyslog, TResolv}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

