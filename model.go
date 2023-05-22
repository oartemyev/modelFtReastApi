package ftRestApi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ekomobile/dadata/v2/api/model"
)

type CheckReq struct {
	ID               int64          `json:"id"`
	ExternalUserID   string         `json:"external_user_id"`
	CityFias         string         `json:"city_fias"`
	CreatedOn        string         `json:"created_on"`
	Name             string         `json:"name"`
	Phone            string         `json:"phone"`
	Email            string         `json:"email"`
	City             string         `json:"city"`
	Address          string         `json:"address"`
	Apt              int64          `json:"apt"`
	Price            float64        `json:"price"`
	DeliveryPrice    float64        `json:"delivery_price"`
	DeliveryID       string         `json:"delivery_id"`
	DeliveryName     string         `json:"delivery_name"`
	Promocode        string         `json:"promocode"`
	LoyaltyCard      string         `json:"loyalty_card"`
	PickupLocationID string         `json:"pickup_location_id"`
	PaymentID        string         `json:"payment_id"`
	PaymentName      string         `json:"payment_name"`
	DeliveryComment  string         `json:"delivery_comment"`
	Items            []CheckItemReq `json:"items"`
}

type CheckItemReq struct {
	Name     string  `json:"name"`
	OfferID  string  `json:"offer_id"`
	Quantity float64 `json:"quantity"`
	Price    float64 `json:"price"`
	Subtotal float64 `json:"subtotal"`
}

/*
	БЛОК АВТОРИЗАЦИИ
*/

type AuthResponse interface {
	StatusCode() int
}

type FittinResponse interface {
	StatusCode() int
	MessageError() string
}

type FittinErrorResponse struct {
	Err  ErrorMessage `json:"error"`
	Code int          `json:"-"`
}

type AuthPart1Requset struct {
	UserIdentifier string `json:"userIdentifier"` //{"userIdentifier": "71001001010"}
}

type Auth1Otp struct {
	Timeout    int    `json:"timeout"`
	Message    string `json:"message"`
	CodeLength int    `json:"codeLength"`
}

type AuthPart1Response struct {
	Otp   Auth1Otp `json:"otp"`
	Error string   `json:"error,omitempty"`
	Code  int      `json:"-"`
}

// type AuthErr struct {
// 	Message string `json:"message"`
// }

// type AuthErrPart1Response struct {
// 	Err  AuthErr `json:"error"`
// 	Code int     `json:"-"`
// }

func (e AuthPart1Response) StatusCode() int {
	if e.Code == 0 {
		return http.StatusOK
	}

	return e.Code
}

func (e AuthPart1Response) MessageError() string {
	return e.Error
}

func (e FittinErrorResponse) StatusCode() int {
	if e.Code == 0 {
		return http.StatusOK
	}

	return e.Code
}

func (e FittinErrorResponse) MessageError() string {
	return e.Err.Message
}

type AuthPart2Requset struct {
	UserIdentifier string `json:"phone"` //{"userIdentifier": "71001001010"}
	Otp            string `json:"otp"`
}

type AuthPart2Response struct {
	User  UserSms `json:"user"`
	Error string  `json:"error,omitempty"`
	Code  int     `json:"-"`
}

func (e AuthPart2Response) StatusCode() int {
	if e.Code == 0 {
		return http.StatusOK
	}

	return e.Code
}

func (e AuthPart2Response) MessageError() string {
	return e.Error
}

type User struct {
	ID               string    `json:"external_user_id,omitempty"`
	Name             string    `json:"name,omitempty"`
	Phone            string    `json:"phone,omitempty"`
	Email            string    `json:"email,omitempty"`
	ExtraLoyaltyCard string    `json:"extraLoyaltyCard,omitempty"`
	CardNumber       string    `json:"cardNumber,omitempty"`
	CardBarcode      string    `json:"cardBarcode,omitempty"`
	Bonuses          float64   `json:"bonuses,omitempty"`
	Sex              string    `json:"sex,omitempty"`
	ErrMsg           string    `json:"error,omitempty"`
	Birthday         ShortDate `json:"birthday,omitempty"`
	Code             int       `json:"-"`
}

func (e User) StatusCode() int {
	if e.Code == 0 {
		return http.StatusOK
	}

	return e.Code
}

func (e User) MessageError() string {
	return e.ErrMsg
}

type UserSms struct {
	ID               string  `json:"external_user_id"`
	Name             string  `json:"name"`
	Phone            string  `json:"phone"`
	Email            string  `json:"email"`
	ExtraLoyaltyCard string  `json:"extraLoyaltyCard"`
	CardNumber       string  `json:"cardNumber"`
	Balance          float64 `json:"bonuses"`
}

