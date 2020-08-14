package launchd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/gnue/inst"
)

type ProcessType string

const (
	Background  = ProcessType("Background")
	Standard    = ProcessType("Standard")
	Adaptive    = ProcessType("Adaptive")
	Interactive = ProcessType("Interactive")
)

type Service struct {
	Label                  string
	Disabled               bool                `plist:",omitempty"`
	UserName               string              `plist:",omitempty"`
	GroupName              string              `plist:",omitempty"`
	InetdCompatibility     *InetdCompatibility `plist:"inetdCompatibility,omitempty"`
	LimitLoadToHosts       []string            `plist:",omitempty"`
	LimitLoadFromHosts     []string            `plist:",omitempty"`
	LimitLoadToSessionType string              `plist:",omitempty"`
	Program                string              `plist:",omitempty"`
	ProgramArguments       []string            `plist:",omitempty"`
	EnableGlobbing         bool                `plist:",omitempty"`
	EnableTransactions     bool                `plist:",omitempty"`
	KeepAlive              interface{}         `plist:",omitempty"`
	RunAtLoad              bool                `plist:",omitempty"`
	RootDirectory          string              `plist:",omitempty"`
	WorkingDirectory       string              `plist:",omitempty"`
	EnvironmentVariables   map[string]string   `plist:",omitempty"`
	Umask                  int                 `plist:",omitempty"`
	TimeOut                int                 `plist:",omitempty"`
	ExitTimeOut            int                 `plist:",omitempty"`
	ThrottleInterval       int                 `plist:",omitempty"`
	InitGroups             bool                `plist:",omitempty"`
	WatchPaths             []string            `plist:",omitempty"`
	QueueDirectories       []string            `plist:",omitempty"`
	StartOnMount           bool                `plist:",omitempty"`
	StartInterval          int                 `plist:",omitempty"`
	StartCalendarInterval  interface{}         `plist:",omitempty"`
	StandardInPath         string              `plist:",omitempty"`
	StandardOutPath        string              `plist:",omitempty"`
	StandardErrorPath      string              `plist:",omitempty"`
	Debug                  bool                `plist:",omitempty"`
	WaitForDebugger        bool                `plist:",omitempty"`
	SoftResourceLimits     map[string]int      `plist:",omitempty"`
	HardResourceLimits     *HardResourceLimits `plist:",omitempty"`
	Nice                   int                 `plist:",omitempty"`
	ProcessType            ProcessType         `plist:",omitempty"`
	AbandonProcessGroup    bool                `plist:",omitempty"`
	LowPriorityIO          bool                `plist:",omitempty"`
	LaunchOnlyOnce         bool                `plist:",omitempty"`
	MachServices           interface{}         `plist:",omitempty"`
	Sockets                interface{}         `plist:",omitempty"`
}

type InetdCompatibility struct {
	Wait bool `plist:",omitempty"`
}

type KeepAlive struct {
	SuccessfulExit  bool            `plist:",omitempty"`
	NetworkState    bool            `plist:",omitempty"`
	PathState       map[string]bool `plist:",omitempty"`
	OtherJobEnabled map[string]bool `plist:",omitempty"`
}

type StartCalendarInterval struct {
	Minute  int `plist:",omitempty"`
	Hour    int `plist:",omitempty"`
	Day     int `plist:",omitempty"`
	Weekday int `plist:",omitempty"`
	Month   int `plist:",omitempty"`
}

type HardResourceLimits struct {
	Core              int `plist:",omitempty"`
	CPU               int `plist:",omitempty"`
	Data              int `plist:",omitempty"`
	FileSize          int `plist:",omitempty"`
	MemoryLock        int `plist:",omitempty"`
	NumberOfFiles     int `plist:",omitempty"`
	NumberOfProcesses int `plist:",omitempty"`
	ResidentSetSize   int `plist:",omitempty"`
	Stack             int `plist:",omitempty"`
}

type MachServices struct {
	ResetAtClose     bool `plist:",omitempty"`
	HideUntilCheckIn bool `plist:",omitempty"`
}

type Sockets struct {
	SockType            string      `plist:",omitempty"`
	SockPassive         bool        `plist:",omitempty"`
	SockNodeName        string      `plist:",omitempty"`
	SockServiceName     string      `plist:",omitempty"`
	SockFamily          string      `plist:",omitempty"`
	SockProtocol        string      `plist:",omitempty"`
	SockPathName        string      `plist:",omitempty"`
	SecureSocketWithKey string      `plist:",omitempty"`
	SockPathMode        int         `plist:",omitempty"`
	Bonjour             interface{} `plist:",omitempty"`
	MulticastGroup      string      `plist:",omitempty"`
}

func InstallAction(fname string, loc inst.Locate) error {
	var d do
	if loc == inst.Global {
		d = do("sudo")
	}

	return d.launchctl("load", "-w", fname)
}

func UninstallAction(fname string, loc inst.Locate) error {
	var d do
	if loc == inst.Global {
		d = do("sudo")
	}

	return d.launchctl("unload", fname)
}

type do string

func (d do) launchctl(args ...string) error {
	var cmd *exec.Cmd

	if d == do("sudo") {
		args = append([]string{"launchctl"}, args...)
		cmd = exec.Command("sudo", args...)
	} else {
		cmd = exec.Command("launchctl", args...)
	}

	b, err := cmd.CombinedOutput()
	if err == nil && 0 < len(b) {
		err = fmt.Errorf("launchctl: %v", strings.TrimSpace(string(b)))
	}

	return err
}
