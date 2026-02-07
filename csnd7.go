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

void cMessage(CSOUND *csound, const char *msg)
{
  csoundMessage(csound, "%s", msg);
}

void cMessageS(CSOUND *csound, int32_t attr, char *msg)
{
  csoundMessageS(csound, attr, "%s", msg);
}
*/
import "C"

import (
	"fmt"
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

type MYFLT float64

type CsoundParams struct {
	Odebug            C.int32_t // debug flag
	Sfread            C.int32_t // sound input read flag
	Sfwrite           C.int32_t // sound output write flag (-s)
	Filetyp           C.int32_t // soundfile type code
	Inbufsamps        C.int32_t // input buffer size in samples
	Outbufsamps       C.int32_t // output buffer size in samples
	Informat          C.int32_t // input soundfile format
	Outformat         C.int32_t // output soundfile format
	SndfileSampleSize C.int32_t // sample size
	Displays          C.int32_t // displays flag
	Graphsoff         C.int32_t // graphs flag
	Postscript        C.int32_t // postscript graphs flag
	Msglevel          C.int32_t // message level (-m)
	Beatmode          C.int32_t // beat mode
	OMaxLag           C.int32_t // hardware buffer size (samples)
	Linein            C.int32_t // linevents flag (-L)
	RTevents          C.int32_t // realtime events flag (scoreless, -L, -F, -M)
	Midiin            C.int32_t // midi input flag (-M)
	FMidiin           C.int32_t // midi file input flag (-F)
	RMidiin           C.int32_t // remote events flag
	Ringbell          C.int32_t // ringbell flag
	Termifend         C.int32_t // terminate on MIDI file input flag (-T)
	Tewrt_hdr         C.int32_t // rewrite header flag
	Heartbeat         C.int32_t // heartbeat flag
	Gen01defer        C.int32_t // GEN01 defer allocation flag
	CmdTempo          C.double  // tempo value (-t)
	Sr_override       MYFLT     // sampling rate override (-r)
	Kr_override       MYFLT     // control rate override (-k)
	Nchnls_override   C.int32_t // nchnls override
	Nchnls_i_override C.int32_t // nchnls_i override
	Infilename        *C.char   // input file name (-i)
	Outfilename       *C.char   // output file name (-o)
	Linename          *C.char   // line events source (-L)
	Midiname          *C.char   // MIDI input device name (-M)
	FMidiname         *C.char   // MIDI input file name (-F)
	Midioutname       *C.char   // MIDI output device name (-Q)
	FMidioutname      *C.char   // MIDI output file name
	MidiKey           C.int32_t // MIDI key pfield mapping
	MidiKeyCps        C.int32_t // MIDI key-cps pfield mapping
	MidiKeyOct        C.int32_t // MIDI key-oct pfield mapping
	MidiKeyPch        C.int32_t // MIDI key-pch pfield mapping
	MidiVelocity      C.int32_t // MIDI vel pfield mapping
	MidiVelocityAmp   C.int32_t // MIDI vel-amp pfield mapping
	NoDefaultPaths    C.int32_t // default paths flag
	NumThreads        C.int32_t // multicore number of threads (-j)
	SyntaxCheckOnly   C.int32_t // syntax check only flag
	RunUnitTests      C.int32_t // run unit tests flag
	UseCsdLineCounts  C.int32_t // csd line nums option
	SampleAccurate    C.int32_t // sample accurate flag
	Realtime          C.int32_t // realtime priority flag
	E0dbfs_override   MYFLT     // 0dbfs override
	Daemon            C.int32_t // daemon mode flag
	Quality           C.double  // OGG encoding quality
	Ksmps_override    C.int32_t // ksmps override
	Fft_lib           C.int32_t // FFT library option
	Echo              C.int32_t // UDP echo commands flag
	Limiter           MYFLT     // audio output limiter option
	Sr_default        MYFLT     // default sampling rate
	Kr_default        MYFLT     // default control rate
	Mp3_mode          C.int32_t // MP3 encoding mode
	Redef             C.int32_t // instr redefinition flag
	Error_deprecated  C.int32_t // error on deprecated opcodes
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

/*
 * Instantiation
 */

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
func (csound *CSOUND) Destroy() {
	C.csoundDestroy(csound.Cs)
	csound.Cs = nil
}

/*
 * Attributes
 */

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

/*
 * Returns the number of audio channels in the Csound instance.
 * If isInput is false, the value of nchnls is returned,
 * otherwise nchnls_i.
 */
func (csound CSOUND) Channels(isInput bool) int {
	return int(C.csoundGetChannels(csound.Cs, cbool(isInput)))
}

/*
 * Returns the 0dBFS level of the spin/spout buffers.
 */
func (csound CSOUND) Get0dbFS() MYFLT {
	return MYFLT(C.csoundGet0dBFS(csound.Cs))
}

/*
 * Returns the A4 frequency reference
 */
func (csound CSOUND) A4() MYFLT {
	return MYFLT(C.csoundGetA4(csound.Cs))
}

/*
 * Returns the current performance time in sample frames
 */
func (csound CSOUND) CurrentTimeSamples() int {
	return int(C.csoundGetCurrentTimeSamples(csound.Cs))
}

/*
 * Returns the size of MYFLT in bytes.
 */
func (csound CSOUND) SizeOfMYFLT() int {
	return int(C.csoundGetSizeOfMYFLT())
}

/*
 * Returns host data.
 */
func (csound CSOUND) HostData() unsafe.Pointer {
	return C.csoundGetHostData(csound.Cs)
}

/*
 * Sets host data.
 */
func (csound CSOUND) SetHostData(hostData unsafe.Pointer) {
	C.csoundSetHostData(csound.Cs, hostData)
}

/*
 * Returns the total error count of the current performance.
 */
func (csound CSOUND) ErrCnt() int {
	return int(C.csoundErrCnt(csound.Cs))
}

/*
 * Get the value of environment variable 'name', searching
 * in this order: local environment of 'csound' (if not NULL), variables
 * set with csound.SetGlobalEnv(), and system environment variables.
 * If 'csound' is not NULL, should be called after csound.Compile().
 * Return value is "" if the variable is not set.
 */
func (csound CSOUND) Env(name string) string {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return C.GoString(C.csoundGetEnv(csound.Cs, cname))
}

/*
 * Set the global value of environment variable 'name' to 'value',
 * or delete variable if 'value' is "".
 * It is not safe to call this function while any Csound instances
 * are active.
 * Returns zero on success.
 */
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

/*
 * Set csound options (flag). Returns CSOUND_SUCCESS on success.
 * This needs to be called after Create() and before any code is
 * compiled. Multiple options are allowed in one string.
 */
func (csound CSOUND) SetOption(option string) int {
	coption := C.CString(option)
	defer C.free(unsafe.Pointer(coption))
	return int(C.csoundSetOption(csound.Cs, coption))
}

/*
 * Get the current set of parameters from a CSOUND instance in
 * a CSOUND_PARAMS structure.
 */
func (csound CSOUND) Params() *CsoundParams {
	p := unsafe.Pointer(C.csoundGetParams(csound.Cs))
	return (*CsoundParams)(p)
}

/*
 * Returns whether Csound is set to print debug messages sent through the
 * DebugMsg() internal API function. Anything different to 0 means true.
 */
func (csound CSOUND) Debug() bool {
	return C.csoundGetDebug(csound.Cs) != 0
}

/*
 * Sets whether Csound prints debug messages from the DebugMsg() internal
 * API function. Anything different to 0 means true.
 */
func (csound CSOUND) SetDebug(debug bool) {
	C.csoundSetDebug(csound.Cs, cbool(debug))
}

/*
 * If val > 0, sets the internal variable holding the system HW sr.
 * Returns the stored value containing the system HW sr.
 */
func (csound CSOUND) SystemSr(val MYFLT) MYFLT {
	return MYFLT(C.csoundSystemSr(csound.Cs, cMYFLT(val)))
}

/*
 * Retrieves a module name and type ("audio" or "midi").
 * Given a number Modules are added to list as csound loads them.
 * Returns CSOUND_SUCCESS on success and CSOUND_ERROR if module
 * number was not found
 *
 *      for n :=0; ; n++ {
 *          name, type, err := csound.Module(n)
 *          if err != CSOUND_SUCCESS {
 *              break
 *          }
 *          fmt.printf("Module %d:  %s (%s) \n", n, name, type)
 *
 */
func (csound CSOUND) Module(n int) (name, type_ string, err int) {
	var cname, ctype *C.char
	cerror := C.csoundGetModule(csound.Cs, C.int32_t(n), &cname, &ctype)
	name = C.GoString(cname)
	type_ = C.GoString(ctype)
	err = int(cerror)
	return
}

/*
 * This function can be called to obtain a list of available
 * input or output audio devices available (depending on the backend
 * module used).
 *
 *     list := csound.AudioDevList(true)
 *     for i := range list {
 *         csound.Message(" %d: %s (%s)\n",
 *             i, list[i].DeviceId, list[i].DeviceName);
 *
 */
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

/*
 * Returns the Csound message level (from 0 to 231).
 */
func (csound CSOUND) MessageLevel() int {
	return int(C.csoundGetMessageLevel(csound.Cs))
}

/*
 * Sets the Csound message level (from 0 to 231).
 */
func (csound CSOUND) SetMessageLevel(messageLevel int) {
	C.csoundSetMessageLevel(csound.Cs, C.int32_t(messageLevel))
}

/*
 * Compilation and Performance
 */

/*
 * Compiles Csound input files (such as an orchestra and score, or CSD)
 * as directed by the supplied command-line arguments,
 * but does not perform them. Returns a non-zero error code on failure.
 * In this mode, the sequence of calls should be as follows:
 *
 *     csound.Compile(args)
 *     csound.Start()
 *     for !csound.PerformKsmps() {}
 *     csound.Reset()
 */
func (csound CSOUND) Compile(args []string) int {
	argc := C.int32_t(len(args))
	argv := make([]*C.char, argc)
	for i, arg := range args {
		argv[i] = C.CString(arg)
		defer C.free(unsafe.Pointer(argv[i]))
	}
	return int(C.csoundCompile(csound.Cs, argc, &argv[0]))
}

/*
 * Parse, and compile the given orchestra from an ASCII string,
 * also evaluating any global space code (i-time only)
 * in synchronous or asynchronous (async = true) mode.
 *
 *     orc := "instr 1 \n a1 = rand(0dbfs/4) \n out(a1) \n"
 *     csound.CompileOrc(orc, false);
 */
func (csound CSOUND) CompileOrc(code string, async bool) int {
	ccode := C.CString(code)
	defer C.free(unsafe.Pointer(ccode))
	return int(C.csoundCompileOrc(csound.Cs, ccode, cbool(async)))
}

/*
 * Compiles a Csound input file (CSD, .csd file) or a txt string
 * containing the CSD code, in synchronous or asynchronous (async = true) mode.
 * Returns a non-zero error code on failure.
 *
 * If csound.Start is called before csound.CompileCSD, the <CsOptions>
 * element is ignored (but csound.SetOption can be called any number of
 * times), the <CsScore> element is dispatched as score events (e.g.
 * as it is done by csoundEventString())
 *
 *     csound.SetOption("-an_option")
 *     csound.SetOption("-another_option")
 *     csound.Start()
 *     csound.CompileCSD(csd_filename, 0, false)
 *     for {
 *       csound.PerformKsmps()
 *       // Something to break out of the loop
 *       // when finished here...
 *     }
 *
 * NB: this function can be called repeatedly during performance to
 * replace or add new instruments and events.
 *
 * But if csound.CompileCsd is called before csound.Start, the <CsOptions>
 * element is used, the <CsScore> section is pre-processed and dispatched
 * normally, and performance terminates when the score terminates.
 *
 *     csound.CompileCSD(csound, csd_filename, 0, false)
 *     csound.Start()
 *     for !csound.PerformKsmps() {
 *     }
 *
 * if mode = 1, csd contains a full CSD code (rather than a filename).
 * This is convenient when it is desirable to package the csd as part of
 * an application or a multi-language piece.
 */
func (csound CSOUND) CompileCSD(csd string, mode int, async bool) int {
	csdName := C.CString(csd)
	defer C.free(unsafe.Pointer(csdName))
	return int(C.csoundCompileCSD(csound.Cs, csdName, C.int32_t(mode), cbool(async)))
}

/*
 * Prepares Csound for performance. Normally called after compiling
 * a csd file or an orc file, in which case score preprocessing is
 * performed and performance terminates when the score terminates.
 *
 * However, if called before compiling a csd file or an orc file,
 * score preprocessing is not performed and "i" statements are dispatched
 * as real-time events, the <CsOptions> tag is ignored, and performance
 * continues indefinitely or until ended using the API.
 */
func (csound CSOUND) Start() int {
	return int(C.csoundStart(csound.Cs))
}

/*
 * Senses input events, and performs one block of
 * audio output containing ksmps frames. csound.Start() must be called first.
 * Returns false during performance, and true when performance is finished.
 * If called until it returns true, will perform an entire score.
 * Enables external software to control the execution of Csound,
 * and to synchronize performance with audio input and output.
 */
func (csound CSOUND) PerformKsmps() bool {
	return C.csoundPerformKsmps(csound.Cs) != 0
}

/*
 * Resets all internal memory and state in preparation for a new performance.
 * Enables external software to run successive Csound performances
 * without reloading Csound.
 */
func (csound CSOUND) Reset() {
	C.csoundReset(csound.Cs)
}

/*
 * Audio I/O
 */

/*
 * Returns the address of the Csound audio input working buffer (spin).
 * Enables external software to write audio into Csound before calling
 * csound.PerformKsmps.
 */
func (csound CSOUND) Spin() []MYFLT {
	buffer := (*MYFLT)(C.csoundGetSpin(csound.Cs))
	length := csound.Ksmps() * csound.Channels(true)
	slice := []MYFLT(unsafe.Slice(buffer, length))
	return slice
}

/*
 * Returns the address of the Csound audio output working buffer (spout).
 * Enables external software to read audio from Csound after calling
 * csound.PerformKsmps.
 */
func (csound CSOUND) Spout() []MYFLT {
	buffer := (*MYFLT)(C.csoundGetSpout(csound.Cs))
	length := csound.Ksmps() * csound.Channels(false)
	slice := []MYFLT(unsafe.Slice(buffer, length))
	return slice
}

/*
 * Csound Messages and Text
 */

/*
 * Displays an informational message.
 */
func (cs CSOUND) Message(format string, vals ...any) {
	s := fmt.Sprintf(format, vals...)
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C.cMessage(cs.Cs, cstr)
}

/*
 * Channels, Control, and Events
 */

/*
 * Sets the value of control channel identified by name.
 */
func (csound CSOUND) SetControlChannel(name string, val MYFLT) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.csoundSetControlChannel(csound.Cs, cname, cMYFLT(val))
}

/*
 * Schedule new score or realtime event(s) as a string
 * Two operation modes are supported:
 * - Score events: any calls before csound.Start() add the string events to
 * the score (before pre-processing) (async should be set to false).
 * - Realtime events: after the engine starts, string events are added to
 * the realtime event queue.
 *
 * Multiple events separated by newlines are possible
 * and score preprocessing (carry, etc) is applied.
 * optionally run asynchronously (async = true)
 */
func (csound CSOUND) EventString(message string, async bool) {
	cmsg := C.CString(message)
	defer C.free(unsafe.Pointer(cmsg))
	C.csoundEventString(csound.Cs, cmsg, cbool(async))
}

/*
 * Tables
 */

/*
 * Returns the length of a function table (not including the guard point),
 * or -1 if the table does not exist.
 */
func (csound CSOUND) TableLenght(tableNum int) int {
	return int(C.csoundTableLength(csound.Cs, C.int32_t(tableNum)))
}

/*
 * Returns a slice of table 'tableNum' and the table length
 * (not including the guard point).
 * If the table does not exist, the slice is set to nil and
 * -1 is returned as the table length..
 * NB: this function and the slice returned are not threadsafe
 */
func (csound CSOUND) Table(tableNum int) ([]MYFLT, int) {
	var tablePtr *MYFLT
	length := int(C.csoundGetTable(csound.Cs, cppMYFLT(&tablePtr), C.int32_t(tableNum)))
	if length == -1 {
		return nil, -1
	}
	slice := []MYFLT(unsafe.Slice(tablePtr, length))
	return slice, length

}