type UserUpdate struct {
	ID       string    `json:"external_user_id"`
	Name     string    `json:"name"`
	Phone    string    `json:"phone"`
	Email    string    `json:"email"`
	Sex      string    `json:"sex"`
	Birthday ShortDate `json:"birthday"`
}

type UserCardBind struct {
	ID         string `json:"external_user_id"`
	CardNumber string `json:"cardNumber"`
	Pin        string `json:"pin"`
}

type UserQueryNewCard struct {
	ID string `json:"external_user_id"`
}

type UserReturn struct {
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Sex         string    `json:"sex"`
	Birthday    ShortDate `json:"birthday"`
	CardNumber  string    `json:"cardNumber"`
	CardBarcode string    `json:"cardBarcode"`
	Bonuses     float64   `json:"bonuses"`
}

type ShopQueryByAddr struct {
	City    string `json:"city"`
	Address string `json:"address"`
	Apt     string `json:"apt"`
}

type ProfilResponse struct {
	M    map[string]interface{}
	Code int
}

func (c ProfilResponse) MarshalJSON() ([]byte, error) {
	//return []byte(fmt.Sprintf(`"%s"`, c.Time.Format(layout))), nil
	return json.Marshal(c.M)
}

func (c *ProfilResponse) UnmarshalJSON(b []byte) (err error) {
	// s := strings.Trim(string(b), `"`) // remove quotes
	// if s == "null" {
	// 	return
	// }
	// c.Time, err = time.Parse(layoutShortDate, s)
	err = json.Unmarshal(b, &c.M)
	return
}

func (e ProfilResponse) StatusCode() int {
	if e.Code == 0 {
		return http.StatusOK
	}

	return e.Code
}

type AvailabilityRequest struct {
	City     string   `json:"city" xml:"city"`
	CityFias string   `json:"city_fias" xml:"city_fias"`
	OfferID  []string `json:"offer_ids" xml:"offer_ids"`
}

type ErrorMessage struct {
	Message string `json:"message,omitempty"`
}

type AvailabilityResponse struct {
	Warehouses   []AvailShopRsp  `json:"warehouses,omitempty"`
	Availability []AvailOfferRsp `json:"availability,omitempty"`
	Err          ErrorMessage    `json:"error,omitempty"`
	Code         int             `json:"-"`
}

