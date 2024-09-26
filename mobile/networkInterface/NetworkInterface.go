package networkInterface

//type Interface struct {
//    Index        int
//    MTU          int
//    Name         string
//    HardwareAddr HardwareAddr
//    Flags        Flags
//}
//
//type HardwareAddr []byte
//type Flags uint

//func Interfaces() ([]net.Interface, error) {
//	var interfaces []net.Interface
//	var err error
//
//	if runtime.GOOS == "android" {
//		AndroidInterfaces, err := MeBubbleNetwork.Interfaces()
//		if err != nil {
//			return nil, err
//		}
//
//		interfaces = AndroidInterfaces.([]net.Interface)
//
//	} else {
//		interfaces, err = net.Interfaces()
//		if err != nil {
//			return nil, err
//		}
//	}
//	return interfaces, nil
//}

// ManualConvert pretty stupid but works.
//func ManualConvert(I []MeBubbleNetwork.Interface) (n []net.Interface) {
//    for i, _ := range I {
//        var sn net.Interface
//        sn.Name = I[i].Name
//        sn.Index = I[i].Index
//        //sn.Flags = I[i].Flags
//            }
//}
