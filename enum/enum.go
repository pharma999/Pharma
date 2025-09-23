package enum

// UserStatus represents the status of a user account
type UserStatus string

const (
	Active    UserStatus = "ACTIVE"
	Inactive  UserStatus = "INACTIVE"
	Suspended UserStatus = "SUSPENDED"
)

// BlockStatus represents the block status of a user
type BlockStatus string

const (
	Blocked   BlockStatus = "BLOCKED"
	Unblocked BlockStatus = "UNBLOCKED"
)


// Gender represents
type Gender string

const (
	Male   Gender = "MALE"
	Female Gender = "FEMALE"
	Other  Gender = "OTHER"
)


// UserServiceStatus represents the service subscription status of a user
type UserServiceStatus string

const (
	Subscribed   UserServiceStatus = "SUBSCRIBED"
	Unsubscribed UserServiceStatus = "UNSUBSCRIBED"
	Trial        UserServiceStatus = "TRIAL"
)


// ServiceStatus represents the number of service a user has
type ServiceStatus string

const (
	New       ServiceStatus = "NEW"
	Secondary ServiceStatus = "SECONDARY"
	Multiple  ServiceStatus = "MULTIPLE"
)
