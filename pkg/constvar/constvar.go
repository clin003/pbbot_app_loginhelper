package constvar

import (
	"fmt"
	"time"
)

func APPDesc() string {
	return fmt.Sprintf("交流QQ群: 1051824036, 网站：www.lyhuilin.com \nCopyright ©2018-%d LYHUILIN Team. All Rights Reserved", time.Now().Year())
}
func APPDesc404() string {
	return fmt.Sprintf("慧林淘友交流QQ群：153690156 ，网站：www.lyhuilin.com (Copyright ©2018-%d LYHUILIN Team All Rights Reserved)(Error API route.)", time.Now().Year())
}
