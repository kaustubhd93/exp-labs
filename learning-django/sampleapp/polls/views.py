from django.shortcuts import render, get_object_or_404
from django.http import HttpResponse, Http404, HttpResponseRedirect
from django.urls import reverse

from .models import Question

# Create your views here.
def index(request):
    latest_questions = Question.objects.order_by("-pub_date")[:5]
    context = {"latest_questions" : latest_questions}
    return render(request, "polls/index.html", context)

# Made this just for experimenting.
def sample_app_view(request):
    return HttpResponse("This view is from another path.")

def detail(request, question_id):
    question = get_object_or_404(Question, pk=question_id)
    context = {"question" : question}
    return render(request, "polls/detail.html", context)

def results(request, question_id):
    question = get_object_or_404(Question, pk=question_id)
    return render(request, "polls/results.html", {"question": question})


def votes(request, question_id):
    question = get_object_or_404(Question, pk=question_id)
    try:
        selected_choice = question.choice_set.get(pk=request.POST["choice"])
    except (KeyError, Choice.DoesNotExist):
        context = {"question": question, "error_message": "You didn't select a choice."}
        return render(request, "polls/detail.html", context)
    else:
        selected_choice.votes += 1
        selected_choice.save()
        return HttpResponseRedirect(reverse("polls:results", args=(question.id,)))
