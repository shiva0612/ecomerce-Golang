# Ecomerce

```jsx
clone the repo
go mod tidy
go run *.go (or) go build . && ./main
```

```jsx
[GIN-debug] POST   /signup                   --> shiva/controller/authenticate.Signup (1 handlers)
[GIN-debug] POST   /login                    --> shiva/controller/authenticate.Login (1 handlers)
[GIN-debug] POST   /refresh-token            --> shiva/controller/authenticate.RefreshToken (2 handlers)
[GIN-debug] GET    /logout                   --> shiva/controller/authenticate.Logout (2 handlers)
[GIN-debug] POST   /deleteme                 --> shiva/controller/authenticate.Delete (2 handlers)
[GIN-debug] GET    /profile                  --> shiva/controller/user.Profile (2 handlers)
[GIN-debug] GET    /orders                   --> shiva/controller/user.ViewOrders (2 handlers)
[GIN-debug] POST   /address/add              --> shiva/controller/user.AddAddress (2 handlers)
[GIN-debug] GET    /address/view             --> shiva/controller/user.GetAddress (2 handlers)
[GIN-debug] POST   /admin/addproducts        --> shiva/controller/ecom.AddProducts (2 handlers)
[GIN-debug] GET    /                         --> shiva/controller/ecom.ProductsList (2 handlers)
[GIN-debug] POST   /cart/add                 --> shiva/controller/ecom.AddProductToCart (2 handlers)
[GIN-debug] POST   /cart/remove              --> shiva/controller/ecom.RemoveProductFromCart (2 handlers)
[GIN-debug] POST   /cart/view                --> shiva/controller/ecom.ViewCart (2 handlers)
[GIN-debug] GET    /cart/checkout            --> shiva/controller/ecom.CheckoutCart (2 handlers)
[GIN-debug] POST   /cart/placeorder          --> shiva/controller/ecom.PlaceOrder (2 handlers)
```

[ecomerce.gif](Ecomerce%2044960d0bed7846da9881204a3fddb4cb/ecomerce.gif)