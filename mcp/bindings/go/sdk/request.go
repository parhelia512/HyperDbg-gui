package sdk

import (
	"cmp"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ddkwork/golibrary/std/mylog"
)

const DefaultHyperdbgServer = "http://127.0.0.1:8888/"

var client = &http.Client{
	Timeout: 15 * time.Second,
	Transport: &http.Transport{
		DisableKeepAlives: true,
	},
}

//	Type type Ordered interface {
//		~int | ~int8 | ~int16 | ~int32 | ~int64 |
//			~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
//			~float32 | ~float64 |
//			~string
//	}
type Type interface {
	cmp.Ordered |
		bool |
		[]byte |
		void
	//moduleInfo |
	//[]moduleInfo |
	//moduleSectionInfo |
	//[]moduleSectionInfo |
	//moduleExport |
	//[]moduleExport |
	//moduleImport |
	//[]moduleImport |
	//memoryBase |
	//disassemblerAddress |
	//disassembleRip |
	//disassembleRipWithSetupIn |
	//assemblerResult |
}
type void any

func request[T Type](endpoint string, params map[string]string) T {
	x64dbgServerURL := DefaultHyperdbgServer
	url := x64dbgServerURL + endpoint

	// 添加查询参数
	if len(params) > 0 {
		query := ""
		for key, value := range params {
			query += key + "=" + value + "&"
		}
		url += "?" + strings.TrimSuffix(query, "&")
	}

	resp := mylog.Check2(client.Get(url))
	defer func() {
		mylog.Check(resp.Body.Close())
	}()

	body := mylog.Check2(io.ReadAll(resp.Body))

	if resp.StatusCode != http.StatusOK {
		mylog.Check(fmt.Sprintf("Error %d: %s", resp.StatusCode, string(body)))
	}

	str := strings.TrimSpace(string(body))
	base := 10
	if strings.HasPrefix(str, "0x") {
		base = 16
	}
	str = strings.TrimPrefix(str, "0x")
	var zero T
	switch v := any(zero).(type) {
	case bool:
		if strings.EqualFold(str, "true") {
			return any(true).(T)
		}
		if strings.EqualFold(str, "false") {
			return any(false).(T)
		}
	case []byte:
		b := mylog.Check2(hex.DecodeString(str))
		return any(b).(T)
	case int:

		value := mylog.Check2(strconv.ParseInt(str, base, 64))
		return any(value).(T)
	case int8:

		value := mylog.Check2(strconv.ParseInt(str, base, 8))
		return any(value).(T)
	case int16:

		value := mylog.Check2(strconv.ParseInt(str, base, 16))
		return any(value).(T)
	case int32:

		value := mylog.Check2(strconv.ParseInt(str, base, 32))
		return any(value).(T)
	case int64:

		value := mylog.Check2(strconv.ParseInt(str, base, 64))
		return any(value).(T)
	case uint:

		value := mylog.Check2(strconv.ParseUint(str, base, 64))
		return any(value).(T)
	case uint8:

		value := mylog.Check2(strconv.ParseUint(str, base, 8))
		return any(value).(T)
	case uint16:

		value := mylog.Check2(strconv.ParseUint(str, base, 16))
		return any(value).(T)
	case uint32:

		value := mylog.Check2(strconv.ParseUint(str, base, 32))
		return any(value).(T)
	case uint64:

		value := mylog.Check2(strconv.ParseUint(str, base, 64))
		return any(value).(T)
	case uintptr:

		value := mylog.Check2(strconv.ParseUint(str, base, 64))
		return any(value).(T)
	case float32:
		value := mylog.Check2(strconv.ParseFloat(str, 32))
		return any(value).(T)
	case float64:
		value := mylog.Check2(strconv.ParseFloat(str, 64))
		return any(value).(T)
	case string:
		return any(str).(T)

		//todo 处理cpp服务端的字段返回 0x12345678 这种格式，我估计json会解码失败
	//case moduleInfo:
	//	mylog.Check(json.Unmarshal(body, &v))
	//	return any(v).(T)
	//case []moduleInfo:
	//	mylog.Check(json.Unmarshal(body, &v))
	//	return any(v).(T)
	//case moduleSectionInfo:
	//	mylog.Check(json.Unmarshal(body, &v))
	//	return any(v).(T)
	//case []moduleSectionInfo:
	//	mylog.Check(json.Unmarshal(body, &v))
	//	return any(v).(T)
	//case moduleExport:
	//	mylog.Check(json.Unmarshal(body, &v))
	//	return any(v).(T)
	//case []moduleExport:
	//	mylog.Check(json.Unmarshal(body, &v))
	//	return any(v).(T)
	//case moduleImport:
	//	mylog.Check(json.Unmarshal(body, &v))
	//	return any(v).(T)
	//case []moduleImport:
	//	mylog.Check(json.Unmarshal(body, &v))
	//	return any(v).(T)
	//case memoryBase:
	//	mylog.Check(json.Unmarshal(body, &v))
	//	return any(v).(T)
	//case disassemblerAddress:
	//	mylog.Check(json.Unmarshal(body, &v))
	//	return any(v).(T)
	//case disassembleRip:
	//	mylog.Check(json.Unmarshal(body, &v))
	//	return any(v).(T)
	//case disassembleRipWithSetupIn:
	//	mylog.Check(json.Unmarshal(body, &v))
	//	return any(v).(T)
	//case assemblerResult:
	//	mylog.Check(json.Unmarshal(body, &v))
	//	return any(v).(T)
	case void:
		v = v
		return any(nil).(T)

	}
	panic("not support type")
}

type (
	PDEBUGGER_DT_COMMAND_OPTIONS      struct{}
	GUEST_REGS                        struct{}
	GUEST_EXTRA_REGISTERS             struct{}
	DEBUGGER_READ_MEMORY_ADDRESS_MODE byte
	DEBUGGER_READ_MEMORY_TYPE         byte
	DEBUGGER_READ_READING_TYPE        byte
)
