from django.shortcuts import render, get_object_or_404
from django.http import HttpResponse, Http404, HttpResponseRedirect
from django.urls import reverse
from django.views import generic

from .models import Question, Choice

# Create your views here.
# Below code is commented to show clear migration to Generic views and
# tutorial progress.
# def index(request):
#     latest_questions = Question.objects.order_by("-pub_date")[:5]
#     context = {"latest_questions" : latest_questions}
#     return render(request, "polls/index.html", context)

# def detail(request, question_id):
#     question = get_object_or_404(Question, pk=question_id)
#     context = {"question" : question}
#     return render(request, "polls/detail.html", context)
#
# def results(request, question_id):
#     question = get_object_or_404(Question, pk=question_id)
#     return render(request, "polls/results.html", {"question": question})

class IndexView(generic.ListView):
    # Variables "template_name", "context_object_name" and method "get_queryset"
    # are a part of the ListView Syntax. By using these variables we are specifying which template
    # to use and what should be the object name of the context.
    # get_queryset will get the list of items to display.
    template_name = "polls/index.html"
    context_object_name = "latest_questions"

    def get_queryset(self):
        return Question.objects.order_by("-pub_date")[:5]

class DetailView(generic.DetailView):
    model = Question
    template_name = "polls/detail.html"

class ResultsView(generic.DetailView):
    model = Question
    template_name = "polls/results.html"

# This view cannot be made generic.
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
        # Using reverse function in order to avoid hardcoding of redirecting url.
        return HttpResponseRedirect(reverse("polls:results", args=(question.id,)))

# Made this just for experimenting.
def sample_app_view(request):
    return HttpResponse("This view is from another path.")
