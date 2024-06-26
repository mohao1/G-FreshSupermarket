// Code generated by goctl. DO NOT EDIT.
package types

type UserLoginReq struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type UserRegisterReq struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type GetShopReq struct {
	City string `json:"city"`
}

type ProductListReq struct {
	ShopId      string `json:"shopId"`
	ProductType string `json:"productType"`
	Quantity    int64  `json:"quantity"` //存在多少
	Limit       int64  `json:"limit"`    //需要多少
}

type LowProductListReq struct {
	ShopId      string `json:"shopId"`
	ProductType string `json:"productType"`
	Quantity    int64  `json:"quantity"` //存在多少
	Limit       int64  `json:"limit"`    //需要多少
}

type DetailedProductReq struct {
	ProductId string `json:"productId"`
	Type      string `json:"type"`
}

type AnnouncementListReq struct {
	ShopId string `json:"shopId"`
}

type GetUserInfo struct {
}

type TotalPriceReq struct {
	ProductId       []string `json:"productId"`
	ProductQuantity []int64  `json:"productQuantity"`
}

type PostOrderReq struct {
	ShopId          string   `json:"shopId"`
	ProductId       []string `json:"productId"`
	ProductQuantity []int64  `json:"productQuantity"`
	DeliveryTime    string   `json:"deliveryTime"`
	Notes           string   `json:"notes"`
}

type OrderListReq struct {
	Limit string `json:"limit"`
}

type GetDetailedOrderReq struct {
	OrderNumber string `json:"orderNumber"`
}

type OverOrderReq struct {
	OrderNumber string `json:"orderNumber"`
}

type CancellationOverOrderReq struct {
	OrderNumber string `json:"orderNumber"`
}

type ConfirmOrderReq struct {
	OrderNumber string `json:"orderNumber"`
}

type DataResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type StaffLoginReq struct {
	StaffId  string `json:"staffId"`
	PassWord string `json:"password"`
}

type ManageLoginReq struct {
	StaffId  string `json:"staffId"`
	PassWord string `json:"password"`
}

type PersonalInfoReq struct {
}

type GetUnreceivedOrderReq struct {
	Limit int64 `json:"limit"`
}

type ReceivingOrderReq struct {
	OrderNumber string `json:"orderNumber"`
}

type UnReceivingOrderReq struct {
	OrderNumber string `json:"orderNumber"`
}

type GetReceivedOrderReq struct {
	Limit int64 `json:"limit"`
}

type GetCompleteOrderReq struct {
	Limit int64 `json:"limit"`
}

type GetOrderDetailsReq struct {
	OrderNumber string `json:"orderNumber"`
}

type GetPositionListReq struct {
}

type GetEmployeeListReq struct {
}

type SetEmployeeReq struct {
	NewStaffId string `json:"newStaffId"`
	StaffName  string `json:"staffName"`
	PositionId string `json:"positionId"`
	PassWord   string `json:"passWord"`
}

type DeleteEmployeeReq struct {
	DeleteStaffId string `json:"deleteStaffId"`
}

type UpDateEmployeeReq struct {
	UpdateStaffId string `json:"updateStaffId"`
	StaffName     string `json:"staffName"`
	PositionId    string `json:"positionId"`
	PassWord      string `json:"passWord"`
}

type GetProductTypeListReq struct {
}

type GetProductListReq struct {
	Limit int64 `json:"limit"`
}

type GetProductDetailsReq struct {
	ProductId string `json:"productId"`
}

type ReleaseProductReq struct {
	ProductName     string `json:"productName"`
	ProductTitle    string `json:"productTitle"`
	ProductTypeId   string `json:"productTypeId"`
	ProductQuantity int64  `json:"productQuantity"`
	ProductPicture  string `json:"productPicture"`
	Price           string `json:"price"`
	ProductSize     int64  `json:"productSize"`
	Producer        string `json:"producer"`
}

type UpDateProductReq struct {
	ProductId       string `json:"productId"`
	ProductName     string `json:"productName"`
	ProductTitle    string `json:"productTitle"`
	ProductTypeId   string `json:"productTypeId"`
	ProductQuantity int64  `json:"productQuantity"`
	ProductPicture  string `json:"productPicture"`
	Price           string `json:"price"`
	ProductSize     int64  `json:"productSize"`
	Producer        string `json:"producer"`
}

type DeleteProductReq struct {
	ProductId string `json:"productId"`
}

type GetLowProductReq struct {
	Limit int64 `json:"limit"`
}

type GetLowProductDetailsReq struct {
	ProductId string `json:"productId"`
}

type ReleaseLowProductReq struct {
	ProductName     string `json:"productName"`
	ProductTitle    string `json:"productTitle"`
	ProductTypeId   string `json:"productTypeId"`
	ProductQuantity int64  `json:"productQuantity"`
	ProductPicture  string `json:"productPicture"`
	Price           string `json:"price"`
	ProductSize     int64  `json:"productSize"`
	Producer        string `json:"producer"`
	Quota           int64  `json:"quota"`
	StartTime       int64  `json:"startTime"`
	EndTime         int64  `json:"endTime"`
}

