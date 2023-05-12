from typing import Any, Dict
from django.forms import ValidationError
from django.shortcuts import get_object_or_404, render
from django.contrib.auth import authenticate, login as dj_login, logout as dj_logout
from django.db import transaction
from django.core.exceptions import SuspiciousOperation
from pydantic import BaseModel

# Create your views here.

from django.http import Http404, HttpResponse

from knt_backend.models import Product, Receipt, Transaction, User
from knt_backend.utils import load_model, redirect_to_next


def index(request):
    return HttpResponse("Hello, world. You're at the knt index.")


def calculate_cart_total(request) -> int:
    cart = request.session.get("cart", {})
    total = 0
    for product_id, quantity in cart.items():
        product = Product.objects.get(id=product_id)
        total += product.price * quantity

    return total


def products_table(request):
    products = Product.objects.all()
    template = "products.html"

    return render(
        request,
        template,
        {
            "products": products,
            "cart": {int(k): v for k, v in request.session.get("cart", {}).items()},
            "message": request.GET.get("message", ""),
            "total": calculate_cart_total(request),
        },
    )


def login(request):
    username = request.POST.get("username")
    password = request.POST.get("password")

    user = authenticate(username=username, password=password)

    if user is None:
        return redirect_to_next(request, message="Invalid login")
    else:
        dj_login(request, user)

    return redirect_to_next(request, message="Logged in!")


def logout(request):
    dj_logout(request)
    return redirect_to_next(request, message="Logged out!")


def product(request, product_id):
    product = get_object_or_404(Product, pk=product_id)

    return HttpResponse(
        "You're looking at product %s. name: %s, price: %s "
        % (product_id, product.name, product.price)
    )


def add_to_cart(request, product_id: int):
    # make sure product exists

    _ = get_object_or_404(Product, pk=product_id)

    # adds the product to the user's session
    # if the product is already in the user's session, increment the quantity
    # if the product is not in the user's session, add it with quantity 1

    cart: Dict[str, int] = request.session.get("cart", {})
    cart.setdefault(str(product_id), 0)
    cart[str(product_id)] += 1
    request.session["cart"] = cart

    return redirect_to_next(request)


def reset_cart(request):
    request.session["cart"] = {}
    return redirect_to_next(request)


def checkout(request):
    # user purchases multiple products

    cart = request.session.get("cart", {})

    if not cart:
        raise ValidationError("Cart is empty")

    if not request.user.is_authenticated:
        raise ValidationError("User is not authenticated")

    cart = {int(k): v for k, v in cart.items()}

    products = Product.objects.filter(id__in=cart.keys())

    total_cost = calculate_cart_total(request)

    # check if user has enough money
    user: User = User.objects.get(id=request.user.id)

    if user.balance < total_cost:
        raise SuspiciousOperation("User does not have enough money")

    user_balance_before = user.balance
    # update user balance
    user.balance -= total_cost

    user_balance_after = user.balance
    user.save()

    receipt_text = "\n".join(
        [
            f"{product.name} x{cart[product.id]}: {product.price * cart[product.id]}"
            for product in products
        ]
    )
    receipt = Receipt(data=receipt_text)
    receipt.save(force_insert=True)

    # create transaction
    new_transaction = Transaction.objects.create(
        user_id=user.id,
        starting_balance=user_balance_before,
        delta_balance=total_cost,
        final_balance=user_balance_after,
        receipt_id=receipt.id,
        ref="idk lol",
    )

    new_transaction.save()

    # clear cart
    request.session["cart"] = {}

    # logout user
    logout(request)

    return redirect_to_next(request, message="Checkout successful!")
