package businessLogic

import (
	"6fg-app-api/lib"
	"fmt"
	"os"
)

func GenerateEquipmentQR(equipId int) ([]byte, error) {
	qr, err := lib.GenerateQR(fmt.Sprintf("%s/equipment/%d", os.Getenv("DOMEIN"), equipId))
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	return qr, nil
}
