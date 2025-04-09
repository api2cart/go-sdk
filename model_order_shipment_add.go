/*
API2Cart OpenAPI

API2Cart

API version: 1.1
Contact: contact@api2cart.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the OrderShipmentAdd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderShipmentAdd{}

// OrderShipmentAdd struct for OrderShipmentAdd
type OrderShipmentAdd struct {
	// Defines the order for which the shipment will be created
	OrderId *string `json:"order_id,omitempty"`
	// Store Id
	StoreId *string `json:"store_id,omitempty"`
	// This parameter is used for selecting a warehouse where you need to set/modify a product quantity.
	WarehouseId *string `json:"warehouse_id,omitempty"`
	// Defines company name that provide tracking of shipment
	ShipmentProvider *string `json:"shipment_provider,omitempty"`
	// Define shipping method
	ShippingMethod *string `json:"shipping_method,omitempty"`
	// Defines items in the order that will be shipped
	Items []OrderShipmentAddItemsInner `json:"items,omitempty"`
	// Send notifications to customer after shipment was created
	SendNotifications *bool `json:"send_notifications,omitempty"`
	// Defines shipment's tracking numbers that have to be added</br> How set tracking numbers to appropriate carrier:<ul><li>tracking_numbers[]=a2c.demo1,a2c.demo2 - set default carrier</li><li>tracking_numbers[<b>carrier_id</b>]=a2c.demo - set appropriate carrier</li></ul>To get the list of carriers IDs that are available in your store, use the <a href = \"https://api2cart.com/docs/#/cart/CartInfo\">cart.info</a > method
	TrackingNumbers []OrderShipmentAddTrackingNumbersInner `json:"tracking_numbers,omitempty"`
	// This parameter is used for adjust stock.
	AdjustStock *bool `json:"adjust_stock,omitempty"`
	// If the value is 'true' and order exist in our cache, we will use order.info from cache to prepare shipment items.
	EnableCache *bool `json:"enable_cache,omitempty"`
	// Defines custom tracking link
	TrackingLink *string `json:"tracking_link,omitempty"`
	// Defines shipment's status
	IsShipped *bool `json:"is_shipped,omitempty"`
	// Disable or enable check process status. Please note that the response will be slower due to additional requests to the store.
	CheckProcessStatus *bool `json:"check_process_status,omitempty"`
	// Use the latest platform API version
	UseLatestApiVersion *bool `json:"use_latest_api_version,omitempty"`
}

// NewOrderShipmentAdd instantiates a new OrderShipmentAdd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderShipmentAdd() *OrderShipmentAdd {
	this := OrderShipmentAdd{}
	var sendNotifications bool = false
	this.SendNotifications = &sendNotifications
	var adjustStock bool = false
	this.AdjustStock = &adjustStock
	var enableCache bool = false
	this.EnableCache = &enableCache
	var isShipped bool = true
	this.IsShipped = &isShipped
	var checkProcessStatus bool = false
	this.CheckProcessStatus = &checkProcessStatus
	var useLatestApiVersion bool = false
	this.UseLatestApiVersion = &useLatestApiVersion
	return &this
}

// NewOrderShipmentAddWithDefaults instantiates a new OrderShipmentAdd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderShipmentAddWithDefaults() *OrderShipmentAdd {
	this := OrderShipmentAdd{}
	var sendNotifications bool = false
	this.SendNotifications = &sendNotifications
	var adjustStock bool = false
	this.AdjustStock = &adjustStock
	var enableCache bool = false
	this.EnableCache = &enableCache
	var isShipped bool = true
	this.IsShipped = &isShipped
	var checkProcessStatus bool = false
	this.CheckProcessStatus = &checkProcessStatus
	var useLatestApiVersion bool = false
	this.UseLatestApiVersion = &useLatestApiVersion
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OrderShipmentAdd) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipmentAdd) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OrderShipmentAdd) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *OrderShipmentAdd) SetOrderId(v string) {
	o.OrderId = &v
}

// GetStoreId returns the StoreId field value if set, zero value otherwise.
func (o *OrderShipmentAdd) GetStoreId() string {
	if o == nil || IsNil(o.StoreId) {
		var ret string
		return ret
	}
	return *o.StoreId
}

// GetStoreIdOk returns a tuple with the StoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipmentAdd) GetStoreIdOk() (*string, bool) {
	if o == nil || IsNil(o.StoreId) {
		return nil, false
	}
	return o.StoreId, true
}

// HasStoreId returns a boolean if a field has been set.
func (o *OrderShipmentAdd) HasStoreId() bool {
	if o != nil && !IsNil(o.StoreId) {
		return true
	}

	return false
}

// SetStoreId gets a reference to the given string and assigns it to the StoreId field.
func (o *OrderShipmentAdd) SetStoreId(v string) {
	o.StoreId = &v
}

// GetWarehouseId returns the WarehouseId field value if set, zero value otherwise.
func (o *OrderShipmentAdd) GetWarehouseId() string {
	if o == nil || IsNil(o.WarehouseId) {
		var ret string
		return ret
	}
	return *o.WarehouseId
}

// GetWarehouseIdOk returns a tuple with the WarehouseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipmentAdd) GetWarehouseIdOk() (*string, bool) {
	if o == nil || IsNil(o.WarehouseId) {
		return nil, false
	}
	return o.WarehouseId, true
}

// HasWarehouseId returns a boolean if a field has been set.
func (o *OrderShipmentAdd) HasWarehouseId() bool {
	if o != nil && !IsNil(o.WarehouseId) {
		return true
	}

	return false
}

// SetWarehouseId gets a reference to the given string and assigns it to the WarehouseId field.
func (o *OrderShipmentAdd) SetWarehouseId(v string) {
	o.WarehouseId = &v
}

// GetShipmentProvider returns the ShipmentProvider field value if set, zero value otherwise.
func (o *OrderShipmentAdd) GetShipmentProvider() string {
	if o == nil || IsNil(o.ShipmentProvider) {
		var ret string
		return ret
	}
	return *o.ShipmentProvider
}

// GetShipmentProviderOk returns a tuple with the ShipmentProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipmentAdd) GetShipmentProviderOk() (*string, bool) {
	if o == nil || IsNil(o.ShipmentProvider) {
		return nil, false
	}
	return o.ShipmentProvider, true
}

// HasShipmentProvider returns a boolean if a field has been set.
func (o *OrderShipmentAdd) HasShipmentProvider() bool {
	if o != nil && !IsNil(o.ShipmentProvider) {
		return true
	}

	return false
}

// SetShipmentProvider gets a reference to the given string and assigns it to the ShipmentProvider field.
func (o *OrderShipmentAdd) SetShipmentProvider(v string) {
	o.ShipmentProvider = &v
}

// GetShippingMethod returns the ShippingMethod field value if set, zero value otherwise.
func (o *OrderShipmentAdd) GetShippingMethod() string {
	if o == nil || IsNil(o.ShippingMethod) {
		var ret string
		return ret
	}
	return *o.ShippingMethod
}

// GetShippingMethodOk returns a tuple with the ShippingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipmentAdd) GetShippingMethodOk() (*string, bool) {
	if o == nil || IsNil(o.ShippingMethod) {
		return nil, false
	}
	return o.ShippingMethod, true
}

// HasShippingMethod returns a boolean if a field has been set.
func (o *OrderShipmentAdd) HasShippingMethod() bool {
	if o != nil && !IsNil(o.ShippingMethod) {
		return true
	}

	return false
}

// SetShippingMethod gets a reference to the given string and assigns it to the ShippingMethod field.
func (o *OrderShipmentAdd) SetShippingMethod(v string) {
	o.ShippingMethod = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *OrderShipmentAdd) GetItems() []OrderShipmentAddItemsInner {
	if o == nil || IsNil(o.Items) {
		var ret []OrderShipmentAddItemsInner
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipmentAdd) GetItemsOk() ([]OrderShipmentAddItemsInner, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *OrderShipmentAdd) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []OrderShipmentAddItemsInner and assigns it to the Items field.
func (o *OrderShipmentAdd) SetItems(v []OrderShipmentAddItemsInner) {
	o.Items = v
}

// GetSendNotifications returns the SendNotifications field value if set, zero value otherwise.
func (o *OrderShipmentAdd) GetSendNotifications() bool {
	if o == nil || IsNil(o.SendNotifications) {
		var ret bool
		return ret
	}
	return *o.SendNotifications
}

// GetSendNotificationsOk returns a tuple with the SendNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipmentAdd) GetSendNotificationsOk() (*bool, bool) {
	if o == nil || IsNil(o.SendNotifications) {
		return nil, false
	}
	return o.SendNotifications, true
}

// HasSendNotifications returns a boolean if a field has been set.
func (o *OrderShipmentAdd) HasSendNotifications() bool {
	if o != nil && !IsNil(o.SendNotifications) {
		return true
	}

	return false
}

// SetSendNotifications gets a reference to the given bool and assigns it to the SendNotifications field.
func (o *OrderShipmentAdd) SetSendNotifications(v bool) {
	o.SendNotifications = &v
}

// GetTrackingNumbers returns the TrackingNumbers field value if set, zero value otherwise.
func (o *OrderShipmentAdd) GetTrackingNumbers() []OrderShipmentAddTrackingNumbersInner {
	if o == nil || IsNil(o.TrackingNumbers) {
		var ret []OrderShipmentAddTrackingNumbersInner
		return ret
	}
	return o.TrackingNumbers
}

// GetTrackingNumbersOk returns a tuple with the TrackingNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipmentAdd) GetTrackingNumbersOk() ([]OrderShipmentAddTrackingNumbersInner, bool) {
	if o == nil || IsNil(o.TrackingNumbers) {
		return nil, false
	}
	return o.TrackingNumbers, true
}

// HasTrackingNumbers returns a boolean if a field has been set.
func (o *OrderShipmentAdd) HasTrackingNumbers() bool {
	if o != nil && !IsNil(o.TrackingNumbers) {
		return true
	}

	return false
}

// SetTrackingNumbers gets a reference to the given []OrderShipmentAddTrackingNumbersInner and assigns it to the TrackingNumbers field.
func (o *OrderShipmentAdd) SetTrackingNumbers(v []OrderShipmentAddTrackingNumbersInner) {
	o.TrackingNumbers = v
}

// GetAdjustStock returns the AdjustStock field value if set, zero value otherwise.
func (o *OrderShipmentAdd) GetAdjustStock() bool {
	if o == nil || IsNil(o.AdjustStock) {
		var ret bool
		return ret
	}
	return *o.AdjustStock
}

// GetAdjustStockOk returns a tuple with the AdjustStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipmentAdd) GetAdjustStockOk() (*bool, bool) {
	if o == nil || IsNil(o.AdjustStock) {
		return nil, false
	}
	return o.AdjustStock, true
}

// HasAdjustStock returns a boolean if a field has been set.
func (o *OrderShipmentAdd) HasAdjustStock() bool {
	if o != nil && !IsNil(o.AdjustStock) {
		return true
	}

	return false
}

// SetAdjustStock gets a reference to the given bool and assigns it to the AdjustStock field.
func (o *OrderShipmentAdd) SetAdjustStock(v bool) {
	o.AdjustStock = &v
}

// GetEnableCache returns the EnableCache field value if set, zero value otherwise.
func (o *OrderShipmentAdd) GetEnableCache() bool {
	if o == nil || IsNil(o.EnableCache) {
		var ret bool
		return ret
	}
	return *o.EnableCache
}

// GetEnableCacheOk returns a tuple with the EnableCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipmentAdd) GetEnableCacheOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableCache) {
		return nil, false
	}
	return o.EnableCache, true
}

// HasEnableCache returns a boolean if a field has been set.
func (o *OrderShipmentAdd) HasEnableCache() bool {
	if o != nil && !IsNil(o.EnableCache) {
		return true
	}

	return false
}

// SetEnableCache gets a reference to the given bool and assigns it to the EnableCache field.
func (o *OrderShipmentAdd) SetEnableCache(v bool) {
	o.EnableCache = &v
}

// GetTrackingLink returns the TrackingLink field value if set, zero value otherwise.
func (o *OrderShipmentAdd) GetTrackingLink() string {
	if o == nil || IsNil(o.TrackingLink) {
		var ret string
		return ret
	}
	return *o.TrackingLink
}

// GetTrackingLinkOk returns a tuple with the TrackingLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipmentAdd) GetTrackingLinkOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingLink) {
		return nil, false
	}
	return o.TrackingLink, true
}

// HasTrackingLink returns a boolean if a field has been set.
func (o *OrderShipmentAdd) HasTrackingLink() bool {
	if o != nil && !IsNil(o.TrackingLink) {
		return true
	}

	return false
}

// SetTrackingLink gets a reference to the given string and assigns it to the TrackingLink field.
func (o *OrderShipmentAdd) SetTrackingLink(v string) {
	o.TrackingLink = &v
}

// GetIsShipped returns the IsShipped field value if set, zero value otherwise.
func (o *OrderShipmentAdd) GetIsShipped() bool {
	if o == nil || IsNil(o.IsShipped) {
		var ret bool
		return ret
	}
	return *o.IsShipped
}

// GetIsShippedOk returns a tuple with the IsShipped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipmentAdd) GetIsShippedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsShipped) {
		return nil, false
	}
	return o.IsShipped, true
}

// HasIsShipped returns a boolean if a field has been set.
func (o *OrderShipmentAdd) HasIsShipped() bool {
	if o != nil && !IsNil(o.IsShipped) {
		return true
	}

	return false
}

// SetIsShipped gets a reference to the given bool and assigns it to the IsShipped field.
func (o *OrderShipmentAdd) SetIsShipped(v bool) {
	o.IsShipped = &v
}

// GetCheckProcessStatus returns the CheckProcessStatus field value if set, zero value otherwise.
func (o *OrderShipmentAdd) GetCheckProcessStatus() bool {
	if o == nil || IsNil(o.CheckProcessStatus) {
		var ret bool
		return ret
	}
	return *o.CheckProcessStatus
}

// GetCheckProcessStatusOk returns a tuple with the CheckProcessStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipmentAdd) GetCheckProcessStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.CheckProcessStatus) {
		return nil, false
	}
	return o.CheckProcessStatus, true
}

// HasCheckProcessStatus returns a boolean if a field has been set.
func (o *OrderShipmentAdd) HasCheckProcessStatus() bool {
	if o != nil && !IsNil(o.CheckProcessStatus) {
		return true
	}

	return false
}

// SetCheckProcessStatus gets a reference to the given bool and assigns it to the CheckProcessStatus field.
func (o *OrderShipmentAdd) SetCheckProcessStatus(v bool) {
	o.CheckProcessStatus = &v
}

// GetUseLatestApiVersion returns the UseLatestApiVersion field value if set, zero value otherwise.
func (o *OrderShipmentAdd) GetUseLatestApiVersion() bool {
	if o == nil || IsNil(o.UseLatestApiVersion) {
		var ret bool
		return ret
	}
	return *o.UseLatestApiVersion
}

// GetUseLatestApiVersionOk returns a tuple with the UseLatestApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderShipmentAdd) GetUseLatestApiVersionOk() (*bool, bool) {
	if o == nil || IsNil(o.UseLatestApiVersion) {
		return nil, false
	}
	return o.UseLatestApiVersion, true
}

// HasUseLatestApiVersion returns a boolean if a field has been set.
func (o *OrderShipmentAdd) HasUseLatestApiVersion() bool {
	if o != nil && !IsNil(o.UseLatestApiVersion) {
		return true
	}

	return false
}

// SetUseLatestApiVersion gets a reference to the given bool and assigns it to the UseLatestApiVersion field.
func (o *OrderShipmentAdd) SetUseLatestApiVersion(v bool) {
	o.UseLatestApiVersion = &v
}

func (o OrderShipmentAdd) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderShipmentAdd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.StoreId) {
		toSerialize["store_id"] = o.StoreId
	}
	if !IsNil(o.WarehouseId) {
		toSerialize["warehouse_id"] = o.WarehouseId
	}
	if !IsNil(o.ShipmentProvider) {
		toSerialize["shipment_provider"] = o.ShipmentProvider
	}
	if !IsNil(o.ShippingMethod) {
		toSerialize["shipping_method"] = o.ShippingMethod
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.SendNotifications) {
		toSerialize["send_notifications"] = o.SendNotifications
	}
	if !IsNil(o.TrackingNumbers) {
		toSerialize["tracking_numbers"] = o.TrackingNumbers
	}
	if !IsNil(o.AdjustStock) {
		toSerialize["adjust_stock"] = o.AdjustStock
	}
	if !IsNil(o.EnableCache) {
		toSerialize["enable_cache"] = o.EnableCache
	}
	if !IsNil(o.TrackingLink) {
		toSerialize["tracking_link"] = o.TrackingLink
	}
	if !IsNil(o.IsShipped) {
		toSerialize["is_shipped"] = o.IsShipped
	}
	if !IsNil(o.CheckProcessStatus) {
		toSerialize["check_process_status"] = o.CheckProcessStatus
	}
	if !IsNil(o.UseLatestApiVersion) {
		toSerialize["use_latest_api_version"] = o.UseLatestApiVersion
	}
	return toSerialize, nil
}

type NullableOrderShipmentAdd struct {
	value *OrderShipmentAdd
	isSet bool
}

func (v NullableOrderShipmentAdd) Get() *OrderShipmentAdd {
	return v.value
}

func (v *NullableOrderShipmentAdd) Set(val *OrderShipmentAdd) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderShipmentAdd) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderShipmentAdd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderShipmentAdd(val *OrderShipmentAdd) *NullableOrderShipmentAdd {
	return &NullableOrderShipmentAdd{value: val, isSet: true}
}

func (v NullableOrderShipmentAdd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderShipmentAdd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


