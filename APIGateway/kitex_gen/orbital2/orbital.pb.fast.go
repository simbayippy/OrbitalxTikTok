// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package orbital2

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *Person) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Person[number], err)
}

func (x *Person) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Person) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Age, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *Request) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Request[number], err)
}

func (x *Request) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Message, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Response) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Response[number], err)
}

func (x *Response) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Message, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *EncodedMessage) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_EncodedMessage[number], err)
}

func (x *EncodedMessage) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Method, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *EncodedMessage) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Data, offset, err = fastpb.ReadBytes(buf, _type)
	return offset, err
}

func (x *Person) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *Person) fastWriteField1(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetName())
	return offset
}

func (x *Person) fastWriteField2(buf []byte) (offset int) {
	if x.Age == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetAge())
	return offset
}

func (x *Request) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *Request) fastWriteField1(buf []byte) (offset int) {
	if x.Message == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetMessage())
	return offset
}

func (x *Response) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *Response) fastWriteField1(buf []byte) (offset int) {
	if x.Message == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetMessage())
	return offset
}

func (x *EncodedMessage) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *EncodedMessage) fastWriteField1(buf []byte) (offset int) {
	if x.Method == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetMethod())
	return offset
}

func (x *EncodedMessage) fastWriteField2(buf []byte) (offset int) {
	if len(x.Data) == 0 {
		return offset
	}
	offset += fastpb.WriteBytes(buf[offset:], 2, x.GetData())
	return offset
}

func (x *Person) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *Person) sizeField1() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetName())
	return n
}

func (x *Person) sizeField2() (n int) {
	if x.Age == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetAge())
	return n
}

func (x *Request) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *Request) sizeField1() (n int) {
	if x.Message == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetMessage())
	return n
}

func (x *Response) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *Response) sizeField1() (n int) {
	if x.Message == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetMessage())
	return n
}

func (x *EncodedMessage) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *EncodedMessage) sizeField1() (n int) {
	if x.Method == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetMethod())
	return n
}

func (x *EncodedMessage) sizeField2() (n int) {
	if len(x.Data) == 0 {
		return n
	}
	n += fastpb.SizeBytes(2, x.GetData())
	return n
}

var fieldIDToName_Person = map[int32]string{
	1: "Name",
	2: "Age",
}

var fieldIDToName_Request = map[int32]string{
	1: "Message",
}

var fieldIDToName_Response = map[int32]string{
	1: "Message",
}

var fieldIDToName_EncodedMessage = map[int32]string{
	1: "Method",
	2: "Data",
}
