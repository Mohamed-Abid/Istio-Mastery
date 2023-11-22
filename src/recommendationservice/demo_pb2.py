# source: demo.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\ndemo.proto\x12\x0bhipstershop\"0\n\x08\x43\x61rtItem\x12\x12\n\nproduct_id\x18\x01 \x01(\t\x12\x10\n\x08quantity\x18\x02 \x01(\x05\"F\n\x0e\x41\x64\x64ItemRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12#\n\x04item\x18\x02 \x01(\x0b\x32\x15.hipstershop.CartItem\"#\n\x10\x45mptyCartRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\"!\n\x0eGetCartRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\"=\n\x04\x43\x61rt\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12$\n\x05items\x18\x02 \x03(\x0b\x32\x15.hipstershop.CartItem\"\x07\n\x05\x45mpty\"B\n\x1aListRecommendationsRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x13\n\x0bproduct_ids\x18\x02 \x03(\t\"2\n\x1bListRecommendationsResponse\x12\x13\n\x0bproduct_ids\x18\x01 \x03(\t\"\x84\x01\n\x07Product\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\x0f\n\x07picture\x18\x04 \x01(\t\x12%\n\tprice_usd\x18\x05 \x01(\x0b\x32\x12.hipstershop.Money\x12\x12\n\ncategories\x18\x06 \x03(\t\">\n\x14ListProductsResponse\x12&\n\x08products\x18\x01 \x03(\x0b\x32\x14.hipstershop.Product\"\x1f\n\x11GetProductRequest\x12\n\n\x02id\x18\x01 \x01(\t\"&\n\x15SearchProductsRequest\x12\r\n\x05query\x18\x01 \x01(\t\"?\n\x16SearchProductsResponse\x12%\n\x07results\x18\x01 \x03(\x0b\x32\x14.hipstershop.Product\"^\n\x0fGetQuoteRequest\x12%\n\x07\x61\x64\x64ress\x18\x01 \x01(\x0b\x32\x14.hipstershop.Address\x12$\n\x05items\x18\x02 \x03(\x0b\x32\x15.hipstershop.CartItem\"8\n\x10GetQuoteResponse\x12$\n\x08\x63ost_usd\x18\x01 \x01(\x0b\x32\x12.hipstershop.Money\"_\n\x10ShipOrderRequest\x12%\n\x07\x61\x64\x64ress\x18\x01 \x01(\x0b\x32\x14.hipstershop.Address\x12$\n\x05items\x18\x02 \x03(\x0b\x32\x15.hipstershop.CartItem\"(\n\x11ShipOrderResponse\x12\x13\n\x0btracking_id\x18\x01 \x01(\t\"a\n\x07\x41\x64\x64ress\x12\x16\n\x0estreet_address\x18\x01 \x01(\t\x12\x0c\n\x04\x63ity\x18\x02 \x01(\t\x12\r\n\x05state\x18\x03 \x01(\t\x12\x0f\n\x07\x63ountry\x18\x04 \x01(\t\x12\x10\n\x08zip_code\x18\x05 \x01(\x05\"<\n\x05Money\x12\x15\n\rcurrency_code\x18\x01 \x01(\t\x12\r\n\x05units\x18\x02 \x01(\x03\x12\r\n\x05nanos\x18\x03 \x01(\x05\"8\n\x1eGetSupportedCurrenciesResponse\x12\x16\n\x0e\x63urrency_codes\x18\x01 \x03(\t\"N\n\x19\x43urrencyConversionRequest\x12 \n\x04\x66rom\x18\x01 \x01(\x0b\x32\x12.hipstershop.Money\x12\x0f\n\x07to_code\x18\x02 \x01(\t\"\x90\x01\n\x0e\x43reditCardInfo\x12\x1a\n\x12\x63redit_card_number\x18\x01 \x01(\t\x12\x17\n\x0f\x63redit_card_cvv\x18\x02 \x01(\x05\x12#\n\x1b\x63redit_card_expiration_year\x18\x03 \x01(\x05\x12$\n\x1c\x63redit_card_expiration_month\x18\x04 \x01(\x05\"e\n\rChargeRequest\x12\"\n\x06\x61mount\x18\x01 \x01(\x0b\x32\x12.hipstershop.Money\x12\x30\n\x0b\x63redit_card\x18\x02 \x01(\x0b\x32\x1b.hipstershop.CreditCardInfo\"(\n\x0e\x43hargeResponse\x12\x16\n\x0etransaction_id\x18\x01 \x01(\t\"R\n\tOrderItem\x12#\n\x04item\x18\x01 \x01(\x0b\x32\x15.hipstershop.CartItem\x12 \n\x04\x63ost\x18\x02 \x01(\x0b\x32\x12.hipstershop.Money\"\xbf\x01\n\x0bOrderResult\x12\x10\n\x08order_id\x18\x01 \x01(\t\x12\x1c\n\x14shipping_tracking_id\x18\x02 \x01(\t\x12)\n\rshipping_cost\x18\x03 \x01(\x0b\x32\x12.hipstershop.Money\x12.\n\x10shipping_address\x18\x04 \x01(\x0b\x32\x14.hipstershop.Address\x12%\n\x05items\x18\x05 \x03(\x0b\x32\x16.hipstershop.OrderItem\"V\n\x1cSendOrderConfirmationRequest\x12\r\n\x05\x65mail\x18\x01 \x01(\t\x12\'\n\x05order\x18\x02 \x01(\x0b\x32\x18.hipstershop.OrderResult\"\xa3\x01\n\x11PlaceOrderRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x15\n\ruser_currency\x18\x02 \x01(\t\x12%\n\x07\x61\x64\x64ress\x18\x03 \x01(\x0b\x32\x14.hipstershop.Address\x12\r\n\x05\x65mail\x18\x05 \x01(\t\x12\x30\n\x0b\x63redit_card\x18\x06 \x01(\x0b\x32\x1b.hipstershop.CreditCardInfo\"=\n\x12PlaceOrderResponse\x12\'\n\x05order\x18\x01 \x01(\x0b\x32\x18.hipstershop.OrderResult\"!\n\tAdRequest\x12\x14\n\x0c\x63ontext_keys\x18\x01 \x03(\t\"*\n\nAdResponse\x12\x1c\n\x03\x61\x64s\x18\x01 \x03(\x0b\x32\x0f.hipstershop.Ad\"(\n\x02\x41\x64\x12\x14\n\x0credirect_url\x18\x01 \x01(\t\x12\x0c\n\x04text\x18\x02 \x01(\t2\xca\x01\n\x0b\x43\x61rtService\x12<\n\x07\x41\x64\x64Item\x12\x1b.hipstershop.AddItemRequest\x1a\x12.hipstershop.Empty\"\x00\x12;\n\x07GetCart\x12\x1b.hipstershop.GetCartRequest\x1a\x11.hipstershop.Cart\"\x00\x12@\n\tEmptyCart\x12\x1d.hipstershop.EmptyCartRequest\x1a\x12.hipstershop.Empty\"\x00\x32\x83\x01\n\x15RecommendationService\x12j\n\x13ListRecommendations\x12\'.hipstershop.ListRecommendationsRequest\x1a(.hipstershop.ListRecommendationsResponse\"\x00\x32\x83\x02\n\x15ProductCatalogService\x12G\n\x0cListProducts\x12\x12.hipstershop.Empty\x1a!.hipstershop.ListProductsResponse\"\x00\x12\x44\n\nGetProduct\x12\x1e.hipstershop.GetProductRequest\x1a\x14.hipstershop.Product\"\x00\x12[\n\x0eSearchProducts\x12\".hipstershop.SearchProductsRequest\x1a#.hipstershop.SearchProductsResponse\"\x00\x32\xaa\x01\n\x0fShippingService\x12I\n\x08GetQuote\x12\x1c.hipstershop.GetQuoteRequest\x1a\x1d.hipstershop.GetQuoteResponse\"\x00\x12L\n\tShipOrder\x12\x1d.hipstershop.ShipOrderRequest\x1a\x1e.hipstershop.ShipOrderResponse\"\x00\x32\xb7\x01\n\x0f\x43urrencyService\x12[\n\x16GetSupportedCurrencies\x12\x12.hipstershop.Empty\x1a+.hipstershop.GetSupportedCurrenciesResponse\"\x00\x12G\n\x07\x43onvert\x12&.hipstershop.CurrencyConversionRequest\x1a\x12.hipstershop.Money\"\x00\x32U\n\x0ePaymentService\x12\x43\n\x06\x43harge\x12\x1a.hipstershop.ChargeRequest\x1a\x1b.hipstershop.ChargeResponse\"\x00\x32h\n\x0c\x45mailService\x12X\n\x15SendOrderConfirmation\x12).hipstershop.SendOrderConfirmationRequest\x1a\x12.hipstershop.Empty\"\x00\x32\x62\n\x0f\x43heckoutService\x12O\n\nPlaceOrder\x12\x1e.hipstershop.PlaceOrderRequest\x1a\x1f.hipstershop.PlaceOrderResponse\"\x00\x32H\n\tAdService\x12;\n\x06GetAds\x12\x16.hipstershop.AdRequest\x1a\x17.hipstershop.AdResponse\"\x00\x62\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'demo_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  _CARTITEM._serialized_start=27
  _CARTITEM._serialized_end=75
  _ADDITEMREQUEST._serialized_start=77
  _ADDITEMREQUEST._serialized_end=147
  _EMPTYCARTREQUEST._serialized_start=149
  _EMPTYCARTREQUEST._serialized_end=184
  _GETCARTREQUEST._serialized_start=186
  _GETCARTREQUEST._serialized_end=219
  _CART._serialized_start=221
  _CART._serialized_end=282
  _EMPTY._serialized_start=284
  _EMPTY._serialized_end=291
  _LISTRECOMMENDATIONSREQUEST._serialized_start=293
  _LISTRECOMMENDATIONSREQUEST._serialized_end=359
  _LISTRECOMMENDATIONSRESPONSE._serialized_start=361
  _LISTRECOMMENDATIONSRESPONSE._serialized_end=411
  _PRODUCT._serialized_start=414
  _PRODUCT._serialized_end=546
  _LISTPRODUCTSRESPONSE._serialized_start=548
  _LISTPRODUCTSRESPONSE._serialized_end=610
  _GETPRODUCTREQUEST._serialized_start=612
  _GETPRODUCTREQUEST._serialized_end=643
  _SEARCHPRODUCTSREQUEST._serialized_start=645
  _SEARCHPRODUCTSREQUEST._serialized_end=683
  _SEARCHPRODUCTSRESPONSE._serialized_start=685
  _SEARCHPRODUCTSRESPONSE._serialized_end=748
  _GETQUOTEREQUEST._serialized_start=750
  _GETQUOTEREQUEST._serialized_end=844
  _GETQUOTERESPONSE._serialized_start=846
  _GETQUOTERESPONSE._serialized_end=902
  _SHIPORDERREQUEST._serialized_start=904
  _SHIPORDERREQUEST._serialized_end=999
  _SHIPORDERRESPONSE._serialized_start=1001
  _SHIPORDERRESPONSE._serialized_end=1041
  _ADDRESS._serialized_start=1043
  _ADDRESS._serialized_end=1140
  _MONEY._serialized_start=1142
  _MONEY._serialized_end=1202
  _GETSUPPORTEDCURRENCIESRESPONSE._serialized_start=1204
  _GETSUPPORTEDCURRENCIESRESPONSE._serialized_end=1260
  _CURRENCYCONVERSIONREQUEST._serialized_start=1262
  _CURRENCYCONVERSIONREQUEST._serialized_end=1340
  _CREDITCARDINFO._serialized_start=1343
  _CREDITCARDINFO._serialized_end=1487
  _CHARGEREQUEST._serialized_start=1489
  _CHARGEREQUEST._serialized_end=1590
  _CHARGERESPONSE._serialized_start=1592
  _CHARGERESPONSE._serialized_end=1632
  _ORDERITEM._serialized_start=1634
  _ORDERITEM._serialized_end=1716
  _ORDERRESULT._serialized_start=1719
  _ORDERRESULT._serialized_end=1910
  _SENDORDERCONFIRMATIONREQUEST._serialized_start=1912
  _SENDORDERCONFIRMATIONREQUEST._serialized_end=1998
  _PLACEORDERREQUEST._serialized_start=2001
  _PLACEORDERREQUEST._serialized_end=2164
  _PLACEORDERRESPONSE._serialized_start=2166
  _PLACEORDERRESPONSE._serialized_end=2227
  _ADREQUEST._serialized_start=2229
  _ADREQUEST._serialized_end=2262
  _ADRESPONSE._serialized_start=2264
  _ADRESPONSE._serialized_end=2306
  _AD._serialized_start=2308
  _AD._serialized_end=2348
  _CARTSERVICE._serialized_start=2351
  _CARTSERVICE._serialized_end=2553
  _RECOMMENDATIONSERVICE._serialized_start=2556
  _RECOMMENDATIONSERVICE._serialized_end=2687
  _PRODUCTCATALOGSERVICE._serialized_start=2690
  _PRODUCTCATALOGSERVICE._serialized_end=2949
  _SHIPPINGSERVICE._serialized_start=2952
  _SHIPPINGSERVICE._serialized_end=3122
  _CURRENCYSERVICE._serialized_start=3125
  _CURRENCYSERVICE._serialized_end=3308
  _PAYMENTSERVICE._serialized_start=3310
  _PAYMENTSERVICE._serialized_end=3395
  _EMAILSERVICE._serialized_start=3397
  _EMAILSERVICE._serialized_end=3501
  _CHECKOUTSERVICE._serialized_start=3503
  _CHECKOUTSERVICE._serialized_end=3601
  _ADSERVICE._serialized_start=3603
  _ADSERVICE._serialized_end=3675
# @@protoc_insertion_point(module_scope)