type UpDateLowProductReq struct {
	ProductId       string `json:"productId"`
	ProductName     string `json:"productName"`
	ProductTitle    string `json:"productTitle"`
	ProductTypeId   string `json:"productTypeId"`
	ProductQuantity int64  `json:"productQuantity"`
	ProductPicture  string `json:"productPicture"`
	Price           string `json:"price"`
	ProductSize     int64  `json:"productSize"`
	Producer        string `json:"producer"`
	Quota           int64  `json:"quota"`
	StartTime       int64  `json:"startTime"`
	EndTime         int64  `json:"endTime"`
}

type DeleteLowProductReq struct {
	ProductId string `json:"productId"`
}

type GetOverOrderReq struct {
	Limit int64 `json:"limit"`
}

type OverOrderHandleReq struct {
	OrderNumber string `json:"orderNumber"`
	TypeNumber  int64  `json:"typeNumber"`
}

type GetSalesDataReq struct {
}

type GetOrderDataReq struct {
	GetTime int64 `json:"getTime"`
}

type PostAnnouncementReq struct {
	NoticeTitle string `json:"noticeTitle"`
}

type GetAnnouncementListReq struct {
}

type UpdateAnnouncementReq struct {
	NoticeId    string `json:"noticeId"`
	NoticeTitle string `json:"noticeTitle"`
}

type DeleteAnnouncementReq struct {
	NoticeId string `json:"noticeId"`
}

type BmsDataResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type AdminLoginReq struct {
	AdminName string `json:"adminName"`
	PassWord  string `json:"passWord"`
}

type UpDatePassWordReq struct {
	PassWord    string `json:"passWord"`
	NewPassWord string `json:"newPassWord"`
}

type GetShopListReq struct {
	Limit int64 `json:"limit"`
}

type PostShopListReq struct {
	ShopId      string `json:"shopId"`
	ShopName    string `json:"shopName"`
	ShopAddress string `json:"shopAddress"`
	ShopCity    string `json:"shopCity"`
}

type PostAdminReq struct {
	StaffId   string `json:"staffId"`
	StaffName string `json:"staffName"`
	Password  string `json:"password"`
}

type GetAdminReq struct {
}

type DeleteAdminReq struct {
	StaffId string `json:"staffId"`
}

type PostShopAdminReq struct {
	StaffId string `json:"staffId"`
	ShopId  string `json:"shopId"`
}

type DeleteShopAdminReq struct {
	ShopId string `json:"shopId"`
}

type GetShopAdminReq struct {
	ShopId string `json:"shopId"`
}

type UpDateShopReq struct {
	ShopId      string `json:"shopId"`
	ShopName    string `json:"shopName"`
	ShopAddress string `json:"shopAddress"`
	ShopCity    string `json:"shopCity"`
}

type DeleteShopReq struct {
	ShopId string `json:"shopId"`
}

type GetProductTypeListAdminReq struct {
	Limit int64 `json:"limit"`
}

type PostProductTypeReq struct {
	ProductTypeName string `json:"productTypeName"`
	ProductTypeUnit string `json:"productTypeUnit"`
}

type UpDateProductTypeReq struct {
	ProductTypeId   string `json:"productTypeId"`
	ProductTypeName string `json:"productTypeName"`
	ProductTypeUnit string `json:"productTypeUnit"`
}

type DeleteProductTypeReq struct {
	ProductTypeId string `json:"productTypeId"`
}

type GetAllShopStaffListReq struct {
	Limit int64 `json:"limit"`
}

type GetShopStaffListReq struct {
	ShopId string `json:"shopId"`
	Limit  int64  `json:"limit"`
}

type GetShopAllStaffSumListReq struct {
}

type GetUserListReq struct {
	Limit int64 `json:"limit"`
}

type GetPositionListAdminReq struct {
	Limit int64 `json:"limit"`
}

type PostPositionReq struct {
	PositionName  string `json:"positionName"`
	PositionGrade int64  `json:"positionGrade"`
}

type UpDatePositionReq struct {
	PositionId    string `json:"positionId"`
	PositionName  string `json:"positionName"`
	PositionGrade int64  `json:"positionGrade"`
}

type DeletePositionReq struct {
	PositionId string `json:"positionId"`
}

type GetShopSumReq struct {
}

type GetUserSumReq struct {
}

type GetShopLowProductListReq struct {
	ShopId string `json:"shopId"`
	Limit  int64  `json:"limit"`
}

type GetShopLowProductSumReq struct {
	ShopId string `json:"shopId"`
}

type GetShopProductListReq struct {
	ShopId string `json:"shopId"`
	Limit  int64  `json:"limit"`
}

type GetShopProductSumReq struct {
	ShopId string `json:"shopId"`
}

type GetLowProductSumReq struct {
}

type GetProductSumReq struct {
}

type GetShopSalesRecordsSumReq struct {
	ShopId string `json:"shopId"`
}

type GetShopSalesRecordsListReq struct {
	ShopId string `json:"shopId"`
}

type GetShopTimeOrderSumReq struct {
	ShopId  string `json:"shopId"`
	TopTime int64  `json:"topTime"`
	EndTime int64  `json:"endTime"`
}

type GetShopOrderSumReq struct {
	ShopId string `json:"shopId"`
}

type GetOrderSumReq struct {
}

type GetNewUserSumToDayReq struct {
}

type AmsDataResp struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
