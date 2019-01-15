package billing

import (
	"encoding/json"
	"github.com/ProtocolONE/payone-repository/tools"
)

type JsonProject struct {
	Id string `json:"id"`
	*Project
}

func (m Project) MarshalJSON() ([]byte, error) {
	return json.Marshal(&JsonProject{Id: tools.ByteToObjectId(m.Id).Hex(), Project: (*Project)(&m)})
}
