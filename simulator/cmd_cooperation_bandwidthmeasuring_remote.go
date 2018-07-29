package simulator

import "context"

var CmdCooperationBandwidthMeasuringRemote = &CommandSpec{
	[]Token{TCooperation, TBandwidthMeasuring, TRemote}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoCooperationBandwidthMeasuringRemote = &CommandSpec{
	[]Token{TNo, TCooperation, TBandwidthMeasuring, TRemote}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

