package service

import (
	"GeoAPI/internal/service/models"
	"encoding/json"
	"log"
)

type Convertor interface {
	ConversionFromSourceToSearch() (search []*models.AddressSearch, err error)
	ConversionFromSourceToGeo() (geo *models.AddressGeo, err error)
}

type ServiceConvertor struct {
	data []byte
	models.AddressSearch
}

func NewServiceConvertor(dataInByte []byte) ServiceConvertor {
	return ServiceConvertor{
		data: dataInByte,
	}
}

func (sc *ServiceConvertor) ConversionFromSourceToSearch() (search []*models.AddressSearch, err error) {
	log.Println(string(sc.data))
	if err := json.Unmarshal(sc.data, &search); err != nil {
		return search, err
	}
	return search, nil
}

func (sc *ServiceConvertor) ConversionFromSourceToGeo() (geo *models.AddressGeo, err error) {
	log.Println(string(sc.data))
	if err := json.Unmarshal(sc.data, &geo); err != nil {
		return geo, err
	}

	return geo, nil
}
