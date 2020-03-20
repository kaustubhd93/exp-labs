from django.contrib import admin
from .models import Question,Choice

class ChoicesOfQuestion(admin.TabularInline):
    model = Choice
    # An empty field with option to add choice will be present.
    extra = 1

class QuestionAdmin(admin.ModelAdmin):
    list_display = ("question_text", "pub_date")
    fieldsets = [
        (None, {"fields": ["question_text"]}),
        ("Date Information", {"fields": ["pub_date"]}),
    ]
    # Below is an InlineModelAdmin object which
    inlines = [ChoicesOfQuestion]

admin.site.register(Question, QuestionAdmin)
#admin.site.register(Choice)
