package simulator

import "context"

var CmdIpInterfaceIntrusionDetectionReport = &CommandSpec{
	[]Token{TIp, TL2Interface, TIntrusion, TDetection, TReport}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpIntrusionDetectionReport = &CommandSpec{
	[]Token{TIp, TPp, TIntrusion, TDetection, TReport}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpTunnelIntrusionDetectionReport = &CommandSpec{
	[]Token{TIp, TTunnel, TIntrusion, TDetection, TReport}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceIntrusionDetectionReport = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TIntrusion, TDetection, TReport}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpIntrusionDetectionReport = &CommandSpec{
	[]Token{TNo, TIp, TPp, TIntrusion, TDetection, TReport}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTunnelIntrusionDetectionReport = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TIntrusion, TDetection, TReport}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

