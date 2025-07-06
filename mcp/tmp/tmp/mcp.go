package sdk
import (
"fmt"
"strconv"
"encoding/hex"
"encoding/json"
)
type debuggers struct {}
func (debuggers) VmxSupportDetection() bool {
	return request[bool]("VmxSupportDetection", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) CpuReadVendorString(vendor_string string, )  {
	 request[void]("CpuReadVendorString", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) LoadVmmModule() int {
	return request[int]("LoadVmmModule", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) UnloadVmm() int {
	return request[int]("UnloadVmm", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) InstallVmmDriver() int {
	return request[int]("InstallVmmDriver", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) UninstallVmmDriver() int {
	return request[int]("UninstallVmmDriver", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) StopVmmDriver() int {
	return request[int]("StopVmmDriver", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) Interpreter(command string, ) int {
	return request[int]("Interpreter", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) TestCommandParser(command string, number_of_tokens uint32, tokens_list , failed_token_num , failed_token_position , ) bool {
	return request[bool]("TestCommandParser", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) TestCommandParserShowTokens(command string, )  {
	 request[void]("TestCommandParserShowTokens", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) ShowSignature()  {
	 request[void]("ShowSignature", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) SetTextMessageCallback()  {
	 request[void]("SetTextMessageCallback", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) SetTextMessageCallbackUsingSharedBuffer()  {
	 request[void]("SetTextMessageCallbackUsingSharedBuffer", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) UnsetTextMessageCallback()  {
	 request[void]("UnsetTextMessageCallback", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) ScriptReadFileAndExecuteCommandline(argc int, argv string, ) int {
	return request[int]("ScriptReadFileAndExecuteCommandline", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) ContinuePreviousCommand() bool {
	return request[bool]("ContinuePreviousCommand", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) CheckMultilineCommand(current_command string, reset bool, ) bool {
	return request[bool]("CheckMultilineCommand", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) ConnectLocalDebugger()  {
	 request[void]("ConnectLocalDebugger", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) ConnectRemoteDebugger(ip string, port string, ) bool {
	return request[bool]("ConnectRemoteDebugger", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) Continue()  {
	 request[void]("Continue", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) Pause()  {
	 request[void]("Pause", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) SetBreakPoint(address uint64, pid uint32, tid uint32, core_numer uint32, )  {
	 request[void]("SetBreakPoint", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) SetCustomDriverPath(driver_file_path string, driver_name string, ) bool {
	return request[bool]("SetCustomDriverPath", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) UseDefaultDriverPath()  {
	 request[void]("UseDefaultDriverPath", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) ReadMemory(target_address uint64, memory_type , reading_Type , pid , size , get_address_mode , address_mode , target_buffer_to_store , return_length , ) bool {
	return request[bool]("ReadMemory", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) ShowMemoryOrDisassemble(style , address uint64, memory_type , reading_type , pid uint32, size uint32, dt_details , )  {
	 request[void]("ShowMemoryOrDisassemble", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) ReadAllRegisters(guest_registers , extra_registers , ) bool {
	return request[bool]("ReadAllRegisters", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) xxxxxxxxxxxxxxxxxxx(xxxxxxxxxxxxxx , xxxxxxxxxx , )  {
	return request[]("xxxxxxxxxxxxxxxxxxx", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) ReadTargetRegister(xxxxxxxxxxxxxx , xxxxxxxxxx , ) bool {
	return request[bool]("ReadTargetRegister", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) WriteTargetRegister(register_id , value uint64, ) bool {
	return request[bool]("WriteTargetRegister", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) RegisterShowAll() bool {
	return request[bool]("RegisterShowAll", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) RegisterShowTargetRegister(register_id , ) bool {
	return request[bool]("RegisterShowTargetRegister", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) WriteMemory(destination_address , memory_type , process_id uint32, source_address , number_of_bytes uint32, ) bool {
	return request[bool]("WriteMemory", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) DebuggerGetKernelBase() uint64 {
	return request[uint64]("DebuggerGetKernelBase", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) DebugRemoteDeviceUsingComPort(port_name string, baudrate , pause_after_connection bool, ) bool {
	return request[bool]("DebugRemoteDeviceUsingComPort", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) DebugRemoteDeviceUsingNamedPipe(named_pipe string, pause_after_connection bool, ) bool {
	return request[bool]("DebugRemoteDeviceUsingNamedPipe", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) DebugCloseRemoteDebugger() bool {
	return request[bool]("DebugCloseRemoteDebugger", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) DebugCurrentDeviceUsingComPort(port_name string, baudrate , ) bool {
	return request[bool]("DebugCurrentDeviceUsingComPort", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) UdAttachToProcess(path , ) bool {
	return request[bool]("UdAttachToProcess", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) UdAttachToProcess(path , arguments , ) bool {
	return request[bool]("UdAttachToProcess", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) AssembleGetLength(assembly_code string, start_address uint64, length , ) bool {
	return request[bool]("AssembleGetLength", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) Assemble(assembly_code string, start_address uint64, buffer_to_store_assembled_data , buffer_size uint32, ) bool {
	return request[bool]("Assemble", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) SetupPathForFileName(filename string, file_location string, buffer_len uint32, check_file_existence bool, ) bool {
	return request[bool]("SetupPathForFileName", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) SteppingInstrumentationStepIn() bool {
	return request[bool]("SteppingInstrumentationStepIn", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) SteppingRegularStepIn() bool {
	return request[bool]("SteppingRegularStepIn", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) SteppingStepOver() bool {
	return request[bool]("SteppingStepOver", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) SteppingInstrumentationStepInForTracking() bool {
	return request[bool]("SteppingInstrumentationStepInForTracking", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) SteppingStepOverForGu(last_instruction bool, ) bool {
	return request[bool]("SteppingStepOverForGu", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) GetLocalApic(local_apic , is_using_x2apic , ) bool {
	return request[bool]("GetLocalApic", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) GetIoApic(io_apic , ) bool {
	return request[bool]("GetIoApic", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) GetIdtEntry(idt_packet , ) bool {
	return request[bool]("GetIdtEntry", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HwdbgScriptRunScript(script , instance_filepath_to_read , hardware_script_file_path_to_save , initial_bram_buffer_size uint32, ) bool {
	return request[bool]("HwdbgScriptRunScript", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) ScriptEngineWrapperTestParserForHwdbg(Expr , )  {
	 request[void]("ScriptEngineWrapperTestParserForHwdbg", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) EnableTransparentMode(ProcessId uint32, ProcessName string, IsProcessId bool, ) bool {
	return request[bool]("EnableTransparentMode", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) DisableTransparentMode() bool {
	return request[bool]("DisableTransparentMode", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
