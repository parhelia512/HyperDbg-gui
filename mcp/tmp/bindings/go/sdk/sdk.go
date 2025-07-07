package sdk
import (
"encoding/hex"
"encoding/json"
"fmt"
"strconv"
)
type debugger struct {}
func (debugger) VmxSupportDetection() bool {
	return request[bool]("VmxSupportDetection",nil)
}
func (debugger) CpuReadVendorString(vendor_string string)  {
	 request[void]("CpuReadVendorString",map[string]string{"vendor_string":vendor_string
})
}
func (debugger) LoadVmmModule() int {
	return request[int]("LoadVmmModule",nil)
}
func (debugger) UnloadVmm() int {
	return request[int]("UnloadVmm",nil)
}
func (debugger) InstallVmmDriver() int {
	return request[int]("InstallVmmDriver",nil)
}
func (debugger) UninstallVmmDriver() int {
	return request[int]("UninstallVmmDriver",nil)
}
func (debugger) StopVmmDriver() int {
	return request[int]("StopVmmDriver",nil)
}
func (debugger) Interpreter(command string) int {
	return request[int]("Interpreter",map[string]string{"command":command
})
}
func (debugger) TestCommandParser(command string, number_of_tokens uint32, tokens_list , failed_token_num , failed_token_position ) bool {
	return request[bool]("TestCommandParser",map[string]string{"command":command,
"number_of_tokens":number_of_tokens,
"tokens_list":,
"failed_token_num":,
"failed_token_position":
})
}
func (debugger) TestCommandParserShowTokens(command string)  {
	 request[void]("TestCommandParserShowTokens",map[string]string{"command":command
})
}
func (debugger) ShowSignature()  {
	 request[void]("ShowSignature",nil)
}
func (debugger) SetTextMessageCallback()  {
	 request[void]("SetTextMessageCallback",nil)
}
func (debugger) SetTextMessageCallbackUsingSharedBuffer()  {
	 request[void]("SetTextMessageCallbackUsingSharedBuffer",nil)
}
func (debugger) UnsetTextMessageCallback()  {
	 request[void]("UnsetTextMessageCallback",nil)
}
func (debugger) ScriptReadFileAndExecuteCommandline(argc int, argv string) int {
	return request[int]("ScriptReadFileAndExecuteCommandline",map[string]string{"argc":argc,
"argv":argv
})
}
func (debugger) ContinuePreviousCommand() bool {
	return request[bool]("ContinuePreviousCommand",nil)
}
func (debugger) CheckMultilineCommand(current_command string, reset bool) bool {
	return request[bool]("CheckMultilineCommand",map[string]string{"current_command":current_command,
"reset":str(int(reset))
})
}
func (debugger) ConnectLocalDebugger()  {
	 request[void]("ConnectLocalDebugger",nil)
}
func (debugger) ConnectRemoteDebugger(ip string, port string) bool {
	return request[bool]("ConnectRemoteDebugger",map[string]string{"ip":ip,
"port":port
})
}
func (debugger) Continue()  {
	 request[void]("Continue",nil)
}
func (debugger) Pause()  {
	 request[void]("Pause",nil)
}
func (debugger) SetBreakPoint(address uint64, pid uint32, tid uint32, core_numer uint32)  {
	 request[void]("SetBreakPoint",map[string]string{"address":address,
"pid":pid,
"tid":tid,
"core_numer":core_numer
})
}
func (debugger) SetCustomDriverPath(driver_file_path string, driver_name string) bool {
	return request[bool]("SetCustomDriverPath",map[string]string{"driver_file_path":driver_file_path,
"driver_name":driver_name
})
}
func (debugger) UseDefaultDriverPath()  {
	 request[void]("UseDefaultDriverPath",nil)
}
func (debugger) ReadMemory(target_address uint64, memory_type , reading_Type , pid uint32, size uint32, get_address_mode bool, address_mode , target_buffer_to_store , return_length ) bool {
	return request[bool]("ReadMemory",map[string]string{"target_address":target_address,
"memory_type":,
"reading_Type":,
"pid":pid,
"size":size,
"get_address_mode":str(int(get_address_mode)),
"address_mode":,
"target_buffer_to_store":,
"return_length":
})
}
func (debugger) ShowMemoryOrDisassemble(style , address uint64, memory_type , reading_type , pid uint32, size uint32, dt_details )  {
	 request[void]("ShowMemoryOrDisassemble",map[string]string{"style":,
"address":address,
"memory_type":,
"reading_type":,
"pid":pid,
"size":size,
"dt_details":
})
}
func (debugger) ReadAllRegisters(guest_registers , extra_registers ) bool {
	return request[bool]("ReadAllRegisters",map[string]string{"guest_registers":,
"extra_registers":
})
}
func (debugger) xxxxxxxxxxxxxxxxxxx(xxxxxxxxxxxxxx , xxxxxxxxxx )  {
	return request[]("xxxxxxxxxxxxxxxxxxx",map[string]string{"xxxxxxxxxxxxxx":,
"xxxxxxxxxx":
})
}
func (debugger) ReadTargetRegister(xxxxxxxxxxxxxx , xxxxxxxxxx ) bool {
	return request[bool]("ReadTargetRegister",map[string]string{"xxxxxxxxxxxxxx":,
"xxxxxxxxxx":
})
}
func (debugger) WriteTargetRegister(register_id , value uint64) bool {
	return request[bool]("WriteTargetRegister",map[string]string{"register_id":,
"value":value
})
}
func (debugger) RegisterShowAll() bool {
	return request[bool]("RegisterShowAll",nil)
}
func (debugger) RegisterShowTargetRegister(register_id ) bool {
	return request[bool]("RegisterShowTargetRegister",map[string]string{"register_id":
})
}
func (debugger) WriteMemory(destination_address , memory_type , process_id uint32, source_address , number_of_bytes uint32) bool {
	return request[bool]("WriteMemory",map[string]string{"destination_address":,
"memory_type":,
"process_id":process_id,
"source_address":,
"number_of_bytes":number_of_bytes
})
}
func (debugger) DebuggerGetKernelBase() uint64 {
	return request[uint64]("DebuggerGetKernelBase",nil)
}
func (debugger) DebugRemoteDeviceUsingComPort(port_name string, baudrate , pause_after_connection bool) bool {
	return request[bool]("DebugRemoteDeviceUsingComPort",map[string]string{"port_name":port_name,
"baudrate":,
"pause_after_connection":str(int(pause_after_connection))
})
}
func (debugger) DebugRemoteDeviceUsingNamedPipe(named_pipe string, pause_after_connection bool) bool {
	return request[bool]("DebugRemoteDeviceUsingNamedPipe",map[string]string{"named_pipe":named_pipe,
"pause_after_connection":str(int(pause_after_connection))
})
}
func (debugger) DebugCloseRemoteDebugger() bool {
	return request[bool]("DebugCloseRemoteDebugger",nil)
}
func (debugger) DebugCurrentDeviceUsingComPort(port_name string, baudrate ) bool {
	return request[bool]("DebugCurrentDeviceUsingComPort",map[string]string{"port_name":port_name,
"baudrate":
})
}
func (debugger) UdAttachToProcess(path string) bool {
	return request[bool]("UdAttachToProcess",map[string]string{"path":path
})
}
func (debugger) UdAttachToProcess(path string, arguments string) bool {
	return request[bool]("UdAttachToProcess",map[string]string{"path":path,
"arguments":arguments
})
}
func (debugger) AssembleGetLength(assembly_code string, start_address uint64, length ) bool {
	return request[bool]("AssembleGetLength",map[string]string{"assembly_code":assembly_code,
"start_address":start_address,
"length":
})
}
func (debugger) Assemble(assembly_code string, start_address uint64, buffer_to_store_assembled_data , buffer_size uint32) bool {
	return request[bool]("Assemble",map[string]string{"assembly_code":assembly_code,
"start_address":start_address,
"buffer_to_store_assembled_data":,
"buffer_size":buffer_size
})
}
func (debugger) SetupPathForFileName(filename string, file_location string, buffer_len uint32, check_file_existence bool) bool {
	return request[bool]("SetupPathForFileName",map[string]string{"filename":filename,
"file_location":file_location,
"buffer_len":buffer_len,
"check_file_existence":str(int(check_file_existence))
})
}
func (debugger) SteppingInstrumentationStepIn() bool {
	return request[bool]("SteppingInstrumentationStepIn",nil)
}
func (debugger) SteppingRegularStepIn() bool {
	return request[bool]("SteppingRegularStepIn",nil)
}
func (debugger) SteppingStepOver() bool {
	return request[bool]("SteppingStepOver",nil)
}
func (debugger) SteppingInstrumentationStepInForTracking() bool {
	return request[bool]("SteppingInstrumentationStepInForTracking",nil)
}
func (debugger) SteppingStepOverForGu(last_instruction bool) bool {
	return request[bool]("SteppingStepOverForGu",map[string]string{"last_instruction":str(int(last_instruction))
})
}
func (debugger) GetLocalApic(local_apic , is_using_x2apic ) bool {
	return request[bool]("GetLocalApic",map[string]string{"local_apic":,
"is_using_x2apic":
})
}
func (debugger) GetIoApic(io_apic ) bool {
	return request[bool]("GetIoApic",map[string]string{"io_apic":
})
}
func (debugger) GetIdtEntry(idt_packet ) bool {
	return request[bool]("GetIdtEntry",map[string]string{"idt_packet":
})
}
func (debugger) HwdbgScriptRunScript(script , instance_filepath_to_read , hardware_script_file_path_to_save , initial_bram_buffer_size uint32) bool {
	return request[bool]("HwdbgScriptRunScript",map[string]string{"script":,
"instance_filepath_to_read":,
"hardware_script_file_path_to_save":,
"initial_bram_buffer_size":initial_bram_buffer_size
})
}
func (debugger) ScriptEngineWrapperTestParserForHwdbg(Expr )  {
	 request[void]("ScriptEngineWrapperTestParserForHwdbg",map[string]string{"Expr":
})
}
func (debugger) EnableTransparentMode(ProcessId uint32, ProcessName string, IsProcessId bool) bool {
	return request[bool]("EnableTransparentMode",map[string]string{"ProcessId":ProcessId,
"ProcessName":ProcessName,
"IsProcessId":str(int(IsProcessId))
})
}
func (debugger) DisableTransparentMode() bool {
	return request[bool]("DisableTransparentMode",nil)
}
