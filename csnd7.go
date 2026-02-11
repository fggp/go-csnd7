package csnd7

/*
#cgo CFLAGS: -DUSE_DOUBLE=1
#cgo CFLAGS: -I /usr/local/include
#cgo linux CFLAGS: -DLINUX=1
#cgo LDFLAGS: -lcsound64

#include <csound/csound.h>

CS_AUDIODEVICE *getAudioDevList(CSOUND *csound, int n, int isOutput)
{
  CS_AUDIODEVICE *devs = (CS_AUDIODEVICE *)malloc(n*sizeof(CS_AUDIODEVICE));
  csoundGetAudioDevList(csound, devs, isOutput);
  return devs;
}

void getAudioDev(CS_AUDIODEVICE *devs, int i, char **pname, char **pid, char **pmodule,
                 int32_t *nchnls, int32_t *flag)
{
  CS_AUDIODEVICE dev = devs[i];
  *pname = dev.device_name;
  *pid = dev.device_id;
  *pmodule = dev.rt_module;
  *nchnls = dev.max_nchnls;
  *flag = dev.isOutput;
}

CS_MIDIDEVICE *getMidiDevList(CSOUND *csound, int n, int isOutput)
{
  CS_MIDIDEVICE *devs = (CS_MIDIDEVICE *)malloc(n*sizeof(CS_MIDIDEVICE));
  csoundGetMIDIDevList(csound, devs, isOutput);
  return devs;
}

void getMidiDev(CS_MIDIDEVICE *devs, int i, char **pname, char** piname, char **pid,
                char **pmodule, int32_t *flag)
{
  CS_MIDIDEVICE dev = devs[i];
  *pname = dev.device_name;
  *piname = dev.interface_name;
  *pid = dev.device_id;
  *pmodule = dev.midi_module;
  *flag = dev.isOutput;
}

int32_t getInt32Val(int32_t *lst, int i)
{
  return lst[i];
}

void cMessage(CSOUND *csound, const char *msg)
{
  csoundMessage(csound, "%s", msg);
}

void cMessageS(CSOUND *csound, int32_t attr, char *msg)
{
  csoundMessageS(csound, attr, "%s", msg);
}

controlChannelHints_t *getControlChannelInfo(controlChannelInfo_t *list, int i,
                                             char **name, int *type)
{
  *name = list[i].name;
  *type = list[i].type;
  return &list[i].hints;
}
*/
import "C"

import (
	"bytes"
	"fmt"
	"html/template"
	"unsafe"
)

func cbool(flag bool) C.int32_t {
	if flag {
		return 1
	}
	return 0
}

func cMYFLT(val MYFLT) C.double {
	return C.double(val)
}

func cpMYFLT(pval *MYFLT) *C.double {
	return (*C.double)(pval)
}

func cppMYFLT(ppval **MYFLT) **C.double {
	return (**C.double)(unsafe.Pointer(ppval))
}

// Error Definitions
const (
	CSOUND_SUCCESS        = 0  // Completed successfully.
	CSOUND_ERROR          = -1 // Unspecified failure.
	CSOUND_INITIALIZATION = -2 // Failed during initialization.
	CSOUND_PERFORMANCE    = -3 // Failed during performance.
	CSOUND_MEMORY         = -4 // Failed to allocate requested memory.
	CSOUND_SIGNAL         = -5 // Termination requested by SIGINT or SIGTERM.
)

// Encapsulates an opaque pointer to a Csound instance
type CSOUND struct {
	Cs (*C.CSOUND)
}

type STRINGDAT *C.STRINGDAT
type ARRAYDAT *C.ARRAYDAT
type PVSDAT *C.PVSDAT

type MYFLT float64

type CsoundParams struct {
	Odebug            bool    // debug flag
	Sfread            bool    // sound input read flag
	Sfwrite           bool    // sound output write flag (-s)
	Filetyp           int     // soundfile type code
	Inbufsamps        int     // input buffer size in samples
	Outbufsamps       int     // output buffer size in samples
	Informat          int     // input soundfile format
	Outformat         int     // output soundfile format
	SndfileSampleSize int     // sample size
	Displays          bool    // displays flag
	Graphsoff         bool    // graphs flag
	Postscript        bool    // postscript graphs flag
	Msglevel          int     // message level (-m)
	Beatmode          bool    // beat mode
	OMaxLag           int     // hardware buffer size (samples)
	Linein            bool    // linevents flag (-L)
	RTevents          bool    // realtime events flag (scoreless, -L, -F, -M)
	Midiin            bool    // midi input flag (-M)
	FMidiin           bool    // midi file input flag (-F)
	RMidiin           bool    // remote events flag
	Ringbell          bool    // ringbell flag
	Termifend         bool    // terminate on MIDI file input flag (-T)
	Rewrt_hdr         bool    // rewrite header flag
	Heartbeat         int     // heartbeat flag
	Gen01defer        bool    // GEN01 defer allocation flag
	CmdTempo          float64 // tempo value (-t)
	Sr_override       MYFLT   // sampling rate override (-r)
	Kr_override       MYFLT   // control rate override (-k)
	Nchnls_override   int     // nchnls override
	Nchnls_i_override int     // nchnls_i override
	Infilename        string  // input file name (-i)
	Outfilename       string  // output file name (-o)
	Linename          string  // line events source (-L)
	Midiname          string  // MIDI input device name (-M)
	FMidiname         string  // MIDI input file name (-F)
	Midioutname       string  // MIDI output device name (-Q)
	FMidioutname      string  // MIDI output file name
	MidiKey           int     // MIDI key pfield mapping
	MidiKeyCps        int     // MIDI key-cps pfield mapping
	MidiKeyOct        int     // MIDI key-oct pfield mapping
	MidiKeyPch        int     // MIDI key-pch pfield mapping
	MidiVelocity      int     // MIDI vel pfield mapping
	MidiVelocityAmp   int     // MIDI vel-amp pfield mapping
	NoDefaultPaths    bool    // default paths flag
	NumThreads        int     // multicore number of threads (-j)
	SyntaxCheckOnly   bool    // syntax check only flag
	RunUnitTests      bool    // run unit tests flag
	UseCsdLineCounts  bool    // csd line nums option
	SampleAccurate    bool    // sample accurate flag
	Realtime          bool    // realtime priority flag
	E0dbfs_override   MYFLT   // 0dbfs override
	Daemon            bool    // daemon mode flag
	Quality           float64 // OGG encoding quality
	Ksmps_override    int     // ksmps override
	Fft_lib           int     // FFT library option
	Echo              bool    // UDP echo commands flag
	Limiter           MYFLT   // audio output limiter option
	Sr_default        MYFLT   // default sampling rate
	Kr_default        MYFLT   // default control rate
	Mp3_mode          int     // MP3 encoding mode
	Redef             bool    // instr redefinition flag
	Error_deprecated  bool    // error on deprecated opcodes
}

