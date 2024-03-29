package storage

const (
	ID            = 0
	NodeStorageID = 1
	GroupID       = 2
	Name          = 3
	NodeSAndVMID  = 4
	Lock          = 5
	UpdateName    = 100
	UpdateNodeS   = 101
	UpdateGroup   = 102
	UpdateAll     = 110
)

//type Storage struct {
//	gorm.Model
//	VMID          uint   `json:"vm_id"`
//	NodeStorageID uint   `json:"node_storage_id"`
//	GroupID       uint   `json:"group_id"`
//	Name          string `json:"name"`
//	Type          uint   `json:"type"`
//	FileType      uint   `json:"file_type"`
//	MaxCapacity   uint   `json:"max_capacity"`
//	UUID          string `json:"path"`
//	ReadOnly      *bool  `json:"readonly"`
//	Comment       string `json:"comment"`
//	Lock          *bool  `json:"lock"`
//}

type Result struct {
	Status  bool      `json:"status"`
	Error   string    `json:"error"`
	Storage []Storage `json:"storage"`
}

type ResultOne struct {
	Status  bool    `json:"status"`
	Error   string  `json:"error"`
	Storage Storage `json:"storage"`
}

type ResultDatabase struct {
	Err     error
	Storage []Storage
}

var Path = make(chan string)

//type Storage struct {
//	Info       gateway.Info `json:"info"`           //Info
//	Mode       uint         `json:"modestorageTmp"` //0:Manual 1:From ImaCon
//	FromImaCon ImaCon       `json:"from_imacon"`    //Imageをpullする際に使用するURL
//	VMName     string       `json:"vm_name"`        //VMNameがある場合は、Pathに追加する
//	Type       uint         `json:"type"`           //1: CDROM 2:Floppy (no support) 10:BootDev(VirtIO) 11: BootDev(SATA) 12: BootDev(IDE)
//	FileType   uint         `json:"filetype"`       //0:qcow2 1:img
//	PathType   uint         `json:"path_type"`      //node側のストレージの種類 0~9:SSD 10~19:HDD 20~29:NVMe 100~109:SSD(iSCSI) 110~119:SSD(iSCSI) 120~129:NVme(iSCSI)
//	Path       string       `json:"path"`           //node側のパス
//	Capacity   uint         `json:"capacity"`       //容量
//	ReadOnly   bool         `json:"readonly"`       //Readonlyであるか
//	Boot       uint         `json:"boot"`
//}

type Storage struct {
	Mode     uint    `json:"modestorageTmp"` //0:Manual 1:From ImaCon
	SrcSrv   SSHAuth `json:"src_srv"`        //src_srvをpullする際に使用するURL
	DstSrv   SSHAuth `json:"dst_srv"`        //dst_srvをpullする際に使用するURL
	VMName   string  `json:"vm_name"`        //VMNameがある場合は、Pathに追加する
	Type     uint    `json:"type"`           //1: CDROM 2:Floppy (no support) 10:BootDev(VirtIO) 11: BootDev(SATA) 12: BootDev(IDE)
	FileType uint    `json:"filetype"`       //0:qcow2 1:img
	PathType uint    `json:"path_type"`      //node側のストレージの種類 0~9:SSD 10~19:HDD 20~29:NVMe 100~109:SSD(iSCSI) 110~119:SSD(iSCSI) 120~129:NVme(iSCSI)
	Path     string  `json:"path"`           //node側のパス
	Capacity uint    `json:"capacity"`       //容量
	ReadOnly bool    `json:"readonly"`       //Readonlyであるか
	Boot     uint    `json:"boot"`
}

type VMStorage struct {
	Type     uint   `json:"type"`      //0:BootDev(VirtIO) 1: CDROM 2:Floppy (no support) 11: BootDev(SATA) 12: BootDev(IDE)
	FileType uint   `json:"file_type"` //0:qcow2 1:raw
	Path     string `json:"path"`      //node側のパス or storage type(hdd1,hdd2,ssd1,ssd2,nvme1,nvme2)
	ReadOnly bool   `json:"readonly"`  //Readonlyであるか
	Boot     uint   `json:"boot"`
}

type SSHAuth struct {
	IP   string `json:"ip"`
	User string `json:"user"`
	Pass string `json:"pass"`
	Path string `json:"path"`
}

type Convert struct {
	SrcFile string `json:"src_file"`
	SrcType string `json:"src_type"`
	DstFile string `json:"dst_file"`
	DstType string `json:"dst_type"`
}

type GenerateStorageXml struct {
	Storage       VMStorage
	Number        uint
	PCISlot       uint
	AddressNumber uint
}

type FileTransfer struct {
	URL         string
	CurrentSize int64
	AllSize     int64
}

type Auth struct {
	IP   string
	Port uint
	User string
	Pass string
}
