package sdk
import (
"encoding/hex"
"encoding/json"
"fmt"
"strconv"
)
type debuggers struct {}
func (debuggers) VmxSupportDetection() bool {
	return request[bool]("VmxSupportDetection", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) CpuReadVendorString(vendor_string string, ) void {
	return request[void]("CpuReadVendorString", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgLoadVmmModule() int {
	return request[int]("HyperDbgLoadVmmModule", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgUnloadVmm() int {
	return request[int]("HyperDbgUnloadVmm", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgInstallVmmDriver() int {
	return request[int]("HyperDbgInstallVmmDriver", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgUninstallVmmDriver() int {
	return request[int]("HyperDbgUninstallVmmDriver", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgStopVmmDriver() int {
	return request[int]("HyperDbgStopVmmDriver", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgInterpreter(command string, ) int {
	return request[int]("HyperDbgInterpreter", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgTestCommandParser(command string, number_of_tokens uint32, tokens_list , failed_token_num , failed_token_position , ) bool {
	return request[bool]("HyperDbgTestCommandParser", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgTestCommandParserShowTokens(command string, ) void {
	return request[void]("HyperDbgTestCommandParserShowTokens", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgShowSignature() void {
	return request[void]("HyperDbgShowSignature", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) SetTextMessageCallback() void {
	return request[void]("SetTextMessageCallback", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) SetTextMessageCallbackUsingSharedBuffer() void {
	return request[void]("SetTextMessageCallbackUsingSharedBuffer", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) UnsetTextMessageCallback() void {
	return request[void]("UnsetTextMessageCallback", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
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
func (debuggers) ConnectLocalDebugger() void {
	return request[void]("ConnectLocalDebugger", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) ConnectRemoteDebugger(ip string, port string, ) bool {
	return request[bool]("ConnectRemoteDebugger", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) Continue() void {
	return request[void]("Continue", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) Pause() void {
	return request[void]("Pause", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) SetBreakPoint(address uint64, pid uint32, tid uint32, core_numer uint32, ) void {
	return request[void]("SetBreakPoint", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) SetCustomDriverPath(driver_file_path string, driver_name string, ) bool {
	return request[bool]("SetCustomDriverPath", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) UseDefaultDriverPath() void {
	return request[void]("UseDefaultDriverPath", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgReadMemory(target_address uint64, memory_type , reading_Type , pid , size , get_address_mode , address_mode , target_buffer_to_store , return_length , ) bool {
	return request[bool]("HyperDbgReadMemory", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgShowMemoryOrDisassemble(style , address uint64, memory_type , reading_type , pid uint32, size uint32, dt_details , ) void {
	return request[void]("HyperDbgShowMemoryOrDisassemble", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgReadAllRegisters(guest_registers , extra_registers , ) bool {
	return request[bool]("HyperDbgReadAllRegisters", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) xxxxxxxxxxxxxxxxxxx(xxxxxxxxxxxxxx , xxxxxxxxxx , )  {
	return request[]("xxxxxxxxxxxxxxxxxxx", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgReadTargetRegister(xxxxxxxxxxxxxx , xxxxxxxxxx , ) bool {
	return request[bool]("HyperDbgReadTargetRegister", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgWriteTargetRegister(register_id , value uint64, ) bool {
	return request[bool]("HyperDbgWriteTargetRegister", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgRegisterShowAll() bool {
	return request[bool]("HyperDbgRegisterShowAll", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgRegisterShowTargetRegister(register_id , ) bool {
	return request[bool]("HyperDbgRegisterShowTargetRegister", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgWriteMemory(destination_address , memory_type , process_id uint32, source_address , number_of_bytes uint32, ) bool {
	return request[bool]("HyperDbgWriteMemory", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) DebuggerGetKernelBase() uint64 {
	return request[uint64]("DebuggerGetKernelBase", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgDebugRemoteDeviceUsingComPort(port_name string, baudrate , pause_after_connection bool, ) bool {
	return request[bool]("HyperDbgDebugRemoteDeviceUsingComPort", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgDebugRemoteDeviceUsingNamedPipe(named_pipe string, pause_after_connection bool, ) bool {
	return request[bool]("HyperDbgDebugRemoteDeviceUsingNamedPipe", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgDebugCloseRemoteDebugger() bool {
	return request[bool]("HyperDbgDebugCloseRemoteDebugger", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgDebugCurrentDeviceUsingComPort(port_name string, baudrate , ) bool {
	return request[bool]("HyperDbgDebugCurrentDeviceUsingComPort", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) UdAttachToProcess(path , ) bool {
	return request[bool]("UdAttachToProcess", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) UdAttachToProcess(path , arguments , ) bool {
	return request[bool]("UdAttachToProcess", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgAssembleGetLength(assembly_code string, start_address uint64, length , ) bool {
	return request[bool]("HyperDbgAssembleGetLength", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgAssemble(assembly_code string, start_address uint64, buffer_to_store_assembled_data , buffer_size uint32, ) bool {
	return request[bool]("HyperDbgAssemble", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
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
func (debuggers) HyperDbgGetLocalApic(local_apic , is_using_x2apic , ) bool {
	return request[bool]("HyperDbgGetLocalApic", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgGetIoApic(io_apic , ) bool {
	return request[bool]("HyperDbgGetIoApic", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgGetIdtEntry(idt_packet , ) bool {
	return request[bool]("HyperDbgGetIdtEntry", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HwdbgScriptRunScript(script , instance_filepath_to_read , hardware_script_file_path_to_save , initial_bram_buffer_size uint32, ) bool {
	return request[bool]("HwdbgScriptRunScript", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) ScriptEngineWrapperTestParserForHwdbg(Expr , ) void {
	return request[void]("ScriptEngineWrapperTestParserForHwdbg", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgEnableTransparentMode(ProcessId uint32, ProcessName string, IsProcessId bool, ) bool {
	return request[bool]("HyperDbgEnableTransparentMode", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgDisableTransparentMode() bool {
	return request[bool]("HyperDbgDisableTransparentMode", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