const paramsString = `
Debug: {{.Odebug}}

Sound input read:              {{.Sfread}}
Sound output write:            {{.Sfwrite}}
Sound File type code:          {{.Filetyp}}
Input buffer size in samples:  {{.Inbufsamps}}
Output buffer size in samples: {{.Outbufsamps}}
Input soundfile format:        {{.Informat}}
Output soundfile format:       {{.Outformat}}
Sample size:                   {{.SndfileSampleSize}}

Displays:   {{.Displays}}
Graphs off: {{.Graphsoff}}
Postscript: {{.Postscript}}

Message level:                  {{.Msglevel}}
Beat mode:                      {{.Beatmode}}
Hardware buffer size (samples): {{.OMaxLag}}

Line events:                  {{.Linein}}
Realtime events:              {{.RTevents}}
MIDI input:                   {{.Midiin}}
MIDI file input:              {{.FMidiin}}
Remote events:                {{.RMidiin}}
Ringbell:                     {{.Ringbell}}
Terminate on MIDI file input: {{.Termifend}}
Rewrite header:               {{.Rewrt_hdr}}
Heart beat:                   {{.Heartbeat}}
GEN01 defer allocation        {{.Gen01defer}}

Tempo value:            {{.CmdTempo}}
Sampling rate override: {{.Sr_override}}
Control rate override:  {{.Kr_override}}
nchnls override:        {{.Nchnls_override}}
nchnls_i override:      {{.Nchnls_i_override}}

Input file name:    {{.Infilename}}
Output file name:   {{.Outfilename}}
Line events source: {{.Linename}}
MIDI input device:  {{.Midiname}}
MIDI input file:    {{.FMidiname}}
MIDI output device: {{.Midioutname}}
MIDI output file:   {{.FMidioutname}}

MIDI key pfield mapping:     {{.MidiKey}}
MIDI key-cps pfield mapping: {{.MidiKeyCps}}
MIDI key-oct pfield mapping: {{.MidiKeyOct}}
MIDI key-pch pfield mapping: {{.MidiKeyPch}}
MIDI vel pfield mapping:     {{.MidiVelocity}}
MIDI vel-amp pfield mapping: {{.MidiVelocityAmp}}

No default paths:            {{.NoDefaultPaths}}
Multicore number of threads: {{.NumThreads}}
Syntax check only:           {{.SyntaxCheckOnly}}
Run unit tests:              {{.RunUnitTests}}
Csd line numbers:            {{.UseCsdLineCounts}}
Sample accurate:             {{.SampleAccurate}}
Realtime priority:           {{.Realtime}}
0dfbs override:              {{.E0dbfs_override}}
Daemon mode:                 {{.Daemon}}

OGG encoding quality:        {{.Quality}}
ksmps override:              {{.Ksmps_override}}
FFT library option:          {{.Fft_lib}}
UDP echo commandes:          {{.Echo}}
Audio output limiter:        {{.Limiter}}
Default sampling rate:       {{.Sr_default}}
Default control rate:        {{.Kr_default}}
Mp3 encoding mode:           {{.Mp3_mode}}
Instr redefinition:          {{.Redef}}
Error on deprecated opcodes: {{.Error_deprecated}}
`

func (params CsoundParams) String() string {
	var buf bytes.Buffer
	t := template.Must(template.New("params").Parse(paramsString))
	t.Execute(&buf, params)
	return buf.String()
}

type CsoundAudioDevice struct {
	DeviceName string
	DeviceId   string
	RtModule   string
	MaxNchnls  int
	IsOutput   bool
}

func (dev CsoundAudioDevice) String() string {
	return fmt.Sprintf("(%s, %s, %s, %d, %t)", dev.DeviceName, dev.DeviceId,
		dev.RtModule, dev.MaxNchnls, dev.IsOutput)
}

type CsoundMidiDevice struct {
	DeviceName    string
	InterfaceName string
	DeviceId      string
	MidiModule    string
	IsOutput      bool
}

func (dev CsoundMidiDevice) String() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %t)", dev.DeviceName,
		dev.InterfaceName, dev.DeviceId, dev.MidiModule, dev.IsOutput)
}

// Constants used by the bus interface (csoundGetChannelPtr() etc.).
const (
	CSOUND_CONTROL_CHANNEL   = 1
	CSOUND_AUDIO_CHANNEL     = 2
	CSOUND_STRING_CHANNEL    = 3
	CSOUND_PVS_CHANNEL       = 4
	CSOUND_VAR_CHANNEL       = 5
	CSOUND_ARRAY_CHANNEL     = 6
	CSOUND_CHANNEL_TYPE_MASK = 15
	CSOUND_INPUT_CHANNEL     = 16
	CSOUND_OUTPUT_CHANNEL    = 32
)

// Control Channel Behavior
const (
	CSOUND_CONTROL_CHANNEL_NO_HINTS = 0
	CSOUND_CONTROL_CHANNEL_INT      = 1
	CSOUND_CONTROL_CHANNEL_LIN      = 2
	CSOUND_CONTROL_CHANNEL_EXP      = 3
)

// This structure holds the parameter hints for control channels.
type ControlChannelHints struct {
	Behav      int
	Dflt       MYFLT
	Min        MYFLT
	Max        MYFLT
	X          int
	Y          int
	Width      int
	Height     int
	Attributes string // This member must be set explicitly to NULL if not used
}

type ControlChannelInfo struct {
	Name  string
	Type  int
	Hints ControlChannelHints
}

const (
	CS_INSTR_EVENT = iota
	CS_TABLE_EVENT
	CS_END_EVENT
	CS_ADV_EVENT
)

//////////////////
// Instantiation
//////////////////

// Initialize Csound library with specific flags. This function is called
// internally by csoundCreate(), so there is// generally no need to use it
// explicitly unless you need to avoid default initialization that sets
// signal handlers and atexit() callbacks.
// Return value is zero on success, positive if initialisation was
// done already, and negative on error.
func Initialize(flags int) int {
	return int(C.csoundInitialize(C.int32_t(flags)))
}

