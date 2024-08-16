from flask import Flask, render_template, url_for, request, redirect
import hashlib
import os
import requests

config = {
    "port": os.environ.get('PORT', 5000),
    "debug": os.environ.get('DEBUG', False),
    "backend": os.environ.get('BACKEND', 'backend'),
    "backendport": os.environ.get('BACKENDPORT', 1234)
}

def generate_id_key(text):
    hash_object = hashlib.sha256(text.encode())
    hash_hex = hash_object.hexdigest()
    key = int(hash_hex[:8], 16)
    key_str = f"{key:08}"
    return key_str[::-1]

def check_user(name, id):
    backend = config["backend"]
    backendport = config["backendport"]
    response = requests.get(f"http://{backend}:{backendport}/api/user?name={name}&id={id}")
    if response.status_code == 200:
        data = response.json()
        if data.get("exists", False):
            return 0
        else:
            return 1
    elif response.status_code == 403:
        return 2
    else:
        return 1

def create_user(name, id):
    backend = config["backend"]
    backendport = config["backendport"]
    url = f"http://{backend}:{backendport}/api/create_user"
    payload = {"name": name, "id": id}
    headers = {"Content-Type": "application/json"}

    response = requests.post(url, json=payload, headers=headers)

    if response.status_code == 200:
        data = response.json()
        return True, data.get("message", "User created successfully")
    else:
        return False, f"Failed to create user: {response.status_code}"

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
        exist = check_user(name, id)
        if exist == 0:
            return render_template('index.html')
        elif exist == 1:
            success, message = create_user(name, str(id))
            if success:
                return render_template('index.html')
            else:
                return message
        else:
            return "403"

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=config["port"], debug=config["debug"])
