package billing

import "github.com/ProtocolONE/payone-repository/pkg/constant"

func (m *PaymentMethod) isBankCard() bool {
	return m.Group == constant.PaymentSystemGroupAliasBankCard
}