// Creates an instance of Csound.  Returns an opaque pointer that
// must be passed to most Csound API functions.  The hostData
// parameter can be nil, or it can be a pointer to any sort of
// data; this pointer can be accessed from the Csound instance
// that is passed to callback routines.
// If not an empty string the opcodedir parameter sets an override
// for the plugin module/opcode directory search.
func Create(hostData unsafe.Pointer, opcodeDir string) CSOUND {
	var cs (*C.CSOUND)
	if opcodeDir == "" {
		cs = C.csoundCreate(hostData, nil)
	} else {
		copcodeDir := C.CString(opcodeDir)
		defer C.free(unsafe.Pointer(copcodeDir))
		cs = C.csoundCreate(nil, copcodeDir)
	}
	return CSOUND{cs}
}

// Destroy an instance of Csound.
func (csound CSOUND) Destroy() {
	C.csoundDestroy(csound.Cs)
	csound.Cs = nil
}

////////////////
// Attributes
////////////////

// Returns the version number
func (csound CSOUND) Version() string {
	n := int(C.csoundGetVersion())
	l1, l3 := n/1000, n%1000
	l2 := l3 / 10
	l3 %= 10
	return fmt.Sprintf("%d.%02d.%d", l1, l2, l3)
}

// Returns the number of audio sample frames per second.
func (csound CSOUND) Sr() MYFLT {
	return MYFLT(C.csoundGetSr(csound.Cs))
}

// Returns the number of control samples per second.
func (csound CSOUND) Kr() MYFLT {
	return MYFLT(C.csoundGetKr(csound.Cs))
}

// Returns the audio vector size in frames (= sr/kr).
func (csound CSOUND) Ksmps() int {
	return int(C.csoundGetKsmps(csound.Cs))
}

// Returns the number of audio channels in the Csound instance.
// If isInput is false, the value of nchnls is returned,
// otherwise nchnls_i.
func (csound CSOUND) Channels(isInput bool) int {
	return int(C.csoundGetChannels(csound.Cs, cbool(isInput)))
}

// Returns the 0dBFS level of the spin/spout buffers.
func (csound CSOUND) Get0dbFS() MYFLT {
	return MYFLT(C.csoundGet0dBFS(csound.Cs))
}

// Returns the A4 frequency reference
func (csound CSOUND) A4() MYFLT {
	return MYFLT(C.csoundGetA4(csound.Cs))
}

// Returns the current performance time in sample frames
func (csound CSOUND) CurrentTimeSamples() int {
	return int(C.csoundGetCurrentTimeSamples(csound.Cs))
}

// Returns the size of MYFLT in bytes.
func (csound CSOUND) SizeOfMYFLT() int {
	return int(C.csoundGetSizeOfMYFLT())
}

// Returns host data.
func (csound CSOUND) HostData() unsafe.Pointer {
	return C.csoundGetHostData(csound.Cs)
}

// Sets host data.
func (csound CSOUND) SetHostData(hostData unsafe.Pointer) {
	C.csoundSetHostData(csound.Cs, hostData)
}

// Returns the total error count of the current performance.
func (csound CSOUND) ErrCnt() int {
	return int(C.csoundErrCnt(csound.Cs))
}

// Get the value of environment variable 'name', searching
// in this order: local environment of 'csound' (if not NULL), variables
// set with csound.SetGlobalEnv(), and system environment variables.
// If 'csound' is not NULL, should be called after csound.Compile().
// Return value is "" if the variable is not set.
func (csound CSOUND) Env(name string) string {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return C.GoString(C.csoundGetEnv(csound.Cs, cname))
}

// Set the global value of environment variable 'name' to 'value',
// or delete variable if 'value' is "".
// It is not safe to call this function while any Csound instances
// are active.
// Returns zero on success.
func (csound CSOUND) SetGlobalEnv(name, value string) int {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var cvalue *C.char
	if len(value) == 0 {
		cvalue = nil
	} else {
		cvalue = C.CString(value)
		defer C.free(unsafe.Pointer(cvalue))
	}
	return int(C.csoundSetGlobalEnv(cname, cvalue))
}

// Set csound options (flag). Returns CSOUND_SUCCESS on success.
// This needs to be called after Create() and before any code is
// compiled. Multiple options are allowed in one string.
func (csound CSOUND) SetOption(option string) int {
	coption := C.CString(option)
	defer C.free(unsafe.Pointer(coption))
	return int(C.csoundSetOption(csound.Cs, coption))
}

// Get the current set of parameters from a CSOUND instance in
// a CSOUND_PARAMS structure.
func (csound CSOUND) Params() CsoundParams {
	p := C.csoundGetParams(csound.Cs)
	var params CsoundParams

	params.Odebug = p.odebug != 0
	params.Sfread = p.sfread != 0
	params.Sfwrite = p.sfwrite != 0
	params.Filetyp = int(p.filetyp)
	params.Inbufsamps = int(p.inbufsamps)
	params.Outbufsamps = int(p.outbufsamps)
	params.Informat = int(p.informat)
	params.Outformat = int(p.outformat)
	params.SndfileSampleSize = int(p.sndfileSampleSize)
	params.Displays = p.displays != 0
	params.Graphsoff = p.graphsoff != 0
	params.Postscript = p.postscript != 0
	params.Msglevel = int(p.msglevel)
	params.Beatmode = p.Beatmode != 0
	params.OMaxLag = int(p.oMaxLag)
	params.Linein = p.Linein != 0
	params.RTevents = p.RTevents != 0
	params.Midiin = p.Midiin != 0
	params.FMidiin = p.FMidiin != 0
	params.RMidiin = p.RMidiin != 0
	params.Ringbell = p.ringbell != 0
	params.Termifend = p.termifend != 0
	params.Rewrt_hdr = p.rewrt_hdr != 0
	params.Heartbeat = int(p.heartbeat)
	params.Gen01defer = p.gen01defer != 0
	params.CmdTempo = float64(p.cmdTempo)
	params.Sr_override = MYFLT(p.sr_override)
	params.Kr_override = MYFLT(p.kr_override)
	params.Nchnls_override = int(p.nchnls_override)
	params.Nchnls_i_override = int(p.nchnls_i_override)
	params.Infilename = C.GoString(p.infilename)
	params.Outfilename = C.GoString(p.outfilename)
	params.Linename = C.GoString(p.Linename)
	params.Midiname = C.GoString(p.Midiname)
	params.FMidiname = C.GoString(p.FMidiname)
	params.Midioutname = C.GoString(p.Midioutname)
	params.FMidioutname = C.GoString(p.FMidioutname)
	params.MidiKey = int(p.midiKey)
	params.MidiKeyCps = int(p.midiKeyCps)
	params.MidiKeyOct = int(p.midiKeyOct)
	params.MidiKeyPch = int(p.midiKeyPch)
	params.MidiVelocity = int(p.midiVelocity)
	params.MidiVelocityAmp = int(p.midiVelocityAmp)
	params.NoDefaultPaths = p.noDefaultPaths != 0
	params.NumThreads = int(p.numThreads)
	params.SyntaxCheckOnly = p.syntaxCheckOnly != 0
	params.RunUnitTests = p.runUnitTests != 0
	params.UseCsdLineCounts = p.useCsdLineCounts != 0
	params.SampleAccurate = p.sampleAccurate != 0
	params.Realtime = p.realtime != 0
	params.E0dbfs_override = MYFLT(p.e0dbfs_override)
	params.Daemon = p.daemon != 0
	params.Quality = float64(p.quality)
	params.Ksmps_override = int(p.ksmps_override)
	params.Fft_lib = int(p.fft_lib)
	params.Echo = p.echo != 0
	params.Limiter = MYFLT(p.limiter)
	params.Sr_default = MYFLT(p.sr_default)
	params.Kr_default = MYFLT(p.kr_default)
	params.Mp3_mode = int(p.mp3_mode)
	params.Redef = p.redef != 0
	params.Error_deprecated = p.error_deprecated != 0

	return params
}

