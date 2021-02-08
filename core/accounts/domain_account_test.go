package accounts

import (
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
	"github.com/red_envelope/services"
)

func TestAccountDomain_Create(t *testing.T) {
	dto := services.AccountDTO{
		UserId:   ksuid.New().Next().String(),
		Username: "测试用户",
		Balance:  decimal.NewFromFloat(0),
		Status:   1,
	}
	domain := new(accountDomain)
	Convey("账户创建", t, func() {
		rdto, err := domain.Create(dto)
		So(err, ShouldBeNil)
		So(rdto, ShouldNotBeNil)
		So(rdto.Balance.String(), ShouldEqual, dto.Balance.String())
		So(rdto.UserId, ShouldEqual, dto.UserId)
		So(rdto.Username, ShouldEqual, dto.Username)
		So(rdto.Status, ShouldEqual, dto.Status)
	})
}

func TestAccountDomain_Transfer(t *testing.T) {
	domain := accountDomain{}
	Convey("转账测试", t, func() {
		//10个账户，交易主体账户要有余额
		accounts := make([]*services.AccountDTO, 0)
		adto := make([]*services.AccountDTO, 0)
		size := 10
		for i := 0; i < size; i++ {
			account := &services.AccountDTO{
				UserId:       ksuid.New().Next().String(),
				Username:     "测试用户" + strconv.Itoa(i+1),
				Balance:  decimal.NewFromFloat(100),
				Status:   1,
			}
			adto = append(adto, account)
			//账户创建
			acDto, err := domain.Create(*account)
			So(err, ShouldBeNil)
			So(acDto, ShouldNotBeNil)

			So(acDto.Balance.String(), ShouldEqual, account.Balance.String())
			So(acDto.UserId, ShouldEqual, account.UserId)
			So(acDto.Username, ShouldEqual, account.Username)
			So(acDto.Status, ShouldEqual, account.Status)
			accounts = append(accounts, acDto)
		}
		for i:=0;i<size;i++{
			adto[i] = accounts[i]
		}

		//转账操作验证
		//1. 余额充足，金额转入其他账户
		Convey("余额充足，金额转入其他账户", func() {
			amount := decimal.NewFromFloat(1)
			body := services.TradeParticipator{
				AccountNo: adto[1].AccountNo,
				UserId:    adto[1].UserId,
				Username:  adto[1].Username,
			}
			target := services.TradeParticipator{
				AccountNo: adto[2].AccountNo,
				UserId:    adto[2].UserId,
				Username:  adto[2].Username,
			}
			dto := services.AccountTransferDTO{
				TradeBody:   body,
				TradeTarget: target,
				TradeNo:     ksuid.New().Next().String(),
				Amount:      amount,
				ChangeType:  services.ChangeType(-1),
				ChangeFlag:  services.FlagTransferOut,
				Decs:        "转账",
			}
			status, err := domain.Transfer(dto)
			So(err, ShouldBeNil)
			So(status, ShouldEqual, services.TransferedStatusSuccess)
			//实际余额更新后的预期值
			a2 := domain.GetAccount(adto[1].AccountNo)
			So(a2, ShouldNotBeNil)
			So(a2.Balance.String(),
				ShouldEqual,
				adto[1].Balance.Sub(amount).String())

		})
		//2. 余额不足，金额转出
		Convey("余额不足，金额转出", func() {
			amount := adto[3].Balance
			amount = amount.Add(decimal.NewFromFloat(200))
			body := services.TradeParticipator{
				AccountNo: adto[3].AccountNo,
				UserId:    adto[3].UserId,
				Username:  adto[3].Username,
			}
			target := services.TradeParticipator{
				AccountNo: adto[4].AccountNo,
				UserId:    adto[4].UserId,
				Username:  adto[4].Username,
			}
			dto := services.AccountTransferDTO{
				TradeBody:   body,
				TradeTarget: target,
				TradeNo:     ksuid.New().Next().String(),
				Amount:      amount,
				ChangeType:  services.ChangeType(-1),
				ChangeFlag:  services.FlagTransferOut,
				Decs:        "转账",
			}
			status, err := domain.Transfer(dto)
			So(err, ShouldNotBeNil)
			So(status, ShouldEqual, services.TransferedStatusSufficientFunds)
			//实际余额更新后的预期值
			a2 := domain.GetAccount(adto[3].AccountNo)
			So(a2, ShouldNotBeNil)
			So(a2.Balance.String(),
				ShouldEqual,
				adto[3].Balance.String())

		})
		//3. 充值
		Convey("充值", func() {
			amount := decimal.NewFromFloat(11.1)
			body := services.TradeParticipator{
				AccountNo: adto[5].AccountNo,
				UserId:    adto[5].UserId,
				Username:  adto[5].Username,
			}
			target := services.TradeParticipator{
				AccountNo: adto[6].AccountNo,
				UserId:    adto[6].UserId,
				Username:  adto[6].Username,
			}
			dto := services.AccountTransferDTO{
				TradeBody:   body,
				TradeTarget: target,
				TradeNo:     ksuid.New().Next().String(),
				Amount:      amount,
				ChangeType:  services.AccountStoreValue,
				ChangeFlag:  services.FlagTransferIn,
				Decs:        "储值",
			}
			status, err := domain.Transfer(dto)
			So(err, ShouldBeNil)
			So(status, ShouldEqual, services.TransferedStatusSuccess)
			//实际余额更新后的预期值
			a2 := domain.GetAccount(adto[5].AccountNo)
			So(a2, ShouldNotBeNil)
			So(a2.Balance.String(),
				ShouldEqual,
				adto[5].Balance.Add(amount).String())

		})
	})

}

