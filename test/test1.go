package test

import "casino_common/proto/ddproto"

func slice()  {
	var array = [3]int32{1,2,3}
	//fmt.Print(reflect.TypeOf(array).String())
	arr := array[0:]
	_ = append(arr,arr...)
	var a *ddproto.GroupMemberInfo = nil
	_ := a.GetUid()
}