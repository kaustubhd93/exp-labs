from django.shortcuts import render, get_object_or_404
from django.http import HttpResponse, Http404

from .models import Question

# Create your views here.
def index(request):
    latest_questions = Question.objects.order_by("-pub_date")[:5]
    #latest_questions_str = ", ".join([question.question_text for question in latest_questions])
    context = {"latest_questions" : latest_questions}
    return render(request, "polls/index.html", context)

def sample_app_view(request):
    return HttpResponse("This view is from another path.")

def detail(request, question_id):
    # try:
    #     question = Question.objects.get(pk=question_id)
    # except Question.DoesNotExist:
    #     raise Http404("Question does not exist.")
    question = get_object_or_404(Question, pk=question_id)
    context = {"question" : question}
    return render(request, "polls/detail.html", context)

def results(request, question_id):
    return HttpResponse("You are looking at the results of question {}".format(question_id))

def votes(request, question_id):
    return HttpResponse("You are voting on question {}.".format(question_id))
