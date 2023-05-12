from django.urls import path

from . import views

app_name = "knt"
urlpatterns = [
    path("", views.index, name="index"),
    path("products/", views.products_table, name="products"),
    path("products/<int:product_id>/", views.product, name="product"),
    path(
        "products/<int:product_id>/add_to_cart/", views.add_to_cart, name="add_to_cart"
    ),
    path("products/reset_cart", views.reset_cart, name="reset_cart"),
    path("products/checkout/", views.checkout, name="checkout"),
    path("login/", views.login, name="login"),
]
