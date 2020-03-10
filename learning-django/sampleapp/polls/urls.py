"""
Add new urls in the below format
path("route", views.func_name, name="func_name", kwargs)
kwargs and name is optional.
"""
from django.urls import path
from . import views

# This is the application namespace.
# This enables us to refer app specific templates in url blocks.
app_name = "polls"
urlpatterns = [
    path("", views.index, name="index"),
    path("another/", views.sample_app_view, name="sample_app_view"),
    path("<int:question_id>/", views.detail, name="detail"),
    path("<int:question_id>/results", views.results, name="results"),
    path("<int:question_id>/votes", views.votes, name="votes"),
]
