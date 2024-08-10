from flask import Flask, render_template, url_for

app = Flask(__name__)

@app.route('/')
def index():
    return render_template("index.html")

@app.route('/<string:name>/<int:id>')
def workflow(name, id):
    return "GET " + name + " CODE " + str(id)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=True)
