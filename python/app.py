from flask import Flask, render_template, url_for, request, redirect
import hashlib

def generate_id_key(text):
    hash_object = hashlib.sha256(text.encode())
    hash_hex = hash_object.hexdigest()
    key = int(hash_hex[:8], 16)
    key_str = f"{key:08}"
    return key_str[::-1]

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
    return 'GET ' + name + ' CODE ' + str(id)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=True)
