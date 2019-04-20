from flask import render_template
from flask_appbuilder.models.sqla.interface import SQLAInterface
from flask_appbuilder import ModelView, BaseView, expose
from app import appbuilder, db

"""
    Create your Views::


    class MyModelView(ModelView):
        datamodel = SQLAInterface(MyModel)


    Next, register your Views::


    appbuilder.add_view(MyModelView, "My View", icon="fa-folder-open-o", category="My Category", category_icon='fa-envelope')
"""

"""
    Application wide 404 error handler
"""
class MyView(BaseView):
    route_base = "/myview"

    @expose('/method1/<string:paramOne>')
    @has_access
    def methodOne(self, paramOne):
        res = "What up.... {}".format(paramOne)
        return res

    @expose('/method2/<string:paramOne>')
    @has_access
    def methodTwo(self, paramOne):
        res = "This is from a different mentthod {} see? Get my point.".format(paramOne)
        return res

appbuilder.add_view_no_menu(MyView())

@appbuilder.app.errorhandler(404)
def page_not_found(e):
    return render_template('404.html', base_template=appbuilder.base_template, appbuilder=appbuilder), 404

db.create_all()


