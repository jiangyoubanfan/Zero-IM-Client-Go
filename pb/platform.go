package pb

type (
	PlatformID int32
)

const (
	IOSPlatformID      PlatformID = 1
	AndroidPlatformID  PlatformID = 2
	WindowsPlatformID  PlatformID = 3
	OSXPlatformID      PlatformID = 4
	WebPlatformID      PlatformID = 5
	MiniWebPlatformID  PlatformID = 6
	LinuxPlatformID    PlatformID = 7
	IOSPlatformStr                = "IOS"
	AndroidPlatformStr            = "Android"
	WindowsPlatformStr            = "Windows"
	OSXPlatformStr                = "OSX"
	WebPlatformStr                = "Web"
	MiniWebPlatformStr            = "MiniWeb"
	LinuxPlatformStr              = "Linux"
)

func (p PlatformID) Int32() int32 {
	return int32(p)
}

func (p PlatformID) String() string {
	switch p {
	case IOSPlatformID:
		return IOSPlatformStr
	case AndroidPlatformID:
		return AndroidPlatformStr
	case WindowsPlatformID:
		return WindowsPlatformStr
	case OSXPlatformID:
		return OSXPlatformStr
	case WebPlatformID:
		return WebPlatformStr
	case MiniWebPlatformID:
		return MiniWebPlatformStr
	}
	return LinuxPlatformStr
}
