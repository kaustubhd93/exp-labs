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
# Below code is commented to show clear migration to Generic views and
# tutorial progress.
# urlpatterns = [
#     path("", views.index, name="index"),
#     path("another/", views.sample_app_view, name="sample_app_view"),
#     path("<int:question_id>/", views.detail, name="detail"),
#     path("<int:question_id>/results", views.results, name="results"),
#     path("<int:question_id>/votes", views.votes, name="votes"),
# ]

urlpatterns = [
    path("", views.IndexView.as_view(), name="index"),
    path("another/", views.sample_app_view, name="sample_app_view"),
    # The DetailView generic view expects the primary key value captured from the URL 
    # to be called "pk", so we’ve changed question_id to pk for the generic views.
    path("<int:pk>/", views.DetailView.as_view(), name="detail"),
    path("<int:pk>/results", views.ResultsView.as_view(), name="results"),
    path("<int:question_id>/votes", views.votes, name="votes"),
]
