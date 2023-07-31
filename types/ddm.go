package types

import (
	"fmt"
	"strings"
)

// Mobile 手机号 132****7986
type Mobile string

func (m Mobile) MarshalJSON() ([]byte, error) {
	if len(m) != 11 {
		return []byte(`"` + m + `"`), nil
	}

	v := fmt.Sprintf("%s****%s", m[:3], m[len(m)-4:])
	return []byte(`"` + v + `"`), nil
}

// BankCard 银行卡号 622888******5676
type BankCard string

func (bc BankCard) MarshalJSON() ([]byte, error) {
	if len(bc) > 19 || len(bc) < 16 {
		return []byte(`"` + bc + `"`), nil
	}

	v := fmt.Sprintf("%s******%s", bc[:6], bc[len(bc)-4:])
	return []byte(`"` + v + `"`), nil
}

// IDCard 身份证号 1******7
type IDCard string

func (card IDCard) MarshalJSON() ([]byte, error) {
	if len(card) != 18 {
		return []byte(`"` + card + `"`), nil
	}

	v := fmt.Sprintf("%s******%s", card[:1], card[len(card)-1:])
	return []byte(`"` + v + `"`), nil
}

// IDName 姓名 *飞
type IDName string

func (name IDName) MarshalJSON() ([]byte, error) {
	if len(name) < 1 {
		return []byte(`""`), nil
	}
	nameRune := []rune(name)
	v := fmt.Sprintf("*%s", string(nameRune[1:]))
	return []byte(`"` + v + `"`), nil
}

// PassWord 密码 ******
type PassWord string

func (pw PassWord) MarshalJSON() ([]byte, error) {
	v := "******"
	return []byte(`"` + v + `"`), nil
}

// Email 邮箱 l***w@gmail.com
type Email string

func (e Email) MarshalJSON() ([]byte, error) {
	if !strings.Contains(string(e), "@") {
		return []byte(`"` + e + `"`), nil
	}
	split := strings.Split(string(e), "@")
	if len(split[0]) < 1 || len(split[1]) < 1 {
		return []byte(`"` + e + `"`), nil
	}
	v := fmt.Sprintf("%s***%s", split[0][:1], split[0][len(split[0])-1:])
	return []byte(`"` + v + "@" + split[1] + `"`), nil
}
