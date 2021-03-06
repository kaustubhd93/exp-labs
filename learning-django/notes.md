# Directory structure of admin control/project/site
```
└── sampleapp               --just a container.
    ├── db.sqlite3
    ├── manage.py           --command line utility to interact with Django project.
    └── sampleapp
        ├── __init__.py
        ├── settings.py     --configuration of django project.
        ├── urls.py         --URL declarations for django project.
        └── wsgi.py         
```

# "include" in urls.py of project settings

The include() function allows referencing other URLconfs. Whenever Django encounters include(), it chops off whatever part of the URL matched up to that point and then sends the remaining string to the included URLconf for further processing.

# How to add more APIs ?

Create a view in views.py, reference and map the view in urls.py of the app. The main URLconf of the app has already been included in the project urls.py so no need to change anything there.

# How to map views in urls.py ?

Add new urls in the below format in the list urlpatterns.  
```python
path("route", views.func_name, name="func_name", kwargs)
```
kwargs and name is optional.

# Installed APPS defaults in main setting.py

- django.contrib.admin             --Admin site.  
- django.contrib.auth              --Authentication system
- django.contrib.contenttypes      --A framework for content types
- django.contrib.sessions          --A session framework        
- django.contrib.messages          --A messagin framework
- django.contrib.staticfiles       --A framework for managing static files.

# models.py

- It is the blueprint of the tables in your database.

# migrations in django

- Migrations are how Django stores changes to your models(and thus your database schema)-they're just files on disk.
- The "migrate" command takes all migrations that haven't been applied(Django tracks every migration by using a special table in the database called django_migrations) and runs them against the database essentially, synchronizing the changes you made to your models with the schema in the database.
- Migrations can let us change models over time, as we develop our project in real time without the need to delete database or tables and make new ones -- database can be upgraded live without any loss of data.

# 3 steps to modify models

- Change code in models.py
- Run `python manage.py makemigrations <app_name>` (Creates migrations)
- Run `python manage.py migrate` (Applies migration to database)

# How do views and urls work ?

- When one requests http://mysite.com/4 (4 here being the question id) Django loads the mysite.urls module because it is pointed by the ROOT_URLCONF settings. It finds the variable named urlpatterns and traverses the patterns in order. After finding relevant match, it strips off the matching part (in this case http://mysite.com) and sends the remaining text "/4" to "polls.urls" URLconf for further processing. There it matches <int:question_id>/ and calls detail() view.
- The question_id=4 part comes from <int:question_id>. Using angle brackets "captures" part of the URL and sends it as a keyword argument to the view function. The question_id> part of the string defines the name that will be used to identify the matched pattern, and the <int: part is a converter that determines what patterns should match this part of the URL path. If we do not use the int part any random keyword will be matched like  "http://mysite.com/fdsfkds".

# Syntax of templates in Django

- A template is rendered with a context(a dict). Rendering replaces variables with their values which are looked up in the context and executes tags. Everything else is output as is.
- The syntax involves 4 constructs: Variables, Tags, Filters, Comments.

### Variables:

- Variables are surrounded by {{ and }} as shown below. Where first_name and last_name are keys of a dictionary. On rendering, the sentence will have values of first_name and last_name respectively.
> My first name is {{ first_name }} and my last name is {{ last_name }}  

- Dictionary lookup, attribute lookup and list-index lookups are implemented with a dot notation.
>{{ sample_dict.key }}  
{{ my_object.attribute }}  
{{ my_list.0 }}  

### Tags

- Tags provide arbitrary logic in the rendering process.
- It can output content, serve as control structure eg: "if" condition, "for" loop; grab content from database and even enable access to other template tags.
- This is the entire reference for what we can use in tags: https://docs.djangoproject.com/en/2.2/ref/templates/builtins/#ref-templates-builtins-tags
- They are written in format shown below.  
> {% tag_stuff %}  

### Filters

- Do not go by it's name, Filters actually *transform* the values of variables and tag arguments.
- They look like this:
> {{ topic|title }}  
With a context {"topic": "Current state of democracy in India "}  

- Reference for built in filters: https://docs.djangoproject.com/en/2.2/ref/templates/builtins/#ref-templates-builtins-filters

### Comments

> {# This won't be rendered. #}  

# Reading arguments from the request object.

- request.POST or request.GET is a dictionary like object that lets you access submitted data by key name.
- In case of POST method submitted data has to come from a form.
- It’s possible that a request can come in via POST with an empty POST dictionary – if, say, a form is requested via the POST HTTP method but does not include form data. Therefore, you shouldn’t use ```if request.POST``` to check for use of the POST method; instead, use ```if request.method == "POST"```


# Class based views - Generic Views

- The views "detail()" "results()" written in polls app are simple and also redundant. If we look closely, it's the same code as in "index()" view. To tackle this Django provides a shortcut, called the "Generic views" system.
- Generic views abstract the common patterns to the point where you don't even have to write python code to write an app.
- Main intro reference: https://docs.djangoproject.com/en/2.2/topics/class-based-views/
- The views here follow a particular syntax.
- Use "template_name" to use a specific template.
- Use "context_object_name" to use a specific object name for your context.
- Use "model" to mention which model will the view operate on.
- Reference for ListView: https://docs.djangoproject.com/en/2.2/ref/class-based-views/generic-display/#listview
- Reference for DetailView: https://docs.djangoproject.com/en/2.2/ref/class-based-views/generic-display/#detailview
- By default, the DetailView generic view uses a template called ```<app name>/<model name>_detail.html```. In our case, it would use the template "polls/question_detail.html". The template_name attribute is used to tell Django to use a specific template name instead of the autogenerated default template name. We also specify the template_name for the results list view – this ensures that the results view and the detail view have a different appearance when rendered, even though they’re both a DetailView behind the scenes.
- Similarly, the ListView generic view uses a default template called ```<app name>/<model name>_list.html```; we use template_name to tell ListView to use our existing "polls/index.html" template.
- For DetailView the question variable is provided automatically – since we’re using a Django model (Question), Django is able to determine an appropriate name for the context variable. However, for ListView, the automatically generated context variable is question_list. To override this we provide the context_object_name attribute, specifying that we want to use latest_question_list instead.
