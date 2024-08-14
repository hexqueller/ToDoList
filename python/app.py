from flask import Flask, render_template, url_for, request, redirect
import hashlib
import os
import requests

config = {
    "port": os.environ.get('PORT', 5000),
    "debug": os.environ.get('DEBUG', False)
}

def generate_id_key(text):
    hash_object = hashlib.sha256(text.encode())
    hash_hex = hash_object.hexdigest()
    key = int(hash_hex[:8], 16)
    key_str = f"{key:08}"
    return key_str[::-1]

def check_user_exists(name):
    backend = "backend"
    port = "1234"
    response = requests.get(f"http://{backend}:{port}/api/user?name={name}")
    if response.status_code == 200:
        data = response.json()
        return data.get("exists", False)
    return False

app = Flask(__name__)

@app.route('/', methods=['GET', 'POST'])
def index():
    if request.method == 'POST':
        text = request.form['text']
        id = generate_id_key(text)
        return redirect(url_for('workflow', name=text, id=id))
    else:
        return render_template('index.html')

@app.route('/<string:name>/<int:id>')
def workflow(name, id):
    if str(id) != generate_id_key(name):
        return "403"
    else:
        if check_user_exists(name):
            return render_template('index.html')
        else:
            return render_template('master.html')

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=config["port"], debug=config["debug"])
