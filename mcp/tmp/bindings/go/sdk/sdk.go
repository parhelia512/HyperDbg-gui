package sdk
import (
"fmt"
"strconv"
"encoding/hex"
"encoding/json"
)
type debugger struct {}
func (debugger) VmxSupportDetection() bool {
	return request[bool]("VmxSupportDetection", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) CpuReadVendorString(vendor_string string, )  {
	 request[void]("CpuReadVendorString", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) LoadVmmModule() int {
	return request[int]("LoadVmmModule", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) UnloadVmm() int {
	return request[int]("UnloadVmm", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) InstallVmmDriver() int {
	return request[int]("InstallVmmDriver", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) UninstallVmmDriver() int {
	return request[int]("UninstallVmmDriver", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) StopVmmDriver() int {
	return request[int]("StopVmmDriver", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) Interpreter(command string, ) int {
	return request[int]("Interpreter", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) TestCommandParser(command string, number_of_tokens uint32, tokens_list , failed_token_num , failed_token_position , ) bool {
	return request[bool]("TestCommandParser", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) TestCommandParserShowTokens(command string, )  {
	 request[void]("TestCommandParserShowTokens", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) ShowSignature()  {
	 request[void]("ShowSignature", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) SetTextMessageCallback()  {
	 request[void]("SetTextMessageCallback", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) SetTextMessageCallbackUsingSharedBuffer()  {
	 request[void]("SetTextMessageCallbackUsingSharedBuffer", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) UnsetTextMessageCallback()  {
	 request[void]("UnsetTextMessageCallback", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) ScriptReadFileAndExecuteCommandline(argc int, argv string, ) int {
	return request[int]("ScriptReadFileAndExecuteCommandline", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) ContinuePreviousCommand() bool {
	return request[bool]("ContinuePreviousCommand", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) CheckMultilineCommand(current_command string, reset bool, ) bool {
	return request[bool]("CheckMultilineCommand", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) ConnectLocalDebugger()  {
	 request[void]("ConnectLocalDebugger", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) ConnectRemoteDebugger(ip string, port string, ) bool {
	return request[bool]("ConnectRemoteDebugger", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) Continue()  {
	 request[void]("Continue", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) Pause()  {
	 request[void]("Pause", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) SetBreakPoint(address uint64, pid uint32, tid uint32, core_numer uint32, )  {
	 request[void]("SetBreakPoint", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) SetCustomDriverPath(driver_file_path string, driver_name string, ) bool {
	return request[bool]("SetCustomDriverPath", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) UseDefaultDriverPath()  {
	 request[void]("UseDefaultDriverPath", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) ReadMemory(target_address uint64, memory_type , reading_Type , pid uint32, size uint32, get_address_mode bool, address_mode , target_buffer_to_store , return_length , ) bool {
	return request[bool]("ReadMemory", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) ShowMemoryOrDisassemble(style , address uint64, memory_type , reading_type , pid uint32, size uint32, dt_details , )  {
	 request[void]("ShowMemoryOrDisassemble", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) ReadAllRegisters(guest_registers , extra_registers , ) bool {
	return request[bool]("ReadAllRegisters", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) xxxxxxxxxxxxxxxxxxx(xxxxxxxxxxxxxx , xxxxxxxxxx , )  {
	return request[]("xxxxxxxxxxxxxxxxxxx", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) ReadTargetRegister(xxxxxxxxxxxxxx , xxxxxxxxxx , ) bool {
	return request[bool]("ReadTargetRegister", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) WriteTargetRegister(register_id , value uint64, ) bool {
	return request[bool]("WriteTargetRegister", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) RegisterShowAll() bool {
	return request[bool]("RegisterShowAll", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) RegisterShowTargetRegister(register_id , ) bool {
	return request[bool]("RegisterShowTargetRegister", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) WriteMemory(destination_address , memory_type , process_id uint32, source_address , number_of_bytes uint32, ) bool {
	return request[bool]("WriteMemory", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) DebuggerGetKernelBase() uint64 {
	return request[uint64]("DebuggerGetKernelBase", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) DebugRemoteDeviceUsingComPort(port_name string, baudrate , pause_after_connection bool, ) bool {
	return request[bool]("DebugRemoteDeviceUsingComPort", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) DebugRemoteDeviceUsingNamedPipe(named_pipe string, pause_after_connection bool, ) bool {
	return request[bool]("DebugRemoteDeviceUsingNamedPipe", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) DebugCloseRemoteDebugger() bool {
	return request[bool]("DebugCloseRemoteDebugger", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) DebugCurrentDeviceUsingComPort(port_name string, baudrate , ) bool {
	return request[bool]("DebugCurrentDeviceUsingComPort", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) UdAttachToProcess(path string, ) bool {
	return request[bool]("UdAttachToProcess", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) UdAttachToProcess(path string, arguments string, ) bool {
	return request[bool]("UdAttachToProcess", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) AssembleGetLength(assembly_code string, start_address uint64, length , ) bool {
	return request[bool]("AssembleGetLength", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) Assemble(assembly_code string, start_address uint64, buffer_to_store_assembled_data , buffer_size uint32, ) bool {
	return request[bool]("Assemble", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) SetupPathForFileName(filename string, file_location string, buffer_len uint32, check_file_existence bool, ) bool {
	return request[bool]("SetupPathForFileName", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) SteppingInstrumentationStepIn() bool {
	return request[bool]("SteppingInstrumentationStepIn", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) SteppingRegularStepIn() bool {
	return request[bool]("SteppingRegularStepIn", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) SteppingStepOver() bool {
	return request[bool]("SteppingStepOver", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) SteppingInstrumentationStepInForTracking() bool {
	return request[bool]("SteppingInstrumentationStepInForTracking", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) SteppingStepOverForGu(last_instruction bool, ) bool {
	return request[bool]("SteppingStepOverForGu", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) GetLocalApic(local_apic , is_using_x2apic , ) bool {
	return request[bool]("GetLocalApic", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) GetIoApic(io_apic , ) bool {
	return request[bool]("GetIoApic", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) GetIdtEntry(idt_packet , ) bool {
	return request[bool]("GetIdtEntry", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) HwdbgScriptRunScript(script , instance_filepath_to_read , hardware_script_file_path_to_save , initial_bram_buffer_size uint32, ) bool {
	return request[bool]("HwdbgScriptRunScript", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) ScriptEngineWrapperTestParserForHwdbg(Expr , )  {
	 request[void]("ScriptEngineWrapperTestParserForHwdbg", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) EnableTransparentMode(ProcessId uint32, ProcessName string, IsProcessId bool, ) bool {
	return request[bool]("EnableTransparentMode", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debugger) DisableTransparentMode() bool {
	return request[bool]("DisableTransparentMode", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
