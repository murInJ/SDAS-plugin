// Code generated by thriftgo (0.3.8). DO NOT EDIT.

package api

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type SourceMsg struct {
	DataType string `thrift:"data_type,1" frugal:"1,default,string" json:"data_type"`
	Data     []byte `thrift:"data,2" frugal:"2,default,binary" json:"data"`
	Ntp      int64  `thrift:"ntp,3" frugal:"3,default,i64" json:"ntp"`
}

func NewSourceMsg() *SourceMsg {
	return &SourceMsg{}
}

func (p *SourceMsg) InitDefault() {
	*p = SourceMsg{}
}

func (p *SourceMsg) GetDataType() (v string) {
	return p.DataType
}

func (p *SourceMsg) GetData() (v []byte) {
	return p.Data
}

func (p *SourceMsg) GetNtp() (v int64) {
	return p.Ntp
}
func (p *SourceMsg) SetDataType(val string) {
	p.DataType = val
}
func (p *SourceMsg) SetData(val []byte) {
	p.Data = val
}
func (p *SourceMsg) SetNtp(val int64) {
	p.Ntp = val
}

var fieldIDToName_SourceMsg = map[int16]string{
	1: "data_type",
	2: "data",
	3: "ntp",
}

func (p *SourceMsg) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 3:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_SourceMsg[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *SourceMsg) ReadField1(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.DataType = v
	}
	return nil
}
func (p *SourceMsg) ReadField2(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadBinary(); err != nil {
		return err
	} else {
		p.Data = []byte(v)
	}
	return nil
}
func (p *SourceMsg) ReadField3(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.Ntp = v
	}
	return nil
}

func (p *SourceMsg) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("SourceMsg"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *SourceMsg) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("data_type", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.DataType); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *SourceMsg) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("data", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteBinary([]byte(p.Data)); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *SourceMsg) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("ntp", thrift.I64, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.Ntp); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *SourceMsg) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SourceMsg(%+v)", *p)

}

func (p *SourceMsg) DeepEqual(ano *SourceMsg) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.DataType) {
		return false
	}
	if !p.Field2DeepEqual(ano.Data) {
		return false
	}
	if !p.Field3DeepEqual(ano.Ntp) {
		return false
	}
	return true
}

func (p *SourceMsg) Field1DeepEqual(src string) bool {

	if strings.Compare(p.DataType, src) != 0 {
		return false
	}
	return true
}
func (p *SourceMsg) Field2DeepEqual(src []byte) bool {

	if bytes.Compare(p.Data, src) != 0 {
		return false
	}
	return true
}
func (p *SourceMsg) Field3DeepEqual(src int64) bool {

	if p.Ntp != src {
		return false
	}
	return true
}