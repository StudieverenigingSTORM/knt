import traceback
from typing import Dict, Type, TypeVar
from django.forms import ValidationError
from django.http import HttpResponseRedirect
import urllib.parse

from pydantic import BaseModel


T = TypeVar("T", bound=BaseModel)


def load_model(request, model: Type[T]) -> T:
    body: str = request.body.decode("utf-8")
    try:
        return model.parse_raw(body)
    except Exception as e:
        traceback.print_exc()
        raise ValidationError("Invalid request body")


def redirect_to_next(request, **kwargs) -> HttpResponseRedirect:
    next = request.GET.get("next", "/")
    if next[-1] != "?":
        next += "?"

    for k, v in kwargs.items():
        if next[-1] != "?" and next[-1] != "&":
            next += "&"
        k = urllib.parse.quote_plus(k)
        v = urllib.parse.quote_plus(v)
        next += f"{k}={v}"

    return HttpResponseRedirect(next)
