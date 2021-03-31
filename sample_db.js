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
