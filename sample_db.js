let users = {
  fullName: FULLNAME,
  email: EMAIL,
  password: PASSWORD,
  token: TOKEN,
};
let categories = [
  {
    name: "electronics",
    parent: "",
    category: "/electronics",
  },
  {
    name: "embedded",
    parent: "electronics",
    category: "/electronics/embedded",
  },
];

let order = {
  orderID: ORDERID,
  userID: userID,
  subTOTAL: SUBTOTAL,
  productsOrder: {
    0: {
      productID: PRODUCTID,
      color: COLOR,
      size: SIZE,
      quantity: QUANTITY,
    },
    1: {
      productID: PRODUCTID,
      color: COLOR,
      size: SIZE,
      quantity: QUANTITY,
    },
  },
};
	// productObject = &models.Product{
	// 	Name:           "Test product",
	// 	AddDate:        time.Now(),
	// 	Categories:     []string{"cate1", "cate2"},
	// 	ParentCategory: "men",
	// 	Images:         []string{"eK35LtEU5Uo6VutAs3iC4P.jpg", "ytM5EvDn3dF4H3qCpYL66Z.jpg"},
	// 	Brand:          "Zara",
	// 	Price:          99,
	// 	Sizes:          []string{"X", "XL"},
	// 	Colors:         []string{"Red", "Green"},
	// 	Quantity:       9,
	// 	Sold:           2,
	// 	Description:    "Desc",
	// }

  //5/4/2021 
  //Admin Login
  //Admin Add product
  //User profile
  //List product
  //Product details
  //Cart
  //Orders
  //Admin orders
  //Admin Edit/Remove product