type AvailShopRsp struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Address string  `json:"address"`
	City    string  `json:"city"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Online  bool    `json:"online"`
	Subway  string  `json:"subway"`
	Mall    string  `json:"mall"`
}

type AvailOfferRsp struct {
	OfferID     string  `json:"offer_id"`
	WarehouseID string  `json:"warehouse_id"`
	Quantity    float64 `json:"quantity"`
}

func (e AvailabilityResponse) StatusCode() int {
	if e.Code == 0 {
		return http.StatusOK
	}

	return e.Code
}

type BasketItem struct {
	Name     string  `json:"name" xml:"@Name"`
	OfferID  string  `json:"offer_id" xml:"@ArticleID"`
	Quantity float64 `json:"quantity" xml:"@Quantity"`
	Price    float64 `json:"price" xml:"@PriceGT"`
	OldPrice float64 `json:"old_price" xml:"@PriceNoGT"`
	Subtotal float64 `json:"subtotal" xml:"@Subtotal"`
}

func (p BasketItem) ToXml() string {
	str := "<BasketItem " +
		`Name="` + p.Name + `" ` +
		`ArticleID="` + p.OfferID + `" ` +
		`Quantity="` + fmt.Sprintf("%.2f", p.Quantity) + `" ` +
		`PriceGT="` + fmt.Sprintf("%.2f", p.Price) + `" ` +
		`PriceNoGT="` + fmt.Sprintf("%.2f", p.OldPrice) + `" ` +
		`Subtotal="` + fmt.Sprintf("%.2f", p.Subtotal) + `" />`

	return str
}

type Basket struct {
	StorageID        string       `json:"storage_id"`
	ExternalUserID   string       `json:"external_user_id"`
	CityFias         string       `json:"city_fias"`
	City             string       `json:"city"`
	DeliveryID       string       `json:"delivery_id"`
	PaymentID        string       `json:"payment_id"`
	PickupLocationID string       `json:"pickupLocationId"`
	Promocode        string       `json:"promocode"`
	LoyaltyCard      string       `json:"loyalty_card"`
	Bonuses          float64      `json:"bonuses"`
	Items            []BasketItem `json:"items"`
}

func (p Basket) ToXml() string {
	str := "<items>"
	for _, v := range p.Items {
		str = str + v.ToXml()
	}
	str = str + "</items>"

	return str
}

type BasketItemResp struct {
	Name             string  `json:"name"`
	OfferID          string  `json:"offer_id"`
	Quantity         float64 `json:"quantity"`
	Price            float64 `json:"price"`
	OldPrice         float64 `json:"old_price"`
	Subtotal         float64 `json:"subtotal"`
	AppliedPromocode string  `json:"applied_promocode"`
}

type BasketResponse struct {
	Success             bool             `json:"success"`
	TotalPrice          float64          `json:"total_price"`
	TotalDiscount       float64          `json:"total_discount"`
	LoyaltyCardDiscount float64          `json:"loyalty_card_discount"`
	PromocodeDiscount   float64          `json:"promocode_discount"`
	TotalOldDiscount    float64          `json:"total_old_discount"`
	BonusesDiscount     float64          `json:"bonuses_discount"`
	PromoActionType     int              `json:"promo_action_type"`
	Items               []BasketItemResp `json:"items,omitempty"`
	ErrMsg              string           `json:"error,omitempty"`
	Code                int              `json:"-"`
}

func (e BasketResponse) StatusCode() int {
	if e.Code == 0 {
		return http.StatusOK
	}

	return e.Code
}

func (e BasketResponse) MessageError() string {
	return e.ErrMsg
}

type OrderItemRequest struct {
	Name     string  `json:"name"`
	OfferID  string  `json:"offer_id"`
	Quantity float64 `json:"quantity"`
	Price    float64 `json:"price"`
	Subtotal float64 `json:"subtotal"`
}

type OrderRequest struct {
	ID               int                `json:"id" xml:"id"`
	StorageID        string             `json:"storage_id" xml:"storage_id"`
	ExternalUserID   string             `json:"external_user_id" xml:"external_user_id"`
	CityFias         string             `json:"city_fias" xml:"city_fias"`
	CreatedOn        string             `json:"created_on" xml:"created_on"`
	Name             string             `json:"name" xml:"name"`
	Phone            string             `json:"phone" xml:"phone"`
	Email            string             `json:"email" xml:"email"`
	City             string             `json:"city" xml:"city"`
	Address          string             `json:"address" xml:"address"`
	Apt              int                `json:"apt" xml:"apt"`
	AddressData      model.Address      `json:"addressData" xml:"addressData"`
	Price            float64            `json:"price" xml:"price"`
	Bonuses          float64            `json:"bonuses" xml:"bonuses"`
	DeliveryPrice    float64            `json:"delivery_price" xml:"delivery_price"`
	DeliveryID       string             `json:"delivery_id" xml:"delivery_id"`
	DeliveryName     string             `json:"delivery_name" xml:"delivery_name"`
	Promocode        string             `json:"promocode" xml:"promocode"`
	LoyaltyCard      string             `json:"loyalty_card" xml:"loyalty_card"`
	PickupLocationID string             `json:"pickup_location_id" xml:"pickup_location_id"`
	PaymentID        string             `json:"payment_id" xml:"payment_id"`
	PaymentName      string             `json:"payment_name" xml:"payment_name"`
	DeliveryComment  string             `json:"delivery_comment" xml:"delivery_comment"`
	DeliveryDate     string             `json:"delivery_date_interval_id" xml:"delivery_date_interval_id"`
	DeliveryTime     string             `json:"delivery_time_interval_id" xml:"delivery_time_interval_id"`
	Items            []OrderItemRequest `json:"items" xml:"-"`
}

func (p OrderItemRequest) ToXml() string {
	str := "<OrderItemReq " +
		`Name="` + p.Name + `" ` +
		`ArticleID="` + p.OfferID + `" ` +
		`Quantity="` + fmt.Sprintf("%.2f", p.Quantity) + `" ` +
		`Price="` + fmt.Sprintf("%.2f", p.Price) + `" ` +
		`Subtotal="` + fmt.Sprintf("%.2f", p.Subtotal) + `" />`

	return str
}

func (p OrderRequest) ToXml() string {
	str := "<items>"
	for _, v := range p.Items {
		str = str + v.ToXml()
	}
	str = str + "</items>"

	return str
}

/*
Для создания заказа
*/
type OrderItemResp struct {
	OfferID  string  `json:"offer_id"`
	Price    float64 `json:"price"`
	Quantity float64 `json:"quantity"`
	Discount float64 `json:"discount"`
	Subtotal float64 `json:"subtotal"`
}

type OrdersResp struct {
	Success       bool            `json:"success"`
	ID            int             `json:"id"`
	ExternalID    string          `json:"external_id"`
	Error         int             `json:"error"`
	ErrorText     string          `json:"error_text,omitempty"`
	Price         float64         `json:"price"`
	DeliveryPrice float64         `json:"delivery_price"`
	FullPrice     float64         `json:"full_price"`
	Discount      float64         `json:"discount"`
	Items         []OrderItemResp `json:"items"`
}

type OrderResp struct {
	Orders []OrdersResp `json:"orders"`
	Error  string       `json:"error,omitempty"`
	Code   int          `json:"-"`
}

func (e OrderResp) StatusCode() int {
	if e.Code == 0 {
		return http.StatusOK
	}

	return e.Code
}

func (e OrderResp) MessageError() string {
	return e.Error
}

// Способы доставки
type DeliveryWayRequest struct {
	ExternalUserID string                   `json:"external_user_id"`
	LoyaltyCard    string                   `json:"loyalty_card"`
	AddressData    model.Address            `json:"addressData"`
	Items          []DeliveryWayItemRequest `json:"items"`
}

type DeliveryWayItemRequest struct {
	Name     string  `json:"name"`
	OfferID  string  `json:"offer_id"`
	Quantity float64 `json:"quantity"`
}

type DeliveryWayItem struct {
	Id           string                 `json:"id"`    // regular - Доставка курьером, pickup - Самовывоз из магазина
	Title        string                 `json:"title"` // Название типа доставки
	Description  string                 `json:"description"`
	Tip          string                 `json:"type"`            // delivery или pickup
	Price        float64                `json:"price,omitempty"` // стоимость доставки
	Min          int64                  `json:"min"`             // Минимальное и максимальное количество дней для доставки (0 - сегодня, 1 - завтра и так далее)
	Max          int64                  `json:"max"`             // Минимальное и максимальное количество дней для доставки (0 - сегодня, 1 - завтра и так далее)
	DateInterval []DeliveryDateInterval `json:"date_intervals,omitempty"`
	Locations    []DeliveryLocation     `json:"locations,omitempty"`
}

type DeliveryTimeInterval struct {
	Id    string `json:"id"`    // "10-23"
	Title string `json:"title"` //"10:00-23:00"
}

type DeliveryDateInterval struct {
	Id           string                 `json:"id"`        // "2022-10-20"
	Title        string                 `json:"title"`     //"20 октября""
	SubTitle     string                 `json:"sub_title"` // "чт"
	TimeInterval []DeliveryTimeInterval `json:"time_intervals"`
}

type DeliveryLocation struct {
	FirmaID int64   `json:"id"`
	Name    string  `json:"title"`
	Address string  `json:"address"`
	City    string  `json:"city"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}

