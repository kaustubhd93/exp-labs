from django.shortcuts import render
from django.http import HttpResponse

# Create your views here.
def index(request):
    return HttpResponse("Hello, this is the first view.")

def sample_app_view(request):
    return HttpResponse("This view is from another path.")
