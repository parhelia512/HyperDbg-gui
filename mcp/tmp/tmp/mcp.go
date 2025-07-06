package sdk

import "fmt"

type debuggers struct {}
func (debuggers) VmxSupportDetection() BOOLEAN {
	return request[BOOLEAN]("VmxSupportDetection", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) CpuReadVendorString(vendor_string CHAR *, ) VOID {
	return request[VOID]("CpuReadVendorString", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgLoadVmmModule() INT {
	return request[INT]("HyperDbgLoadVmmModule", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgUnloadVmm() INT {
	return request[INT]("HyperDbgUnloadVmm", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgInstallVmmDriver() INT {
	return request[INT]("HyperDbgInstallVmmDriver", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgUninstallVmmDriver() INT {
	return request[INT]("HyperDbgUninstallVmmDriver", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgStopVmmDriver() INT {
	return request[INT]("HyperDbgStopVmmDriver", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgInterpreter(command CHAR *, ) INT {
	return request[INT]("HyperDbgInterpreter", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgTestCommandParser(command CHAR *, number_of_tokens UINT32, tokens_list CHAR **, failed_token_num UINT32 *, failed_token_position UINT32 *, ) BOOLEAN {
	return request[BOOLEAN]("HyperDbgTestCommandParser", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgTestCommandParserShowTokens(command CHAR *, ) VOID {
	return request[VOID]("HyperDbgTestCommandParserShowTokens", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgShowSignature() VOID {
	return request[VOID]("HyperDbgShowSignature", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) SetTextMessageCallback() VOID {
	return request[VOID]("SetTextMessageCallback", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) SetTextMessageCallbackUsingSharedBuffer() VOID {
	return request[VOID]("SetTextMessageCallbackUsingSharedBuffer", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) UnsetTextMessageCallback() VOID {
	return request[VOID]("UnsetTextMessageCallback", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) ScriptReadFileAndExecuteCommandline(argc INT, argv CHAR *, ) INT {
	return request[INT]("ScriptReadFileAndExecuteCommandline", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) ContinuePreviousCommand() BOOLEAN {
	return request[BOOLEAN]("ContinuePreviousCommand", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) CheckMultilineCommand(current_command CHAR *, reset BOOLEAN, ) BOOLEAN {
	return request[BOOLEAN]("CheckMultilineCommand", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) ConnectLocalDebugger() VOID {
	return request[VOID]("ConnectLocalDebugger", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) ConnectRemoteDebugger(ip const CHAR *, port const CHAR *, ) BOOLEAN {
	return request[BOOLEAN]("ConnectRemoteDebugger", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) Continue() VOID {
	return request[VOID]("Continue", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) Pause() VOID {
	return request[VOID]("Pause", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) SetBreakPoint(address UINT64, pid UINT32, tid UINT32, core_numer UINT32, ) VOID {
	return request[VOID]("SetBreakPoint", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) SetCustomDriverPath(driver_file_path CHAR *, driver_name CHAR *, ) BOOLEAN {
	return request[BOOLEAN]("SetCustomDriverPath", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) UseDefaultDriverPath() VOID {
	return request[VOID]("UseDefaultDriverPath", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgReadMemory(target_address UINT64, memory_type DEBUGGER_READ_MEMORY_TYPE, reading_Type xxxxxxxxxx, pid xxxxxxxxxx, size xxxxxxxxxx, get_address_mode xxxxxxxxxx, address_mode xxxxxxxxxx, target_buffer_to_store xxxxxxxxxx, return_length xxxxxxxxxx, ) BOOLEAN {
	return request[BOOLEAN]("HyperDbgReadMemory", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgShowMemoryOrDisassemble(style DEBUGGER_SHOW_MEMORY_STYLE, address UINT64, memory_type DEBUGGER_READ_MEMORY_TYPE, reading_type DEBUGGER_READ_READING_TYPE, pid UINT32, size UINT32, dt_details PDEBUGGER_DT_COMMAND_OPTIONS, ) VOID {
	return request[VOID]("HyperDbgShowMemoryOrDisassemble", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgReadAllRegisters(guest_registers xxxxxxxxxxxxxxx, extra_registers xxxxxxxxxx, ) BOOLEAN {
	return request[BOOLEAN]("HyperDbgReadAllRegisters", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) xxxxxxxxxxxxxxxxxxx(xxxxxxxxxxxxxx xxxxxxxxxxxxxxx, xxxxxxxxxx xxxxxxxxxx, ) xxxxxxxxxxxxxxxxx {
	return request[xxxxxxxxxxxxxxxxx]("xxxxxxxxxxxxxxxxxxx", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgReadTargetRegister(xxxxxxxxxxxxxx xxxxxxxxxxxxxxx, xxxxxxxxxx xxxxxxxxxx, ) BOOLEAN {
	return request[BOOLEAN]("HyperDbgReadTargetRegister", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgWriteTargetRegister(register_id REGS_ENUM, value UINT64, ) BOOLEAN {
	return request[BOOLEAN]("HyperDbgWriteTargetRegister", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgRegisterShowAll() BOOLEAN {
	return request[BOOLEAN]("HyperDbgRegisterShowAll", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgRegisterShowTargetRegister(register_id REGS_ENUM, ) BOOLEAN {
	return request[BOOLEAN]("HyperDbgRegisterShowTargetRegister", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgWriteMemory(destination_address PVOID, memory_type DEBUGGER_EDIT_MEMORY_TYPE, process_id UINT32, source_address PVOID, number_of_bytes UINT32, ) BOOLEAN {
	return request[BOOLEAN]("HyperDbgWriteMemory", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) DebuggerGetKernelBase() UINT64 {
	return request[UINT64]("DebuggerGetKernelBase", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgDebugRemoteDeviceUsingComPort(port_name const CHAR *, baudrate DWORD, pause_after_connection BOOLEAN, ) BOOLEAN {
	return request[BOOLEAN]("HyperDbgDebugRemoteDeviceUsingComPort", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgDebugRemoteDeviceUsingNamedPipe(named_pipe const CHAR *, pause_after_connection BOOLEAN, ) BOOLEAN {
	return request[BOOLEAN]("HyperDbgDebugRemoteDeviceUsingNamedPipe", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgDebugCloseRemoteDebugger() BOOLEAN {
	return request[BOOLEAN]("HyperDbgDebugCloseRemoteDebugger", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgDebugCurrentDeviceUsingComPort(port_name const CHAR *, baudrate DWORD, ) BOOLEAN {
	return request[BOOLEAN]("HyperDbgDebugCurrentDeviceUsingComPort", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) UdAttachToProcess(path const WCHAR *, ) BOOLEAN {
	return request[BOOLEAN]("UdAttachToProcess", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) UdAttachToProcess(path const WCHAR *, arguments const WCHAR *, ) BOOLEAN {
	return request[BOOLEAN]("UdAttachToProcess", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgAssembleGetLength(assembly_code const CHAR *, start_address UINT64, length UINT32 *, ) BOOLEAN {
	return request[BOOLEAN]("HyperDbgAssembleGetLength", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgAssemble(assembly_code const CHAR *, start_address UINT64, buffer_to_store_assembled_data PVOID, buffer_size UINT32, ) BOOLEAN {
	return request[BOOLEAN]("HyperDbgAssemble", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) SetupPathForFileName(filename const CHAR *, file_location CHAR *, buffer_len UINT32, check_file_existence BOOLEAN, ) BOOLEAN {
	return request[BOOLEAN]("SetupPathForFileName", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) SteppingInstrumentationStepIn() BOOLEAN {
	return request[BOOLEAN]("SteppingInstrumentationStepIn", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) SteppingRegularStepIn() BOOLEAN {
	return request[BOOLEAN]("SteppingRegularStepIn", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) SteppingStepOver() BOOLEAN {
	return request[BOOLEAN]("SteppingStepOver", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) SteppingInstrumentationStepInForTracking() BOOLEAN {
	return request[BOOLEAN]("SteppingInstrumentationStepInForTracking", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) SteppingStepOverForGu(last_instruction BOOLEAN, ) BOOLEAN {
	return request[BOOLEAN]("SteppingStepOverForGu", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgGetLocalApic(local_apic PLAPIC_PAGE, is_using_x2apic BOOLEAN *, ) BOOLEAN {
	return request[BOOLEAN]("HyperDbgGetLocalApic", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgGetIoApic(io_apic IO_APIC_ENTRY_PACKETS *, ) BOOLEAN {
	return request[BOOLEAN]("HyperDbgGetIoApic", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgGetIdtEntry(idt_packet INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS *, ) BOOLEAN {
	return request[BOOLEAN]("HyperDbgGetIdtEntry", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HwdbgScriptRunScript(script const CHAR * , instance_filepath_to_read const CHAR * , hardware_script_file_path_to_save const CHAR * , initial_bram_buffer_size UINT32, ) BOOLEAN {
	return request[BOOLEAN]("HwdbgScriptRunScript", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) ScriptEngineWrapperTestParserForHwdbg(Expr const char *, ) VOID {
	return request[VOID]("ScriptEngineWrapperTestParserForHwdbg", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgEnableTransparentMode(ProcessId UINT32, ProcessName CHAR *, IsProcessId BOOLEAN, ) BOOLEAN {
	return request[BOOLEAN]("HyperDbgEnableTransparentMode", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debuggers) HyperDbgDisableTransparentMode() BOOLEAN {
	return request[BOOLEAN]("HyperDbgDisableTransparentMode", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