type DeliveryWayResp struct {
	Dev   []DeliveryWayItem `json:"deliveries"`
	Error string            `json:"error,omitempty"`
	Code  int               `json:"-"`
}

func (e DeliveryWayResp) StatusCode() int {
	if e.Code == 0 {
		return http.StatusOK
	}

	return e.Code
}

func (e DeliveryWayResp) MessageError() string {
	return e.Error
}

type IntervalDelyvery struct {
	City      string `json:"city"`
	TimeStart string `json:"time_start"`
	TimeEnd   string `json:"time_end"`
}

type ArrayIntervalDelyvery []IntervalDelyvery

type OrderListRequest struct {
	ExternalUserID string `json:"external_user_id"`
	Phone          string `json:"phone"`
}

type OrderListItem struct {
	OfferID string  `json:"offer_id"`
	Price   float64 `json:"price"`
	Count   float64 `json:"count"`
}

type OrderListItemSql struct {
	ID string `json:"id"`
	OrderListItem
}

type OrderList struct {
	ID            string          `json:"id"`
	StorageID     string          `json:"storage_id"`
	Price         float64         `json:"price"`
	TotalPrice    float64         `json:"total_price"`
	CreatedAt     string          `json:"created_at"`
	StatusText    string          `json:"status_text"`
	DetailText    string          `json:"detail_text"`
	UserName      string          `json:"user_name"`
	DeliveryPrice float64         `json:"delivery_price"`
	DeliveryName  string          `json:"delivery_name"`
	PaymentStatus string          `json:"payment_status"`
	PaymentName   string          `json:"payment_name"`
	Completed     bool            `json:"completed"`
	ItemsPrice    float64         `json:"items_price"`
	Items         []OrderListItem `json:"items"`
}

type OrderListResponse struct {
	Orders []OrderList `json:"orders"`
	Error  string      `json:"error,omitempty"`
	Code   int         `json:"-"`
}

func (e OrderListResponse) StatusCode() int {
	if e.Code == 0 {
		return http.StatusOK
	}

	return e.Code
}

func (e OrderListResponse) MessageError() string {
	return e.Error
}

type CheckListItem struct {
	Checkdate string `json:"checkdate"`
	Amount    int    `json:"amount"`
	ID        int    `json:"id"`
}

type CheckListLoya struct {
	Checks []CheckListItem `json:"checks"`
}
