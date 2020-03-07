"""
Add new urls in the below format
path("route", views.func_name, name="func_name", kwargs)
kwargs and name is optional.
"""
from django.urls import path
from . import views

urlpatterns = [
    path("", views.index, name="index"),
    path("another/", views.sample_app_view, name="sample_app_view"),
]
