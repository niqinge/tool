package reflects

import (
    "errors"
    "fmt"
    "reflect"
)

var (
    ErrReflectResult2Count = errors.New("方法返回参数不是2个")
    ErrReflectResult1Count = errors.New("方法返回参数不是1个")
    ErrReflectResult3Count = errors.New("方法返回参数不是3个")
    ErrReflectResult4Count = errors.New("方法返回参数不是4个")
)

func Call(typ interface{}, methodName string, params ...interface{}) []reflect.Value {
    val := reflect.ValueOf(typ)

    callReq := make([]reflect.Value, len(params))
    for i := range params {
        callReq[i] = reflect.ValueOf(params[i])
    }

    // 调用SetName方法
    return val.MethodByName(methodName).Call(callReq)
}

// 反射调用结构体对象方法, 并返回2个结果值
func Call2Result(typ interface{}, methodName string, params ...interface{}) (interface{}, interface{}) {
    val := reflect.ValueOf(typ)

    callReq := make([]reflect.Value, len(params))
    for i := range params {
        callReq[i] = reflect.ValueOf(params[i])
    }

    // 调用方法
    result := val.MethodByName(methodName).Call(callReq)

    if len(result) != 2 {
        return nil, fmt.Errorf("%s.%s%s", val.Type().Name(), methodName, ErrReflectResult2Count)
    }

    return result[0].Interface(), result[1].Interface()
}

// 反射调用结构体对象方法, 并返回1个结果值
func Call1Result(typ interface{}, methodName string, params ...interface{}) interface{} {
    val := reflect.ValueOf(typ)

    callReq := make([]reflect.Value, len(params))
    for i := range params {
        callReq[i] = reflect.ValueOf(params[i])
    }

    // 调用方法
    result := val.MethodByName(methodName).Call(callReq)

    if len(result) != 1 {
        return fmt.Errorf("%s.%s%s", val.Type().Name(), methodName, ErrReflectResult1Count)
    }

    return result[0].Interface()
}

// 反射调用结构体对象方法, 并返回3个结果值
func Call3Result(typ interface{}, methodName string, params ...interface{}) (interface{}, interface{}, interface{}) {
    val := reflect.ValueOf(typ)

    callReq := make([]reflect.Value, len(params))
    for i := range params {
        callReq[i] = reflect.ValueOf(params[i])
    }

    // 调用方法
    result := val.MethodByName(methodName).Call(callReq)

    if len(result) != 3 {
        return nil, nil, fmt.Errorf("%s.%s%s", val.Type().Name(), methodName, ErrReflectResult3Count)
    }

    return result[0].Interface(), result[1].Interface(), result[2].Interface()
}

// 反射调用结构体对象方法, 并返回4个结果值
func Call4Result(typ interface{}, methodName string, params ...interface{}) (interface{}, interface{}, interface{}, interface{}) {
    val := reflect.ValueOf(typ)

    callReq := make([]reflect.Value, len(params))
    for i := range params {
        callReq[i] = reflect.ValueOf(params[i])
    }

    // 调用方法
    result := val.MethodByName(methodName).Call(callReq)

    if len(result) != 4 {
        return nil, nil, nil, fmt.Errorf("%s.%s%s", val.Type().Name(), methodName, ErrReflectResult4Count)
    }

    return result[0].Interface(), result[1].Interface(), result[2].Interface(), result[3].Interface()
}
