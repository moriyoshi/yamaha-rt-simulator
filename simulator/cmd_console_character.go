package simulator

import "context"

var CmdConsoleCharacter = &CommandSpec{
	[]Token{TConsole, TCharacter, TCharsetName}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if len(tis) < 3 {
			return WrongNumberOfArgs
		}
		for _, nls := range sess.Spec.Charsets {
			for _, v := range nls.Variants {
				if v == tis[2].Image {
					sess.Charset = nls
					sess.ResetEncoder()
					return nil
				}
			}
		}
		return IllegalKeyword
	},
}
