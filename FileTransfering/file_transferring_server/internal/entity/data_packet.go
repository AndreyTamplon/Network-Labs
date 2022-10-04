package entity

import (
	"encoding/json"
)

type DataPacket struct {
	Length     int    `json:"length"`
	Data       []byte `json:"data"`
	LastPacket bool   `json:"last_packet"`
}

func NewPacket(data []byte) DataPacket {
	return DataPacket{
		Length: len(data),
		Data:   data,
	}
}

func (p *DataPacket) Marshal() ([]byte, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (p *DataPacket) Unmarshal(data []byte) error {
	err := json.Unmarshal(data, p)
	if err != nil {
		return err
	}
	return nil
}