// Returns whether Csound is set to print debug messages sent through the
// DebugMsg() internal API function. Anything different to 0 means true.
func (csound CSOUND) Debug() bool {
	return C.csoundGetDebug(csound.Cs) != 0
}

// Sets whether Csound prints debug messages from the DebugMsg() internal
// API function. Anything different to 0 means true.
func (csound CSOUND) SetDebug(debug bool) {
	C.csoundSetDebug(csound.Cs, cbool(debug))
}

// If val > 0, sets the internal variable holding the system HW sr.
// Returns the stored value containing the system HW sr.
func (csound CSOUND) SystemSr(val MYFLT) MYFLT {
	return MYFLT(C.csoundSystemSr(csound.Cs, cMYFLT(val)))
}

// Retrieves a module name and type ("audio" or "midi").
// Given a number Modules are added to list as csound loads them.
// Returns CSOUND_SUCCESS on success and CSOUND_ERROR if module
// number was not found
//
//	for n :=0; ; n++ {
//	    name, type, err := csound.Module(n)
//	    if err != CSOUND_SUCCESS {
//	        break
//	    }
//	    fmt.printf("Module %d:  %s (%s) \n", n, name, type)
func (csound CSOUND) Module(n int) (name, type_ string, err int) {
	var cname, ctype *C.char
	cerror := C.csoundGetModule(csound.Cs, C.int32_t(n), &cname, &ctype)
	name = C.GoString(cname)
	type_ = C.GoString(ctype)
	err = int(cerror)
	return
}

// This function can be called to obtain a list of available
// input or output audio devices available (depending on the backend
// module used).
//
//	list := csound.AudioDevList(true)
//	for i := range list {
//	    csound.Message(" %d: %s (%s)\n",
//	        i, list[i].DeviceId, list[i].DeviceName);
func (csound CSOUND) AudioDevList(isOutput bool) []CsoundAudioDevice {
	cflag := cbool(isOutput)
	n := C.csoundGetAudioDevList(csound.Cs, nil, cflag)
	devs := C.getAudioDevList(csound.Cs, n, cflag)
	defer C.free(unsafe.Pointer(devs))
	var list = make([]CsoundAudioDevice, int(n))
	var name, id, module *C.char
	var nchnls, isOut C.int32_t
	for i := range list {
		C.getAudioDev(devs, C.int(i), &name, &id, &module, &nchnls, &isOut)
		list[i].DeviceName = C.GoString(name)
		list[i].DeviceId = C.GoString(id)
		list[i].RtModule = C.GoString(module)
		list[i].MaxNchnls = int(nchnls)
		list[i].IsOutput = (isOut == 1)
	}
	return list
}

// This function can be called to obtain a list of available
// input or output midi devices. (see also AudioDevList())
func (csound CSOUND) MidiDevList(isOutput bool) []CsoundMidiDevice {
	cflag := cbool(isOutput)
	n := C.csoundGetMIDIDevList(csound.Cs, nil, cflag)
	devs := C.getMidiDevList(csound.Cs, n, cflag)
	defer C.free(unsafe.Pointer(devs))
	var list = make([]CsoundMidiDevice, int(n))
	var name, iname, id, module *C.char
	var isOut C.int32_t
	for i := range list {
		C.getMidiDev(devs, C.int(i), &name, &iname, &id, &module, &isOut)
		list[i].DeviceName = C.GoString(name)
		list[i].InterfaceName = C.GoString(iname)
		list[i].DeviceId = C.GoString(id)
		list[i].MidiModule = C.GoString(module)
		list[i].IsOutput = (isOut == 1)
	}
	return list
}

// Returns the Csound message level (from 0 to 231).
func (csound CSOUND) MessageLevel() int {
	return int(C.csoundGetMessageLevel(csound.Cs))
}

// Sets the Csound message level (from 0 to 231).
func (csound CSOUND) SetMessageLevel(messageLevel int) {
	C.csoundSetMessageLevel(csound.Cs, C.int32_t(messageLevel))
}

////////////////////////////////
// Compilation and Performance
////////////////////////////////

// Compiles Csound input files (such as an orchestra and score, or CSD)
// as directed by the supplied command-line arguments,
// but does not perform them. Returns a non-zero error code on failure.
// In this mode, the sequence of calls should be as follows:
//
//	csound.Compile(args)
//	csound.Start()
//	for !csound.PerformKsmps() {}
//	csound.Reset()
func (csound CSOUND) Compile(args []string) int {
	argc := C.int32_t(len(args))
	argv := make([]*C.char, argc)
	for i, arg := range args {
		argv[i] = C.CString(arg)
		defer C.free(unsafe.Pointer(argv[i]))
	}
	return int(C.csoundCompile(csound.Cs, argc, &argv[0]))
}

