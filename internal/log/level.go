package log

type Level int

const LevelDebug = Level(1)
const LevelInfo = Level(2)
const LevelWarn = Level(3)
const LevelFatal = Level(4)

const LevelDebugString = "debug"
const LevelInfoString = "info"
const LevelWarnString = "warn"
const LevelFatalString = "fatal"

var LevelToString = map[Level]string{
	LevelDebug: LevelDebugString,
	LevelInfo:  LevelInfoString,
	LevelWarn:  LevelWarnString,
	LevelFatal: LevelFatalString,
}
