package vpc

// DnsStatus the model 'DnsStatus'
type DnsStatus string

// List of DnsStatus
const (
	DnsStatusClose   DnsStatus = "close"
	DnsStatusWait    DnsStatus = "wait"
	DnsStatusSyncing DnsStatus = "syncing"
	DnsStatusOpen    DnsStatus = "open"
	DnsStatusClosing DnsStatus = "closing"
)

// All allowed values of DnsStatus enum
var AllowedDnsStatusEnumValues = []DnsStatus{
	"close",
	"wait",
	"syncing",
	"open",
	"closing",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DnsStatus) IsValid() bool {
	for _, existing := range AllowedDnsStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