// Parse, and compile the given orchestra from an ASCII string,
// also evaluating any global space code (i-time only)
// in synchronous or asynchronous (async = true) mode.
//
//	orc := "instr 1 \n a1 = rand(0dbfs/4) \n out(a1) \n"
//	csound.CompileOrc(orc, false);
func (csound CSOUND) CompileOrc(code string, async bool) int {
	ccode := C.CString(code)
	defer C.free(unsafe.Pointer(ccode))
	return int(C.csoundCompileOrc(csound.Cs, ccode, cbool(async)))
}

// Parse and compile an orchestra given on a string, synchronously,
// evaluating any global space code (i-time only).
// On SUCCESS it returns a value passed to the
// 'return' opcode in global space.
//
//	 code := `
//	        i1 = 2 + 2
//	        return i1
//	        `
//	 retval := csound.EvalCode(code)
//
//	If the code fails to evaluate, the return value is always 0.
func (csound CSOUND) EvalCode(code string) MYFLT {
	ccode := C.CString(code)
	defer C.free(unsafe.Pointer(ccode))
	return MYFLT(C.csoundEvalCode(csound.Cs, ccode))
}

// Compiles a Csound input file (CSD, .csd file) or a txt string
// containing the CSD code, in synchronous or asynchronous (async = true) mode.
// Returns a non-zero error code on failure.
//
// If csound.Start is called before csound.CompileCSD, the <CsOptions>
// element is ignored (but csound.SetOption can be called any number of
// times), the <CsScore> element is dispatched as score events (e.g.
// as it is done by csoundEventString())
//
//	csound.SetOption("-an_option")
//	csound.SetOption("-another_option")
//	csound.Start()
//	csound.CompileCSD(csd_filename, 0, false)
//	for {
//	  csound.PerformKsmps()
//	  // Something to break out of the loop
//	  // when finished here...
//	}
//
// NB: this function can be called repeatedly during performance to
// replace or add new instruments and events.
//
// But if csound.CompileCsd is called before csound.Start, the <CsOptions>
// element is used, the <CsScore> section is pre-processed and dispatched
// normally, and performance terminates when the score terminates.
//
//	csound.CompileCSD(csound, csd_filename, 0, false)
//	csound.Start()
//	for !csound.PerformKsmps() {
//	}
//
// if mode = 1, csd contains a full CSD code (rather than a filename).
// This is convenient when it is desirable to package the csd as part of
// an application or a multi-language piece.
func (csound CSOUND) CompileCSD(csd string, mode int, async bool) int {
	csdName := C.CString(csd)
	defer C.free(unsafe.Pointer(csdName))
	return int(C.csoundCompileCSD(csound.Cs, csdName, C.int32_t(mode), cbool(async)))
}

// Prepares Csound for performance. Normally called after compiling
// a csd file or an orc file, in which case score preprocessing is
// performed and performance terminates when the score terminates.
//
// However, if called before compiling a csd file or an orc file,
// score preprocessing is not performed and "i" statements are dispatched
// as real-time events, the <CsOptions> tag is ignored, and performance
// continues indefinitely or until ended using the API.
func (csound CSOUND) Start() int {
	return int(C.csoundStart(csound.Cs))
}

// Senses input events, and performs one block of
// audio output containing ksmps frames. csound.Start() must be called first.
// Returns false during performance, and true when performance is finished.
// If called until it returns true, will perform an entire score.
// Enables external software to control the execution of Csound,
// and to synchronize performance with audio input and output.
func (csound CSOUND) PerformKsmps() bool {
	return C.csoundPerformKsmps(csound.Cs) != 0
}

// Run utility with the specified name and command line arguments.
// Should be called after loading utility plugins.
// Use csound.Reset() to clean up after calling this function.
// Returns zero if the utility was run successfully.
func (csound CSOUND) RunUtility(name string, args []string) int {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	argc := C.int32_t(len(args))
	argv := make([]*C.char, argc)
	for i, arg := range args {
		argv[i] = C.CString(arg)
		defer C.free(unsafe.Pointer(argv[i]))
	}
	return int(C.csoundRunUtility(csound.Cs, cname, argc, &argv[0]))
}

// Resets all internal memory and state in preparation for a new performance.
// Enables external software to run successive Csound performances
// without reloading Csound.
func (csound CSOUND) Reset() {
	C.csoundReset(csound.Cs)
}

//////////////
// Audio I/O
//////////////

// Returns the address of the Csound audio input working buffer (spin).
// Enables external software to write audio into Csound before calling
// csound.PerformKsmps.
func (csound CSOUND) Spin() []MYFLT {
	buffer := (*MYFLT)(C.csoundGetSpin(csound.Cs))
	length := csound.Ksmps() * csound.Channels(true)
	slice := []MYFLT(unsafe.Slice(buffer, length))
	return slice
}

// Returns the address of the Csound audio output working buffer (spout).
// Enables external software to read audio from Csound after calling
// csound.PerformKsmps.
func (csound CSOUND) Spout() []MYFLT {
	buffer := (*MYFLT)(C.csoundGetSpout(csound.Cs))
	length := csound.Ksmps() * csound.Channels(false)
	slice := []MYFLT(unsafe.Slice(buffer, length))
	return slice
}

/////////////////////////////
// Csound Messages and Text
/////////////////////////////

// Displays an informational message.
func (csound CSOUND) Message(format string, vals ...any) {
	s := fmt.Sprintf(format, vals...)
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C.cMessage(csound.Cs, cstr)
}

// Print message with special attributes (see msg_attr.h for the list of
// available attributes). With attr=0, csoundMessageS() is identical to
// csoundMessage().
func (csound CSOUND) MessageS(attr int, format string, vals ...any) {
	s := fmt.Sprintf(format, vals...)
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C.cMessageS(csound.Cs, C.int32_t(attr), cstr)
}

// Creates a buffer for storing messages printed by Csound.
// Should be called after creating a Csound instance and the buffer
// can be freed by calling csound.DestroyMessageBuffer() before
// deleting the Csound instance.
// If 'toStdOut' is true, the messages are also printed to
// stdout and stderr (depending on the type of the message),
// in addition to being stored in the buffer.
// Using the message buffer ties up the internal message callback, so
// csound.SetMessageCallback should not be called after creating the
// message buffer.
func (csound CSOUND) CreateMessageBuffet(toStdOut bool) {
	C.csoundCreateMessageBuffer(csound.Cs, cbool(toStdOut))
}

