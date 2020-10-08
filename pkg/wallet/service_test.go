package wallet
import (
	"github.com/iamgafurov/wallet_panic/pkg/types"
	"testing"
	)
		
func TestService_Repeat_success(t *testing.T){
	srv := &Service{}
	srv.RegisterAccount(types.Phone("90999390"))
	account,err := srv.FindAccountByID(1)
	if err != nil {
		t.Errorf("Repeat(): cant't register account,error = %v",err)
	}
	err = srv.Deposit(account.ID , types.Money(30_000))
	if err != nil{
		t.Errorf("Repeat(): cant't register account,error = %v",err)
	}
	payment,er := srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Repeat(): cant't register account,error = %v",err)
	}
	payment,err =srv.Repeat(payment.ID)
	if err !=nil {
		t.Errorf("Repeat(): cant't register account,error = %v",err)
	}
}

func TestService_Repeat_fail(t *testing.T){
	srv := &Service{}
	srv.RegisterAccount(types.Phone("90999390"))
	account,err := srv.FindAccountByID(1)
	if err != nil {
		t.Errorf("Repeat(): cant't register account,error = %v",err)
	}
	err = srv.Deposit(account.ID , types.Money(30_000))
	if err != nil{
		t.Errorf("Repeat(): cant't register account,error = %v",err)
	}
	payment,er := srv.Pay(account.ID,types.Money(25000),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Repeat(): cant't register account,error = %v",err)
	}
	payment,err =srv.Repeat(payment.ID)
	if err ==nil {
		t.Errorf("Repeat(): cant't register account,error = %v",err)
	}
}


func TestService_FavoritePayment_success(t *testing.T){
	srv := &Service{}
	srv.RegisterAccount(types.Phone("90999390"))
	account,err := srv.FindAccountByID(1)
	if err != nil {
		t.Errorf("Favorite(): cant't register account,error = %v",err)
	}
	err = srv.Deposit(account.ID , types.Money(30_000))
	if err != nil{
		t.Errorf("Favorite(): deposit,error = %v",err)
	}
	payment,er := srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Favorite(): cant't pay,error = %v",err)
	}
	_,err =srv.FavoritePayment(payment.ID, "test name")
	if err !=nil {
		t.Errorf("Favorite(): cant't Favorite,error = %v",err)
	}
}

func TestService_PayFromFavorite_success(t *testing.T){
	srv := &Service{}
	srv.RegisterAccount(types.Phone("90999390"))
	account,err := srv.FindAccountByID(1)
	if err != nil {
		t.Errorf("Favorite(): cant't register account,error = %v",err)
	}
	err = srv.Deposit(account.ID , types.Money(30_000))
	if err != nil{
		t.Errorf("Favorite(): deposit,error = %v",err)
	}
	payment,er := srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Favorite(): cant't pay,error = %v",err)
	}
	favorite,er :=srv.FavoritePayment(payment.ID, "test name")
	if er !=nil {
		t.Errorf("Favorite(): cant't Favorite,error = %v",err)
	}
	_,err =srv.PayFromFavorite(favorite.ID)
	if err !=nil {
		t.Errorf("Favorite(): cant't Favorite,error = %v",err)
	} 
}


func TestService_FindFavoriteByID_success(t *testing.T){
	srv := &Service{}
	srv.RegisterAccount(types.Phone("90999390"))
	account,err := srv.FindAccountByID(1)
	if err != nil {
		t.Errorf("Favorite(): cant't register account,error = %v",err)
	}
	err = srv.Deposit(account.ID , types.Money(30_000))
	if err != nil{
		t.Errorf("Favorite(): deposit,error = %v",err)
	}
	payment,er := srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Favorite(): cant't pay,error = %v",err)
	}
	favorite,er :=srv.FavoritePayment(payment.ID,"test name")
	if er !=nil {
		t.Errorf("Favorite(): cant't Favorite,error = %v",err)
	}
	_,err =srv.FindFavoriteByID(favorite.ID)
	if err !=nil {
		t.Errorf("Favorite(): cant't Favorite,error = %v",err)
	} 
}

func TestService_Pay_fail(t *testing.T){
	srv := &Service{}
	srv.RegisterAccount(types.Phone("90999390"))
	account,err := srv.FindAccountByID(1)
	if err != nil {
		t.Errorf("Favorite(): cant't register account,error = %v",err)
	}
	err = srv.Deposit(account.ID , types.Money(1_000))
	if err != nil{
		t.Errorf("Favorite(): deposit,error = %v",err)
	}
	_,er := srv.Pay(account.ID,types.Money(150000),types.PaymentCategory("food"))
	if er ==nil {
		t.Errorf("Favorite(): cant't pay,error = %v",err)
	}
}
func TestService_Pay_success(t *testing.T){
	srv := &Service{}
	srv.RegisterAccount(types.Phone("90999390"))
	account,err := srv.FindAccountByID(1)
	if err != nil {
		t.Errorf("Favorite(): cant't register account,error = %v",err)
	}
	err = srv.Deposit(account.ID , types.Money(18_000))
	if err != nil{
		t.Errorf("Favorite(): deposit,error = %v",err)
	}
	_,er := srv.Pay(account.ID,types.Money(15000),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Favorite(): cant't pay,error = %v",err)
	}
}
func TestService_Reject_success(t *testing.T){
	srv := &Service{}
	srv.RegisterAccount(types.Phone("90999390"))
	account,err := srv.FindAccountByID(1)
	if err != nil {
		t.Errorf("Reject(): cant't register account,error = %v",err)
	}
	err = srv.Deposit(account.ID , types.Money(30_000))
	if err != nil{
		t.Errorf("Reject(): deposit,error = %v",err)
	}
	payment,er := srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Reject(): cant't pay,error = %v",err)
	}
	err =srv.Reject(payment.ID)
	if err !=nil {
		t.Errorf("Reject(): cant't reject,error = %v",err)
	}
}

func TestService_Reject_fail(t *testing.T){
	srv := &Service{}
	srv.RegisterAccount(types.Phone("90999390"))
	account,err := srv.FindAccountByID(1)
	if err != nil {
		t.Errorf("Reject(): cant't register account,error = %v",err)
	}
	err = srv.Deposit(account.ID , types.Money(30_000))
	if err != nil{
		t.Errorf("Reject(): deposit,error = %v",err)
	}
	_,er := srv.Pay(account.ID,types.Money(100),types.PaymentCategory("food"))
	if er !=nil {
		t.Errorf("Reject(): cant't pay,error = %v",err)
	}
	err =srv.Reject("-1")
	if err ==nil {
		t.Errorf("Reject(): cant't reject,error = %v",err)
	}
}

func TestService_FinAccountByID_fail(t *testing.T){
	srv := &Service{}
	srv.RegisterAccount(types.Phone("90999390"))
	_,err :=srv.FindAccountByID(int64(-1))
	if err ==nil {
		t.Errorf("Reject(): cant't reject,error = %v",err)
	}
}