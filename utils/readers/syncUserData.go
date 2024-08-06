package readers

import (
	"os"
	"projects/config"
)

func SyncUserData() {
	UserStore = FReaderUser(config.USER_FILE, os.O_RDONLY|os.O_CREATE)
}