// Returns the first message from the buffer.
func (csound CSOUND) FirstMessage() string {
	return C.GoString(C.csoundGetFirstMessage(csound.Cs))
}

// Returns the attribute parameter (see msg_attr.h) of the first message
// in the buffer.
func (csound CSOUND) FirstMessageAttr() int {
	return int(C.csoundGetFirstMessageAttr(csound.Cs))
}

// Removes the first message from the buffer.
func (csound CSOUND) PopFirstMessage() {
	C.csoundPopFirstMessage(csound.Cs)
}

// Returns the number of pending messages in the buffer.
func (csound CSOUND) MessageCnt() int {
	return int(C.csoundGetMessageCnt(csound.Cs))
}

// Releases all memory used by the message buffer.
func (csound CSOUND) DestroyMessageBuffer() {
	C.csoundDestroyMessageBuffer(csound.Cs)
}

//////////////////////////////////
// Channels, Control, and Events
//////////////////////////////////

// Returns the var type for a channel name or "" if the channel
// was not found.
// Currently supported channel var types are "k" (control), "a" (audio),
// "S" (string), "f" (pvs), and "[" (array).
func (csound CSOUND) ChannelVarTypeName(name string) string {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return C.GoString(C.csoundGetChannelVarTypeName(csound.Cs, cname))
}

// Returns a list of allocated channels in a slice of ControlChannelInfo.
// A controlChannelInfo struct contains the channel characteristics.
// The err value is the number of channels, which may be zero if there
// are none, or CSOUND_MEMORY if there is not enough memory for allocating
// the list. In the case of no channels or an error, the slice is set to nil.
// Notes: the caller is responsible for freeing the list returned in *lst
// with csoundDeleteChannelList(). The name pointers may become invalid
// after calling csound.Reset().
func (csound CSOUND) ListChannels() ([]ControlChannelInfo, error) {
	var clist *C.controlChannelInfo_t
	n := int(C.csoundListChannels(csound.Cs, &clist))
	switch n {
	case CSOUND_MEMORY:
		return nil, fmt.Errorf("Not enough memory for allocating channels list")
	case 0:
		return nil, nil
	}
	var list = make([]ControlChannelInfo, n)
	var cname *C.char
	var ctype C.int32_t
	for i := range list {
		hints := C.getControlChannelInfo(clist, C.int(i), &cname, &ctype)
		list[i].Name = C.GoString(cname)
		list[i].Type = int(ctype)
		list[i].Hints.Behav = int(hints.behav)
		list[i].Hints.Dflt = MYFLT(hints.dflt)
		list[i].Hints.Min = MYFLT(hints.min)
		list[i].Hints.Max = MYFLT(hints.max)
		list[i].Hints.X = int(hints.x)
		list[i].Hints.Y = int(hints.y)
		list[i].Hints.Width = int(hints.width)
		list[i].Hints.Height = int(hints.height)
		list[i].Hints.Attributes = C.GoString(hints.attributes)
	}
	C.csoundDeleteChannelList(csound.Cs, clist)
	return list, nil
}

// Set parameters hints for a control channel. These hints have no internal
// function but can be used by front ends to construct GUIs or to constrain
// values. See the ControlChannelHints structure for details.
// Returns zero on success, or a non-zero error code on failure:
//
//	CSOUND_ERROR:  the channel does not exist, is not a control channel,
//	               or the specified parameters are invalid
//	CSOUND_MEMORY: could not allocate memory
func (csound CSOUND) SetControlChannelHints(name string, hints ControlChannelHints) int {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var chints C.controlChannelHints_t
	chints.behav = C.controlChannelBehavior(hints.Behav)
	chints.dflt = cMYFLT(hints.Dflt)
	chints.min = cMYFLT(hints.Min)
	chints.max = cMYFLT(hints.Max)
	chints.x = C.int32_t(hints.X)
	chints.y = C.int32_t(hints.Y)
	chints.width = C.int32_t(hints.Width)
	chints.height = C.int32_t(hints.Height)
	if len(hints.Attributes) > 0 {
		chints.attributes = C.CString(hints.Attributes)
	}
	return int(C.csoundSetControlChannelHints(csound.Cs, cname, chints))
}

// Return special parameters (assuming there are any) of a control channel,
// previously set with csound.SetControlChannelHints() or the chnparams
// opcode.
// If the channel exists, is a control channel, the channel hints
// are stored in the ControlChannelHints structure.
//
// The return value is zero if the channel exists and is a control
// channel, otherwise, an error code is returned.
func (csound CSOUND) ControlChannelHints(name string) (ControlChannelHints, int) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var chints C.controlChannelHints_t
	ret := C.csoundGetControlChannelHints(csound.Cs, cname, &chints)
	var hints ControlChannelHints
	if ret == 0 {
		hints.Behav = int(chints.behav)
		hints.Dflt = MYFLT(chints.dflt)
		hints.Min = MYFLT(chints.min)
		hints.Max = MYFLT(chints.max)
		hints.X = int(chints.x)
		hints.Y = int(chints.y)
		hints.Width = int(chints.width)
		hints.Height = int(chints.height)
		hints.Attributes = C.GoString(chints.attributes)
		if chints.attributes != nil {
			defer C.free(unsafe.Pointer(chints.attributes))
		}
	}
	return hints, int(ret)
}

// Locks access to the channel allowing access to data in
// a threadsafe manner.
func (csound CSOUND) LockChannel(channel string) {
	cchan := C.CString(channel)
	defer C.free(unsafe.Pointer(cchan))
	C.csoundLockChannel(csound.Cs, cchan)
}

// Unlocks access to the channel, allowing access to data from
// elsewhere.
func (csound CSOUND) UnlockChannel(channel string) {
	cchan := C.CString(channel)
	defer C.free(unsafe.Pointer(cchan))
	C.csoundUnlockChannel(csound.Cs, cchan)
}

// Retrieves the value of control channel identified by *name.
// If the err argument is not NULL, the error (or success) code
// finding or accessing the channel is stored in it.
func (csound CSOUND) ControlChannel(name string) (MYFLT, int) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var cerr C.int32_t
	cval := C.csoundGetControlChannel(csound.Cs, cname, &cerr)
	if cerr != 0 {
		return 0, int(cerr)
	}
	return MYFLT(cval), 0
}

// Sets the value of control channel identified by name.
func (csound CSOUND) SetControlChannel(name string, val MYFLT) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.csoundSetControlChannel(csound.Cs, cname, cMYFLT(val))
}

