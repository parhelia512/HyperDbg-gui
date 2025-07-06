package mcp

import (
	"strings"
	"testing"

	"github.com/ddkwork/golibrary/std/stream"
)

func TestName(t *testing.T) {
	for s := range stream.ReadFileToLines("export/export.cpp") {
		if strings.HasPrefix(s, "    return ") {
			s = strings.TrimSpace(s)
			s = strings.TrimPrefix(s, "return")
			println(s)
		}
	}

	//b := string(mylog.Check2(os.ReadFile("export/export.cpp")))
	//split := strings.Split(b, "hyperdbg_u_")
	//mylog.Struct(split)
}

/* todo
 *  make api meta json file and all commands json file
 *  decode all error code text and send to client
VmxSupportDetection();
HyperDbgLoadVmmModule();
HyperDbgUnloadVmm();
HyperDbgInstallVmmDriver();
HyperDbgUninstallVmmDriver();
HyperDbgStopVmmDriver();
HyperDbgInterpreter(command);
HyperDbgTestCommandParser(command, number_of_tokens, tokens_list, failed_token_num, failed_token_position);
HyperDbgTestCommandParserShowTokens(command);
SetTextMessageCallbackUsingSharedBuffer(handler);
ScriptReadFileAndExecuteCommandline(argc, argv);
ContinuePreviousCommand();
CheckMultilineCommand(current_command, reset);
ConnectRemoteDebugger(ip, port);
HyperDbgReadMemory(target_address, memory_type, reading_Type, pid, size, get_address_mode, address_mode, target_buffer_to_store, return_length);
HyperDbgReadAllRegisters(guest_registers, extra_registers);
HyperDbgReadTargetRegister(register_id, target_register);
HyperDbgWriteTargetRegister(register_id, value);
HyperDbgRegisterShowAll();
HyperDbgRegisterShowTargetRegister(register_id);
HyperDbgWriteMemory(destination_address, memory_type, process_id, source_address, number_of_bytes);
DebuggerGetKernelBase();
HyperDbgDebugRemoteDeviceUsingComPort(port_name, baudrate, pause_after_connection);
HyperDbgDebugRemoteDeviceUsingNamedPipe(named_pipe, pause_after_connection);
HyperDbgDebugCloseRemoteDebugger();
HyperDbgDebugCurrentDeviceUsingComPort(port_name, baudrate);
UdAttachToProcess(NULL,
UdAttachToProcess(NULL,
HyperDbgAssembleGetLength(assembly_code, start_address, length);
HyperDbgAssemble(assembly_code, start_address, buffer_to_store_assembled_data, buffer_size);
SetupPathForFileName(filename, file_location, buffer_len, check_file_existence);
SteppingInstrumentationStepIn();
SteppingRegularStepIn();
SteppingStepOver();
SteppingInstrumentationStepInForTracking();
SteppingStepOverForGu(last_instruction);
HyperDbgGetLocalApic(local_apic, is_using_x2apic);
HyperDbgGetIoApic(io_apic);
HyperDbgGetIdtEntry(idt_packet);
HwdbgScriptRunScript(script,
HyperDbgEnableTransparentMode(ProcessId, ProcessName, IsProcessId);
HyperDbgDisableTransparentMode();
*/
