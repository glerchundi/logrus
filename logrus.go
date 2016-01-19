package logrus

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"io"
)

var mainLogger = NewLogger()

func init() {
	mainLogger.SetTag(os.Args[0])
}

type Logger struct {
	// Create a new instance of the logger. You can have any number of instances.
	log *logrus.Logger
	tag string
}

func NewLogger() *Logger {
	logger := &Logger{log:logrus.New()}
	logger.log.Formatter = logger
	return logger
}

func (l *Logger) Format(e *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Format(time.RFC3339)
	hostname, _ := os.Hostname()
	return []byte(
		fmt.Sprintf(
			"%s %s %s[%d]: %s %s\n",
			timestamp, hostname, l.tag, os.Getpid(), strings.ToUpper(e.Level.String()), e.Message,
		),
	), nil
}

func (l *Logger) SetTag(tag string) {
	l.tag = tag
}

func (l *Logger) SetLevel(level string) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		Fatal(fmt.Sprintf(`not a valid level: "%s"`, level))
	}
	l.log.Level = lvl
}

func (l *Logger) SetOutput(w io.Writer) {
	l.log.Out = w
}

func (l *Logger) SetFormatter(f logrus.Formatter) {
	l.log.Formatter = f
}

func (l *Logger) Debugf(format string, args ...interface{}) { l.log.Debugf(format, args...) }
func (l *Logger) Infof(format string, args ...interface{}) { l.log.Infof(format, args...) }
func (l *Logger) Printf(format string, args ...interface{}) { l.log.Printf(format, args...) }
func (l *Logger) Warnf(format string, args ...interface{}) { l.log.Warnf(format, args...) }
func (l *Logger) Errorf(format string, args ...interface{}) { l.log.Errorf(format, args...) }
func (l *Logger) Fatalf(format string, args ...interface{}) { l.log.Fatalf(format, args...) }
func (l *Logger) Panicf(format string, args ...interface{}) { l.log.Panicf(format, args...) }

func (l *Logger) Debug(args ...interface{}) { l.log.Debug(args...) }
func (l *Logger) Info(args ...interface{}) { l.log.Info(args...) }
func (l *Logger) Print(args ...interface{}) { l.log.Print(args...) }
func (l *Logger) Warn(args ...interface{}) { l.log.Warn(args...) }
func (l *Logger) Error(args ...interface{}) { l.log.Error(args...) }
func (l *Logger) Fatal(args ...interface{}) { l.log.Fatal(args...) }
func (l *Logger) Panic(args ...interface{}) { l.log.Panic(args...) }

func (l *Logger) Debugln(args ...interface{}) { l.log.Debugln(args...) }
func (l *Logger) Infoln(args ...interface{}) { l.log.Infoln(args...) }
func (l *Logger) Println(args ...interface{}) { l.log.Println(args...) }
func (l *Logger) Warnln(args ...interface{}) { l.log.Warnln(args...) }
func (l *Logger) Errorln(args ...interface{}) { l.log.Errorln(args...) }
func (l *Logger) Fatalln(args ...interface{}) { l.log.Fatalln(args...) }
func (l *Logger) Panicln(args ...interface{}) { l.log.Panicln(args...) }

func Debugf(format string, args ...interface{}) { mainLogger.Debugf(format, args...) }
func Infof(format string, args ...interface{}) { mainLogger.Infof(format, args...) }
func Printf(format string, args ...interface{}) { mainLogger.Printf(format, args...) }
func Warnf(format string, args ...interface{}) { mainLogger.Warnf(format, args...) }
func Errorf(format string, args ...interface{}) { mainLogger.Errorf(format, args...) }
func Fatalf(format string, args ...interface{}) { mainLogger.Fatalf(format, args...) }
func Panicf(format string, args ...interface{}) { mainLogger.Panicf(format, args...) }

func Debug(args ...interface{}) { mainLogger.Debug(args...) }
func Info(args ...interface{}) { mainLogger.Info(args...) }
func Print(args ...interface{}) { mainLogger.Print(args...) }
func Warn(args ...interface{}) { mainLogger.Warn(args...) }
func Error(args ...interface{}) { mainLogger.Error(args...) }
func Fatal(args ...interface{}) { mainLogger.Fatal(args...) }
func Panic(args ...interface{}) { mainLogger.Panic(args...) }

func Debugln(args ...interface{}) { mainLogger.Debugln(args...) }
func Infoln(args ...interface{}) { mainLogger.Infoln(args...) }
func Println(args ...interface{}) { mainLogger.Println(args...) }
func Warnln(args ...interface{}) { mainLogger.Warnln(args...) }
func Errorln(args ...interface{}) { mainLogger.Errorln(args...) }
func Fatalln(args ...interface{}) { mainLogger.Fatalln(args...) }
func Panicln(args ...interface{}) { mainLogger.Panicln(args...) }