// Copies the audio channel identified by name into slice
// samples which should contain enough memory for ksmps MYFLTs
func (csound CSOUND) AudioChannel(name string, samples []MYFLT) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.csoundGetAudioChannel(csound.Cs, cname, cpMYFLT(&samples[0]))
}

// Sets the audio channel identified by name with data from slice
// samples which should contain at least ksmps MYFLTs
func (csound CSOUND) SetAudioChannel(name string, samples []MYFLT) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.csoundSetAudioChannel(csound.Cs, cname, cpMYFLT(&samples[0]))
}

// Returns a copy of the string channel identified by name.
func (csound CSOUND) StringChannel(name string) string {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	length := int(C.csoundGetChannelDatasize(csound.Cs, cname))
	buf := make([]byte, length)
	C.csoundGetStringChannel(csound.Cs, cname, (*C.char)(unsafe.Pointer(&buf[0])))
	return string(buf)
}

// Sets the string channel identified by name with stringVal.
func (csound CSOUND) SetStringChannel(name, stringVal string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cstrval := C.CString(stringVal)
	defer C.free(unsafe.Pointer(cstrval))
	C.csoundSetStringChannel(csound.Cs, cname, cstrval)
}

// Create and initialise an array channel with a given array type
//   - "a" (audio sigs): each item is a ksmps-size MYFLT array
//   - "i" (init vars): each item is a MYFLT
//   - "S" (strings): each item is a STRINGDAT (see csound.StringData() and
//     csound.SetStringData())
//   - "k" (control sigs): each item is a MYFLT
//     sizes - sizes for each dimension
//
// returns the ARRAYDAT for the requested channel or NULL on error
// NB: if the channel exists and has already been initialised,
// this function is a non-op.
func (csound CSOUND) InitArrayChannel(name, type_ string, sizes []int) ARRAYDAT {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	ctype := C.CString(type_)
	defer C.free(unsafe.Pointer(ctype))
	dimensions := len(sizes)
	csizes := make([]C.int32_t, dimensions)
	for i := range dimensions {
		csizes[i] = C.int32_t(sizes[i])
	}
	return C.csoundInitArrayChannel(csound.Cs, cname, ctype, C.int32_t(dimensions),
		&csizes[0])
}

// Get the type of data the ARRAYDAT adat, returning
//   - "a" (audio sigs): each item is a ksmps-size MYFLT array
//   - "i" (init vars): each item is a MYFLT
//   - "S" (strings): each item is a STRINGDAT (see csound.GetStringData() and
//     csound.SetStringData())
//   - "k" (control sigs): each item is a MYFLT
func (csound CSOUND) ArrayDataType(adat ARRAYDAT) string {
	return C.GoString(C.csoundArrayDataType(adat))
}

// Get the dimensions of the ARRAYDAT adat.
func (csound CSOUND) ArrayDataDimensions(adat ARRAYDAT) int {
	return int(C.csoundArrayDataDimensions(adat))
}

// Get the sizes of each dimension of the ARRAYDAT adat.
func (csound CSOUND) ArrayDataSizes(adat ARRAYDAT) []int {
	dimensions := int(C.csoundArrayDataDimensions(adat))
	sizes := make([]int, dimensions)
	pcsizes := C.csoundArrayDataSizes(adat)
	for i := range dimensions {
		sizes[i] = int(C.getInt32Val(pcsizes, C.int(i)))
	}
	return sizes
}

// Set the data in the ARRAYDAT adat.
func (csound CSOUND) SetArrayData(adat ARRAYDAT, data unsafe.Pointer) {
	C.csoundSetArrayData(adat, data)
}

// Get the data from the ARRAYDAT adat.
func (csound CSOUND) ArrayData(adat ARRAYDAT) unsafe.Pointer {
	return C.csoundGetArrayData(adat)
}

// Get a string from a STRINGDAT structure
func (csound CSOUND) StringData(sdata STRINGDAT) string {
	return C.GoString(C.csoundGetStringData(csound.Cs, sdata))
}

// Set a STRINGDAT structure with a string
func (csound CSOUND) SetStringData(sdata STRINGDAT, str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.csoundSetStringData(csound.Cs, sdata, cstr)
}

// Create/initialise an Fsig channel with
// size - FFT analysis size
// overlap - analysis overlap size
// winsize - analysis window size
// wintype - analysis window type (see pvsdat types enumeration)
// format - analysis data format (see pvsdat format enumeration)
// returns the PVSDAT for the requested channel or NULL on error
// NB: if the channel exists and has already been initialised,
// this function is a non-op.
func (csound CSOUND) InitPvsChannel(name string, size, overlap, winsize,
	wintype, format int) PVSDAT {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return C.csoundInitPvsChannel(csound.Cs, cname, C.int32_t(size),
		C.int32_t(overlap), C.int32_t(winsize), C.int32_t(wintype), C.int32_t(format))
}

// Get the analysis FFT size used by the PVSDAT pvsdat.
func (csound CSOUND) PvsDataFFTSize(pvsdat PVSDAT) int {
	return int(C.csoundPvsDataFFTSize(pvsdat))
}

// Get the analysis overlap size used by the PVSDAT pvsdat.
func (csound CSOUND) PvsDataOverlap(pvsdat PVSDAT) int {
	return int(C.csoundPvsDataOverlap(pvsdat))
}

// Get the analysis window size used by the PVSDAT pvsdat.
func (csound CSOUND) PvsDataWindowSize(pvsdat PVSDAT) int {
	return int(C.csoundPvsDataWindowSize(pvsdat))
}

// Get the analysis data format used by the PVSDAT pvsdat.
func (csound CSOUND) PvsDataFormat(pvsdat PVSDAT) int {
	return int(C.csoundPvsDataFormat(pvsdat))
}

// Get the current framecount from PVSDAT pvsdat.
func (csound CSOUND) PvsDataFramecount(pvsdat PVSDAT) int {
	return int(C.csoundPvsDataFramecount(pvsdat))
}

// Returns the size of data stored in a channel; for string channels
// this might change if the channel space gets reallocated
// Since string variables use dynamic memory allocation
// this function can be called to get the space required for
// csound.GetStringChannel().
func (csound CSOUND) ChannelDataSize(name string) int {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return int(C.csoundGetChannelDatasize(csound.Cs, cname))
}

