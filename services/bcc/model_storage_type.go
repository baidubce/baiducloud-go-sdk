package bcc

// StorageType the model 'StorageType'
type StorageType string

// List of StorageType
const (
	StorageTypeEnhancedSsdPl0       StorageType = "enhanced_ssd_pl0"
	StorageTypeEnhancedSsdPl1       StorageType = "enhanced_ssd_pl1"
	StorageTypeSsdEnhanced          StorageType = "SSD_Enhanced"
	StorageTypeEnhancedSsdPl2       StorageType = "enhanced_ssd_pl2"
	StorageTypeEnhancedSsdPl3       StorageType = "enhanced_ssd_pl3"
	StorageTypeCloudHp1             StorageType = "cloud_hp1"
	StorageTypePremiumSsd           StorageType = "premium_ssd"
	StorageTypeHp1                  StorageType = "hp1"
	StorageTypeSsd                  StorageType = "ssd"
	StorageTypeHdd                  StorageType = "hdd"
	StorageTypeElasticEphemeralDisk StorageType = "elastic_ephemeral_disk"
	StorageTypeLocal                StorageType = "local"
	StorageTypeSata                 StorageType = "sata"
	StorageTypeLocalSsd             StorageType = "local_ssd"
	StorageTypeLocalHdd             StorageType = "local_hdd"
	StorageTypeStd1                 StorageType = "std1"
	StorageTypeLocalNvme            StorageType = "local_nvme"
	StorageTypeNvme                 StorageType = "nvme"
)

// All allowed values of StorageType enum
var AllowedStorageTypeEnumValues = []StorageType{
	"enhanced_ssd_pl0",
	"enhanced_ssd_pl1",
	"SSD_Enhanced",
	"enhanced_ssd_pl2",
	"enhanced_ssd_pl3",
	"cloud_hp1",
	"premium_ssd",
	"hp1",
	"ssd",
	"hdd",
	"elastic_ephemeral_disk",
	"local",
	"sata",
	"local_ssd",
	"local_hdd",
	"std1",
	"local_nvme",
	"nvme",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StorageType) IsValid() bool {
	for _, existing := range AllowedStorageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
