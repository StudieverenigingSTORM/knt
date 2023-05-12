from typing import Any, Dict
from django.forms import ValidationError
from django.shortcuts import get_object_or_404, render
from django.contrib.auth import authenticate
from django.db import transaction
from django.core.exceptions import SuspiciousOperation
from pydantic import BaseModel

# Create your views here.

from django.http import Http404, HttpResponse

from knt_backend.models import Product, Receipt, Transaction, User
from knt_backend.utils import load_model, redirect_to_next


def index(request):
    return HttpResponse("Hello, world. You're at the knt index.")


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
        },
    )


def login(request):
    user_id = request.POST.get("user_id")
    password = request.POST.get("password")

    user = authenticate(username=user_id, password=password)

    if user is None:
        return redirect_to_next(request, message="Invalid login")
    else:
        request.session["user_id"] = user.pk

    return redirect_to_next(request, message="Logged in!")


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


@transaction.atomic
def checkout(request):
    # user purchases multiple products

    cart = request.session.get("cart", {})

    if not cart:
        raise ValidationError("Cart is empty")

    if not request.user.is_authenticated:
        raise ValidationError("User is not authenticated")

    cart = {int(k): v for k, v in cart.items()}

    products = Product.objects.filter(id__in=cart.keys())

    total_cost = 0
    for product in products:
        total_cost += product.price * cart[product.id]

    # check if user has enough money
    user: User = request.user

    if user.balance < total_cost:
        raise SuspiciousOperation("User does not have enough money")

    user_balance_before = user.balance
    # update user balance
    user.balance -= total_cost
    user.save()

    user_balance_after = user.balance

    receipt_text = "\n".join(
        [
            f"{product.name} x{cart[product.id]}: {product.price * cart[product.id]}"
            for product in products
        ]
    )

    receipt = Receipt.objects.create(data=receipt_text)

    # create transaction
    Transaction.objects.create(
        user_id=user.id,
        starting_balance=user_balance_before,
        delta_balance=total_cost,
        final_balance=user_balance_after,
        receipt_id=receipt.id,
        ref="idk lol",
    )

    # clear cart
    request.session["cart"] = {}

    # logout user
    del request.session["user_id"]

    return HttpResponse("You bought some stuff!")