// Schedule a new realtime event. 'type_' is the event type
// type 0 - instrument instance     CS_INSTR_EVENT
// type 1 - function table instance CS_TABLE_EVENT
// type 2 - end event               CS_END_EVENT
// event parameters is a MYFLT slice with the event parameters (p-fields)
// optionally run asynchronously (async = true)
// NB: This is non-op before csound.Start() is called.
func (csound CSOUND) Event(type_ int, params []MYFLT, async bool) {
	C.csoundEvent(csound.Cs, C.int32_t(type_), cpMYFLT(&params[0]),
		C.int32_t(len(params)), cbool(async))
}

// Schedule new score or realtime event(s) as a string
// Two operation modes are supported:
// - Score events: any calls before csound.Start() add the string events to
// the score (before pre-processing) (async should be set to false).
// - Realtime events: after the engine starts, string events are added to
// the realtime event queue.
//
// Multiple events separated by newlines are possible
// and score preprocessing (carry, etc) is applied.
// optionally run asynchronously (async = true)
func (csound CSOUND) EventString(message string, async bool) {
	cmsg := C.CString(message)
	defer C.free(unsafe.Pointer(cmsg))
	C.csoundEventString(csound.Cs, cmsg, cbool(async))
}

// Get the instrument number for a given instrument name string
// for use in numeric parameters list (csound.Event()).
// Returns the instrument number or -1 if not found.
func (csound CSOUND) InstrNumber(name string) int {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return int(C.csoundGetInstrNumber(csound.Cs, cname))
}

// Set the ASCII code of the most recent key pressed.
// This value is used by the 'sensekey' opcode if a callback
// for returning keyboard events is not set (see
// csound.RegisterKeyboardCallback()).
func (csound CSOUND) KeyPress(c byte) {
	C.csoundKeyPress(csound.Cs, C.char(c))
}

///////////
// Tables
///////////

// Returns the length of a function table (not including the guard point),
// or -1 if the table does not exist.
func (csound CSOUND) TableLength(tableNum int) int {
	return int(C.csoundTableLength(csound.Cs, C.int32_t(tableNum)))
}

// Returns a slice of table 'tableNum' and the table length
// (not including the guard point).
// If the table does not exist, the slice is set to nil and
// -1 is returned as the table length..
// NB: this function and the slice returned are not threadsafe
func (csound CSOUND) Table(tableNum int) ([]MYFLT, int) {
	var tablePtr *MYFLT
	length := int(C.csoundGetTable(csound.Cs, cppMYFLT(&tablePtr), C.int32_t(tableNum)))
	if length == -1 {
		return nil, -1
	}
	slice := []MYFLT(unsafe.Slice(tablePtr, length))
	return slice, length
}

// Returns a slice of MYFLTs containing the arguments used to generate
// function table 'tableNum'.
// If the table does not exist, an error is returned.
// NB: the argument list starts with the GEN number and is followed by
// its parameters. eg. f 1 0 1024 10 1 0.5  yields the list {10.0,1.0,0.5}
// This function and the slice returned are not threadsafe
func (csound CSOUND) TableArgs(tableNum int) ([]MYFLT, error) {
	var argsPtr *MYFLT
	length := int(C.csoundGetTableArgs(csound.Cs, cppMYFLT(&argsPtr),
		C.int32_t(tableNum)))
	if length == -1 {
		return nil, fmt.Errorf("Function table %d does not exist", tableNum)
	}
	slice := []MYFLT(unsafe.Slice(argsPtr, length))
	return slice, nil

}

// Copies an array stored in data to the function table
// number given by tableNum, which should exist in the engine.
// The input array should be at least as long as the table
// size plus one (guard point required).
// This function is threadsafe and can also be run asynchronously
func (csound CSOUND) TableCopyIn(tableNum int, data []MYFLT, async bool) {
	C.csoundTableCopyIn(csound.Cs, C.int32_t(tableNum), cpMYFLT(&data[0]),
		cbool(async))
}

// Copies a function table number given by tableNum,
// which should exist in the engine, into the slice data,
// and have enough space to accommodate the array size.
// This function is threadsafe and can also be run asynchronously
func (csound CSOUND) TableCopyOut(tableNum int, data []MYFLT, async bool) {
	C.csoundTableCopyOut(csound.Cs, C.int32_t(tableNum), cpMYFLT(&data[0]),
		cbool(async))
}

///////////////////
// Score Handling
///////////////////

// Returns the current score time in seconds
// since the beginning of performance.
func (csound CSOUND) ScoreTime() float64 {
	return float64(C.csoundGetScoreTime(csound.Cs))
}

// Sets whether Csound score events are performed or not, independently
// of real-time MIDI events (see csound.SetScorePending()).
func (csound CSOUND) IsScorePending() bool {
	return C.csoundIsScorePending(csound.Cs) != 0
}

// Sets whether Csound score events are performed or not (real-time
// events will continue to be performed). Can be used by external software,
// such as a VST host, to turn off performance of score events (while
// continuing to perform real-time events), for example to
// mute a Csound score while working on other tracks of a piece, or
// to play the Csound instruments live.
func (csound CSOUND) SetScorePending(pending bool) {
	C.csoundSetScorePending(csound.Cs, cbool(pending))
}

// Returns the score time beginning at which score events will
// actually immediately be performed (see csound.SetScoreOffsetSeconds()).
func (csound CSOUND) ScoreOffsetSeconds() {
	C.csoundGetScoreOffsetSeconds(csound.Cs)
}

// Csound score events prior to the specified time are not performed, and
// performance begins immediately at the specified time (real-time events
// will continue to be performed as they are received).
// Can be used by external software, such as a VST host,
// to begin score performance midway through a Csound score,
// for example to repeat a loop in a sequencer, or to synchronize
// other events with the Csound score.
func (csound CSOUND) SetScoreOffsetSeconds(time MYFLT) {
	C.csoundSetScoreOffsetSeconds(csound.Cs, C.MYFLT(time))
}

// Rewinds a compiled Csound score to the time specified with
// csound.SetScoreOffsetSeconds().
func (csound CSOUND) RewindScore() {
	C.csoundRewindScore(csound.Cs)
}

////////////
// Opcodes
////////////

// Loads all plugins from a given directory. Generally called
// immediately after csound.Create()
// to make new opcodes/modules available for compilation and performance.
func (csound CSOUND) LoadPlugins(dir string) int {
	cdir := C.CString(dir)
	defer C.free(unsafe.Pointer(cdir))
	return int(C.csoundLoadPlugins(csound.Cs, cdir))
}
