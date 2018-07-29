package simulator

import "context"

var CmdShowIpIntrusionDetection = &CommandSpec{
	[]Token{TShow, TIp, TIntrusion, TDetection}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowIpIntrusionDetectionPp = &CommandSpec{
	[]Token{TShow, TIp, TIntrusion, TDetection, TPp}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowIpIntrusionDetectionTunnel = &CommandSpec{
	[]Token{TShow, TIp, TIntrusion, TDetection, TTunnel}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
