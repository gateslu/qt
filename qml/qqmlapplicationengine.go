package qml

//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QQmlApplicationEngine struct {
	QQmlEngine
}

type QQmlApplicationEngine_ITF interface {
	QQmlEngine_ITF
	QQmlApplicationEngine_PTR() *QQmlApplicationEngine
}

func PointerFromQQmlApplicationEngine(ptr QQmlApplicationEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlApplicationEngine_PTR().Pointer()
	}
	return nil
}

func NewQQmlApplicationEngineFromPointer(ptr unsafe.Pointer) *QQmlApplicationEngine {
	var n = new(QQmlApplicationEngine)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QQmlApplicationEngine_") {
		n.SetObjectName("QQmlApplicationEngine_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlApplicationEngine) QQmlApplicationEngine_PTR() *QQmlApplicationEngine {
	return ptr
}

func NewQQmlApplicationEngine(parent core.QObject_ITF) *QQmlApplicationEngine {
	defer qt.Recovering("QQmlApplicationEngine::QQmlApplicationEngine")

	return NewQQmlApplicationEngineFromPointer(C.QQmlApplicationEngine_NewQQmlApplicationEngine(core.PointerFromQObject(parent)))
}

func NewQQmlApplicationEngine3(filePath string, parent core.QObject_ITF) *QQmlApplicationEngine {
	defer qt.Recovering("QQmlApplicationEngine::QQmlApplicationEngine")

	return NewQQmlApplicationEngineFromPointer(C.QQmlApplicationEngine_NewQQmlApplicationEngine3(C.CString(filePath), core.PointerFromQObject(parent)))
}

func NewQQmlApplicationEngine2(url core.QUrl_ITF, parent core.QObject_ITF) *QQmlApplicationEngine {
	defer qt.Recovering("QQmlApplicationEngine::QQmlApplicationEngine")

	return NewQQmlApplicationEngineFromPointer(C.QQmlApplicationEngine_NewQQmlApplicationEngine2(core.PointerFromQUrl(url), core.PointerFromQObject(parent)))
}

func (ptr *QQmlApplicationEngine) Load2(filePath string) {
	defer qt.Recovering("QQmlApplicationEngine::load")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_Load2(ptr.Pointer(), C.CString(filePath))
	}
}

func (ptr *QQmlApplicationEngine) Load(url core.QUrl_ITF) {
	defer qt.Recovering("QQmlApplicationEngine::load")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_Load(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlApplicationEngine) LoadData(data core.QByteArray_ITF, url core.QUrl_ITF) {
	defer qt.Recovering("QQmlApplicationEngine::loadData")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_LoadData(ptr.Pointer(), core.PointerFromQByteArray(data), core.PointerFromQUrl(url))
	}
}

func (ptr *QQmlApplicationEngine) ConnectObjectCreated(f func(object *core.QObject, url *core.QUrl)) {
	defer qt.Recovering("connect QQmlApplicationEngine::objectCreated")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_ConnectObjectCreated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "objectCreated", f)
	}
}

func (ptr *QQmlApplicationEngine) DisconnectObjectCreated() {
	defer qt.Recovering("disconnect QQmlApplicationEngine::objectCreated")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_DisconnectObjectCreated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "objectCreated")
	}
}

//export callbackQQmlApplicationEngineObjectCreated
func callbackQQmlApplicationEngineObjectCreated(ptrName *C.char, object unsafe.Pointer, url unsafe.Pointer) {
	defer qt.Recovering("callback QQmlApplicationEngine::objectCreated")

	var signal = qt.GetSignal(C.GoString(ptrName), "objectCreated")
	if signal != nil {
		signal.(func(*core.QObject, *core.QUrl))(core.NewQObjectFromPointer(object), core.NewQUrlFromPointer(url))
	}

}

func (ptr *QQmlApplicationEngine) DestroyQQmlApplicationEngine() {
	defer qt.Recovering("QQmlApplicationEngine::~QQmlApplicationEngine")

	if ptr.Pointer() != nil {
		C.QQmlApplicationEngine_DestroyQQmlApplicationEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}