package components

import (
	"log"
	"strconv"
	"unisun/api/classroom-listener/src/models"
	"unisun/api/classroom-listener/src/ports"
)

type MappingClassRoomPriceAdap struct {
	Service ports.ServicePrice
}

func NewMappingClassRoomPriceAdap(service ports.ServicePrice) *MappingClassRoomPriceAdap {
	return &MappingClassRoomPriceAdap{
		Service: service,
	}
}

func (srv *MappingClassRoomPriceAdap) MappingClassRoomPrice(value models.ClassRoom) (*models.ClassRoom, error) {
	result, err := srv.Service.GetClassRoomPrice(strconv.FormatInt(value.Id, 10))
	if err != nil {
		log.Panic(err)
	}
	value.Prices.RegularPrice = result.Data.RegularPrice
	value.Prices.SpecialPrice = result.Data.SpecialPrice
	return &value, nil
}
